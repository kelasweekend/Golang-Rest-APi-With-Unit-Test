package Activity

import (
	Config "go-restapi-mysql/config"

	_ "github.com/go-sql-driver/mysql"
)

func GetAll(activity *[]Activity) (err error) {
	if err = Config.DB.Find(activity).Error; err != nil {
		return err
	}
	return nil
}
