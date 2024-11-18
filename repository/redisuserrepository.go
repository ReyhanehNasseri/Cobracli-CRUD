package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"learn_1/model"

	"github.com/go-redis/redis/v8"
)

type RedisUserRepository struct {
	RDB *redis.Client
}

var ctx = context.Background()

func (redisRepo *RedisUserRepository) Connect() error {
	redisRepo.RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	_, err := redisRepo.RDB.Ping(ctx).Result()
	return err
}

func (redisRepo *RedisUserRepository) InsertUser(user model.User) (model.User, error) {
	userJSON, _ := json.Marshal(user)
	err2 := redisRepo.RDB.Set(ctx, user.ID, userJSON, 0).Err()
	return user, err2
}

func (redisRepo *RedisUserRepository) UpdateUser(id string, newFirstname string, newLastname string) error {
	newuser := model.User{}
	newuser.ID = id
	newuser.Firstname = newFirstname
	newuser.Lastname = newLastname
	newuserJSON, _ := json.Marshal(newuser)
	err := redisRepo.RDB.Set(ctx, id, newuserJSON, 0).Err()
	return err
}

func (redisRepo *RedisUserRepository) DeleteUser(id string) error {
	result := redisRepo.RDB.Del(ctx, id)
	fmt.Println("User deleted!")
	return result.Err()
}

func (redisRepo *RedisUserRepository) Read(id string) error {
	userData, err := redisRepo.RDB.Get(ctx, id).Result()
	if err == nil {
		fmt.Printf("User: %#v\n", userData)
	}
	return err
}

func (redisRepo *RedisUserRepository) Read_all() {}
