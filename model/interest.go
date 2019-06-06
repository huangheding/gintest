package model

import "fmt"

// 定义interrest类型结构
type Cfg_interest struct {
	Id       string `gorm:"primary_key";not null; json:"id"`
	Name     string `gorm:"name;type:varchar(500)";not null; json:"name"`
	FullName string `gorm:"fullName;type:varchar(500)"; json:"full_name"`
	ParentID string `gorm:"parentID;type:int(11)";not null; json:"parent_id"`
	Code     string `gorm:"code;type:int(10)";not null; json:"code"`
	Sort     string `gorm:"sort;type:int(11)";not null; json:"sort"`

	Childs []*Cfg_interest `json:"child"`
}

//查询cfg_interest表中所有数据
func (i Cfg_interest) ArrangeInterest() ([]*Cfg_interest, error) {
	//全局禁用表名复数
	db.SingularTable(true)
	interests := make([]*Cfg_interest, 0, 10)
	item := make([]*Cfg_interest, 0, 10)
	//赋值item
	if err := db.Find(&item).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, v := range item {
		//最上层父级
		if v.ParentID == "0" {
			interests = append(interests, v)
			linkedlist(item, v)
		}
	}

	//defer rows.Close()
	return interests, nil
}

//生成struct链表
func linkedlist(data []*Cfg_interest, ptemp *Cfg_interest) {
	//find child
	childs := make([]*Cfg_interest, 0, 10)
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
