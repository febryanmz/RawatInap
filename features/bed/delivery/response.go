package delivery

import "github.com/KamarRS-App/KamarRS-App/features/bed"

type BedResponse struct {
	ID              uint   `json:"id"`
	NamaTempatTidur string `json:"nama_tempat_tidur"`
	Ruangan         string `json:"ruangan"`
	Kelas           string `json:"kelas"`
	Status          string `json:"status"`
	HospitalID      uint   `json:"hospital_id"`
}

// -----------------Bed--------------------------------
func FromCore(dataCore bed.BedCore) BedResponse {
	return BedResponse{
		ID:              dataCore.ID,
		NamaTempatTidur: dataCore.NamaTempatTidur,
		Ruangan:         dataCore.Ruangan,
		Kelas:           dataCore.Kelas,
		Status:          dataCore.Status,
		HospitalID:      dataCore.HospitalID,
	}
}

// data dari core ke response
func FromCoreList(dataCore []bed.BedCore) []BedResponse {
	var dataResponse []BedResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse
}