package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "lenslocked_dev"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
	Color string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDb.Ping()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	name, email, color := getInfo()

	u := User{
		Name:  name,
		Email: email,
		Color: color,
	}

	fmt.Printf("%+v\n", u)
}

func getInfo() (name, email, color string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	fmt.Println("What is your email?")
	email, _ = reader.ReadString('\n')
	fmt.Println("What is your favourite color?")
	color, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	color = strings.TrimSpace(color)
	return name, email, color
}
