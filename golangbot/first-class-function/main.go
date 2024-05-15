package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(students []student, fn func(student) bool) []student {
	var result []student
	for _, v := range students {
		if fn(v) == true {
			result = append(result, v)
		}
	}
	return result
}
func main() {
	s1 := student{
		firstName: "Ben",
		lastName:  "Gunn",
		grade:     "B",
		country:   "Myanmar",
	}
	s2 := student{
		firstName: "Jeff",
		lastName:  "Clay",
		grade:     "A",
		country:   "Thailand",
	}
	s := []student{s1, s2}
	fn := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(fn)
}
