package service

import "github.com/VelVit24/models"

func (s *Service) GetLoaders() ([]models.Loader, error) {
	return s.repo.SelectLoaders()
}
func (s *Service) GetManualLoaders() ([]models.ManualLoader, error) {
	return s.repo.SelectManualLoaders()
}
func (s *Service) GetLoaderImage(id int) (string, string, error) {
	return s.repo.SelectLoaderImage(id)
}
func (s *Service) GetManualLoaderImage(id int) (string, string, error) {
	return s.repo.SelectManualLoaderImage(id)
}
