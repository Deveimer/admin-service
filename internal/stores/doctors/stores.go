package doctors

import "gorm.io/gorm"

type DoctorStore struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *DoctorStore {
	return &DoctorStore{DB: db}
}
