package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
)

func main() {
	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()
	e.GET("/student/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, "wrong id")
		}

		var sid, groupID int
		var firstname, lastname, gender, birthDate, groupName string

		q := `
			SELECT s.id, s.firstname, s.lastname, s.gender, s.birth_date::text, s.group_id, g.name
			FROM students s
			JOIN groups g ON g.id = s.group_id
			WHERE s.id = $1
		`

		err = db.QueryRow(context.Background(), q, id).Scan(
			&sid, &firstname, &lastname, &gender, &birthDate, &groupID, &groupName,
		)
		if err != nil {
			return c.JSON(404, "student not found")
		}

		return c.JSON(200, map[string]any{
			"id":         sid,
			"firstname":  firstname,
			"lastname":   lastname,
			"gender":     gender,
			"birth_date": birthDate,
			"group_id":   groupID,
			"group_name": groupName,
		})
	})
	
	e.GET("/all_class_schedule", func(c echo.Context) error {
		rows, err := db.Query(context.Background(),
			"SELECT id, subject, day_of_week, lesson_time, group_id, faculty_id FROM schedule ORDER BY id")
		if err != nil {
			return c.JSON(500, err.Error())
		}
		defer rows.Close()

		result := []map[string]any{}

		for rows.Next() {
			var id, groupID, facultyID int
			var subject, day, timeStr string

			err := rows.Scan(&id, &subject, &day, &timeStr, &groupID, &facultyID)
			if err != nil {
				return c.JSON(500, err.Error())
			}

			result = append(result, map[string]any{
				"id":          id,
				"subject":     subject,
				"day_of_week": day,
				"lesson_time": timeStr,
				"group_id":    groupID,
				"faculty_id":  facultyID,
			})
		}

		return c.JSON(http.StatusOK, result)
	})

	e.GET("/schedule/group/:id", func(c echo.Context) error {
		groupID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, "wrong group id")
		}

		rows, err := db.Query(context.Background(),
			"SELECT id, subject, day_of_week, lesson_time, group_id, faculty_id FROM schedule WHERE group_id=$1 ORDER BY id",
			groupID)
		if err != nil {
			return c.JSON(500, err.Error())
		}
		defer rows.Close()

		result := []map[string]any{}

		for rows.Next() {
			var id, gid, facultyID int
			var subject, day, timeStr string

			err := rows.Scan(&id, &subject, &day, &timeStr, &gid, &facultyID)
			if err != nil {
				return c.JSON(500, err.Error())
			}

			result = append(result, map[string]any{
				"id":          id,
				"subject":     subject,
				"day_of_week": day,
				"lesson_time": timeStr,
				"group_id":    gid,
				"faculty_id":  facultyID,
			})
		}

		return c.JSON(200, result)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
