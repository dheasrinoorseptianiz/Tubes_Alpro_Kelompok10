// Aplikasi Perpustakaan
// Deskripsi: Aplikasi ini digunakan oleh admin perpustakaan untuk mencatat buku-buku yang akan dipinjam, tanggal peminjaman,
// tanggal pengembalian serta denda apabila melewati batas waktu.
// Spesifikasi:
// a. Pengguna bisa menambahkan, mengubah (edit), dan menghapus databuku.
// b. Pengguna bisa menambahkan, mengubah (edit), dan menghapus data peminjaman buku.
// c. Pengguna bisa melihat data buku secara terurut berdasarkan berbagai macan kategori.
// d. Pengguna bisa melakukan pencarian buku dengan kata kunci tertentu.
// e. Aplikasi memiliki perhitungan tarif peminjaman dan denda keterlambatan.
// f. Aplikasi bisa menampilkan buku yang sedang dipinjam, dan juga 5 buku terfavorit atau yang paling sering dipinjam.

package main

import "fmt"

const nmax = 100

type Buku struct {
	Judul    string
	Penulis  string
	Tahun    int
	Jumlah   int
	Kategori string
}

type Pinjaman struct {
	Judul           string
	NamaPeminjam    string
	TglPeminjaman   string
	LamaPeminjaman  int
	TglPengembalian string
	TarifDihitung   bool
}

type BukuJumlah struct {
	Judul  string
	Jumlah int
}

var jumlahBuku int
var jumlahPinjaman int

type daftarBuku [nmax]Buku
type daftarPinjaman [nmax]Pinjaman
type bukuJumlahArray [nmax]BukuJumlah

