package patients

import "gorm.io/gorm"

type PatientStore struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *PatientStore {
	return &PatientStore{DB: db}
}
