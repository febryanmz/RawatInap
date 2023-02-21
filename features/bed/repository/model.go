package repository

import (
	"github.com/KamarRS-App/KamarRS-App/features/bed"

	"gorm.io/gorm"
)

type Bed struct {
	gorm.Model
	NamaTempatTidur string
	Ruangan         string
	Kelas           string
	Status          string
	HospitalID      uint
}

type Hospital struct {
	gorm.Model
	KodeRs            string
	Nama              string
	Foto              string
	Alamat            string
	Provinsi          string
	KabupatenKota     string
	Kecamatan         string
	KodePos           string
	NoTelpon          string
	Email             string
	KelasRs           string
	Pengelola         string
	JumlahTempatTidur string
	StatusPenggunaan  string
	BiayaPendaftaran  string
	Beds              []Bed `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type BedReservation struct {
	gorm.Model
	HospitalId       uint
	StatusPasien     string
	BiayaRegistrasi  int
	OrderId          string
	LinkPembayaran   string
	StatusPembayaran string
	PatientID        uint
	BedID            uint
}

func FromCore(dataCore bed.BedCore) Bed {
	bedGorm := Bed{
		NamaTempatTidur: dataCore.NamaTempatTidur,
		Ruangan:         dataCore.Ruangan,
		Kelas:           dataCore.Kelas,
		Status:          dataCore.Status,
		HospitalID:      dataCore.HospitalID,
	}
	return bedGorm
}


func (dataModel *Bed) ToCore() bed.BedCore {
	return bed.BedCore{
		ID:              dataModel.ID,
		NamaTempatTidur: dataModel.NamaTempatTidur,
		Ruangan:         dataModel.Ruangan,
		Kelas:           dataModel.Kelas,
		Status:          dataModel.Status,
		HospitalID:      dataModel.HospitalID,
	}
}

func ToCoreList(dataModel []Bed) []bed.BedCore {
	var dataCore []bed.BedCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.ToCore())
	}
	return dataCore
}
