package models

import (
  	"time"

  	"golang.org/x/crypto/bcrypt"
  	"gorm.io/gorm"
  )

// User represents a user in the system
type User struct {
  	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
  	Name      string         `json:"name" gorm:"not null"`
  	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
  	Password  string         `json:"-" gorm:"not null"`
  	Role      string         `json:"role" gorm:"default:'user'"`
  	CreatedAt time.Time      `json:"created_at"`
  	UpdatedAt time.Time      `json:"updated_at"`
  	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
  }

// RegisterRequest holds data for user registration
type RegisterRequest struct {
  	Name     string `json:"name" binding:"required,min=2"`
  	Email    string `json:"email" binding:"required,email"`
  	Password string `json:"password" binding:"required,min=6"`
  }

// LoginRequest holds data for user login
type LoginRequest struct {
  	Email    string `json:"email" binding:"required,email"`
  	Password string `json:"password" binding:"required"`
  }

// UpdateUserRequest holds data for updating a user
type UpdateUserRequest struct {
  	Name  string `json:"name"`
  	Email string `json:"email"`
  }

// HashPassword hashes the user's password before storing
func (u *User) HashPassword() error {
  	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
  	if err != nil {
      		return err
      	}
  	u.Password = string(hashed)
  	return nil
  }

// CheckPassword verifies a plain password against the stored hash
func (u *User) CheckPassword(plain string) bool {
  	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
  	return err == nil
  }
