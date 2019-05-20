package model

import (
	"fmt"
	"gintest/common"
)

//定义person类型结构
type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

// 定义interrest类型结构
type Interest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	ParentId int    `json:"parentId"`
	next     *Interest
	//	Code     int    `json:"code"`
	//	Sort     int    `json:"sort"`
}

//查询cfg_interest表中所有数据

func (i Interest) arrangeInterrest() (interests []Interest, err error) {
	rows, err := common.QueryTableData("Select * from cfg_interest")
	if err != nil {
		fmt.Println(err)
		return
	}
	//把数据赋值到interests
	for rows.Next() {
		var interest Interest
		rows.Scan(&interest.Id, &interest.Name, &interest.FullName, &interest.ParentId)
		interests = append(interests, interest)
	}
	return
}

//把interest数据分层

func listInterest(i *Interest) (in Interest) {
	parentId := 9999
	for i != nil {
		if i.Id == parentId {

		}
		i = i.next //移动指针
		parentId = i.Id
	}
	return
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
