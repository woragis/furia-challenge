package user

import (
	"database/sql"

	"github.com/google/uuid"
)

type Repository interface {
	CreateUser(user *User) (uuid.UUID, error)
	CreateStudent(student *Student) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateUser(user *User) (uuid.UUID, error) {
	var id uuid.UUID
	err := r.db.QueryRow(`
		INSERT INTO users (name, email, password_hash, role)
		VALUES ($1, $2, $3, $4) RETURNING user_id
	`, user.Name, user.Email, user.PasswordHash, user.Role).Scan(&id)
	return id, err
}

func (r *repository) CreateStudent(student *Student) error {
	_, err := r.db.Exec(`
		INSERT INTO students (user_id, date_of_birth, grade, location)
		VALUES ($1, $2, $3, $4)
	`, student.UserID, student.DateOfBirth, student.Grade, student.Location)
	return err
}
