package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

var sqlHeader = "INSERT INTO `customers` (`id`, `name`, `economy_id`, `address`, `bin`, `email`, `bank_details`, `created_at`, `updated_at`)"

func main() {
	excelFileName := "./data/customers.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	check(err)

	sqlFileName := "./data/customers.sql"
	file, err := os.Create(sqlFileName)
	check(err)
	defer file.Close()

	file.WriteString(sqlHeader)
	file.Write([]byte("\n"))
	file.WriteString("VALUES")

	for i, row := range xlFile.Sheets[0].Rows {
		if i == 0 {
			continue
		}
		name := clean(row.Cells[1].String())
		addr := clean(row.Cells[5].String())
		cbin := clean(row.Cells[8].String())
		data := fmt.Sprintf("(%d,'%s',3,'%s','%s','','',NOW(),NOW())", i, name, addr, cbin)
		file.Write([]byte("\n\t"))
		file.WriteString(data)
		if i == len(xlFile.Sheets[0].Rows)-1 {
			file.WriteString(";")
		} else {
			file.WriteString(",")
		}
	}

	fmt.Println("Done.")
}

func clean(str string) string {
	return strings.TrimSpace(strings.Trim(str, ""))
}

func check(err error) {
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
