package service

import (
	"belajar-unit-test/entity"
	"belajar-unit-test/repository"
	"errors"
)

type MahasiswaService struct {
	Repository repository.MahasiswaRepository
}

func (service MahasiswaService) Get(Npm int) (*entity.Mahasiswa, error) {
	mahasiswa := service.Repository.FindNpm(Npm)

	if mahasiswa == nil {
		return nil, errors.New("npm not found")
	} else {
		return mahasiswa, nil
	}

}
