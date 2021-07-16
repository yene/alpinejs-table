package main

import (
	"github.com/beego/beego/v2/server/web"
)

var people []Person

type DataJSON struct {
	People      []Person `json:"people"`
	PeopleCount int      `json:"peopleCount"`
}

type PeopleDataController struct {
	web.Controller
}

func (this *PeopleDataController) Get() {
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 2
	}
	offset, _ := this.GetInt("offset")
	if offset > len(people) {
		offset = 0
	}

	low := offset
	high := offset + limit
	if high > len(people) {
		high = len(people)
	}

	s := people[low:high]
	d := DataJSON{People: s, PeopleCount: len(people)}
	this.Data["json"] = &d
	this.ServeJSON()
}

func main() {
	// Generate fake people list
	for i := 0; i < 19; i++ {
		p := NewPerson()
		people = append(people, p)
	}

	web.Router("/v1/people", &PeopleDataController{})

	web.SetStaticPath("/", "public")
	web.Run()

}
