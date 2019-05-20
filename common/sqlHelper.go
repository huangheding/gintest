package common

import (
	"database/sql"
	"reflect"
)

func Fill(param interface{}, rows *sql.Rows) error {
	// v := reflect.ValueOf(param).Elem()
	// t := reflect.TypeOf(param).Elem()

	// fmt.Println(t.NumField())
	// for i := 0; i < t.NumField(); i++ {
	// 	filed := t.Field(i)
	// 	fv := v.Field(i)
	// 	if !fv.CanSet() || fv.Kind() != reflect.String {
	// 		continue
	// 	}
	// 	col := strings.Split(filed.Tag.Get("json"), ",")[0]
	// 	fmt.Println(col)
	// 	fmt.Println(rows)
	// 	fmt.Println(rows.Columns())

	// 	// v := rows.GetField(col)
	// 	v := "123"
	// 	if v != "" {
	// 		fv.SetString(v)
	// 	}
	// }

	// values := make([]sql.RawBytes, 10)
	s := reflect.ValueOf(param).Elem()
	col, err := rows.Columns()
	if err != nil {
		return err
	}
	rs := make([]interface{}, len(col))
	for i := 0; i < len(col); i++ {
		rs[i] = s.Field(i).Addr().Interface()
	}

	err = rows.Scan(rs...)
	if err != nil {
		return err
	}

	return nil
}
