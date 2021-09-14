package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

//function to save user
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil{
		return user, err
	}

	return user, nil
}

//cari user berdasarkan email
func (r *repository) FindByEmail(email string) (User, error){
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil{
		return user, err
	}

	return user, nil
}

//cari user berdasarkan ID
func (r *repository) FindByID(ID int) (User, error){
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil{
		return user, err
	}

	return user, nil
}

//update data user
func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	
	if err != nil{
		return user, err
	}

	return user, nil
}