package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// เชื่อมต่อกับฐานข้อมูล MySQL
	dsn := "root:123456@tcp(127.0.0.1:3309)/gin_db" // ปรับข้อมูลนี้ตาม MySQL ของคุณ
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ตรวจสอบการเชื่อมต่อ
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// สร้าง Fiber app
	app := fiber.New()

	// สร้าง route สำหรับแสดงข้อมูลจาก MySQL
	app.Get("/users", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			return c.Status(500).SendString("Error fetching users")
		}
		defer rows.Close()

		var users []string
		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				return c.Status(500).SendString("Error reading rows")
			}
			users = append(users, fmt.Sprintf("ID: %d, Name: %s", id, name))
		}

		if err := rows.Err(); err != nil {
			return c.Status(500).SendString("Error with rows")
		}

		// ส่งข้อมูลที่ได้กลับไป
		return c.JSON(users)
	})

	// รันแอป
	app.Listen(":3000")
}
