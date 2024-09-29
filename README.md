Buatlah sebuah program Golang yang melakukan perhitungan sederhana menggunakan fungsi sebagai first-class citizen, memanfaatkan interface untuk mendefinisikan perilaku dari operasi matematika, dan menjalankan perhitungan tersebut secara concurrent menggunakan goroutine.

Detail Soal:

1. Buat interface Operation yang memiliki satu metode bernama Calculate() int, yang mendefinisikan operasi matematika yang akan dilakukan.
2. Implementasikan dua struct: Addition dan Multiplication, yang akan mengimplementasikan interface Operation. Keduanya harus memiliki properti a dan b (dua bilangan integer yang akan dioperasikan).
3. Buat fungsi performOperation yang menerima parameter Operation (interface) dan menjalankan metode Calculate() untuk mendapatkan hasilnya.
4. Jalankan fungsi perform Operation untuk setiap operasi (Addition dan Multiplication) secara concurrent menggunakan goroutine.
5. Tampilkan hasil perhitungan dari kedua operasi tersebut.

Contoh Hasil Eksekusi:
Jika a 3 dan b = 5:
Penjumlahan: 8
Perkalian: 15

Petunjuk:
Gunakan first-class citizen dengan menyimpan fungsi performOperation dalam variabel.
Jalankan operasi perhitungan dengan menggunakan goroutine untuk penjumlahan dan perkalian.
Gunakan interface untuk mendefinisikan operasi perhitungan

Contoh Output (eksekusi dilakukan secara concurrent):
Penjumlahan: 8
Perkalian: 15
