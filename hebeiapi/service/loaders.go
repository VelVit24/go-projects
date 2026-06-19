package service

import "github.com/VelVit24/models"

func (s *Service) GetLoaders() ([]models.Loader, error) {
	return s.repo.SelectLoaders()
}
func (s *Service) GetManualLoaders() ([]models.ManualLoader, error) {
	return s.repo.SelectManualLoaders()
}
func (s *Service) GetLoaderImage(id int) (string, error) {
	linux, _, err := s.repo.SelectLoaderImage(id)
	if err != nil {
		return "", err
	}
	return linux, nil
}
func (s *Service) GetManualLoaderImage(id int) (string, error) {
	linux, _, err := s.repo.SelectManualLoaderImage(id)
	if err != nil {
		return "", err
	}
	return linux, nil
}
