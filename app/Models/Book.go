package Models

import (
	Config "go-restapi-mysql/config"

	_ "github.com/go-sql-driver/mysql"
)

// GetAllUsers Fetch all user data
func GetAllBook(book *[]Book) (err error) {
	if err = Config.DB.Find(book).Error; err != nil {
		return err
	}
	return nil
}
