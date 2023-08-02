# dbapi
## 把一些简单的go-leveldb操作封装成函数
### 使用方法：
``` golang
package main

import (
  "github.com/cross10/dbapi"
  "github.com/syndtr/goleveldb/leveldb"
)

var db  *leveldb.DB

func init() {
	var err error
	db, err = leveldb.OpenFile("db/test", nil)
	if err != nil {
		log.Println("连接db失败!")
		panic(err)
	}
}

func main() {
	defer db.Close()

  // 存储
	if err := dbapi.SetData("test", []byte{'a'}, db); err != nil {
		 panic(err)
	}

  // 读取
	has, err := dbapi.HaveKey("test", db)
	if err != nil {
		 panic(err)
	}
  if has {
    fmt.Println(string(has))
  }

}
```
