package tiga

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Nama string
}
type UserManager struct {
	Users map[int]*User
}

func NewUser(ID int, Nama string) *User {
	return &User{ID: ID, Nama: Nama}
}

func (um *UserManager) AddUser(u *User) error {
	_, exists := um.Users[u.ID]
	if exists {
		err := fmt.Sprintf("User dengan ID %d sudah ada\n", u.ID)
		return errors.New(err)
	}
	um.Users[u.ID] = u
	fmt.Printf("User dengan ID %d bertambah\n", u.ID)

	return nil
}

func (um *UserManager) GetUser(ID int) (*User, error) {
	user, exists := um.Users[ID]
	if !exists {
		err := fmt.Sprintf("User dengan ID %d tidak ditemukan\n", ID)
		return nil, errors.New(err)
	}
	return user, nil
}

func (um *UserManager) GetAllUsers() {
	for _, user := range um.Users {
		fmt.Println(*user)
	}
}
