package models

import (
	"database/sql"
	"log"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
)

// ตัวแปรสำหรับฐานข้อมูล
var DB *sql.DB

// ฟังก์ชันเชื่อมต่อฐานข้อมูล
func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

// ฟังก์ชันแฮชรหัสผ่าน
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ฟังก์ชันเพิ่มผู้ใช้ใหม่
func CreateUser(username, password, email string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	// เพิ่มผู้ใช้ใหม่ลงในฐานข้อมูล
	_, err = DB.Exec("INSERT INTO tbl_user (username, password, email) VALUES (?, ?, ?)", username, hashedPassword, email)
	if err != nil {
		return err
	}
	return nil
}

// ฟังก์ชันตรวจสอบรหัสผ่าน
func CheckLogin(username, password string) (bool, error) {
	var storedPassword string
	err := DB.QueryRow("SELECT password FROM tbl_user WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // ไม่มีผู้ใช้
		}
		return false, err
	}

	// ตรวจสอบรหัสผ่าน
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return false, nil // รหัสผ่านไม่ตรง
	}
	return true, nil
}
