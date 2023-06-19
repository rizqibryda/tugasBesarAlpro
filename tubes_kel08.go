// Program yang dikerjakan oleh Rizqi Bhamaryuda dan M. Rizqi Fadhilah dari kelas IT-46-01 dari kelompok 8
package main

import "fmt"

const NMAX int = 1024 // untuk batasan index

type admin struct { // tipe data untuk mengakses sebagai admin
	nama  		string
	nilai 		int
	status 		string
	TINDAKAN	int
}

type major struct { // tipe data untuk jurusan pada tipe buatan dari "Mahasiswa"
	namaJurusan string // berdasarkan dari i_jurusan atau nomor jurusan yang di list
	i_jurusan   int // nomor jurusan
}

type Mahasiswa struct { // tipe data untuk mengakses sebagai user
	nama    	string
	jurusan 	major
	nilai   	int
	status		string
}

type NMahasiswa [NMAX]Mahasiswa // deklarasi sebagai indeks dan batasannya

func main() { // program utama
	var (
		i, kesekian int
		mhs         NMahasiswa
		add         [NMAX]int

		pilihan_admin int
		change        admin
	)
	
	fmt.Println("====================================================================")
	fmt.Println("            APLIKASI PENDAFTARAN CALON MAHASISWA BARU       ")
	fmt.Println("	Dibuat oleh Rizqi Bhamaryuda dan M. Rizqi Fadhilah         ")
	fmt.Println("	     Tugas Besar Algoritma Pemrograman 2023               ")
	fmt.Println("====================================================================")

	fmt.Print("Tekan tombol '1' untuk registrasi atau '2' sebagai admin dan tekan 'Enter' untuk melanjutkan...  ")
	fmt.Scan(&add[i])

	for add[i] == 1 || add[i] == 2 {
		for add[i] == 1 {
			i++

			fmt.Println("===============")
			fmt.Print("\nMasukkan data nama mahasiswa (JANGAN menggunakan spasi): ") // Masukkan data nama camaba
			fmt.Scan(&mhs[i].nama)

			listJurusanMahasiswa()

			fmt.Print("Masukkan data program studi mahasiswa: ") // Masukkan data jurusan untuk camaba
			dataJurusan(&mhs, i)

			fmt.Println("========================================") // Batasan biar enak dilihat
			fmt.Print("Masukkan data nilai mahasiswa: ") // Masukkan data nilai camaba
			nilaiMahasiswa(&mhs, i)

			switchProgram(&kesekian)
			add[i] = kesekian
		}

		for add[i] == 2 {

			fmt.Println("=========================================================")
			fmt.Println("      APLIKASI PENDAFTARAN CALON MAHASISWA BARU         ")
			fmt.Println("                Selamat datang, Admin                  ")
			fmt.Println("=========================================================")
			fmt.Println("---- Menu Admin: ----")
			fmt.Println("0. Keluar dari sebagai admin")
			fmt.Println("1. List mahasiswa baru berdasarkan nama")
			fmt.Println("2. List mahasiswa baru berdasarkan jurusan")
			fmt.Println("3. List mahasiswa baru berdasarkan nilai dan status")
			fmt.Println("4. List mahasiswa baru semuanya")
			fmt.Print("Pilihan: ")
			fmt.Scan(&pilihan_admin) // pengguna diminta oleh dari 4 opsi yang tersedia

			if pilihan_admin >= 1 && pilihan_admin <= 4 {
				listMahasiswa(&pilihan_admin, mhs, i, kesekian) // DILIHAT KEMBALI
				konfirmasi(&pilihan_admin, &kesekian)
				if pilihan_admin == 1 && kesekian != -1 {
					dataMahasiswa(mhs, kesekian)
					gantiHapusMahasiswa(pilihan_admin, &mhs, kesekian, &change, &i)
				}
			}

			if pilihan_admin >= 0 {
				switchProgram(&kesekian)
				add[i] = kesekian
			}
		}
	}
}

