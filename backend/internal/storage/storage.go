package storage

import (
	"database/sql"
	"fmt"
	"todo-rest-api/internal/models"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Dsn string
}


func NewDBConfig(dsn string) *DBConfig {
	return &DBConfig{
		Dsn: dsn,
	}
}

type Storage  interface {
	GetUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id int) error
	GetTodos() ([]models.Todo, error)
	GetTodoByID(id int) (models.Todo, error)
	CreateTodo(todo models.Todo) (models.Todo, error)
	UpdateTodo(todo models.Todo) (models.Todo, error)
	DeleteTodo(id int) error
}



type PostgresStore	struct {
	db *sql.DB
	Storage
}

func NewPostgresStore(username, name , host , password string) (*PostgresStore, error) {

	fmt.Println(username, name, host , password)

	connStr := fmt.Sprintf("postgresql://postgres:password@mydb:5432?sslmode=disable")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,	
	}, nil
	
}