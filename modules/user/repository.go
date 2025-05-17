package user

import (
	"app/database"
	"database/sql"
	"fmt"
)

func Insert(user *User) (int64, error) {
	db := database.Connect()
	defer db.Close()
	query := `
	INSERT INTO Users (Name, Email, PasswordHash, Role, Hash)
	VALUES (?, ?, ?, ?)
	`
	result, err := db.Exec(query, user.Name, user.Email, user.PasswordHash, RoleDefault, user.Hash)
	if err != nil {
		return 0, fmt.Errorf("user.Insert: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("user.Insert: %v", err)
	}
	return id, nil
}

func Update(user *User) error {
	db := database.Connect()
	defer db.Close()
	query := `
	UPDATE Users SET 
		Name = ?,
		Email = ?,
		PasswordHash = ? 
		Role = ?,
		UpdatedAt = NOW(),
		Active = ?,
		Hash = ?,
		EmailVerified = ?
	WHERE ID = ?
	`
	_, err := db.Exec(query, user.Name, user.Email, user.PasswordHash, user.Role, user.Active, user.Hash, user.EmailVerified, user.ID)
	if err != nil {
		return fmt.Errorf("user.Update: %v", err)
	}
	return nil
}

func Find(id int64) (*User, error) {
	db := database.Connect()
	defer db.Close()
	query := `
	SELECT * FROM Users WHERE ID = ?
	`
	var user *User
	err := db.Get(&user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("user.Find: %v", err)
	}
	return user, nil
}

func Select(filter *UserFilter) ([]*User, error) {
	db := database.Connect()
	defer db.Close()
	query := `
	SELECT * FROM Users WHERE 1 = 1
	`
	args := []any{}
	if filter.ID != nil {
		query += " AND ID = ?"
		args = append(args, *filter.ID)
	}
	if filter.Name != nil {
		query += " AND Name = ?"
		args = append(args, *filter.Name)
	}
	if filter.Email != nil {
		query += " AND Email = ?"
		args = append(args, *filter.Email)
	}
	if filter.PasswordHash != nil {
		query += " AND PasswordHash = ?"
		args = append(args, *filter.PasswordHash)
	}
	if filter.Role != nil {
		query += " AND Role = ?"
		args = append(args, *filter.Role)
	}
	if filter.Active != nil {
		query += " AND Active = ?"
		args = append(args, *filter.Active)
	}
	if filter.Hash != nil {
		query += " AND Hash = ?"
		args = append(args, *filter.Hash)
	}
	if filter.EmailVerified != nil {
		query += " AND EmailVerified = ?"
		args = append(args, *filter.EmailVerified)
	}
	rows, err := db.Queryx(query, args)
	if err != nil {
		return nil, fmt.Errorf("user.Select: %v", err)
	}
	var users []*User
	for rows.Next() {
		var user *User
		err := rows.StructScan(&user)
		if err != nil {
			return nil, fmt.Errorf("user.Select: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
