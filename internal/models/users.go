package models

type Users struct {
	ID           int    `db:"id" json:"id" validate:"required"`
	Login        string `db:"login" json:"login" validate:"required"`
	PasswordHash string `db:"password_hash" json:"password_hash" validate:"required"`
	Email        string `db:"email" json:"email" validate:"required"`
}
