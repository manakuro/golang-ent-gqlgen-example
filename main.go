package main

import (
	"context"
	"errors"
	"golang-ent-gqlgen-example/ent"
	"golang-ent-gqlgen-example/ent/user"
	"log"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	loc, err := time.LoadLocation("Local")
	if !errors.Is(err, nil) {
		log.Fatalf("Error: load time location: %v\n", err)
	}
	mc := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "localhost" + ":" + "3307",
		DBName:               "golang_ent_gqlgen",
		AllowNativePasswords: true,
		ParseTime:            true,
		Loc:                  loc,
	}
	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	if err != nil {
		log.Fatalf("Error: mysql client: %v\n", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); !errors.Is(err, nil) {
		log.Fatalf("Error: failed creating schema resources %v\n", err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})
	e.POST("/users", func(c echo.Context) error {
		u, err := client.User.Create().SetName("Bob").SetAge(21).Save(c.Request().Context())
		if !errors.Is(err, nil) {
			log.Fatalf("Error: failed creating user %v\n", err)
		}

		return c.JSON(http.StatusCreated, u)
	})
	e.GET("/users", func(c echo.Context) error {
		us, err := client.User.
			Query().
			Where(user.IDEQ(1)).
			Only(c.Request().Context())

		if !errors.Is(err, nil) {
			log.Fatalf("Error: failed quering users %v\n", err)
		}

		return c.JSON(http.StatusOK, us)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
