package main

import (
	"log"

	_ "github.com/lib/pq"
	//"Task1/database"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}
type Product struct {
	gorm.Model // Embedded struct for standard fields (ID, CreatedAt, UpdatedAt, DeletedAt)
	Code       string
	Price      uint
}

func CreateUser(db *gorm.DB) error {
	user := User{
		Username: "postgres",
		Password: "admin",
	}

	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// func ConnectDB() (*gorm.DB, error) {
// 	dsn := "postgres://postgres:admin@localhost/test?sslmode=disable"
// 	// Replace with your PostgreSQL connection string

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

//		return db, nil
//	}
func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	//defer db.Close()
	//AutoMigrate your database schemas - creates tables based on the defined models
	db.AutoMigrate(&User{})

	err = CreateUser(db)
	if err != nil {
		log.Fatal("Failed to create user:", err)
	}

	log.Println("User created successfully")

}
