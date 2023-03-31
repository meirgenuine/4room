package models

import (
	"database/sql"
	"errors"
)

type Category struct {
	ID   int64
	Name string
}

func GetAllCategories(db *sql.DB) ([]*Category, error) {
	rows, err := db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]*Category, 0)
	for rows.Next() {
		category := new(Category)
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func CreateCategory(db *sql.DB, category *Category) error {
	if category.Name == "" {
		return errors.New("category name is required")
	}

	stmt, err := db.Prepare("INSERT INTO categories (name) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(category.Name)
	if err != nil {
		return err
	}

	category.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
