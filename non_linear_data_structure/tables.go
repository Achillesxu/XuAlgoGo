// Package non_linear_data_structure
// Time    : 2021/5/21 4:43 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package non_linear_data_structure

import "fmt"

// Row class
type Row struct {
	Columns []Column
	Id      int
}

// Column class
type Column struct {
	Id    int
	Value string
}

// Table class
type Table struct {
	Rows        []Row
	Name        string
	ColumnNames []string
}

// PrintTable print table content
func PrintTable(t Table) {
	var rows = t.Rows
	fmt.Println(t.Name)
	for _, row := range rows {
		var cols = row.Columns
		for i, col := range cols {
			fmt.Println(t.ColumnNames[i], col.Id, col.Value)
		}
	}
}
