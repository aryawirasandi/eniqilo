package dataSource

import (
	"fmt"
	"log"
	"time"

	"database/sql"

	"github.com/aryawirasandi/eniqilo/src/domain/entities"
)

func Create(params *entities.CreateUser) (*entities.ResponseCreate[*entities.UserCreated], error) {
	fmt.Println(params, " ", "lah")
	// return nil, nil
	type DataSourceUser struct {
		DB *sql.DB
	}

	i := &DataSourceUser{}

	var id int64
	var createdAt time.Time
	var phoneNumber string
	var Name string

	err := i.DB.QueryRow("INSERT INTO users (name, phonenumber, password, role) VALUES ($1, $2, $3, $4) RETURNING id, phonenumber, name created_at",
		params.Name,
		params.Password,
		params.PhoneNumber,
		params.Role,
	).Scan(&id, &phoneNumber, &Name, &createdAt)

	if err != nil {
		log.Printf("Error Creating User: %s", err)
		return nil, err
	}

	users := &entities.ResponseCreate[*entities.UserCreated]{
		Message: "User registered successfully",
		Data: &entities.UserCreated{
			UserId:      id,
			PhoneNumber: params.PhoneNumber,
			Name:        params.Name,
			AccessToken: "",
		},
	}

	return users, nil
}
