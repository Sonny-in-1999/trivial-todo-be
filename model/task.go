package model

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"` // UserId(외래키)를 통해 User 참조. 종속 설정: 연관된 User 레코드 삭제 시 함께 삭제
	UserId    uint      `json:"user_id" gorm:"not null"`
}

type TaskResponse struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
