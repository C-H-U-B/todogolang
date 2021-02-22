package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func start() {
	todo := TodoList{}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.JSON(200, todo)
		return nil
	})

	e.POST("/todo", func(c echo.Context) error {
		var task TodoItem
		c.Bind(&task)
		todo.Add(task)
		return c.NoContent(200)
	})

	e.DELETE("/todo/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusNotAcceptable, "invalid id")
		}
		todo.Delete(id)
		return c.NoContent(200)
	})

	e.Start(":5000")

}
