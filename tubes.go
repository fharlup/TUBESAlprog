package main

const nmax = 120

type data_parawisata struct {
	tempat_wisata [nmax]data_wis
	hotel         [nmax]dataHotel
}

type dataHotel struct {
	ukuran  int
	harga   int
	nama    string
	bintang int
}
type data_wis struct {
	nama  string
	harga int
}

func inputDatawis(datawis *data_wis) {
	/*
		if admin menginputkan data tempat wisata
		fs data tempat wisata di simpan di tabel data wisata
			berhenti ketika admin mengetikan stop
	*/
	for data
}
