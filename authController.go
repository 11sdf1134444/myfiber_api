package controllers

import (
	"github.com/gofiber/fiber/v2"
	"my_go_fiber/models"
)

// ฟังก์ชันลงทะเบียนผู้ใช้
func Register(c *fiber.Ctx) error {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	// เพิ่มผู้ใช้ใหม่
	err := models.CreateUser(data.Username, data.Password, data.Email)
	if err != nil {
		return c.Status(500).SendString("Error creating user")
	}

	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

// ฟังก์ชันสำหรับเข้าสู่ระบบ
func Login(c *fiber.Ctx) error {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	// ตรวจสอบการเข้าสู่ระบบ
	success, err := models.CheckLogin(data.Username, data.Password)
	if err != nil {
		return c.Status(500).SendString("Error logging in")
	}
	if !success {
		return c.Status(401).SendString("Invalid username or password")
	}

	return c.JSON(fiber.Map{"message": "Login successful"})
}
