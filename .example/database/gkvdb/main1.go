package main

import (
	"fmt"

	"github.com/gogf/gf/database/gkvdb"
)

func main() {
	key := []byte("key")
	//value := []byte("value")

	db := gkvdb.Instance()
	db.SetPath("/tmp/gkvdb/test")
	//db.Set(key, value)
	fmt.Println(db.Get(key))
}
