package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func main() {
	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	e.GET("/students", func(c echo.Context) error {
		rows, err := db.Query(context.Background(),
			"SELECT id, firstname, lastname FROM students")
		if err != nil {
			return c.JSON(500, err.Error())
		}
		defer rows.Close()

		result := []map[string]any{}

		for rows.Next() {
			var id int
			var firstname, lastname string

			rows.Scan(&id, &firstname, &lastname)

			result = append(result, map[string]any{
				"id":        id,
				"firstname": firstname,
				"lastname":  lastname,
			})
		}

		return c.JSON(http.StatusOK, result)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
