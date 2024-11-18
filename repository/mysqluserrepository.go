package repository

import (
	"fmt"
	"log"
	"os"

	"learn_1/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlUserRepository struct {
	DB *gorm.DB
}

func (sqlrepo *MysqlUserRepository) Connect() error {
	pass := os.Getenv("mysqlpassword")
	dsn := "root:" + pass + "@tcp(localhost:3306)/mydbtest"
	var err error
	sqlrepo.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := sqlrepo.DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalln("failed to create entity: ", err)
	}
	return nil
}

func (sqlrepo *MysqlUserRepository) InsertUser(user model.User) (model.User, error) {
	fmt.Println(user)
	result := sqlrepo.DB.Create(&user)

	return user, result.Error
}

func (sqlrepo *MysqlUserRepository) UpdateUser(id string, newFirstname string, newLastname string) error {
	var user model.User
	if err := sqlrepo.DB.First(&user, id).Error; err != nil {
		log.Println(err)
	}

	user.Firstname = newFirstname
	user.Lastname = newLastname

	err := sqlrepo.DB.Save(&user).Error

	return err
}

func (usercrud *MysqlUserRepository) DeleteUser(id string) error {
	result := usercrud.DB.Delete(&model.User{}, id)
	return result.Error
}

func (usercrud *MysqlUserRepository) Read(id string) error {

	var user []model.User
	result := usercrud.DB.First(&user, id)
	for _, attribute := range user {
		fmt.Printf("ID: %s, Firstname: %s , Lastname :%s\n", attribute.ID, attribute.Firstname, attribute.Lastname)
	}
	return result.Error
}

func (usercrud *MysqlUserRepository) Read_all() {

	var Users []model.User
	usercrud.DB.Find(&Users)
	for _, user := range Users {
		fmt.Printf("ID: %s, Firstname: %s , Lastname : %s\n ", user.ID, user.Firstname, user.Lastname)
	}
}
