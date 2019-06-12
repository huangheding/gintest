package model

import (
	"fmt"
	"time"
)

//定义person类型结构
type Person struct {
	ID   string    `gorm:"primary_key;type:int(5);auto_increment" json:"id"`
	Name string    `gorm:"name;type:varchar(100)" json:"name"`
	Age  string    `gorm:"age;type:int(5)" json:"age"`
	Date time.Time `json:"date"`
}

//find
func (p Person) FindPerson() (persons []Person, err error) {
	//这里使用db.Table指定表 是因为struct person映射的表应该是people,所以要手动指定
	//详见model.InitDB()函数
	if err := db.Table("person").Find(&persons).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	// persons = append(persons, p)
	return
}

//update
func (p Person) UpdatePerson() error {
	if err := db.Table("person").Model(&p).Updates(&p).Error; err != nil {
		return err
	}
	return nil
}

//delete
func (p Person) DeletePerson(id string) error {

	if err := db.Table("person").Where("id=?", id).Delete(&p).Error; err != nil {
		return err
	}
	return nil
}

//add
func (p Person) AddPerson() error {
	if err := db.Table("person").Create(&p).Error; err != nil {
		return err
	}
	return nil
}
