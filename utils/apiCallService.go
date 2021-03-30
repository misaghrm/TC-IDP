package utils

import (
	"github.com/kavenegar/kavenegar-go"
	"log"
)

const KavehNegarAPIKEY = "39576370486C516C583565504B6B566A556E62344E6563474D3070386F4A4D46325A68644E7A6B594D52453D"

func SendOtpCode(Phone string, OtpCode string, AppSignatureHash string) {
	api := kavenegar.New(KavehNegarAPIKEY)

	params := &kavenegar.VerifyLookupParam{
		Type: kavenegar.Type_VerifyLookup_Sms,
	}
	log.Println("Kavehnegar 1 ")
	if res, err := api.Verify.Lookup("0"+Phone, "verify", OtpCode, params); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			log.Println(err.Error())
		case *kavenegar.HTTPError:
			log.Println(err.Error())
		default:
			log.Println(err.Error())
		}
	} else {
		log.Println("MessageID 	= ", res.MessageID)
		log.Println("Status    	= ", res.Status)
		//...
	}
	log.Println("Kavehnegar 2 ")

}

//func SendOtpCode (Phone string,OtpCode string,AppSignatureHash string) {
//	//var c []*models.ParamaterMessage
//	_= []*models.ParamaterMessage{{
//		Key:   "token",
//		Value: OtpCode,
//	}}
//	a := models.SendMessageRequest{
//		Recipient:  Phone,
//		Type:       models.GrpcMessageType_General,
//		Parameters: []*models.ParamaterMessage{{
//			Key:   "token",
//			Value: OtpCode,
//		}},
//	}
//	//grpcCall(a)
//	//var conn *grpc.ClientConn
//	_ = alts.NewClientCreds(alts.DefaultClientOptions())
//	conn, err := grpc.Dial("cg-grpc.tazminchi.tzmnch:443",grpc.WithInsecure())
//
//	if err != nil {
//		log.Fatalf("rpc did not connect: %s", err)
//	}
//	defer conn.Close()
//	println(conn.GetState().String())
//
//	d := models.NewMessageGrpcClient(conn)
//
//	response, err := d.SendMessage(context.Background(), &a)
//	if err != nil {
//		log.Printf("Error when calling ComGateWay: %s", err)
//	}
//	println(conn.GetState().String())
//
//	println(response.String())
//	//log.Printf("Response from server: \ncode :%v,\nclaims: %s", response.GetOk(), response.ErrorMessage)
//	//c,_ := json.Marshal(a)
//	//
//	//client := http.DefaultClient
//	//resp,err := client.Post("cg-grpc.tazminchi.tzmnch:443",http.)
//	//f:= http.Request{
//	//	Body: a,
//	//}
//	//client.Do()
//
//}

//func grpcCall(a models.SendMessageRequest) {
//	var conn *grpc.ClientConn
//	conn, err := grpc.Dial("cg-grpc.tazminchi.tzmnch:443",grpc.WithInsecure())
//
//	if err != nil {
//		log.Fatalf("rpc did not connect: %s", err)
//	}
//	defer conn.Close()
//
//	c := models.NewMessageGrpcClient(conn)
//
//	response, err := c.SendMessage(context.Background(), &a)
//	if err != nil {
//		log.Fatalf("Error when calling ComGateWay: %s", err)
//	}
//	log.Printf("Response from server: \ncode :%v,\nclaims: %s", response.GetOk(), response.ErrorMessage)
//
//}
