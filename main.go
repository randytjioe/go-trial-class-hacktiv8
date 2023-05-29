package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

const (
	dbUsername = "root"
	dbPassword = "password"
	dbName     = "mydatabase"
)

func main() {
	var err error

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", dbUsername, dbPassword, dbName))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	Register()
}

func Register() {
	var username, password string

	fmt.Println("=== Register ===")
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	err := saveUser(username, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Registration successful!")
	Login()
}

func Login() {
	var username, password string

	fmt.Println("\n=== Login ===")
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)

	user, err := getUser(username, password)
	if err != nil {
		log.Fatal(err)
	}

	if user != nil {
		fmt.Printf("Selamat Datang, %s!\n", user.Username)
	} else {
		fmt.Println("Username atau Password Anda Salah")
		fmt.Println("Silahkan Login Kembali")
		Login()
	}
}

type User struct {
	Username string
	Password string
}

func saveUser(username, password string) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := db.Exec(query, username, password)
	return err
}

func getUser(username, password string) (*User, error) {
	query := "SELECT username, password FROM users WHERE username = ? AND password = ?"
	row := db.QueryRow(query, username, password)

	user := &User{}
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
