package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	fmt.Println(`^([^/]+?)` +
		// path must be one or more "/element"s
		`(/[^/]+)*` +
		// must end with "/*"
		`/\*`)
	//fqnRegexp := regexp.MustCompile("^([^/]+?)((/[^/]+)+|/)(#[^/]+)(/.+)*")
	//a, e := strconv.ParseInt("", 10, 32)
	//fmt.Println(a)
	//test := make(map[string]int)
	var err error
	dirPath := "/Users/tanx/go/bazel-testlogs"
	dirPath, err = filepath.EvalSymlinks(dirPath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("---")
	fmt.Println(dirPath)
	messages := []string{}
	fmt.Println("=======")
	fmt.Println(strings.Join(messages, "\n"))
	fmt.Println("=======")

}
