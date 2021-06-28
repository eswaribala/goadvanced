package models

type Date struct {
	Day int
	Month int
	Year int
}

type Vehicle struct{
	RegistrationNo string
	Maker string
	DOR *Date
	ChassisNo string
	EngineNo string
	TypeofFuel string
	Color string
}
