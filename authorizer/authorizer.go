package authorizer

import (
	"context"
	"fmt"
	"log"
	"strings"
	"tc-micro-idp/dataManager"
	"tc-micro-idp/jwt"
	. "tc-micro-idp/models"
)

type Server struct {
}

func (s *Server) Authorize(ctx context.Context, req *Request) (*Response, error) {
	req.TrimDomain()
	req.Token = strings.TrimPrefix(req.Token, "Bearer ")
	internalModel, err := jwt.Decrypt(req.GetToken(), strings.ToLower(req.GetIssuer()))
	if err != nil {
		return &Response{
			Code: 403,
		}, nil
	}
	if db.IsBlocked(internalModel.Phone) {
		return &Response{
			Code: 403,
		}, nil
	}
	if req.URL != "" {
		AllowedRoles, err := db.GetRoles(req.GetURL())
		if err != nil {
			log.Println("error roles:", err)
			return &Response{
				Code: 403,
			}, nil
		}
		if f := IsAllowed(AllowedRoles, internalModel.GetRoles()); !f {
			log.Println("error IsAllowed : ", f)
			return &Response{
				Code: 403,
			}, nil
		}
	}
	return &Response{
		Code: 200,
		TokenClaims: &TokenClaim{
			TokenId:        fmt.Sprintf("%v", internalModel.TokenId),
			IssuedAt:       fmt.Sprintf("%v", internalModel.IssuedAt),
			UserId:         internalModel.UserId,
			Phone:          internalModel.Phone,
			RefreshVersion: fmt.Sprintf("%v", internalModel.RefreshVersion),
			EulaVersion:    fmt.Sprintf("%v", internalModel.EulaVersion),
			Issuer:         internalModel.Issuer,
			LifeTime:       internalModel.LifeTime,
			AccessVersion:  internalModel.AccessVersion,
			DeviceId:       internalModel.DeviceId,
			AppSource:      internalModel.AppSource,
			Roles:          internalModel.GetRoles(),
			Audience:       internalModel.Audience,
			Expires:        fmt.Sprintf("%v", internalModel.Expires),
			NotBefore:      fmt.Sprintf("%v", internalModel.NotBefore),
		},
	}, nil
}

func IsAllowed(requiredRoles []string, UserRoles []string) bool {
	if requiredRoles == nil {
		return true
	}
	for i := range requiredRoles {
		for j := range UserRoles {
			if strings.Contains(strings.ToLower(UserRoles[j]), strings.ToLower(requiredRoles[i])) {
				return true
			}
		}
	}
	return false
}
