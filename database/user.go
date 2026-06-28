package database

import (
	"Flaming_Sword_Ministry/models"
)

func CreateUser(user models.User) error {

	query := `
	INSERT INTO users (
		fullname,
		phone,
		gender,
		password,
		role
	)
	VALUES (?, ?, ?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		user.FullName,
		user.Phone,
		user.Gender,
		user.Password,
		user.Role,
	)

	return err
}
func GetUserByPhone(phone string) (models.User, error) {

	var user models.User

	query := `
	SELECT id, fullname, phone, gender, password, role
	FROM users
	WHERE phone = ?
	`

	err := DB.QueryRow(query, phone).Scan(
		&user.ID,
		&user.FullName,
		&user.Phone,
		&user.Gender,
		&user.Password,
		&user.Role,
	)

	return user, err
}
