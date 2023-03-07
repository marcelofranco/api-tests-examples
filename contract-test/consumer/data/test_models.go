package data

import (
	"gorm.io/gorm"
)

type PosgrestTestRepository struct {
	Conn *gorm.DB
}

// NewPostgresRepository returns a new postgres repository
func NewPosgrestTestRepository(pool *gorm.DB) *PostgresRepository {
	db = pool
	return &PostgresRepository{
		Conn: pool,
	}
}

func (u *PosgrestTestRepository) GetStudent(id int) (*Student, error) {
	var student *Student
	student.ID = 1
	student.CPF = "13212312312"
	student.RG = "121231231"
	student.Name = "Test"
	return student, nil
}

func (u *PosgrestTestRepository) CreateStudent(student Student) (uint, error) {
	return 1, nil
}
