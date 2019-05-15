package model

import (
	"fmt"
	"gin_test/common"
)

//定义person类型结构
type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

//
func (p Person) FindPerson() (persons []Person, err error) {
	rows, err := common.QueryTableData("Select * from person")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var person Person
		//遍历表中所有行的信息
		rows.Scan(&person.Id, &person.Name, &person.Age)
		//将person添加到persons中
		persons = append(persons, person)
	}
	defer rows.Close()
	return
}
