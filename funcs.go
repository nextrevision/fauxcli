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

func toInt(v interface{}) int {
	r := v.(*int)
	return *r
}

func toFloat(v interface{}) float64 {
	r := v.(*float64)
	return *r
}

func count(n int) []struct{} {
	return make([]struct{}, n)
}
