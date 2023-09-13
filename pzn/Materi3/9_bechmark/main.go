package main

/*
benchmark
selain unit test, golang juga mendukung melakukan benchmark
benchmark adalah mekanisme menghitung kecepatan performa kode aplikasi kita
benchmark di golang dilakukan dengan secara otomatis melakukan iterasi kode yang kita panggil berkali-kali sampai waktu tertentu
kita tidak perlu menentukan jumlah iterasi dan lamanya, karena itu sudah diatur oleh testing.B bawaan dari package testing

- testing.B -> struct yang digunakan untuk melakukan benchmark
- testing.B mirip dengan testing.T, terdapat function Fail(), FailNow(), Error(), Fatal(), dan lain-lain
- yang membedakan, ada beberapa attribute dan function tambahan yang digunakan untuk melakukan benchmark
- salah satunya atribute attribute N, ini digunakan untuk melakukan total iterasi sebuah benchmark
*/
