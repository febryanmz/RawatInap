package delivery

import (
	bedreservation "github.com/KamarRS-App/KamarRS-App/features/bedReservation"
)

type BedReservationResponse struct {
	ID               uint   `json:"id"`
	StatusPasien     string `json:"status_pasien"`
	BiayaRegistrasi  int    `json:"biaya_registrasi"`
	KodeDaftar       string `json:"Kode_daftar"`
	PaymentMethod    string `json:"payment_method"`
	LinkPembayaran   string `json:"link_pembayaran"`
	TransactionId    string `json:"transaction_id" form:"transaction_id"`
	VirtualAccount   string `json:"virtual_account" form:"virtual_account"`
	BankPenerima     string `json:"bank_penerima" form:"bank_penerima"`
	WaktuKedaluarsa  string `json:"waktu_kedaluarsa" form:"waktu_kedaluarsa"`
	QrString         string `json:"qr_string"`
	StatusPembayaran string `json:"status_pembayaran"`
	HospitalID       uint   `json:"hospital_id"`
	BedID            uint   `json:"bed_id"`
	// Patient          PatientResponse `json:"patient"`
}

type PatientResponse struct {
	ID                    uint   `json:"id"`
	NoKk                  string `json:"nokk"`
	Nik                   string `json:"nik"`
	NamaPasien            string `json:"namapasien"`
	JenisKelamin          string `json:"jeniskelamin"`
	TanggalLahir          string `json:"tanggallahir"`
	Usia                  int    `json:"usia" form:"usia"`
	NamaWali              string `json:"namawali"`
	EmailWali             string `json:"emailwali"`
	NoTelponWali          string `json:"notelponwali"`
	AlamatKtp             string `json:"alamatktp"`
	ProvinsiKtp           string `json:"provinsiktp"`
	KabupatenKotaKtp      string `json:"kabupatenkotaktp"`
	AlamatDomisili        string `json:"alamatdomisili"`
	ProvinsiDomisili      string `json:"provinsidomisili"`
	KabupatenKotaDomisili string `json:"kabupatenkotadomisili"`
	NoBpjs                string `json:"nobpjs"`
	KelasBpjs             string `json:"kelasbpjs"`
	FotoKtp               string `json:"fotoktp"`
	FotoBpjs              string `json:"fotobpjs"`
	// BedReservation          BedReservation
	// CheckupReservation      CheckupReservation
}

// -----------------Bed Reserve from Core--------------------------------
func fromCore(dataCore bedreservation.BedReservationCore) BedReservationResponse {
	return BedReservationResponse{
		ID:               dataCore.ID,
		StatusPasien:     dataCore.StatusPasien,
		BiayaRegistrasi:  dataCore.BiayaRegistrasi,
		KodeDaftar:       dataCore.KodeDaftar,
		PaymentMethod:    dataCore.PaymentMethod,
		LinkPembayaran:   dataCore.LinkPembayaran,
		VirtualAccount:   dataCore.VirtualAccount,
		TransactionId:    dataCore.TransactionId,
		BankPenerima:     dataCore.BankPenerima,
		WaktuKedaluarsa:  dataCore.WaktuKedaluarsa,
		QrString:         dataCore.QrString,
		StatusPembayaran: dataCore.StatusPembayaran,
		HospitalID:       dataCore.HospitalID,
		BedID:            dataCore.BedID,
		// Patient: PatientResponse{
		// 	ID:                    dataCore.Patient.ID,
		// 	NoKk:                  dataCore.Patient.NoKk,
		// 	Nik:                   dataCore.Patient.Nik,
		// 	NamaPasien:            dataCore.Patient.NamaPasien,
		// 	JenisKelamin:          dataCore.Patient.JenisKelamin,
		// 	TanggalLahir:          dataCore.Patient.TanggalLahir,
		// 	Usia:                  dataCore.Patient.Usia,
		// 	NamaWali:              dataCore.Patient.NamaWali,
		// 	EmailWali:             dataCore.Patient.EmailWali,
		// 	NoTelponWali:          dataCore.Patient.NoTelponWali,
		// 	AlamatKtp:             dataCore.Patient.AlamatKtp,
		// 	ProvinsiKtp:           dataCore.Patient.ProvinsiKtp,
		// 	KabupatenKotaKtp:      dataCore.Patient.KabupatenKotaKtp,
		// 	AlamatDomisili:        dataCore.Patient.AlamatDomisili,
		// 	ProvinsiDomisili:      dataCore.Patient.ProvinsiDomisili,
		// 	KabupatenKotaDomisili: dataCore.Patient.KabupatenKotaDomisili,
		// 	NoBpjs:                dataCore.Patient.NoBpjs,
		// 	KelasBpjs:             dataCore.Patient.KelasBpjs,
		// 	FotoKtp:               dataCore.Patient.FotoKtp,
		// 	FotoBpjs:              dataCore.Patient.FotoBpjs,
		// 	// UserID:                dataCore.Patient.UserID,
		// },
	}
}

// data dari core ke response
func fromCoreList(dataCore []bedreservation.BedReservationCore) []BedReservationResponse {
	var dataResponse []BedReservationResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

//-----------------------Patient from Core---------------------------------------

func PatientCoreToPatientRespon(dataCore bedreservation.PatientCore) PatientResponse { // data user core yang ada di controller yang memanggil user repositoCorePatient
	return PatientResponse{
		ID:                    dataCore.ID,
		NoKk:                  dataCore.NoKk,
		Nik:                   dataCore.Nik,
		NamaPasien:            dataCore.NamaPasien,
		JenisKelamin:          dataCore.JenisKelamin,
		TanggalLahir:          dataCore.TanggalLahir,
		Usia:                  dataCore.Usia,
		NamaWali:              dataCore.NamaWali,
		EmailWali:             dataCore.EmailWali,
		NoTelponWali:          dataCore.NoTelponWali,
		AlamatKtp:             dataCore.AlamatKtp,
		ProvinsiKtp:           dataCore.ProvinsiKtp,
		KabupatenKotaKtp:      dataCore.KabupatenKotaKtp,
		AlamatDomisili:        dataCore.AlamatDomisili,
		ProvinsiDomisili:      dataCore.ProvinsiDomisili,
		KabupatenKotaDomisili: dataCore.KabupatenKotaDomisili,
		NoBpjs:                dataCore.NoBpjs,
		KelasBpjs:             dataCore.KelasBpjs,
		FotoKtp:               dataCore.FotoKtp,
		FotoBpjs:              dataCore.FotoBpjs,
	}
}
func ListpatientCoreTopatientRespon(dataCore []bedreservation.PatientCore) []PatientResponse { //data user.core data yang diambil dari entities ke respon struct
	var ResponData []PatientResponse

	for _, value := range dataCore { //memanggil paramete data core yang berisi data user core
		ResponData = append(ResponData, PatientCoreToPatientRespon(value)) // mengambil data mapping dari user core to respon
	}
	return ResponData
}

//----------------------------------------------------------------------------------
