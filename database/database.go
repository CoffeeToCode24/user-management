package database

import (
	"database/sql"

	"github.com/user-managemnet/dtos"
)

var DB *sql.DB

var Users = []dtos.User{{Name: "John", Age: 20, ID: 1}, {Name: "Doe", Age: 30, ID: 2}, {Name: "Smith", Age: 40, ID: 3}, {Name: "Tom", Age: 50, ID: 4}}