func menuUtama() {
	fmt.Println("                                                                    ")
	fmt.Println("----------------------- APLIKASI PERPUSTAKAAN ----------------------")
	fmt.Println("                                                                    ")
	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                              M E N U                             ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
	fmt.Println("║  1. Kelola Buku                                                  ║")
	fmt.Println("║  2. Kelola Peminjaman Buku                                       ║")
	fmt.Println("║  3. Lihat Data Buku                                              ║")
	fmt.Println("║  4. Pencarian Buku                                               ║")
	fmt.Println("║  5. Keluar                                                       ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
}

func kelolaBukuMenu() {
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("                        K E L O L A  B U K U                        ")
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("1. Tambah Data Buku                                                 ")
	fmt.Println("2. Edit Data Buku                                                   ")
	fmt.Println("3. Hapus Data Buku                                                  ")
	fmt.Println("4. Tampilkan Data Buku                                              ")
	fmt.Println("5. Kembali ke Menu Utama                                            ")
	fmt.Println("--------------------------------------------------------------------")
}

func tambahDataBuku(A *daftarBuku, n *int) { // dapat menambahkan data buku (judul, penulis, tahun terbit, jumlah, kategori)
	var buku Buku

	if *n >= nmax {
		fmt.Println("Kapasitas buku penuh.")
	} else {
		fmt.Println("Masukkan data buku:")
		fmt.Print("Judul: ")
		fmt.Scanln(&buku.Judul)
		fmt.Print("Penulis: ")
		fmt.Scanln(&buku.Penulis)
		fmt.Print("Tahun Terbit: ")
		fmt.Scanln(&buku.Tahun)
		fmt.Print("Jumlah: ")
		fmt.Scanln(&buku.Jumlah)
		fmt.Print("Kategori: ")
		fmt.Scanln(&buku.Kategori)

		A[*n] = buku
		*n++
		fmt.Println("Buku berhasil ditambahkan.")
	}
}

// sequential search untuk mengedit data buku
func editDataBuku(A *daftarBuku, pinjaman daftarPinjaman, nBuku, nPinjaman int) { // dapat mengedit data buku dengan mencari judul buku (judul baru, penulis baru, tahun terbit baru, jumlah baru, kategori baru)
	var x string
	var jumlahBaru int
	var found bool

	fmt.Print("Masukkan judul buku yang ingin diubah: ")
	fmt.Scanln(&x)

	for i := 0; i < nBuku; i++ {
		if A[i].Judul == x {
			found = true
			// mengitung jumlah buku yang sedang dipinjam
			jumlahDipinjam := 0
			for j := 0; j < nPinjaman; j++ {
				if pinjaman[j].Judul == x && !pinjaman[j].TarifDihitung {
					jumlahDipinjam++
				}
			}

			// mengecek apakah masih ada buku yang tersedia untuk diedit
			if A[i].Jumlah > jumlahDipinjam {
				fmt.Println("Buku ditemukan. Silakan masukkan data baru untuk buku ini.")

				fmt.Print("Judul baru: ")
				fmt.Scanln(&A[i].Judul)
				fmt.Print("Penulis baru: ")
				fmt.Scanln(&A[i].Penulis)
				fmt.Print("Tahun terbit baru: ")
				fmt.Scanln(&A[i].Tahun)
				fmt.Print("Jumlah baru: ")
				fmt.Scanln(&jumlahBaru)
				fmt.Print("Kategori baru: ")
				fmt.Scanln(&A[i].Kategori)

				// menyesuaikan jumlah buku yang tersisa dengan jumlah yang diedit
				if jumlahBaru >= jumlahDipinjam {
					A[i].Jumlah = jumlahBaru
					fmt.Println("Data buku berhasil diubah.")
				} else {
					fmt.Println("Jumlah baru harus lebih besar atau sama dengan jumlah buku yang sedang dipinjam.")
				}
			} else {
				fmt.Println("Buku dengan judul", x, "tidak ditemukan.")
			}
		}
	}

	if !found {
		fmt.Println("Buku dengan judul", x, "tidak ditemukan.")
	}
}

// procedure untuk mengurutkan buku berdasarkan judul
func sortJudulBuku(A *daftarBuku, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if A[i].Judul > A[j].Judul {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}

// binary search untuk mencari buku berdasarkan judul
func binarySearchBuku(A daftarBuku, n int, judul string) int {
	kiri, kanan := 0, n-1

	for kiri <= kanan {
		tengah := kiri + (kanan-kiri)/2

		if A[tengah].Judul == judul {
			return tengah
		}

		if A[tengah].Judul < judul {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func hapusDataBuku(A *daftarBuku, pinjaman daftarPinjaman, nBuku, nPinjaman *int) { // dapat menghapus data buku dengan mencari judul buku
	var judul string
	var jumlahHapus int

	if *nBuku == 0 {
		fmt.Println("Tidak ada buku yang bisa dihapus.")
	} else {
		fmt.Print("Masukkan judul buku yang ingin dihapus: ")
		fmt.Scanln(&judul)

		// mengurutkan data buku sebelum melakukan binary search
		sortJudulBuku(A, *nBuku)

		index := binarySearchBuku(*A, *nBuku, judul)

		if index == -1 {
			fmt.Println("Buku dengan judul", judul, "tidak ditemukan.")
		} else {
			// menghitung jumlah buku yang sedang dipinjam
			jumlahDipinjam := 0
			for i := 0; i < *nPinjaman; i++ {
				if pinjaman[i].Judul == judul && !pinjaman[i].TarifDihitung {
					jumlahDipinjam++
				}
			}

			// mengecek apakah masih ada buku yang tersedia untuk dihapus
			if A[index].Jumlah > jumlahDipinjam {
				fmt.Printf("Masukkan jumlah buku yang ingin dihapus (maksimal %d): ", A[index].Jumlah-jumlahDipinjam)
				fmt.Scanln(&jumlahHapus)

				if jumlahHapus <= A[index].Jumlah-jumlahDipinjam {
					// mengurangi jumlah buku yang tersedia
					A[index].Jumlah -= jumlahHapus

					// jika jumlah buku menjadi 0 dan tidak ada yang dipinjam, hapus buku dari daftar
					if A[index].Jumlah == 0 && jumlahDipinjam == 0 {
						for i := index; i < *nBuku-1; i++ {
							A[i] = A[i+1]
						}
						(*nBuku)--
					}
					fmt.Println("Buku berhasil dihapus.")
				} else {
					fmt.Println("Jumlah buku yang ingin dihapus melebihi jumlah yang tersedia.")
				}
			} else {
				fmt.Println("Buku dengan judul", judul, "tidak ditemukan.")
			}
		}
	}
}

func tampilkanDataBuku(A daftarBuku, pinjaman daftarPinjaman, nBuku, nPinjaman int) { // dapat menampilkan data buku yang telah ditambahkan atau yang tersedia (judul, penulis, tahun terbit, jumlah, kategori)
	if nBuku == 0 {
		fmt.Println("Tidak ada data buku yang tersedia.")
	} else {
		bukuTersedia := false // penanda jika ada buku yang tersedia
		for i := 0; i < nBuku; i++ {
			buku := A[i]
			jumlahTersedia := buku.Jumlah

			for j := 0; j < nPinjaman; j++ {
				if pinjaman[j].Judul == buku.Judul && !pinjaman[j].TarifDihitung {
					jumlahTersedia--
				}
			}

			if jumlahTersedia > 0 {
				if !bukuTersedia {
					fmt.Println("Berikut data buku:")
				}
				bukuTersedia = true
				fmt.Println("--------------------------------------------------------------------")
				fmt.Println("Judul    : ", buku.Judul)
				fmt.Println("Penulis  : ", buku.Penulis)
				fmt.Println("Tahun    : ", buku.Tahun)
				fmt.Println("Jumlah   : ", jumlahTersedia)
				fmt.Println("Kategori : ", buku.Kategori)
			}
		}

		if !bukuTersedia {
			fmt.Println("Tidak ada buku yang tersedia.")
		} else {
			fmt.Println("--------------------------------------------------------------------")
		}
	}
}

func kelolaPeminjamanBukuMenu() {
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("            K E L O L A  P E M I N J A M A N  B U K U               ")
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("1. Tambah Data Peminjaman Buku                                      ")
	fmt.Println("2. Edit Data Peminjaman Buku                                        ")
	fmt.Println("3. Hapus Data Peminjaman Buku                                       ")
	fmt.Println("4. Tampilkan Data Peminjaman Buku                                   ")
	fmt.Println("5. Hitung tarif Peminjaman Buku dan Denda Keterlambatan             ")
	fmt.Println("6. Lima Buku Terfavorit (Sering di Pinjam)                          ")
	fmt.Println("7. Kembali ke Menu Utama                                            ")
	fmt.Println("--------------------------------------------------------------------")
}

// mencarai judul buku
func cariBukuBerdasarkanJudul(A daftarBuku, n int, judul string) bool {
	for i := 0; i < n; i++ {
		if A[i].Judul == judul {
			return true
		}
	}
	return false
}

// sequential search untuk menambah data peminjaman buku
func tambahDataPeminjamanBuku(A *daftarPinjaman, B *daftarBuku, nPinjaman, nBuku *int) { // dapat menambahkan data peminjaman buku dengan mencari judul buku yang tersedia (nama peminjam, tanggal peminjaman, lama peminjaman)
	var pinjam Pinjaman

	if *nPinjaman >= nmax {
		fmt.Println("Kapasitas pinjaman penuh.")
	} else {
		fmt.Println("Masukkan data peminjaman buku:")
		fmt.Print("Judul Buku: ")
		fmt.Scanln(&pinjam.Judul)

		// mengecek ketersediaan buku
		tersedia := false
		for i := 0; i < *nBuku; i++ {
			if B[i].Judul == pinjam.Judul {
				jumlahTersedia := B[i].Jumlah
				for j := 0; j < *nPinjaman; j++ {
					if A[j].Judul == pinjam.Judul && !A[j].TarifDihitung {
						jumlahTersedia--
					}
				}
				if jumlahTersedia > 0 {
					tersedia = true
				}
			}
		}

		if tersedia {
			fmt.Print("Nama Peminjam: ")
			fmt.Scanln(&pinjam.NamaPeminjam)
			fmt.Print("Tanggal Peminjaman (dd-mm-yyyy): ")
			fmt.Scanln(&pinjam.TglPeminjaman)
			fmt.Print("Lama Peminjaman (hari): ")
			fmt.Scanln(&pinjam.LamaPeminjaman)

			A[*nPinjaman] = pinjam
			*nPinjaman++
			fmt.Println("Data peminjaman berhasil ditambahkan.")
		} else {
			fmt.Println("Buku tidak tersedia atau sedang dipinjam orang lain.")
		}
	}
}

// sequential search untuk mengedit data peminjaman buku
func editDataPeminjamanBuku(A *daftarPinjaman, B *daftarBuku, nPinjaman, nBuku int) { // dapat mengedit data peminjaman buku dengan mencari nama peminjam (nama peminjam baru, tanggal peminjaman baru, lama peminjaman baru)
	var x string
	var judulBaru string

	found := false
	fmt.Print("Masukkan nama peminjam yang ingin diubah: ")
	fmt.Scanln(&x)

	for i := 0; i < nPinjaman; i++ {
		if A[i].NamaPeminjam == x {
			if A[i].TarifDihitung {
				fmt.Println("Data peminjaman dengan nama", x, "tidak ditemukan.")
				found = true // penanda nama ditemukkan
			} else {
				found = true
				fmt.Println("Data peminjaman ditemukan. Silakan masukkan data baru.")
				fmt.Print("Judul Buku baru: ")
				fmt.Scanln(&judulBaru)

				// mengecek ketersediaan buku
				tersedia := false
				for j := 0; j < nBuku; j++ {
					if B[j].Judul == judulBaru {
						jumlahTersedia := B[j].Jumlah
						for k := 0; k < nPinjaman; k++ {
							// mengecualikan peminjaman saat ini dari perhitungan ketersediaan
							if A[k].Judul == judulBaru && !A[k].TarifDihitung && k != i {
								jumlahTersedia--
							}
						}
						if jumlahTersedia > 0 {
							tersedia = true
						}
					}
				}

				if tersedia {
					A[i].Judul = judulBaru
					fmt.Print("Nama Peminjam baru: ")
					fmt.Scanln(&A[i].NamaPeminjam)
					fmt.Print("Tanggal Peminjaman baru (dd-mm-yyyy): ")
					fmt.Scanln(&A[i].TglPeminjaman)
					fmt.Print("Lama Peminjaman baru (hari): ")
					fmt.Scanln(&A[i].LamaPeminjaman)
					fmt.Println("Data peminjaman berhasil diubah.")
				} else {
					fmt.Println("Buku tidak tersedia atau sedang dipinjam orang lain.")
				}
			}
		}
	}

	if !found {
		fmt.Println("Data peminjaman dengan nama", x, "tidak ditemukan.")
	}
}

// procedure untuk mengurutkan nama peminjam
func sortNamaPeminjam(A *daftarPinjaman, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if A[i].NamaPeminjam > A[j].NamaPeminjam {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}

// binary search untuk mencari data peminjaman berdasarkan nama peminjam
func binarySearchPinjaman(A daftarPinjaman, n int, nama string) int {
	kiri, kanan := 0, n-1

	for kiri <= kanan {
		tengah := kiri + (kanan-kiri)/2

		if A[tengah].NamaPeminjam == nama {
			return tengah
		}

		if A[tengah].NamaPeminjam < nama {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func hapusDataPeminjamanBuku(A *daftarPinjaman, n *int) { // dapat menghapus data peminjaman buku dengan mencari nama peminjam
	var nama string

	if *n == 0 {
		fmt.Println("Tidak ada data peminjaman yang bisa dihapus.")
	} else {
		fmt.Print("Masukkan nama peminjam yang ingin dihapus: ")
		fmt.Scanln(&nama)

		// mengurutkan data peminjaman sebelum melakukan pencarian biner
		sortNamaPeminjam(A, *n)

		index := binarySearchPinjaman(*A, *n, nama)

		if index == -1 {
			fmt.Println("Data peminjaman dengan nama", nama, "tidak ditemukan.")
		} else {
			if A[index].TarifDihitung {
				fmt.Println("Data peminjaman dengan nama", nama, "tidak ditemukan.")
			} else {
				for i := index; i < *n-1; i++ {
					A[i] = A[i+1]
				}
				*n--
				fmt.Println("Data peminjaman berhasil dihapus.")
			}
		}
	}
}

func tampilkanDataPeminjaman(A daftarPinjaman, n int) { // dapat menampilkan data peminjaman buku
	if n == 0 {
		fmt.Println("Tidak ada data peminjaman yang tersedia.")
	} else {
		pinjamanTersedia := false // penanda jika ada peminjaman yang belum dikembalikan

		for i := 0; i < n; i++ {
			if !A[i].TarifDihitung {
				if !pinjamanTersedia {
					fmt.Println("Berikut data peminjaman buku:")
				}
				pinjamanTersedia = true
				pinjam := A[i]
				fmt.Println("--------------------------------------------------------------------")
				fmt.Println("Judul              : ", pinjam.Judul)
				fmt.Println("Nama Peminjam      : ", pinjam.NamaPeminjam)
				fmt.Println("Tanggal Peminjaman : ", pinjam.TglPeminjaman)
				fmt.Println("Lama Peminjaman    : ", pinjam.LamaPeminjaman, "hari")
			}
		}

		if !pinjamanTersedia {
			fmt.Println("Tidak ada data peminjaman yang tersedia.")
		} else {
			fmt.Println("--------------------------------------------------------------------")
		}
	}
}

func hitungTarifDanDenda(A *daftarPinjaman, n int) { // dapat menghitung tarif buku yang dipinjam per hari dan menghitung denda keterlambatan per harinya apabila telat mengembalikannya
	var nama string
	var tglPengembalian string
	var totalTarif, totalDenda float64
	var d, m, y, d1, m1, y1 int

	if n == 0 {
		fmt.Println("Tidak ada data peminjaman yang tersedia.")
	} else {
		found := false
		fmt.Print("Masukkan nama peminjam yang ingin dihitung tarif peminjamannya: ")
		fmt.Scanln(&nama)

		for i := 0; i < n; i++ {
			if A[i].NamaPeminjam == nama && !A[i].TarifDihitung {
				found = true
			}
		}

		if !found {
			fmt.Printf("Data peminjaman dengan nama %s tidak ditemukan.\n", nama)
		} else {
			tarifPerHari := 2000.0
			dendaPerHari := 4000.0

			fmt.Print("Masukkan tanggal pengembalian (dd-mm-yyyy): ")
			fmt.Scanln(&tglPengembalian)

			for i := 0; i < n; i++ {
				if A[i].NamaPeminjam == nama {
					totalTarif += float64(A[i].LamaPeminjaman) * tarifPerHari

					fmt.Sscanf(A[i].TglPeminjaman, "%02d-%02d-%04d", &d, &m, &y)
					fmt.Sscanf(tglPengembalian, "%02d-%02d-%04d", &d1, &m1, &y1)

					hariTerlambat := (d1 - d) + (m1-m)*30 + (y1-y)*365

					if hariTerlambat > A[i].LamaPeminjaman {
						totalDenda += float64(hariTerlambat-A[i].LamaPeminjaman) * dendaPerHari
					}
					A[i].TarifDihitung = true
				}
			}

			if totalDenda > 0 {
				fmt.Printf("Denda keterlambatan oleh %s adalah: Rp%.2f\n", nama, totalDenda)
				fmt.Printf("Total tarif peminjaman oleh %s adalah: Rp%.2f\n", nama, totalTarif)
				fmt.Printf("Total tarif peminjaman oleh %s + denda keterlambatan adalah: Rp%.2f\n", nama, totalDenda+totalTarif)
			} else {
				fmt.Printf("Tidak ada denda keterlambatan oleh %s.\n", nama)
				fmt.Printf("Total tarif peminjaman oleh %s adalah: Rp%.2f\n", nama, totalTarif)
			}
		}
	}
}

// selection sort untuk menentukan lima buku terfavorit (sering dipinjam)
func limaBukuTerfavorit(pinjaman daftarPinjaman, nPinjaman int) { // menampilkan 5 buku terfavorit (sering dipinjam)
	var bukuJumlahList bukuJumlahArray
	var nBukuJumlah int

	if nPinjaman == 0 {
		fmt.Println("Tidak ada data peminjaman yang tersedia.")
	} else {
		for i := 0; i < nPinjaman; i++ {
			judul := pinjaman[i].Judul
			ditemukan := false
			for j := 0; j < nBukuJumlah; j++ {
				if bukuJumlahList[j].Judul == judul {
					bukuJumlahList[j].Jumlah++
					ditemukan = true
				}
			}
			if !ditemukan {
				bukuJumlahList[nBukuJumlah] = BukuJumlah{Judul: judul, Jumlah: 1}
				nBukuJumlah++
			}
		}

		// mengurutkan bukuJumlahList berdasarkan jumlah peminjaman secara menurun
		for i := 0; i < nBukuJumlah-1; i++ {
			maxIndex := i
			for j := i + 1; j < nBukuJumlah; j++ {
				if bukuJumlahList[j].Jumlah > bukuJumlahList[maxIndex].Jumlah {
					maxIndex = j
				}
			}
			bukuJumlahList[i], bukuJumlahList[maxIndex] = bukuJumlahList[maxIndex], bukuJumlahList[i]
		}

		fmt.Println("Lima Buku Terfavorit (Sering Dipinjam):")
		for i := 0; i < 5 && i < nBukuJumlah; i++ {
			fmt.Printf("%d. Judul: %s, Dipinjam: %d kali\n", i+1, bukuJumlahList[i].Judul, bukuJumlahList[i].Jumlah)
		}
	}
}

func lihatKategoriDataBukuMenu() {
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("                 K A T E G O R I  D A T A  B U K U                  ")
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("1. Data Buku Berdasarkan Kategori (Ascending)                      ")
	fmt.Println("2. Data Buku Berdasarkan Kategori (Descending)                     ")
	fmt.Println("3. Kembali ke Menu Utama                                           ")
	fmt.Println("--------------------------------------------------------------------")
}

// insertion sort untuk mengurutkan buku berdasarkan kategori secara ascending
func ascendingKategoriBuku(A daftarBuku, pinjaman daftarPinjaman, nBuku, nPinjaman int) { // menampilkan buku berdasarkan kategori secara ascending
	var bukuTersedia daftarBuku
	var nBukuTersedia int

	if nBuku == 0 {
		fmt.Println("Tidak ada data buku yang tersedia.")
	} else {
		for i := 0; i < nBuku; i++ {
			jumlahTersedia := A[i].Jumlah
			for j := 0; j < nPinjaman; j++ {
				if pinjaman[j].Judul == A[i].Judul && !pinjaman[j].TarifDihitung {
					jumlahTersedia--
				}
			}

			if jumlahTersedia > 0 {
				A[i].Jumlah = jumlahTersedia
				bukuTersedia[nBukuTersedia] = A[i]
				nBukuTersedia++
			}
		}

		if nBukuTersedia == 0 {
			fmt.Println("Tidak ada data buku yang tersedia.")
		} else {
			// insertion sort ascending berdasarkan kategori
			for i := 1; i < nBukuTersedia; i++ {
				key := bukuTersedia[i]
				j := i - 1
				for j >= 0 && bukuTersedia[j].Kategori > key.Kategori {
					bukuTersedia[j+1] = bukuTersedia[j]
					j--
				}
				bukuTersedia[j+1] = key
			}

			fmt.Println("Kategori Buku Berdasarkan Ascending:")
			for i := 0; i < nBukuTersedia; i++ {
				fmt.Println("--------------------------------------------------------------------")
				fmt.Println("Kategori : ", bukuTersedia[i].Kategori)
				fmt.Println("Judul    : ", bukuTersedia[i].Judul)
				fmt.Println("Penulis  : ", bukuTersedia[i].Penulis)
				fmt.Println("Tahun    : ", bukuTersedia[i].Tahun)
				fmt.Println("Jumlah   : ", bukuTersedia[i].Jumlah)
			}
			fmt.Println("--------------------------------------------------------------------")
		}
	}
}

// insertion sort untuk mengurutkan buku berdasarkan kategori secara descending
func descendingKategoriBuku(A daftarBuku, pinjaman daftarPinjaman, nBuku, nPinjaman int) { // menampilkan buku berdasarkan kategori secara descending
	var bukuTersedia daftarBuku
	var nBukuTersedia int

	if nBuku == 0 {
		fmt.Println("Tidak ada data buku yang tersedia.")
	} else {
		for i := 0; i < nBuku; i++ {
			jumlahTersedia := A[i].Jumlah
			for j := 0; j < nPinjaman; j++ {
				if pinjaman[j].Judul == A[i].Judul && !pinjaman[j].TarifDihitung {
					jumlahTersedia--
				}
			}

			if jumlahTersedia > 0 {
				A[i].Jumlah = jumlahTersedia
				bukuTersedia[nBukuTersedia] = A[i]
				nBukuTersedia++
			}
		}

		if nBukuTersedia == 0 {
			fmt.Println("Tidak ada data buku yang tersedia.")
		} else {
			// insertion sort descending berdasarkan kategori
			for i := 1; i < nBukuTersedia; i++ {
				key := bukuTersedia[i]
				j := i - 1
				for j >= 0 && bukuTersedia[j].Kategori < key.Kategori {
					bukuTersedia[j+1] = bukuTersedia[j]
					j--
				}
				bukuTersedia[j+1] = key
			}

			fmt.Println("Kategori Buku Berdasarkan Descending:")
			for i := 0; i < nBukuTersedia; i++ {
				fmt.Println("--------------------------------------------------------------------")
				fmt.Println("Kategori : ", bukuTersedia[i].Kategori)
				fmt.Println("Judul    : ", bukuTersedia[i].Judul)
				fmt.Println("Penulis  : ", bukuTersedia[i].Penulis)
				fmt.Println("Tahun    : ", bukuTersedia[i].Tahun)
				fmt.Println("Jumlah   : ", bukuTersedia[i].Jumlah)
			}
			fmt.Println("--------------------------------------------------------------------")
		}
	}
}

func pencarianBukuMenu() {
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("                    P E N C A R I A N  B U K U                      ")
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("1. Cari berdasarkan Penulis                                         ")
	fmt.Println("2. Cari berdasarkan Tahun Terbit                                    ")
	fmt.Println("3. Cari berdasarkan Kategori                                        ")
	fmt.Println("4. Kembali ke Menu Utama                                            ")
	fmt.Println("--------------------------------------------------------------------")
}

// sequential search untuk mencari buku berdasarkan penulis
func cariBukuBerdasarkanPenulis(A daftarBuku, pinjaman daftarPinjaman, nBuku, nPinjaman int) { // dapat menampilkan buku yang dicari berdasarkan penulis
	var penulis string

	if nBuku == 0 {
		fmt.Println("Tidak ada data buku yang tersedia.")
	} else {
		found := false

		fmt.Print("Masukkan penulis buku: ")
		fmt.Scanln(&penulis)

		for i := 0; i < nBuku; i++ {
			if A[i].Penulis == penulis {
				jumlahTersedia := A[i].Jumlah
				for j := 0; j < nPinjaman; j++ {
					if pinjaman[j].Judul == A[i].Judul && !pinjaman[j].TarifDihitung {
						jumlahTersedia--
					}
				}

				if jumlahTersedia > 0 {
					if !found {
						fmt.Println("Berikut buku berdasarkan penulis", penulis, ":")
					}
					found = true
					fmt.Println("--------------------------------------------------------------------")
					fmt.Println("Judul    : ", A[i].Judul)
					fmt.Println("Penulis  : ", A[i].Penulis)
					fmt.Println("Tahun    : ", A[i].Tahun)
					fmt.Println("Jumlah   : ", jumlahTersedia)
					fmt.Println("Kategori : ", A[i].Kategori)
				}
			}
		}
		if !found {
			fmt.Println("Buku dengan penulis", penulis, "tidak ditemukan.")
		} else {
			fmt.Println("--------------------------------------------------------------------")
		}
	}
}

// sequential search untuk mencari buku berdasarkan tahun terbit
func cariBukuBerdasarkanTahun(A daftarBuku, pinjaman daftarPinjaman, nBuku, nPinjaman int) { // dapat menampilkan buku yang dicari berdasarkan tahun terbit
	var tahun int

	if nBuku == 0 {
		fmt.Println("Tidak ada data buku yang tersedia.")
	} else {
		found := false

		fmt.Print("Masukkan tahun terbit buku: ")
		fmt.Scanln(&tahun)

		for i := 0; i < nBuku; i++ {
			if A[i].Tahun == tahun {
				jumlahTersedia := A[i].Jumlah
				for j := 0; j < nPinjaman; j++ {
					if pinjaman[j].Judul == A[i].Judul && !pinjaman[j].TarifDihitung {
						jumlahTersedia--
					}
				}

				if jumlahTersedia > 0 {
					if !found {
						fmt.Println("Berikut buku berdasarkan tahun terbit", tahun, ":")
					}
					found = true
					fmt.Println("--------------------------------------------------------------------")
					fmt.Println("Judul    : ", A[i].Judul)
					fmt.Println("Penulis  : ", A[i].Penulis)
					fmt.Println("Tahun    : ", A[i].Tahun)
					fmt.Println("Jumlah   : ", jumlahTersedia)
					fmt.Println("Kategori : ", A[i].Kategori)
				}
			}
		}
		if !found {
			fmt.Println("Buku dengan tahun terbit", tahun, "tidak ditemukan.")
		} else {
			fmt.Println("--------------------------------------------------------------------")
		}
	}
}

// sequential search untuk mencari buku berdasarkan kategori
func cariBukuBerdasarkanKategori(A daftarBuku, pinjaman daftarPinjaman, nBuku, nPinjaman int) { // dapat menampilkan buku yang dicari berdasarkan kategori
	var kategori string

	if nBuku == 0 {
		fmt.Println("Tidak ada data buku yang tersedia.")
	} else {
		found := false

		fmt.Print("Masukkan kategori buku: ")
		fmt.Scanln(&kategori)

		for i := 0; i < nBuku; i++ {
			if A[i].Kategori == kategori {
				jumlahTersedia := A[i].Jumlah
				for j := 0; j < nPinjaman; j++ {
					if pinjaman[j].Judul == A[i].Judul && !pinjaman[j].TarifDihitung {
						jumlahTersedia--
					}
				}

				if jumlahTersedia > 0 {
					if !found {
						fmt.Println("Berikut buku berdasarkan kategori", kategori, ":")
					}
					found = true
					fmt.Println("--------------------------------------------------------------------")
					fmt.Println("Judul    : ", A[i].Judul)
					fmt.Println("Penulis  : ", A[i].Penulis)
					fmt.Println("Tahun    : ", A[i].Tahun)
					fmt.Println("Jumlah   : ", jumlahTersedia)
					fmt.Println("Kategori : ", A[i].Kategori)
				}
			}
		}
		if !found {
			fmt.Println("Buku dengan kategori", kategori, "tidak ditemukan.")
		} else {
			fmt.Println("--------------------------------------------------------------------")
		}
	}
}

func keluar() {
	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║    Terimakasih Telah Menggunakan Aplikasi Perpustakaan Kami!     ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
}

func main() {
	var bukuDaftar daftarBuku
	var pinjamanDaftar daftarPinjaman
	var nBuku int
	var nPinjaman int
	var pilihan int

	// test case, bisa dihapus apabila ingin meng-inputkan manual (satu persatu)
	bukuDaftar[0] = Buku{"LaskarPelangi", "Andrea", 2005, 3, "Novel"}
	bukuDaftar[1] = Buku{"Bumi", "TereLiye", 2014, 2, "Novel"}
	bukuDaftar[2] = Buku{"Hujan", "Sapardi", 1995, 2, "Cerpen"}
	bukuDaftar[3] = Buku{"Soekarno", "Cindy", 1965, 1, "Biografi"}
	bukuDaftar[4] = Buku{"Naruto", "Masashi", 1999, 1, "Komik"}
	bukuDaftar[5] = Buku{"SiJuki", "Faza", 2011, 1, "Komik"}
	nBuku = 6

	for pilihan != 5 {
		menuUtama()
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			var subPilihan int
			for subPilihan != 5 {
				kelolaBukuMenu()
				fmt.Print("Masukkan pilihan Anda: ")
				fmt.Scanln(&subPilihan)
				if subPilihan == 1 {
					tambahDataBuku(&bukuDaftar, &nBuku)
				} else if subPilihan == 2 {
					editDataBuku(&bukuDaftar, pinjamanDaftar, nBuku, nPinjaman)
				} else if subPilihan == 3 {
					hapusDataBuku(&bukuDaftar, pinjamanDaftar, &nBuku, &nPinjaman)
				} else if subPilihan == 4 {
					tampilkanDataBuku(bukuDaftar, pinjamanDaftar, nBuku, nPinjaman)
				} else if subPilihan == 5 {
					subPilihan = 5
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		} else if pilihan == 2 {
			var subPilihan int
			for subPilihan != 7 {
				kelolaPeminjamanBukuMenu()
				fmt.Print("Masukkan pilihan Anda: ")
				fmt.Scanln(&subPilihan)
				if subPilihan == 1 {
					tambahDataPeminjamanBuku(&pinjamanDaftar, &bukuDaftar, &nPinjaman, &nBuku)
				} else if subPilihan == 2 {
					editDataPeminjamanBuku(&pinjamanDaftar, &bukuDaftar, nPinjaman, nBuku)
				} else if subPilihan == 3 {
					hapusDataPeminjamanBuku(&pinjamanDaftar, &nPinjaman)
				} else if subPilihan == 4 {
					tampilkanDataPeminjaman(pinjamanDaftar, nPinjaman)
				} else if subPilihan == 5 {
					hitungTarifDanDenda(&pinjamanDaftar, nPinjaman)
				} else if subPilihan == 6 {
					limaBukuTerfavorit(pinjamanDaftar, nPinjaman)
				} else if subPilihan == 7 {
					subPilihan = 7
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		} else if pilihan == 3 {
			var subPilihan int
			for subPilihan != 3 {
				lihatKategoriDataBukuMenu()
				fmt.Print("Masukkan pilihan Anda: ")
				fmt.Scanln(&subPilihan)
				if subPilihan == 1 {
					ascendingKategoriBuku(bukuDaftar, pinjamanDaftar, nBuku, nPinjaman)
				} else if subPilihan == 2 {
					descendingKategoriBuku(bukuDaftar, pinjamanDaftar, nBuku, nPinjaman)
				} else if subPilihan == 3 {
					subPilihan = 3
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		} else if pilihan == 4 {
			var subPilihan int
			for subPilihan != 4 {
				pencarianBukuMenu()
				fmt.Print("Masukkan pilihan Anda: ")
				fmt.Scanln(&subPilihan)
				if subPilihan == 1 {
					cariBukuBerdasarkanPenulis(bukuDaftar, pinjamanDaftar, nBuku, nPinjaman)
				} else if subPilihan == 2 {
					cariBukuBerdasarkanTahun(bukuDaftar, pinjamanDaftar, nBuku, nPinjaman)
				} else if subPilihan == 3 {
					cariBukuBerdasarkanKategori(bukuDaftar, pinjamanDaftar, nBuku, nPinjaman)
				} else if subPilihan == 4 {
					subPilihan = 4
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		} else if pilihan == 5 {
			keluar()
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
