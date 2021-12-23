package item

type IService interface {
	Create(item Item) error
	GetAll() (map[string]string, error)
	GetOne(key string) (Item, error)
	DeleteAll() error
	DeleteOne(key string) error
}

var _ IService = &Service{}

type Service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return Service{repo: repository}
}

func (s *Service) Create(item Item) error {
	err := s.repo.Create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (map[string]string, error) {
	allItem, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return allItem, nil
}

func (s *Service) GetOne(key string) (Item, error) {
	item, err := s.repo.GetOne(key)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func (s *Service) DeleteAll() error {
	err := s.repo.DeleteAll()
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteOne(key string) error {
	err := s.repo.DeleteOne(key)
	if err != nil {
		return err
	}
	return nil
}
