package employee

type Service interface {
	GetEmployees() ([]Employee, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetEmployees() ([]Employee, error) {
	employees, err := s.repository.FindAll()
	if err != nil {
		return employees, err
	}

	return employees, nil
}
