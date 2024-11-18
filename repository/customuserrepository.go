package repository

import (
	"fmt"
	"learn_1/model"
)

type CustomUserRepository struct {
}

var userredisrepo = new(RedisUserRepository)
var usermysqlrepo = new(MysqlUserRepository)

func (customrepo CustomUserRepository) Connect() error {
	if err1 := usermysqlrepo.Connect(); err1 != nil {
		return err1
	} else {
		fmt.Println("sucssesful connection to my sql !!")
	}
	if err2 := userredisrepo.Connect(); err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("sucssesful connection to redis !!")
	}
	return nil
}

func (customrepo CustomUserRepository) InsertUser(user model.User) (model.User, error) {
	if _, err1 := userredisrepo.InsertUser(user); err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("inserted user in mysql")
	}
	if _, err2 := usermysqlrepo.InsertUser(user); err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("inserted user in redis")
	}
	return user, nil
}

func (customrepo CustomUserRepository) UpdateUser(id string, newFirstname string, newLastname string) error {
	if err1 := usermysqlrepo.UpdateUser(id, newFirstname, newLastname); err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Updatad user in mysql")
	}
	if err2 := userredisrepo.UpdateUser(id, newFirstname, newLastname); err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Updated user in mysql")
	}
	return nil

}

func (customrepo CustomUserRepository) DeleteUser(id string) error {
	if err1 := usermysqlrepo.DeleteUser(id); err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("Deleted user in mysql")
	}
	if err2 := (userredisrepo).DeleteUser(id); err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Deleted user in mysql")
	}
	return nil
}

func (customrepo CustomUserRepository) Read(id string) error {
	err := (userredisrepo).Read(id)
	if err != nil {
		err = (usermysqlrepo).Read(id)
	}
	return err
}

func (customrepo CustomUserRepository) Read_all() {
	(usermysqlrepo).Read_all()
}
