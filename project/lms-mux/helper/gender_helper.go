package helper

// function untuk parse gender ke string kata-kata
func ParseGender(gender string) string {
	if gender == "L" {
		return "Laki-Laki"
	} else {
		return "Perempuan"
	}
}
