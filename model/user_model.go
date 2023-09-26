package model

type CreateUserRequest struct {
	Name                   string `json:"nama"`
	PhoneNumber            string `json:"no_hp"`
	NationalIdentityNumber string `json:"nik"`
}

type CreateUserResponse struct {
	AccountNumber string `json:"no_rekening"`
}
