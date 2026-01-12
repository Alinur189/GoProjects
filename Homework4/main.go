package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"Projectstudents/handlers"
)

func main() {
	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	handlers.RegisterStudentRoutes(e, db)
	handlers.RegisterAttendanceRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
