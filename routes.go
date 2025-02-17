package routes

import (
	"github.com/gofiber/fiber/v2"
	"my_go_fiber/controllers"
)

func Setup(app *fiber.App) {
	// กำหนดเส้นทางสำหรับการลงทะเบียนและการเข้าสู่ระบบ
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
