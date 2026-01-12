package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type AttendanceRequest struct {
	SubjectID int    `json:"subject_id"`
	VisitDay  string `json:"visit_day"` // "07.01.2026"
	Visited   bool   `json:"visited"`
	StudentID int    `json:"student_id"`
}

func RegisterAttendanceRoutes(e *echo.Echo, db *pgxpool.Pool) {
	e.POST("/attendance/subject", func(c echo.Context) error {
		var req AttendanceRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, "invalid json")
		}

		if req.SubjectID <= 0 || req.StudentID <= 0 || req.VisitDay == "" {
			return c.JSON(400, "subject_id, student_id, visit_day are required")
		}

		day, err := time.Parse("02.01.2006", req.VisitDay)
		if err != nil {
			return c.JSON(400, "visit_day format must be DD.MM.YYYY (e.g. 07.01.2026)")
		}

		q := `
        INSERT INTO attendance (subject_id, student_id, visit_day, visited)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (subject_id, student_id, visit_day)
        DO UPDATE SET visited = EXCLUDED.visited
        RETURNING id
    `
		var id int
		err = db.QueryRow(context.Background(), q, req.SubjectID, req.StudentID, day, req.Visited).Scan(&id)
		if err != nil {
			return c.JSON(500, err.Error())
		}

		return c.JSON(201, map[string]any{
			"id":         id,
			"subject_id": req.SubjectID,
			"student_id": req.StudentID,
			"visit_day":  day.Format("2006-01-02"),
			"visited":    req.Visited,
		})
	})

	e.GET("/attendanceBySubjectId/:id", func(c echo.Context) error {
		subjectID, err := strconv.Atoi(c.Param("id"))
		if err != nil || subjectID <= 0 {
			return c.JSON(400, "wrong subject id")
		}

		q := `
        SELECT
            a.id,
            s.id,
            s.firstname,
            s.lastname,
            a.visit_day::text,
            a.visited
        FROM attendance a
        JOIN students s ON s.id = a.student_id
        WHERE a.subject_id = $1
        ORDER BY a.visit_day, s.id
    `

		rows, err := db.Query(context.Background(), q, subjectID)
		if err != nil {
			return c.JSON(500, err.Error())
		}
		defer rows.Close()

		result := []map[string]any{}

		for rows.Next() {
			var attendanceID, studentID int
			var firstname, lastname, visitDay string
			var visited bool

			err := rows.Scan(&attendanceID, &studentID, &firstname, &lastname, &visitDay, &visited)
			if err != nil {
				return c.JSON(500, err.Error())
			}

			result = append(result, map[string]any{
				"attendance_id": attendanceID,
				"student_id":    studentID,
				"firstname":     firstname,
				"lastname":      lastname,
				"visit_day":     visitDay,
				"visited":       visited,
			})
		}

		return c.JSON(http.StatusOK, result)
	})

	e.GET("/attendanceByStudentId/:id", func(c echo.Context) error {
		studentID, err := strconv.Atoi(c.Param("id"))
		if err != nil || studentID <= 0 {
			return c.JSON(400, "wrong student id")
		}

		q := `
		SELECT
			a.id,
			a.subject_id,
			sub.name,
			a.visit_day::text,
			a.visited
		FROM attendance a
		JOIN subjects sub ON sub.id = a.subject_id
		WHERE a.student_id = $1
		ORDER BY a.visit_day, a.subject_id
	`

		rows, err := db.Query(context.Background(), q, studentID)
		if err != nil {
			return c.JSON(500, err.Error())
		}
		defer rows.Close()
		result := []map[string]any{}
		for rows.Next() {
			var attendanceID, subjectID int
			var subjectName, visitDay string
			var visited bool
			if err := rows.Scan(&attendanceID, &subjectID, &subjectName, &visitDay, &visited); err != nil {
				return c.JSON(500, err.Error())
			}
			result = append(result, map[string]any{
				"attendance_id": attendanceID,
				"student_id":    studentID,
				"subject_id":    subjectID,
				"subject_name":  subjectName,
				"visit_day":     visitDay,
				"visited":       visited,
			})
		}
		return c.JSON(http.StatusOK, result)
	})
}
