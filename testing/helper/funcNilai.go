package helper

func Validasi(nilai int) string {
	if nilai > 90 {
		return "Nilai Anda A"
	} else if nilai <= 90 && nilai >= 80 {
		return "Nilai Anda B"
	} else if nilai < 80 && nilai >= 70 {
		return "Nilai Anda C"
	} else {
		return "Nilai Anda D"
	}

}
