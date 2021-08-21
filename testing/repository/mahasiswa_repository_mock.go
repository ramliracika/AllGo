package repository

import (
	"belajar-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type MahasiswaRepositoryMock struct {
	Mock mock.Mock
}

func (repository *MahasiswaRepositoryMock) FindNpm(Npm int) *entity.Mahasiswa {
	argument := repository.Mock.Called(Npm)
	if argument.Get(0) == nil {
		return nil
	}

	result := argument.Get(0).(entity.Mahasiswa)
	return &result

}
