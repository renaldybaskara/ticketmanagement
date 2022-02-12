package main

import (
	"log"
	"ticketmanagement/handler"
	"ticketmanagement/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//koneksi ke database
	dsn := "root:@tcp(127.0.0.1:3306)/ticket_management?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	//Repository Save Struct User ke Database User
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// ruleticketRepository := ruleticket.NewRepository(db)

	// Handler mapping input user ke Struct Input (Service)
	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()

	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser) //api ini aka masuk ke func userHandler.RegisterUser
	router.Run()

	// fmt.Println("Connection to Database is good")
	// membuat variable untuk define array user.user
	// var users []user.User
	//membuat variable length untuk menghitung data yang ada di array berapa banyak
	// length := len(users)
	// fmt.Println(length)

	//untuk melakukan pengecekan di Table Users melalui variable users yang telah terhubung dgn struct yang telah dibuat
	// db.Find(&users)

	// length = len(users)
	// fmt.Println(length)

	// for _, user := range users {
	// 	fmt.Println(user)
	// 	fmt.Println("========================================================")
	// }

	// untuk membuat router agar bisa diakses di web browser
	// 	router := gin.Default()
	// 	router.GET("/handler", handler)
	// 	router.Run()
}

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/ticket_management?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User
// 	db.Find(&users)
// membuat json dgn menampilkan status okey dan menampilkan dari database users
// c.JSON(http.StatusOK, users)

//langkah-langkah register user :
// 1. Input form dari user
// 2. Handler mapping input user ke struct input
// 3. Service mapping struct input ke struct user
// 4. Repository save struct ke db user
// 5. db
// }
