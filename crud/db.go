package crud

import (
	"context"
	"fmt"
	"rucq_api/data_scheme"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const addr = "localhost:5432"
const user = "rucqUser"
const password = "mysecretpassword"
const db_name = "rucqDB"

func connectDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: db_name,
	})
	return db
}

func InitDB() {
	db := connectDB()
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")
	if err := createSchema(db); err != nil {
		panic(err)
	}

	db.Close()
	fmt.Println("Close DB")
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*data_scheme.User)(nil),
		(*data_scheme.Rooms)(nil),
	}
	fmt.Println("Creating DB Schema")
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			fmt.Println("Creating DB Schema:ERROR")
			return err
		}
	}
	fmt.Println("Creating DB Schema:SUCCESS")
	return nil
}

func AddUserDB(user *data_scheme.User) error {
	db := connectDB()
	defer db.Close()

	_, err := db.Model(user).Insert()
	return err
}

func GetUserDB() {

}

func GetMessagesDB() {

}

func AddMessageDB(Message *data_scheme.MesContainer) error {
	return nil
}

func AddRoomDB() {

}

func AddUserToRoomDB() {

}

func GetRoomDB() {

}
