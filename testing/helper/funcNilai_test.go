package helper

import (
	"fmt"
	"runtime"
	"testing"

	//penggunaan testify untuk menampilkan detail saat error
	"github.com/stretchr/testify/assert"  // jika eror akan memanggil fungsi t.Fail()
	"github.com/stretchr/testify/require" // jika eror akan memanggil fungsi t.FailNow()
)

func BenchmarkTableBench(b *testing.B) {
	Ujian := []struct {
		name  string
		nilai int
	}{
		{
			name:  "A",
			nilai: 97,
		},
		{
			name:  "B",
			nilai: 88,
		},
	}
	for _, ujian := range Ujian {
		b.Run(ujian.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Validasi(ujian.nilai)
			}
		})
	}
}

func BenchmarkSubBench(b *testing.B) {
	b.Run("A", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Validasi(90)
		}
	})
	b.Run("B", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Validasi(80)
		}
	})

}

func BenchmarkBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validasi(90)
	}
}

func BenchmarkBasic2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validasi(65)
	}
}

//Table test adalah fungsi unit test dengan for dan struct

func TestTableTest(t *testing.T) {
	Ujian := []struct {
		Name     string
		Expected string
		Param    int
	}{
		{
			Name:     "A",
			Expected: "Nilai Anda A",
			Param:    95,
		},
		{
			Name:     "B",
			Expected: "Nilai Anda B",
			Param:    85,
		},
	}

	for _, nilai := range Ujian {
		t.Run(nilai.Name, func(t *testing.T) {
			result := Validasi(nilai.Param)
			require.Equal(t, nilai.Expected, result)
		})
	}

}

//function subtest adalah menjalankan unit test fungsi di dalam fungsi
func TestSubtest(t *testing.T) {
	// memanggil dengan menambahkan /NilaiA
	t.Run("NilaiA", func(t *testing.T) {
		result := Validasi(95)
		assert.Equal(t, "Nilai Anda A", result, "output harus sama dengan A")

	})

	t.Run("NilaiB", func(t *testing.T) {
		result := Validasi(75)
		assert.Equal(t, "Nilai Anda B", result, "output harus sama dengan B")

	})
}

//function test main untuk running seluruh unit test didalam package
func TestMain(m *testing.M) {
	//BEFORE
	fmt.Println("Before Test")

	m.Run()

	//AFTER
	fmt.Println("After Test")

}

//skip function untuk tidak melanjutkan kode dibawahnya dengan kondisi tertentu
func TestSkip(t *testing.T) {
	// deteksi Sistem operasi
	if runtime.GOOS == "linux" {
		t.Skip("Tidak Bisa Dijalankan di linux")
	}

	result := Validasi(85)
	assert.Equal(t, "Nilai Anda B", result, "output harus sama dengan B")
	fmt.Println("Function Assert Done")

}

func TestFuncAssert(t *testing.T) {
	result := Validasi(85)
	assert.Equal(t, "Nilai Anda B", result, "output harus sama dengan B")
	fmt.Println("Function Assert Done")
}

func TestFuncRequire(t *testing.T) {
	result := Validasi(75)
	require.Equal(t, "Nilai Anda B", result, "output harus sama dengan B")
	fmt.Println("Function Require Done") // jika error tidak akan tampil
}

func TestFuncNilaiA(t *testing.T) {
	result := Validasi(91)

	if result != "Nilai Anda A" {
		//error
		// memakai panic hanya akan menghentikan program, jadi tidak disarankan
		// memakai t.Fail() dapat melanjutkan program namun tidak menampilkan log
		// sangat dianjurkan memakai t.Error karena memanggil t.Fail() dan menampilkan log
		t.Error("Output yang ditampilkan Bukan A")
	}
	fmt.Println("Done Testing Function Nilai A")
}

func TestFuncNilaiD(t *testing.T) {
	result := Validasi(65)

	if result != "Nilai Anda D" {
		//error
		// memakai panic hanya akan menghentikan program, jadi tidak disarankan
		// memakai t.FailNow() tidak dapat melanjutkan program dan tidak menampilkan log
		// sangat dianjurkan memakai t.Fatal karena memanggil t.FailNow() dan menampilkan log
		t.Fatal("Output yang ditampilkan Bukan D")

	}

	fmt.Println("Done Testing Function Nilai D") // jika error ini tidak akan dijalankan
}
