package helper

/*
Mock -> object yang sudah kita program dengan ekspektasi terntentu sehingga ketika dipanggil dia akan menghasilkan data yang sudah kita program di awal

mock adalah  salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object yang memang sulit untuk ditesting

misal kita ingin membuat unit test namun ternyata ada kode program kita yang harus memanggik API call ke third party service. Hal ini sangat sulit untuk ditest, karena unit test kita harus selalu memanggil third party service dan belum tentu responsenya sesuai dengan apa yang kita mau

pada kasus seperti ini sangat cocok menggunakan mock object
*/
func HelloWorld(nama string) string {
	return "Hello " + nama
}
