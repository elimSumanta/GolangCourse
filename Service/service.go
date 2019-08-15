package service

import (
	mod "github.com/elim/GoCourses/Model"
	api "github.com/elim/GoCourses/Repository/APIService"
	repo "github.com/elim/GoCourses/Repository/Database"
)

//Service struct
type Service struct {
	api  api.Repo
	repo repo.Repository
}

//New object of struct services
func New(apiRepo api.Repo, repo repo.Repository) *Service {
	s := Service{
		api:  apiRepo,
		repo: repo,
	}
	return &s
}

func (s *Service) GetCarByIDGerage(id string) []mod.GerageStatus {
	return s.repo.SelectCarByIDGerage(id)
}

func (s *Service) GetCarByIDCar(id string) mod.DetailCarStatus {
	carStatus := s.repo.SelectCarByIDCar(id)

	//Get Detail Data from another server
	detail := s.api.GetDetailData(id)
	//

	//Merge Array Data
	var resp mod.DetailCarStatus
	for _, value := range carStatus {
		resp.OwnerName = value.OwnerName
		resp.CarName = value.CarName
		resp.IDCar = value.IDCar
	}

	for _, detail := range detail {
		resp.GarageName = detail.GarageName
		resp.Longtitude = detail.Longtitude
		resp.Latitude = detail.Latitude
		resp.PositionName = detail.PositionName
	}
	//

	return resp
}

func (s *Service) GetDepedancyData(id string) []mod.DetailPosition {
	return s.api.GetDetailData(id)
}
