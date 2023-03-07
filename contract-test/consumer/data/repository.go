package data

type Repository interface {
	GetStudent(id int) (*Student, error)
	CreateStudent(student Student) (uint, error)
}
