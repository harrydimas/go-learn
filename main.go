// main.go
package main

import (
	"go-learn/handlers"
	"go-learn/models"
	"go-learn/repos"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	db, err := gorm.Open("postgres", "dbname=go_learn user=postgres password=qwerty sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.Item{})
	itemRepo := repos.NewItemRepository(db)
	itemHandler := handlers.NewItemHandler(itemRepo)
	r.POST("/items", itemHandler.CreateItem)
	r.GET("/items", itemHandler.GetItems)
	r.GET("/items/:id", itemHandler.GetItemByID)
	r.PUT("/items/:id", itemHandler.UpdateItem)
	r.DELETE("/items/:id", itemHandler.DeleteItem)
	r.Run(":3000")
}
