package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq"
)

const (
	defaultHTTPPort = "8080"

	driverName = "postgres"
	viewsDir   = "./views"
	viewsExt   = ".html"
	publicDir  = "./public"

	rootPath   = "/"
	updatePath = "/update"
	deletePath = "/delete"
)

func main() {
	// Connect to database
	db, err := sql.Open(driverName, connStr())
	if err != nil {
		log.Fatal(err)
	}

	// Create HTTP server
	app := fiber.New(fiber.Config{
		Views: html.New(viewsDir, viewsExt),
	})
	app.Static(rootPath, publicDir)
	app.Get(rootPath, func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})
	app.Post(rootPath, func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})
	app.Put(updatePath, func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})
	app.Delete(deletePath, func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultHTTPPort
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func connStr() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s/%s", // ?sslmode=disable
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOSTNAME"),
		os.Getenv("DB_DATABASE"),
	)
}