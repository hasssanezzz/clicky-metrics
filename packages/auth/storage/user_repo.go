package storage

import "golang.org/x/crypto/bcrypt"

type UserRepo struct{}

func (r *UserRepo) FindByUsername(username string) (User, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)
	return user, result.Error
}

func (r *UserRepo) FindByEmail(email string) (User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (r *UserRepo) Create(user *User) (*User, error) {
	result := db.Create(user)
	return user, result.Error
}

func (r *UserRepo) VerifyPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
