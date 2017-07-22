package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func generateGo(output *OutputInfo) {
	// 初始化KeyVals
	var keyValsBuf bytes.Buffer
	for i := 0; i < len(output.Keys); i++ {
		r := fmt.Sprintf(goKeyVals, output.Keys[i], output.Vals[i])
		keyValsBuf.WriteString(r)
	}

	// 初始化ValStrs
	var valStrsBuff bytes.Buffer
	for i := 0; i < len(output.Keys); i++ {
		r := fmt.Sprintf(goValStrs, output.Vals[i], output.Strs[i])
		valStrsBuff.WriteString(r)
	}

	var fileData = strings.Replace(goFrame, "@filename", output.XLSXFile, -1)
	fileData = strings.Replace(fileData, "@packagename", output.PackageName, -1)
	fileData = strings.Replace(fileData, "@key-vals", keyValsBuf.String(), -1)
	fileData = strings.Replace(fileData, "@val-strs", valStrsBuff.String(), -1)
	fileData = strings.Replace(fileData, "@classname", output.ClassName, -1)

	file, err := os.Create(output.FileName + ".go")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	file.WriteString(fileData)
}
