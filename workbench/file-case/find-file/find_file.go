package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	sep = string(os.PathSeparator)
)

func main() {
	f := "./vendor/vendor.json"
	fp := path.Dir(f)
	dir_list, e := ioutil.ReadDir(fp)
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	var ld string
	for _, v := range dir_list {
		if strings.HasPrefix(v.Name(), "ld-linux") {
			ld = v.Name()
		}
	}
	fmt.Println(ld)
}
