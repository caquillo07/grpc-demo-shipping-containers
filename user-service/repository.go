package main

import (
	pb "github.com/caquillo07/grpc-demo-shipping-containers/user-service/proto/user"
	"github.com/jinzhu/gorm"
)

type Users []*pb.User

type Repository interface {
	GetAll() (Users, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo UserRepository) GetAll() (Users, error) {
	var users Users
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepository) Create(user *pb.User) error {
	return repo.db.First(&user).Error
}

func (repo UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
