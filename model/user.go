package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"` // ID는 항상 부호없는 정수, pk
	Email     string    `json:"email" gorm:"unique"`  // Email은 로그인에 사용하므로 고유값 설정
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primarykey"`
	Email string `json:"email" gorm:"unique"`
}
