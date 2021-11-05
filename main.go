package main

import (
	"context"
	"errors"
	"golang-ent-gqlgen-example/ent"
	"golang-ent-gqlgen-example/ent/article"
	"golang-ent-gqlgen-example/ent/user"
	"golang-ent-gqlgen-example/graph"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

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

	// Configure the GraphQL server and start
	srv := handler.NewDefaultServer(graph.NewSchema(client))
	{
		e.POST("/query", func(c echo.Context) error {
			srv.ServeHTTP(c.Response(), c.Request())
			return nil
		})

		e.GET("/playground", func(c echo.Context) error {
			playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	})
	e.POST("/users", func(c echo.Context) error {
		// Create an article entity
		a, err := client.Article.Create().
			SetTitle("title 1").
			SetDescription("description 1").
			Save(c.Request().Context())
		if !errors.Is(err, nil) {
			log.Fatalf("Error: failed creating article %v\n", err)
		}

		u, err := client.User.
			Create().
			SetName("Bob").
			SetAge(21).
			AddArticles(a). // Add article to the user
			Save(c.Request().Context())

		if !errors.Is(err, nil) {
			log.Fatalf("Error: failed creating user %v\n", err)
		}

		return c.JSON(http.StatusCreated, u)
	})
	e.GET("/users", func(c echo.Context) error {
		u, err := client.User.
			Query().
			WithArticles().
			Where(user.IDEQ(1)).
			Only(c.Request().Context())

		if !errors.Is(err, nil) {
			log.Fatalf("Error: failed quering users %v\n", err)
		}

		return c.JSON(http.StatusOK, u)
	})

	e.GET("/article/user", func(c echo.Context) error {
		u, err := client.Article.
			Query().
			Where(article.IDEQ(1)).
			QueryUser().
			Only(c.Request().Context())

		if !errors.Is(err, nil) {
			log.Fatalf("Error: failed quering article %v\n", err)
		}

		return c.JSON(http.StatusOK, u)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
