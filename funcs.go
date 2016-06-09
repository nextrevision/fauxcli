package main

import "github.com/Pallinder/go-randomdata"

func randomFullName() string {
	return randomdata.FullName(randomdata.RandomGender)
}

func toString(v interface{}) string {
	r := v.(*string)
	return *r
}

func toBool(v interface{}) bool {
	r := v.(*bool)
	return *r
}
