package main

/*
Context
-> merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout dan sinyal deadline
-> context biasanya dibuat per request (misal setiap ada request masuk ke server web melalui http request)
-> context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar proses
*/

/*
Kenapa perlu mempelajari context?
-> context di golang biasa digunakan untuk mengirim request atau sinyal ke proses lain
-> dengan menggunakan context, ketika kita ingin membatalkan proses, kita cukup mengirim sinyal ke context, maka secara otomatis semua proses akan dibatalkan
-> hampir semua bagian di golang memanfaatkan context, seperti database, hpptserver, http client, dan lain-lain
-> bahkan di google sendiri, ketika menggunakan golang, context wajib digunakan dan selalu dikirim ke setiap fucntion yang dikirim
*/

/*
Cara kerja context
misal ada 3 buah proses, maka masing-masing proses akan memiliki context, proses 1 akan mengirim context ke proses 2 dan proses 3. Misal mau dibatalkan, proses 1 akan mengirim context cancel ke proses 2
*/
