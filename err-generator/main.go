package main

import (
	"flag"
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
)

func main() {
	var inFile = flag.String("infile", "", "输入文件")
	var csharpOutFile = flag.String("csharp_out", "", "csharp文件输出名, 如Test.cs")
	var goOutFile = flag.String("go_out", "", "go文件输出名, 如Test.go")
	var packageName = flag.String("package", "", "包名")
	var className = flag.String("class", "", "类名")
	flag.Parse()

	if *inFile == "" {
		fmt.Printf("infile can not be empty!")
		return
	}

	// 如果没有设置go和csharp的输出文件 则使用infile的文件
	if *goOutFile == "" && *csharpOutFile == "" {
		var pos = strings.LastIndex(*inFile, ".")
		if pos == 0 {
			*goOutFile = *inFile
			*csharpOutFile = *inFile
		} else {
			*goOutFile = (*inFile)[0:pos]
			*csharpOutFile = (*inFile)[0:pos]
		}
	}

	if *packageName == "" {
		*packageName = GoPackage
	}

	if *className == "" {
		*className = CSharpPackage
	}

	// 打开文件
	xlFile, err := xlsx.OpenFile(*inFile)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	keys, vals, strs := extractSheet(xlFile.Sheets)

	outInfo := &OutputInfo{
		XLSXFile:    *inFile,
		PackageName: *packageName,
		ClassName:   *className,
		Keys:        keys,
		Vals:        vals,
		Strs:        strs,
	}

	if *goOutFile != "" {
		outInfo.FileName = *goOutFile
		generateGo(outInfo)
	}

	if *csharpOutFile != "" {
		outInfo.FileName = *csharpOutFile
		generateCsharp(outInfo)
	}
}

func extractSheet(sheets []*xlsx.Sheet) (keys []string, vals []int, strs []string) {
	// 根据索引提取对应的值
	keys = make([]string, 0)
	vals = make([]int, 0)
	strs = make([]string, 0)

	for _, sheet := range sheets {
		// 提取Key、Value、String列的索引
		var keyIndex = -1
		var valueIndex = -1
		var stringIndex = -1
		var firstRow = sheet.Rows[0]
		for i, cell := range firstRow.Cells {
			if cell.String() == "Key" {
				keyIndex = i
			} else if cell.String() == "Value" {
				valueIndex = i
			} else if cell.String() == "String" {
				stringIndex = i
			}
		}

		for n, row := range sheet.Rows {
			// 跳过第一行
			if n == 0 {
				continue
			}
			for i, cell := range row.Cells {
				if keyIndex == i {
					keys = append(keys, cell.String())
				} else if valueIndex == i {
					val, err := cell.Int()
					if err != nil {
						val = 0
					}
					vals = append(vals, val)
				} else if stringIndex == i {
					strs = append(strs, cell.String())
				}
			}
		}
	}
	return
}
