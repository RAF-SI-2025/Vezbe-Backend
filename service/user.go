package service

import (
	"errors"

	"github.com/RAF-SI-2025/Vezbe-Backend/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var users = []model.User{}

func CreateUser(firstName, lastName, email, password string) (model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hash),
	}
	users = append(users, user)
	return user, nil
}

func GetAllUsers() []model.User {
	return users
}

func GetUserByID(id uuid.UUID) (model.User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func UpdateUser(id uuid.UUID, firstName, lastName, email string) (model.User, error) {
	for i, u := range users {
		if u.ID == id {
			users[i].FirstName = firstName
			users[i].LastName = lastName
			users[i].Email = email
			return users[i], nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func DeleteUser(id uuid.UUID) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func ChangePassword(id uuid.UUID, oldPassword, newPassword string) error {
	for i, u := range users {
		if u.ID == id {
			if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(oldPassword)); err != nil {
				return errors.New("invalid credentials")
			}
			hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
			if err != nil {
				return err
			}
			users[i].Password = string(hash)
			return nil
		}
	}
	return errors.New("user not found")
}

func CheckPassword(user model.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}
