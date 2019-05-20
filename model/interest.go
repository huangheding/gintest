package model

import (
	"gintest/common"
)

// 定义interrest类型结构
type Interest struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
	ParentID string `json:"parentID"`
	Code     string `json:"code"`
	Sort     string `json:"sort"`

	Childs []*Interest `json:"child"`
}

//查询cfg_interest表中所有数据
func (i Interest) ArrangeInterest() ([]*Interest, error) {
	rows, err := common.QueryTableData("Select * from cfg_interest")
	if err != nil {
		return nil, err
	}
	interests := make([]*Interest, 0, 10)
	item := make([]*Interest, 0, 10)
	//把数据赋值到interests
	for rows.Next() {
		temp := new(Interest)
		// common.Fill(temp, rows)
		rows.Scan(&temp.Id, &temp.Name, &temp.FullName, &temp.ParentID, &temp.Code, &temp.Sort)
		item = append(item, temp)
	}
	for _, v := range item {
		//最上层父级
		if v.ParentID == "0" {
			interests = append(interests, v)
			linkedlist(item, v)
		}
	}
	return interests, nil
}

//生成struct链表
func linkedlist(data []*Interest, ptemp *Interest) {
	//find child
	childs := make([]*Interest, 0, 10)
	for _, v := range data {
		if v.ParentID == ptemp.Id {
			childs = append(childs, v)
		}
	}

	//append childs 递归
	for _, child := range childs {
		ptemp.Childs = append(ptemp.Childs, child)
		linkedlist(data, child)
	}
}
