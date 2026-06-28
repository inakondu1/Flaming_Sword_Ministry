package database

import "Flaming_Sword_Ministry/models"

func CreateUser(user models.User) error {

	query := `
	INSERT INTO users (fullname, phone, gender, password, role)
	VALUES (?, ?, ?, ?, ?)
	`

	_, err := DB.Exec(query,
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
func GetAllUsers() ([]models.User, error) {

	rows, err := DB.Query(`
		SELECT id, fullname, phone, gender, password, role
		FROM users
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User

		err := rows.Scan(
			&u.ID,
			&u.FullName,
			&u.Phone,
			&u.Gender,
			&u.Password,
			&u.Role,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
