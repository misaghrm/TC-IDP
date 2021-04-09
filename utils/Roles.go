package utils

import (
	"database/sql"
	"tc-micro-idp/models"
)

var RolesConst = []models.Role{
	{
		Id: sql.NullInt64{
			Int64: -1,
			Valid: false,
		},
		Name:      "root",
		Title:     "ریشه",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id: sql.NullInt64{
			Int64: -2,
			Valid: false,
		},
		Name:      "super-admin",
		Title:     "مدیر کل سیستم",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id: sql.NullInt64{
			Int64: -3,
			Valid: false,
		},
		Name:      "admin",
		Title:     "مدیر سیستم",
		IsInHouse: true,
		Visible:   true,
	},
	{
		Id: sql.NullInt64{
			Int64: -4,
			Valid: false,
		},
		Name:      "developer",
		Title:     "برنامه‌نویس",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id: sql.NullInt64{
			Int64: -5,
			Valid: false,
		},
		Name:      "backend-developer",
		Title:     "برنامه‌نویس بکند",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id: sql.NullInt64{
			Int64: -6,
			Valid: false,
		},
		Name:      "frontend-developer",
		Title:     "برنامه‌نویس فرانت",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id: sql.NullInt64{
			Int64: -7,
			Valid: false,
		},
		Name:      "android-developer",
		Title:     "برنامه‌نویس اندروید",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id: sql.NullInt64{
			Int64: -8,
			Valid: false,
		},
		Name:      "devops",
		Title:     "مهندس دوآپس",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id: sql.NullInt64{
			Int64: -9,
			Valid: false,
		},
		Name:      "partner",
		Title:     "شریک تجاری",
		IsInHouse: false,
		Visible:   true,
	},
	{
		Id: sql.NullInt64{
			Int64: -10,
			Valid: false,
		},
		Name:      "supporter",
		Title:     "پشتیبان",
		IsInHouse: true,
		Visible:   true,
	},
	{
		Id: sql.NullInt64{
			Int64: -11,
			Valid: false,
		},
		Name:      "operator",
		Title:     "اپراتور",
		IsInHouse: true,
		Visible:   true,
	},
	{
		Id: sql.NullInt64{
			Int64: -12,
			Valid: false,
		},
		Name:      "end-user",
		Title:     "کاربر برنامه",
		IsInHouse: false,
		Visible:   false,
	},
	{
		Id: sql.NullInt64{
			Int64: -13,
			Valid: false,
		},
		Name:      "android-user",
		Title:     "کاربر اندروید",
		IsInHouse: false,
		Visible:   false,
	},
	{
		Id: sql.NullInt64{
			Int64: -14,
			Valid: false,
		},
		Name:      "pwa-user",
		Title:     "کاربر اپلیکیشن وب",
		IsInHouse: false,
		Visible:   false,
	},
	{
		Id: sql.NullInt64{
			Int64: -15,
			Valid: false,
		},
		Name:      "supplier",
		Title:     "تامین",
		IsInHouse: true,
		Visible:   true,
	}}