func switchProgram(nomor *int) { // mengganti role and keluar dari aplikasi

	fmt.Println("============== OPTIONS ===================")
	fmt.Println("0. Tidak, keluar dari aplikasi")
	fmt.Println("1. Tambahkan calon mahasiswa baru")
	fmt.Println("2. switch to 'admin' role")
	fmt.Print("Masukkan: ")
	fmt.Scan(&*nomor)

	if *nomor == 1 || *nomor == 2 {
	} else if *nomor == 0 {
		fmt.Println("Terima kasih telah menggunakan aplikasi kami.")
	} else {
		fmt.Println("Input yang Anda masukkan di luar pada data kami. Masukkan kembali.")
		switchProgram(nomor)
	}
}

func nilaiMahasiswa(mhs *NMahasiswa, i int) { // menilai apakah mahasiswa tersebut dinyatakan lulus atau tidak

	fmt.Scan(&mhs[i].nilai)
	
	if mhs[i].nilai >= 70 && mhs[i].nilai <= 100 {
		fmt.Println("Dinyatakan LULUS seleksi.")
		mhs[i].status = "LULUS"
	} else if mhs[i].nilai < 70 && mhs[i].nilai >= 0 {
		fmt.Println("Dinyatakan TIDAK LULUS seleksi.")
		mhs[i].status = "TIDAK LULUS"
	} else {
		fmt.Println("Data yang Anda masukkan di luar nalar. Masukkan kembali!")
		nilaiMahasiswa(mhs, i)
	}
}

func listMahasiswa(choiceAdmin *int, mhs NMahasiswa, i int, kesekian int) { // menampilkan list mahasiswa berdasarkan pilihan pengguna

	if *choiceAdmin == 1 {			// menampilkan list data camaba dari nama
		for i = 1; i <= NMAX && mhs[i].nama != ""; i++ {
			fmt.Println(i, "-", mhs[i].nama)
		}
		fmt.Println("==========================")
	} else if *choiceAdmin == 2 {		// menampilkan list data camaba dari jurusan
		for i = 1; i <= NMAX && mhs[i].jurusan.namaJurusan != ""; i++ {
			fmt.Println(i, "-", mhs[i].jurusan.namaJurusan)
		}
		fmt.Println("==========================")
	} else if *choiceAdmin == 3 {		// menampilkan list data camaba dari nilai
		for i = 1; i <= NMAX && mhs[i].nilai != 0; i++ {
			fmt.Println(i, "- Nilai", mhs[i].nilai, "dengan status", mhs[i].status)
		}
		fmt.Println("==========================")
	} else if *choiceAdmin == 4 {		// menampilkan list seluruh data tiap camaba
		for i = 1; i <= NMAX && mhs[i].nama != ""; i++ {
			fmt.Println(i, "-", mhs[i].nama, "dari jurusan", mhs[i].jurusan.namaJurusan, "dinyatakan", mhs[i].status, "dengan nilai", mhs[i].nilai)
		}
	} else if *choiceAdmin == 0 {		// balik ke switchProgram atau options
		fmt.Println("==========================")
		fmt.Println("Sampai jumpa, Admin!")
	} else {
		fmt.Println("Pilihan yang Anda masukkan di luar data kami. Masukkan kembali!")
		listMahasiswa(choiceAdmin, mhs, i, kesekian)
	}
}

func konfirmasi(choiceAdmin *int, kesekian *int) {
	fmt.Println("Apakah Anda ingin melihat detail dari salah satu mahasiswa? ")
	fmt.Println("0. Batal")
	fmt.Println("1. Lihat")

	fmt.Print("Pilihan: ")
	fmt.Scan(&*choiceAdmin)

	if *choiceAdmin == 1 {
		fmt.Println("=================================================================")

		fmt.Print("Nomor berapa yang Anda lihat? (Ketik '-1' untuk membatalkan): ")
		fmt.Scan(&*kesekian)
	}
}

func dataMahasiswa(mhs NMahasiswa, kesekian int) {
	if kesekian != -1 {
		fmt.Println("------------------ Berikut nama calon mahasiswa baru --------------")
		fmt.Println("Nama:", mhs[kesekian].nama)
		fmt.Println("Jurusan:", mhs[kesekian].jurusan.namaJurusan)
		fmt.Println("Nilai:", mhs[kesekian].nilai)
		fmt.Println("Status:", mhs[kesekian].status)
		fmt.Println("===============================================")
	}
}

