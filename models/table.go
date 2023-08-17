package models

import (
	"fmt"
	"github.com/motongxue/MySQL2CSV/conf"
	"strings"
)

type Table struct {
	TableName string
	ColName   []string // 表名统一用字符串存储
	Cols      string
}

func NewTable() (*Table, error) {
	columns := conf.C().MySQL.Columns
	table := conf.C().MySQL.Table
	if columns == "" || table == "" {
		return nil, fmt.Errorf("table or columns is empty, please check again")
	}
	return &Table{
		ColName:   strings.Split(columns, ","),
		Cols:      columns,
		TableName: table,
	}, nil
}
