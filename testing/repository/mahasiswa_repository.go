package repository

import "belajar-unit-test/entity"

type MahasiswaRepository interface {
	FindNpm(Npm int) *entity.Mahasiswa
}
