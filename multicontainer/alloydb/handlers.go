package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type todo struct {
	Item string
}

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("Go to: /list")
}

func listHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string
	rows, err := db.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occurred")
	}
	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newTodo := todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occurred: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newTodo)
	if newTodo.Item != "" {
		_, err := db.Exec("INSERT into todos VALUES ($1)", newTodo.Item)
		if err != nil {
			log.Fatalf("An error occurred while executing query: %v", err)
		}
	}

	return c.Redirect(listPath)
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	oldItem := c.Query("oldItem")
	newItem := c.Query("newItem")
	db.Exec("UPDATE todos SET item=$1 WHERE item=$2", newItem, oldItem)
	return c.Redirect(listPath)
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	todoToDelete := c.Query("item")
	db.Exec("DELETE from todos WHERE item=$1", todoToDelete)
	return c.SendString("deleted")
}

func pingHandler(c *fiber.Ctx, db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		c.SendString("Not Pong: " + err.Error())
	}
	return c.SendString("Pong: ok")
}
