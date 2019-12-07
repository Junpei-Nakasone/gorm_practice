package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	Name  string `json:"name"`
	Email string
}

type RequestParams struct {
	WhereList []WhereColumnList `json:"where_column_list"`
	OrderList []OrderColumnList `json:"order_list"`
	Limit     int               `json:"limit"`
}

type WhereColumnList struct {
	WhereName  string `json:"where_name"`
	WhereParam string `json:"where_param"`
}

type OrderColumnList struct {
	OrderName  string `json:"order_name"`
	OrderParam string `json:"order_param"`
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New User Endpoint Hit")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	fmt.Println(name)
	fmt.Println(email)

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}

func selectUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users []User

	var params []RequestParams
	// メソッドチェーンで繋ぐ

	// WhereColumnList := params[0].WhereList
	// whereName := WhereColumnList[0].WhereName
	// WhereParam := WhereColumnList[0].WhereParam

	limit := params[0].Limit
	OrderColumnList := params[0].OrderList

	// Order条件の長さを変数OrderLengthに格納できる？
	OrderLength := len(OrderColumnList)

	// for文に入る前に変数txを宣言し、即時メソッドを実行するまでtxに検索条件を足していく
	tx := db.Limit(limit)

	for i := 0; i < OrderLength; i++ {
		OrderName := OrderColumnList[i].OrderName
		OrderParam := OrderColumnList[i].OrderParam

		tx = db.Order(OrderName[i] + OrderParam[i])
	}

	// 検索条件を格納した変数txで&users構造体に即時メソッドを実行
	tx.Find(&users)

	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	myRouter.HandleFunc("/select", selectUsers).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func main() {
	fmt.Println("Go ORM Tutorial")

	initialMigration()
	// Handle Subsequent requests
	handleRequests()
}
