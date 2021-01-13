package main

import (
	"server/controller"
	"server/jun"
)

func main() {
	baseGroup := jun.Default()

	bookInfoGroup := baseGroup.Group("/bookInfo")

	bookInfoGroup.Post("/add", controller.AddBookInfo)
	bookInfoGroup.Post("/update", controller.UpdateBookInfo)
	bookInfoGroup.Get("/get", controller.GetBookInfo)
	bookInfoGroup.Post("/delete", controller.DeleteBookInfo)

	baseGroup.Run()
}
