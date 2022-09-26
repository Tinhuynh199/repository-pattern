package main

import (
	"fmt"
	"go-postgres/driver"
	"go-postgres/repository/repoimpl"
	models "go-postgres/model" 
)

//host, port, user, password, dbname
const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "123456"
	dbname   = "code4fun"
)

func main() {
	db := driver.Connect(host, port, user, password, dbname)

	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	UserRepo := repoimpl.NewUserRepo(db.SQL)
	hxt := models.User{
		ID: 1,
		Name: "Huynh Xuan Tin",
		Gender: "Nam",
		Email: "tinhuynhvl@gmail.com",
	}

	hxt1 := models.User{
		ID: 2,
		Name: "Huynh Xuan Tin 1",
		Gender: "Nam",
		Email: "tinhuynhvl1@gmail.com",
	}

	hxt2 := models.User{
		ID: 3,
		Name: "Huynh Xuan Tin 2",
		Gender: "Nam",
		Email: "tinhuynhvl2@gmail.com",
	}

	UserRepo.Insert(hxt)
	UserRepo.Insert(hxt1)
	UserRepo.Insert(hxt2)

	users, _ := UserRepo.Select()
	for i := range users {
		fmt.Println(users[i])
	}
	// fmt.Println("Connecting...")
}
