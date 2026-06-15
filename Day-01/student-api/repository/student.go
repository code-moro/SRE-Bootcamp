package repository

import (
	"context"

	"student-api/database"
	"student-api/models"
)

func AddStudent(student models.Student) error {

	query := `
		INSERT INTO students(name, email)
		VALUES($1, $2)
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		student.Name,
		student.Email,
	)

	return err
}

func GetStudents() ([]models.Student, error) {

	query := `
		SELECT id, name, email
		FROM students
		ORDER BY id
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []models.Student

	for rows.Next() {

		var student models.Student

		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Email,
		)

		if err != nil {
			return nil, err
		}

		students = append(
			students,
			student,
		)
	}

	return students, nil
}

func GetStudent(id int) (*models.Student, error) {

	query := `
		SELECT id, name, email
		FROM students
		WHERE id = $1
	`

	var student models.Student

	err := database.DB.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&student.ID,
		&student.Name,
		&student.Email,
	)

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func UpdateStudent(
	id int,
	student models.Student,
) error {

	query := `
		UPDATE students
		SET name = $1,
		    email = $2
		WHERE id = $3
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		student.Name,
		student.Email,
		id,
	)

	return err
}

func DeleteStudent(id int) error {

	query := `
		DELETE FROM students
		WHERE id = $1
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}