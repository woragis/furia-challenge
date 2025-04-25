package user

import "golang.org/x/crypto/bcrypt"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) RegisterStudent(user *User, student *Student, plainPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hash)
	user.Role = "student"

	id, err := s.repo.CreateUser(user)
	if err != nil {
		return err
	}

	student.UserID = id
	return s.repo.CreateStudent(student)
}
