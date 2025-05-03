package repository

import (
	"awesomeProject2/internal/db"
	"awesomeProject2/internal/utils"
	"awesomeProject2/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type UserRepository struct {
	DB *db.DB
}

func NewUserRepository(database *db.DB) *UserRepository {
	return &UserRepository{DB: database}
}

func (r *UserRepository) Register(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	normalizedUsername := strings.ToLower(user.Username)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = r.DB.Conn.Exec(ctx, `
	INSERT INTO shop_users (username, password, age, identification, email, phone) 
	VALUES ($1, $2, $3, $4, $5, $6)`,
		normalizedUsername, string(hashedPassword), user.Age, user.Identification, user.Email, user.Phone)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return errors.New("username already exists")
		}
		return err
	}

	return nil
}

func (r *UserRepository) Login(username, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var passwordHash string
	err := r.DB.Conn.QueryRow(ctx, "SELECT password FROM shop_users WHERE username=$1", username).Scan(&passwordHash)

	if errors.Is(err, pgx.ErrNoRows) {
		return "", errors.New("invalid username or password")
	} else if err != nil {
		return "", err
	}

	fmt.Println("DB hash:", passwordHash)
	fmt.Println("User password:", password)

	if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(username)
	if err != nil {
		return "", err
	}

	return token, nil
}