func gantiHapusMahasiswa(choiceAdmin int, mhs *NMahasiswa, kesekian int, ubah *admin, i *int) { // mengganti data mahasiswa dari nama/jurusan/nilai
	if choiceAdmin == 1 {
		fmt.Println("Silakan pilih yang jika ingin Anda ubah/hapus:")
		fmt.Println("0. Batal")
		fmt.Println("1. Nama")
		fmt.Println("2. Jurusan")
		fmt.Println("3. Nilai")
		fmt.Println("99. HAPUS")
		fmt.Print("Pilihan: ")
		fmt.Scan(&ubah.TINDAKAN)
	}

	if choiceAdmin == 1 && ubah.TINDAKAN == 1 {
		fmt.Println("Nama sebelumnya adalah →", mhs[kesekian].nama)
		fmt.Print("Masukkan nama yang ingin diganti: ")
		fmt.Scan(&ubah.nama)
		mhs[kesekian].nama = ubah.nama
		fmt.Println("NAMA BERHASIL DIUBAH →", mhs[kesekian].nama)
	} else if choiceAdmin == 1 && ubah.TINDAKAN == 2 {
		fmt.Println("Program studi sebelumnya adalah →", mhs[kesekian].jurusan.namaJurusan)
		listJurusanMahasiswa()
		fmt.Print("Masukkan NOMOR jurusan yang ingin diganti: ")
		dataJurusan(mhs, kesekian)
		fmt.Println("JURUSAN BERHASIL DIUBAH →", mhs[kesekian].jurusan.namaJurusan)
	} else if choiceAdmin == 1 && ubah.TINDAKAN == 3 {
		fmt.Println("Nilai sebelumnya adalah →", mhs[kesekian].nilai)
		fmt.Print("Masukkan berapa nilai yang ingin diganti (0–100): ")
		fmt.Scan(&ubah.nilai)
		fmt.Println()

		mhs[kesekian].nilai = ubah.nilai

		fmt.Println("NILAI BERHASIL DIUBAH →", mhs[kesekian].nilai)
		if mhs[kesekian].nilai >= 70 && mhs[kesekian].nilai <= 100 {
			mhs[kesekian].status = "LULUS"
		} else if mhs[kesekian].nilai < 70 && mhs[kesekian].nilai >= 0 {
			mhs[kesekian].status = "TIDAK LULUS"
		}

		fmt.Println("STATUS →", mhs[kesekian].status)
	} else if choiceAdmin == 1 && ubah.TINDAKAN == 99 {
		for kesekian <= *i {
			mhs[kesekian].nama = mhs[kesekian+1].nama
			mhs[kesekian].jurusan.namaJurusan = mhs[kesekian+1].jurusan.namaJurusan
			mhs[kesekian].nilai = mhs[kesekian+1].nilai

			kesekian++
		}
		*i--
		fmt.Println("DATA MAHASISWA BERHASIL DIHAPUS.")
	} else if ubah.TINDAKAN == 0 {

	} else {
		fmt.Println("Masukkan yang Anda inputkan tidak ada pada data kami, masukkan kembali.")
		fmt.Println("==========================")
		gantiHapusMahasiswa(choiceAdmin, mhs, kesekian, ubah, i)
	}
}

