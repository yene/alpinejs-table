package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/Pallinder/go-randomdata"
)

type Person struct {
	CustomerID  string
	ContractID  string
	Username    string
	Gender      string
	Title       string
	Firstname   string
	Lastname    string
	Email       string
	City        string
	State       string
	Address     string
	PostalCode  string
	Birthdate   string
	PhoneNumber string
	ID          string
	Locale      string
	CreditScore string
}

func NewPerson() Person {
	p := Person{}
	p.CustomerID = fmt.Sprintf("%010d", randomdata.Number(1000000000))
	p.ContractID = randomdata.StringNumber(4, "-")

	gender := rand.Intn(2) // 0/1
	if gender == randomdata.Male {
		p.Gender = "male"
	} else {
		p.Gender = "female"
	}
	p.Username = strings.ToLower(randomdata.SillyName())
	p.Title = randomdata.Title(gender)
	p.Firstname = randomdata.FirstName(gender)
	p.Lastname = randomdata.LastName()
	p.Email = randomdata.Email()
	p.City = randomdata.City()
	p.State = randomdata.State(randomdata.Large)
	p.Address = randomdata.Address()
	p.PostalCode = randomdata.PostalCode("SE")
	p.Birthdate = randomdata.FullDate()
	p.PhoneNumber = randomdata.PhoneNumber()
	p.ID = fmt.Sprintf("%d-%d-%d",
		randomdata.Number(101, 999),
		randomdata.Number(01, 99),
		randomdata.Number(100, 9999),
	)
	p.Locale = randomdata.Locale()
	p.CreditScore = fmt.Sprintf("%d", randomdata.Number(101))
	return p
}
