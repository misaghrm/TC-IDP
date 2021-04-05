package utils

import "tc-micro-idp/models"

//type Role struct {
//	Id        int64
//	Name      string
//	Title     string
//	IsInHouse bool
//}
//var RolesConst = []models.Role{{}}
var RolesConst = []models.Role{
	{
		Id:        -1,
		Name:      "root",
		Title:     "ریشه",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id:        -2,
		Name:      "super-admin",
		Title:     "مدیر کل سیستم",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id:        -3,
		Name:      "admin",
		Title:     "مدیر سیستم",
		IsInHouse: true,
		Visible:   true,
	},
	{
		Id:        -4,
		Name:      "developer",
		Title:     "برنامه‌نویس",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id:        -5,
		Name:      "backend-developer",
		Title:     "برنامه‌نویس بکند",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id:        -6,
		Name:      "frontend-developer",
		Title:     "برنامه‌نویس فرانت",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id:        -7,
		Name:      "android-developer",
		Title:     "برنامه‌نویس اندروید",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id:        -8,
		Name:      "devops",
		Title:     "مهندس دوآپس",
		IsInHouse: true,
		Visible:   false,
	},
	{
		Id:        -9,
		Name:      "partner",
		Title:     "شریک تجاری",
		IsInHouse: false,
		Visible:   true,
	},
	{
		Id:        -10,
		Name:      "supporter",
		Title:     "پشتیبان",
		IsInHouse: true,
		Visible:   true,
	},
	{
		Id:        -11,
		Name:      "operator",
		Title:     "اپراتور",
		IsInHouse: true,
		Visible:   true,
	},
	{
		Id:        -12,
		Name:      "end-user",
		Title:     "کاربر برنامه",
		IsInHouse: false,
		Visible:   false,
	},
	{
		Id:        -13,
		Name:      "android-user",
		Title:     "کاربر اندروید",
		IsInHouse: false,
		Visible:   false,
	},
	{
		Id:        -14,
		Name:      "pwa-user",
		Title:     "کاربر اپلیکیشن وب",
		IsInHouse: false,
		Visible:   false,
	},
	{
		Id:        -15,
		Name:      "supplier",
		Title:     "تامین",
		IsInHouse: true,
		Visible:   true,
	}}
