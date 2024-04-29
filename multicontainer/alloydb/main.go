package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	defaultHTTPPort = "8080"

	driverName = "pgx"
	viewsDir   = "./views"
	viewsExt   = ".html"
	publicDir  = "./public"

	rootPath   = "/"
	listPath   = "/list"
	updatePath = "/update"
	deletePath = "/delete"
	pingPath   = "/ping"
)

func main() {
	// Connect to database
	db, err := sql.Open(driverName, connStr())
	if err != nil {
		log.Fatal(fmt.Errorf("sql.Open: %v", err))
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(7)
	db.SetConnMaxLifetime(1800 * time.Second)

	// Create HTTP server
	app := fiber.New(fiber.Config{
		Views: html.New(viewsDir, viewsExt),
	})

	app.Use(recover.New())
	app.Static(rootPath, publicDir)
	app.Get(rootPath, func(c *fiber.Ctx) error {
		return indexHandler(c)
	})
	app.Get(pingPath, func(c *fiber.Ctx) error {
		return pingHandler(c, db)
	})
	app.Get(listPath, func(c *fiber.Ctx) error {
		return listHandler(c, db)
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
		"host=%s user=%s password=%s port=%s database=%s sslmode=disable",
		os.Getenv("DB_HOSTNAME"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
}
