package Todo

import (
	Config "go-restapi-mysql/config"

	_ "github.com/go-sql-driver/mysql"
)

func GetAll(todo *[]Todo) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}
