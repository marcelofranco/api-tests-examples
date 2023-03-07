package data

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

const dbTimeout = time.Second * 3

var db *gorm.DB

type PostgresRepository struct {
	Conn *gorm.DB
}

// NewPostgresRepository returns a new postgres repository
func NewPostgresRepository(pool *gorm.DB) *PostgresRepository {
	db = pool
	return &PostgresRepository{
		Conn: pool,
	}
}

type Student struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	RG        string    `json:"rg"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *PostgresRepository) GetStudent(id int) (*Student, error) {
	var student *Student
	if err := u.Conn.First(&student, id).Error; err != nil {
		return student, errors.New("student not found")
	}
	return student, nil
}

func (u *PostgresRepository) CreateStudent(student Student) (uint, error) {
	if err := u.Conn.Create(&student).Error; err != nil {
		return 0, errors.New("error creating student")
	}
	return student.ID, nil
}
