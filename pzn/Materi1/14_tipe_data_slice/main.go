package main

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/slices"
	"reflect"
	"sort"
)

// slice di golang
func main() {
	// == 1. membuat slice kosong (belum ada valuenya) ==
	fmt.Println("- 1. membuat slice kosong -")

	var mySlice []int
	fmt.Println("value dari slice kosong adalah :", mySlice, "panjang", len(mySlice), "capacity", cap(mySlice))

	mySlice = append(mySlice, 1, 2)
	for _, item := range mySlice {
		fmt.Printf(" %d | ", item)
	}
	fmt.Println("\npanjang slice di atas adalah", len(mySlice), "capacity", cap(mySlice))

	// == 2. membuat slice dari potongan array yang sudah ada ==
	myArray := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	fmt.Println("\n- 2. membuat slice dari array yang sudah ada -")
	fmt.Println("ini dari myArray adalah", myArray)

	sliceFromArray := myArray[1:5] // membuat slice dari potongan array yg sudah ada
	fmt.Println("value slice potongan array adalah :", sliceFromArray)

	// == 3. membuat slice langsung diisi dengan value ==
	fmt.Println("\n- 3. membuat slice langsung diisi dengan value -")
	sliceLangsuValue := []string{"satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan"}
	fmt.Println("value slice langsung : ", sliceLangsuValue, "len :", len(sliceLangsuValue), "capacity :", cap(sliceLangsuValue))

	// == 4. membuat slice dari slice yang sudah ada ==
	fmt.Println("\n- 4. membuat slice dari slice yang sudah ada -")
	dariSlice := sliceLangsuValue[3:8]
	fmt.Println("dari slice yang sudah ada", dariSlice, "len :", len(dariSlice), "capacity :", cap(dariSlice))

	fmt.Println("dilakukan perubahan merubah 'empat' menjadi 'EMPAT'")
	dariSlice[0] = "EMPAT"
	fmt.Println("slice asli menjadi :", sliceLangsuValue)
	fmt.Println("slice yang diubah menjadi :", dariSlice)

	// == 5. membuat slice menggunakan make() function ==
	fmt.Println("\n- 5. membuat slice menggunakan make() function()")

	makeSlice := make([]int, 4)
	fmt.Println("slice dibuat menggunakan make :", makeSlice, "len :", len(makeSlice), "capacity :", cap(makeSlice))

	makeSliceString := make([]string, 3)
	fmt.Println("slice string dibuat menggunakan make :", makeSliceString, "len :", len(makeSliceString), "capacity :", cap(makeSliceString))

	// == 6. iterasi slice menggunakan for loop ==
	fmt.Println("\n- 6. iterasi slice menggunakan for loop -")

	sliceLoop := []string{"nama", "saya", "adalah", "reo"}

	fmt.Printf("print slice menggunakan for loop : ")
	for i := 0; i < len(sliceLoop); i++ {
		if i == len(sliceLoop)-1 {
			fmt.Printf(sliceLoop[i])
			break
		}
		fmt.Printf("%s, ", sliceLoop[i])
	}

	// == 7. iterasi slice menggunakan for range ==
	fmt.Println("\n\n- 7. iterasi slice menggunakan for range -")
	fmt.Printf("print slice menggunakan for range")
	for i, item := range sliceLoop {
		if i == len(sliceLoop)-1 {
			fmt.Printf(item)
			break
		}
		fmt.Printf("%s, ", item)
	}

	// == 8. merubah value slice ==
	fmt.Println("\n\n- 8. merubah value slice -")
	sliceAwal := []string{"satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh"}
	fmt.Println("value slice awal :", sliceAwal, "len:", len(sliceAwal), "capacity:", cap(sliceAwal))

	sliceBaru := sliceAwal[2:8]
	fmt.Println("value slice baru:", sliceBaru, "len:", len(sliceBaru), "capacity:", cap(sliceBaru))

	fmt.Println("ada perubahan pada sliceBaru : 'lima' menjadi 'LIMA'")
	sliceBaru[2] = "LIMA"
	fmt.Println("slice awal menjadi:", sliceAwal)
	fmt.Println("slice baru menjadi:", sliceBaru)

	// == 9. menambah slice menggunakan append ==
	fmt.Println("\n- 9. menambah value slice menggunakan append -")

	sliceAwal = []string{"satu", "dua", "tiga"}
	fmt.Println("slice awal:", sliceAwal, "len:", len(sliceAwal), "capacity:", cap(sliceAwal))

	fmt.Println("append value 'empat' ke sliceAwal")
	sliceAwal = append(sliceAwal, "empat")
	fmt.Println("slice awal menjadi:", sliceAwal, "len:", len(sliceAwal), "capacity:", cap(sliceAwal))

	sliceBaru = sliceAwal[0:2]
	fmt.Println("slice baru:", sliceBaru, "len:", len(sliceBaru), "capacity:", cap(sliceBaru))

	fmt.Println("append beberapa value ke sliceBaru")
	sliceBaru = append(sliceBaru, "TIGA", "EMPAT", "LIMA", "ENAM")
	fmt.Println("slice baru menjadi:", sliceBaru, "len:", len(sliceBaru), "capacity:", cap(sliceBaru))
	fmt.Println("slice awal menjadi:", sliceAwal, "len:", len(sliceAwal), "capacity:", cap(sliceAwal)) // value slice awal ikut berubah maksimal seusai panjang slicenya

	// == 10. compare slice ==
	fmt.Println("\n- 10. compare slice -")

	sliceSatu := []int{1, 2, 3}
	sliceDua := []int{1, 2, 3}
	fmt.Println("value sliceSatu adalah:", sliceSatu, "len:", len(sliceSatu), "capacity:", cap(sliceSatu))
	fmt.Println("value sliceDua adalah:", sliceDua, "len:", len(sliceDua), "capacity:", cap(sliceDua))

	compareSlice := reflect.DeepEqual(sliceSatu, sliceDua)
	fmt.Println("apakah sliceSatu equals to sliceDua :", compareSlice)

	// == 11. cek item contains in slice ==
	fmt.Println("\n- 11. cek item contains in slice -")

	mySlice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("value slice adalah:", mySlice, "len:", len(mySlice), "capacity:", cap(mySlice))

	isContain := slices.Contains(mySlice, 2)
	fmt.Println("apakah slice itu contains 2 :", isContain)

	// == 12. sorting slice ==
	fmt.Println("\n- 12. sorting slice -")
	mySlice = []int{5, 8, 2, 1, 9}
	fmt.Println("value slice awal adalah:", mySlice)

	fmt.Println("dilakukan sorting slice")
	sort.Ints(mySlice)
	fmt.Println("value slice setelah diurutkan:", mySlice)

	// == 13. menghapus prefix dan suffix slice dengan function trim() ==
	fmt.Println("\n- 13. menghapus prefix dan suffix slice dengan function trim() -")

	iniSlice := []byte{'@', 'R', '@'}
	fmt.Println("value slice byte adalah:", iniSlice, "len:", len(iniSlice), "capacity:", cap(iniSlice))

	fmt.Println("dilakukan pemotongan slice menggunakan trim")
	sliceTrim := bytes.Trim(iniSlice, "@")
	fmt.Println("value slice byte setelah dipotong menjadi:", sliceTrim, "len:", len(sliceTrim), "capacity:", cap(sliceTrim))

	// == 14. memisahkan slice dengan menggunakan split ==
	fmt.Println("\n- 14. memisahkan slice dengan menggunakan Split() -")
	slice1 := []byte{'r', 'e', 'e', 'o'}
	fmt.Printf("value slice1 : %s\n", slice1)

	fmt.Println("dilakukan pemisahan menggunakan Split()")
	result1 := bytes.Split(slice1, []byte("ee"))
	fmt.Printf("hasil setelah diSplit : %s\n", result1)
}
