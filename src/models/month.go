package models

import (
	"latihan-gin/src/config"

	"github.com/jinzhu/gorm"
)

type Month struct {
	gorm.Model
	Name string
	Day  int
}

func MonthList() []Month {
	items := []Month{}
	config.DB.Raw("SELECT * FROM months").Scan(&items)
	return items
}

func MonthDetail(ID int) Month {
	item := Month{}
	config.DB.Raw("SELECT * FROM months WHERE id=?", ID).Scan(&item)
	return item
}

func MonthCreate(Name string, Day int) []Month {
	items := []Month{}
	config.DB.Raw("INSERT INTO months (created_at, updated_at, deleted_at,name,day) VALUES( NULL, NULL, NULL, ?, ?)", Name, Day).Scan(&items)

	return items
}

func MonthUpdate(ID int, Name string, Day int) []Month {
	items := []Month{}
	config.DB.Raw("UPDATE months SET name = ?, day = ?  WHERE id = ?", Name, Day, ID).Scan(&items)
	return items
}

func MonthDelete(Id int) []Month {
	items := []Month{}
	config.DB.Raw("DELETE FROM months WHERE id = ?", Id).Scan(&items)
	return items
}
