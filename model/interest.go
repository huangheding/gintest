package model

import (
	"database/sql"
	"gintest/common"
)

// 定义interrest类型结构
type Interest struct {
	Id       string         `json:"id"`
	Name     sql.NullString `json:"name"`
	FullName sql.NullString `json:"fullName"`
	ParentID string         `json:"parentID"`
	Code     sql.NullString `json:"code"`
	Sort     sql.NullString `json:"sort"`

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
		err := common.Fill(temp, rows)
		if err != nil {
			return nil, err
		}
		item = append(item, temp)
	}
	for _, v := range item {
		//最上层父级
		if v.ParentID == "0" {
			interests = append(interests, v)
			linkedlist(item, v)
		}
	}

	defer rows.Close()
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
