package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=test port=5432 sslmode=prefer"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Person{})

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	personID := "f42f3df3-d82c-4046-975a-1646ec389763"
	personKey := fmt.Sprintf("person:%s", personID)

	val, err := client.Get(context.Background(), personKey).Result()
	if err == redis.Nil {

		fmt.Println("Cache miss. Loading from main source.")

		var person Person
		result := db.First(&person, "id = ?", personID)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				fmt.Println("Record not found in PostgreSQL.")

				emptyResponse := map[string]interface{}{personKey: nil}
				emptyResponseJSON, err := json.Marshal(emptyResponse)
				if err != nil {
					fmt.Printf("failed to marshal empty response: %s \n", err.Error())
					return
				}
				err = client.Set(context.Background(), personKey, emptyResponseJSON, 0).Err()
				if err != nil {
					fmt.Printf("failed to set empty response to redis: %s \n", err.Error())
					return
				}
				val = string(emptyResponseJSON)
			} else {
				log.Fatal(result.Error)
				return
			}
		} else {
			personData, err := json.Marshal(person)
			if err != nil {
				fmt.Printf("failed to marshal: %s \n", err.Error())
				return
			}

			err = client.Set(context.Background(), personKey, personData, 0).Err()
			if err != nil {
				fmt.Printf("failed to set value to redis: %s \n", err.Error())
				return
			}

			val = string(personData)
		}
	} else if err != nil {
		fmt.Printf("failed to get value from redis: %s \n", err.Error())
		return
	} else {
		fmt.Println("Cache hit.")
	}

	// Refresh all data in Redis
	// err = client.FlushDB(context.Background()).Err()
	// if err != nil {
	// 	fmt.Printf("failed to flush redis database: %s \n", err.Error())
	// 	return
	// }

	fmt.Println(val)
}
