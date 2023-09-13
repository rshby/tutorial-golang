package main

/*
=== Parent and Child Context ===
-> context menganut konsep parent child
-> artinya, saat kita membuat context, kita bisa membuat child context yang sudah ada
-> parent context bisa memiliki banyak child, namun child hanya memiliki satu parent
-> konsep ini mirip dengan pewaruisan sifat di pemrograman beorientasi object
*/

/*
=== Hubungan antara parent dan child context ===
-> parent dan child context akan selalu terhubung
-> saat nanti kita melakukan pembatalan context A, maka semua child dan sub childnya dari context A akan ikut dibatalkan
-> Namun jika misal kita membatalkan context B, hanya context B dan semua sub childnya yang dibatalkan, parent context B tidak ikut dibatalkan
-> begitu juga nanti saat kita menyisipkan data ke context A, semua child dan sub childnya bisa mendapatkan data tersebut
-> Namun kita kita menyisipkan data ke context B, hanya context B dan sub childnya yang akan mendapat data, parent context B tidak mendapatkan data
*/

/*
=== Immutable ===
-> context merupakan object yang immutable, artinya setelah context dibuat, dia tidak bisa diubah lagi
-> ketika kita ingin nemabahkan value ke dalam context, atau menambahkan pengaturan timeout dan yang lainnya, itu sebenarnya secara otomatis akan membentuk child context baru, bukan merubah context tersebut
*/

func main() {

}