func listJurusanMahasiswa() {
	fmt.Println("Daftar Jurusan: ")
	fmt.Println("1. Akuntansi")
	fmt.Println("2. Teknik Industri")
	fmt.Println("3. Informatika")
	fmt.Println("4. Rekayasa Perangkat Lunak (RPL)")
	fmt.Println("5. Ilmu Komunikasi")
	fmt.Println("6. Creative Arts")
	fmt.Println("7. Electrical Energy Engineering")
	fmt.Println("8. Teknik Biomedis")
	fmt.Println("9. Hospitality & Culinary Arts")
	fmt.Println("10. Sistem Informasi")
	fmt.Println("11. Digital Supply Chain")
	fmt.Println("12. Teknik Telekomunikasi")
	fmt.Println("13. Teknik Elektro")
	fmt.Println("14. Smart Science and Technology")
	fmt.Println("15. Teknik Komputer")
	fmt.Println("16. Data Sains")
	fmt.Println("17. Leisure Management")
	fmt.Println("18. Administrasi Bisnis")
	fmt.Println("19. Digital Public Relation")
	fmt.Println("20. Digital Content Broadcasting")
	fmt.Println("21. Desain Komunikasi Visual")
	fmt.Println("22. Desain Produk & Inovasi")
	fmt.Println("23. Desain Interior")
	fmt.Println("24. Kriya (Fashion & Textile Design)")
	fmt.Println("25. Sistem Informasi Akuntansi")
	fmt.Println("26. Digital Marketing")
	fmt.Println("27. Digital Creative Multimedia")
	fmt.Println("28. Teknologi Informasi (IT)")
	fmt.Println("========================================")
}

func dataJurusan(mhs *NMahasiswa, i int) { // memasukkan data jurusan mahasiswa

	fmt.Scan(&mhs[i].jurusan.i_jurusan)
	
	switch mhs[i].jurusan.i_jurusan {
	case 1:
		mhs[i].jurusan.namaJurusan = "Akuntansi"
	case 2:
		mhs[i].jurusan.namaJurusan = "Teknik Industri"
	case 3:
		mhs[i].jurusan.namaJurusan = "Informatika"
	case 4:
		mhs[i].jurusan.namaJurusan = "Rekayasa Perangkat Lunak"
	case 5:
		mhs[i].jurusan.namaJurusan = "Ilmu Komunikasi"
	case 6:
		mhs[i].jurusan.namaJurusan = "Creative Arts"
	case 7:
		mhs[i].jurusan.namaJurusan = "Electrical Energy Engineering"
	case 8:
		mhs[i].jurusan.namaJurusan = "Teknik Biomedis"
	case 9:
		mhs[i].jurusan.namaJurusan = "Hospitality & Culinary Arts"
	case 10:
		mhs[i].jurusan.namaJurusan = "Sistem Informasi"
	case 11:
		mhs[i].jurusan.namaJurusan = "Digital Supply Chain"
	case 12:
		mhs[i].jurusan.namaJurusan = "Teknik Telekomunikasi"
	case 13:
		mhs[i].jurusan.namaJurusan = "Teknik Elektro"
	case 14:
		mhs[i].jurusan.namaJurusan = "Smart Science and Technology"
	case 15:
		mhs[i].jurusan.namaJurusan = "Teknik Komputer"
	case 16:
		mhs[i].jurusan.namaJurusan = "Data Sains"
	case 17:
		mhs[i].jurusan.namaJurusan = "Leisure Management"
	case 18:
		mhs[i].jurusan.namaJurusan = "Administrasi Bisnis"
	case 19:
		mhs[i].jurusan.namaJurusan = "Digital Public Relation"
	case 20:
		mhs[i].jurusan.namaJurusan = "Digital Content Broadcasting"
	case 21:
		mhs[i].jurusan.namaJurusan = "Desain Komunikasi Visual"
	case 22:
		mhs[i].jurusan.namaJurusan = "Desain Produk & Inovasi"
	case 23:
		mhs[i].jurusan.namaJurusan = "Desain Interior"
	case 24:
		mhs[i].jurusan.namaJurusan = "Kriya (Fashion & Textile Design)"
	case 25:
		mhs[i].jurusan.namaJurusan = "Sistem Informasi Akuntansi"
	case 26:
		mhs[i].jurusan.namaJurusan = "Digital Marketing"
	case 27:
		mhs[i].jurusan.namaJurusan = "Digital Creative Multimedia"
	case 28:
		mhs[i].jurusan.namaJurusan = "Teknologi Informasi"
	default:
		fmt.Println("Input jurusan yang Anda masukkan tidak ada pada data kami. Masukkan kembali!")
		dataJurusan(mhs, i)
	}
}
