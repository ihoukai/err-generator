package main

// OutputInfo 输出信息
type OutputInfo struct {
	XLSXFile    string   // XLSX文件名
	FileName    string   // 文件名
	PackageName string   // 包名
	ClassName   string   // 类名
	Keys        []string // keys
	Vals        []int    // values
	Strs        []string // strings
}
