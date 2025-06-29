package userservice

import "gorm.io/gorm"

type UserRepository interface {
	GetUsers() ([]User, error)
	PostUser(user User) (User, error)
	PatchUserByID(id uint, user User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) PostUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) PatchUserByID(id uint, user User) (User, error) {
	result := r.db.Model(&User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return User{}, result.Error
	}
	var updatedUser User
	result = r.db.First(&updatedUser, id)
	if result.Error != nil {
		return User{}, result.Error
	}
	return updatedUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	result := r.db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
