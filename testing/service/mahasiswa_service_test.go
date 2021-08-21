package service

import (
	"belajar-unit-test/entity"
	"belajar-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mahasiswaRepository = &repository.MahasiswaRepositoryMock{Mock: mock.Mock{}}
var mahasiswaService = MahasiswaService{Repository: mahasiswaRepository}

func TestNpmNotFound(t *testing.T) {
	//program mock

	mahasiswaRepository.Mock.On("FindNpm", 1).Return(nil)

	npm, err := mahasiswaService.Get(1)
	assert.Nil(t, npm)
	assert.NotNil(t, err)

}

func TestNpmFound(t *testing.T) {
	mahasiswa := entity.Mahasiswa{
		Npm:    20174350,
		Name:   "Ramli",
		Alamat: "Bogor",
	}

	mahasiswaRepository.Mock.On("FindNpm", 2).Return(mahasiswa)

	result, err := mahasiswaService.Get(2)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mahasiswa.Npm, result.Npm)
	assert.Equal(t, mahasiswa.Name, result.Name)
	assert.Equal(t, mahasiswa.Alamat, result.Alamat)
}
