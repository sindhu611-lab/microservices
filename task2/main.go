package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var sqlDB *sql.DB

type employee struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Id      int    `json:"id"`
}

var allemployees []employee

func createuser(w http.ResponseWriter, r *http.Request) {
	Employee := employee{}
	var isdataexists bool = false
	w.Header().Set("content-type", "application/json")
	json.NewDecoder(r.Body).Decode(&Employee)
	fmt.Println(Employee.Company, Employee.Name)
	fmt.Println(sqlDB)

	res, err := sqlDB.Exec("INSERT INTO employees.users (name, company) VALUES (?, ?)", Employee.Name, Employee.Company)
	if err != nil {
		fmt.Println("error inserting data", err)
		return
	}

	userID, err := res.LastInsertId()
	if err != nil {
		fmt.Println("error user", err)
	}

	fmt.Println("User id is ", userID)
	return
	for _, value := range allemployees {
		if value.Name == Employee.Name {
			isdataexists = true
			break
		}
	}
	if isdataexists {
		w.Write([]byte("employee already exists"))
		return
	}
	allemployees = append(allemployees, Employee)
	w.Write([]byte("employee created successfully"))

}
func getallemployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allemployees)
}
func updateuser(w http.ResponseWriter, r *http.Request) {
	Employee := employee{}
	var isdataexists bool = false
	json.NewDecoder(r.Body).Decode(&Employee)
	for i, value := range allemployees {
		if value.Company == Employee.Company {
			allemployees[i].Id = Employee.Id
			allemployees[i].Name = Employee.Name
			isdataexists = true

		}
	}
	if !isdataexists {
		w.Write([]byte("employee does not exist"))
		return
	}
	w.Write([]byte("employee updated successfully"))

}
func deleteuser(w http.ResponseWriter, r *http.Request) {
	Employee := employee{}
	var isdataexists bool = false
	json.NewDecoder(r.Body).Decode(&Employee)
	//allemployees = []employee{}
	for i, value := range allemployees {
		if value.Name == Employee.Name {
			allemployees = append(allemployees[:i], allemployees[i+1:]...)
			isdataexists = true
		}

	}
	if !isdataexists {
		w.Write([]byte("throws an error"))
		return
	}
	w.Write([]byte(" employee information deleted  successfully"))

}

func main() {
	mysqlURL := "user:pass@tcp(127.0.0.1:3306)/employees"
	var err error
	sqlDB, err = sql.Open("mysql", mysqlURL)
	fmt.Println(sqlDB, "105")
	if err != nil {
		fmt.Println("error to the connection database", err)
		panic(err)
	}

	//defer sqlDB.Close()

	router := http.NewServeMux()
	router.HandleFunc("POST /employee", createuser)
	router.HandleFunc("GET /employee", getallemployees)
	router.HandleFunc("PUT /employee", updateuser)
	router.HandleFunc("DELETE /employee", deleteuser)
	http.ListenAndServe(":1234", router)

}
