package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error) //akan save ke struct save jika berhasil jika gagal akan mengeluarkan error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) { //membuat function Save dari type struct repository
	err := r.db.Create(&user).Error //untuk melakukan save ke database

	if err != nil { //logic ini diperlukan dimana jika terjadi error maka akan mengeluarkan error
		return user, err
	}
	return user, nil // jika tidak error maka akan masuk kedalam database users
}
