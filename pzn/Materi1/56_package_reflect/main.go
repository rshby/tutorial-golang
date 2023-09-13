package main

import (
	"fmt"
	"reflect"
)

/*
package reflection
-> untuk melihat struktur kode kita pada saat aplikasi sedang berjalan
-> hal ini bisa dilakukan di golang menggunakan package reflect
-> reflection berguna ketika kita ingin membuat library yg general sehingga mudah digunakan
*/

// membuat struct
type Sample struct {
	Name string `required:"true" max:"10"`
}

// membuat function -> untuk melakukan validasi
func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	// membuat object dari struct Sample
	reo := Sample{"reo"}
	sampleType := reflect.TypeOf(reo)
	structField := sampleType.Field(0)

	fmt.Println(structField.Name)

	required := structField.Tag.Get("required")
	fmt.Println(required)
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	// panggil function -> untuk melihat apakah object reo lolos validasi
	fmt.Println("apakah object reo lolos validasi =", IsValid(reo))
}
