package main

type Student struct {
	ID string `json:"id"`
	Grade string `json:"grade"`
	Name string `json:"name"`
}

var Data []Student
var studentsList =[]Student{
	{"1", "11","Gena"},
	{"2","10","Sveta"},
	{"3","11","ololosha"},
	{"4","8","lol"},
}
func (s Student)Info() string{
	return " id:"+s.ID+" name:"+s.Name+" grade:"+s.Grade
}