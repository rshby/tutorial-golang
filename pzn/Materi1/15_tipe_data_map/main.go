package main

import "fmt"

// data map di golang
func main() {
	// == 1. membuat variabel map kosong (belum diisi key-value) ==
	fmt.Println("- 1. membuat variabel map kosong -")

	var myMap map[string]string // create new nil map variabel
	fmt.Println("isi variabel myMap adalah:", myMap)
	fmt.Println("apakah myMap berisi nil :", myMap == nil)

	fmt.Println("add key-value to nil map")
	myMap = map[string]string{
		"name": "reo",
		"age":  "24",
	}
	fmt.Println("isi variabel myMap adalah:", myMap)

	// == 2. membuat map kosong dengan menggunakan function make() ==
	fmt.Println("\n- 2. membuat map kosong dengan function make() -")

	makeMap := make(map[string]string)
	fmt.Println("isi dari variabel makeMap adalah:", makeMap)
	fmt.Println("apakah makeMap berisi nil:", makeMap == nil)

	// mengisi key-value to map
	fmt.Println("mengisi key-value to map")
	makeMap["name"] = "reo"
	makeMap["address"] = "ragunan, jakarta selatan"
	fmt.Println("isi makeMap setelah diisi:", makeMap)

	// == 3. membuat variabel map langsung diisi dengan key value -
	fmt.Println("\n- 3. membuat variabel langsung diisi dengan key value -")
	nilaiMap := map[string]int{
		"matematika": 90,
		"fisika":     80,
		"biologi":    95,
	}

	fmt.Println("isi dari map adalah:", nilaiMap, "len:", len(nilaiMap))

	// == 4. menambah/mengubah value pada map ==
	fmt.Println("\n- 4. menambah/mengubah value pada map -")

	myMap = map[string]string{
		"nama": "reo sahobby",
	}
	fmt.Println("isi map awal:", myMap, "len:", len(myMap))

	fmt.Println("mengubah dan menambah key-value dari map")
	myMap["nama"] = "muhammad reo sahobby"
	myMap["jenis_kelamin"] = "laki-laki"

	fmt.Println("isi map setelah diubah dan ditambah:", myMap, "len:", len(myMap))

	// == 5. mengubah value map berdasarkan key
	fmt.Println("\n- 5. mengubah value map berdasarkan key -")

	myMap = map[string]string{
		"nama":   "sigit purnomo",
		"alamat": "pancoran",
	}
	fmt.Println("isi map awal:", myMap, "len:", len(myMap))

	// mengubah value dari key nama
	myMap["nama"] = "rizal saputra"
	fmt.Println("isi map setelah diubah:", myMap, "len:", len(myMap))

	// == 6. menghapus map berdasarkan key ==
	fmt.Println("\n- 6. menghapus map berdasarkan key -")

	myMap = map[string]string{
		"nama":   "reo",
		"alamat": "ragunan",
		"wrong":  "salah",
	}
	fmt.Println("isi map awal:", myMap, "len:", len(myMap))

	fmt.Println("menghapus map dengan key 'wrong'")
	delete(myMap, "wrong")
	fmt.Println("isi map setelah dihapus key 'wrong':", myMap, "len:", len(myMap))
}
