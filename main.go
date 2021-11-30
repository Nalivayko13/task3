package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

func main() {
	list,_:=json.Marshal(studentsList)
	fmt.Println(string(list))
	json.Unmarshal(list,&Data)

	r:=mux.NewRouter()
	r.HandleFunc("/students/",GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{key}",GetStudentsById).Methods("GET")
	r.HandleFunc("/students/grade/{key}",GetStudentsByGrade).Methods("GET")
	r.HandleFunc("/",PostNewStudents).Methods("POST")
	r.HandleFunc("/students/{key}", DeleteStudent).Methods("DELETE")
	http.ListenAndServe("localhost:8080", r)

	for _,i:=range Data{
		fmt.Println(i)
	}


}

func PostNewStudents(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var post Student
	_ = json.NewDecoder(r.Body).Decode(&post)
	for _,i:=range Data{
		if id:=len(Data)+1; strconv.Itoa(id)==i.ID{
			post.ID=strconv.Itoa(id+1)
		}
	}
	Data = append(Data, post)
	json.NewEncoder(w).Encode(post)
}

func DeleteStudent(w http.ResponseWriter,r *http.Request){
	for i,data:= range Data {
		if data.ID==mux.Vars(r)["key"]{
			copy(Data[i:], Data[i+1:])
			Data=Data[:len(Data)-1]
		}
	}
	json.NewEncoder(w).Encode(Data)
}


func GetAllStudents(w http.ResponseWriter,r *http.Request){
	//json.NewEncoder(w).Encode(Data)
	t,_:=template.ParseFiles("mocks/students.html")
	t.Execute(w,Data)
}

func GetStudentsById(w http.ResponseWriter, r *http.Request){
	for _,i:= range Data {
		if i.ID==mux.Vars(r)["key"]{
			json.NewEncoder(w).Encode(i.Info())
		}
	}
}
func GetStudentsByGrade(w http.ResponseWriter,r *http.Request){
	for _,i:= range Data {
		if i.Grade==mux.Vars(r)["key"]{
			json.NewEncoder(w).Encode(i.Info())
		}
	}
}

