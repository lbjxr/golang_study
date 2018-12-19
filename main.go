package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"

	"strings"

	"strconv"
)

//substring方法实现
func substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length ||	start > end {
		return "参数值不正确"
	}	

	if start == 0 && end == length {
		return source
	}

	return string(r[start : end])
}


func createExcel(){
	xlsx := excelize.NewFile()

	index := xlsx.NewSheet("Sheet2")

	xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	xlsx.SetCellValue("Sheet1", "B2", 100)

	xlsx.SetActiveSheet(index)

	err := xlsx.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func readExcel(){
	xlsx, err := excelize.OpenFile("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)

	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Println(colCell, "\t")
		}
		fmt.Println()
	}
}

func readTestExcel(){
	xlsx, err := excelize.OpenFile("./test.xlsx")

	if err != nil {
		fmt.Println(err)
		return
	}

	rows := xlsx.GetRows("sheet1")

	for j, row := range rows {
		for i, colCell := range row {
			//fmt.Print(colCell,"\t")
			//或当前行第四列的数据，即D列的值
			if i == 3 {
				fmt.Println(colCell, "\t")

				//对D列值进行分析,查找关键字为“;”
				//如果存在，则取分号后6位数字
				if strings.Index(colCell, ";") != -1 {
					fmt.Println(strings.Index(colCell, ";"))
					index := strings.Index(colCell, ";")
					result := substring(colCell, index + 1, index + 7)

					fmt.Println(result)

					//将结果写到下一列
					cellNum := "E" + strconv.Itoa(j)
					//fmt.Println(cellNum)
					xlsx.SetCellValue("sheet1", string(cellNum), result)
					
				}

			}
		}
		fmt.Println("-------------------------")
	}
	//保存为新的表格
	err = xlsx.SaveAs("./test2.xlsx")

	if err != nil{
		fmt.Println(err)
	}
}

func main() {
	
	//createExcel()

	//readExcel()

	readTestExcel()

}