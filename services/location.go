package services

import (
	"golang-demo-mousehunt/repository"
	"golang-demo-mousehunt/structs"
)

type LocationService interface {
	GetAllLocations() ([]structs.Location, error)
	GetLocation(location structs.Location) (structs.Location, error)
	InsertLocation(location structs.Location) (structs.Location, error)
	UpdateLocation(location structs.Location) (structs.Location, error)
	DeleteLocation(location structs.Location) (structs.Location, error)
}

type locationService struct {
	repository repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) *locationService {
	return &locationService{repo}
}

func (s *locationService) GetAllLocations() ([]structs.Location, error) {
	var locations, err = s.repository.GetAllLocations()
	if err != nil {
		return locations, err
	} else {
		return locations, nil
	}
}

func (s *locationService) GetLocation(location structs.Location) (structs.Location, error) {
	location, err = s.repository.GetLocation(location)
	if err != nil {
		return location, err
	} else {
		return location, nil
	}
}

func (s *locationService) InsertLocation(location structs.Location) (structs.Location, error) {
	location, err = s.repository.InsertLocation(location)
	if err != nil {
		return location, err
	} else {
		return location, nil
	}
}

func (s *locationService) UpdateLocation(location structs.Location) (structs.Location, error) {
	location, err = s.repository.UpdateLocation(location)
	if err != nil {
		return location, err
	} else {
		return location, nil
	}
}

func (s *locationService) DeleteLocation(location structs.Location) (structs.Location, error) {
	location, err = s.repository.DeleteLocation(location)
	if err != nil {
		return location, err
	} else {
		return location, nil
	}
}