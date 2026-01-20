package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(name, email string) (*User, error) {
	u := &User{Name: name, Email: email}
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) GetAll() ([]User, error) {
	return s.repo.GetResults()
}

func (s *Service) Get(id int) (*User, error) {
	return s.repo.GetByID(id)
}
