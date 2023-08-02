# dbapi
## 把一些简单的go-leveldb操作封装成函数
### 使用示例：
``` golang
package main

import (
	"log"

	"github.com/cross10/dbapi"
	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

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

	// 判断key是否存在
	has, err := dbapi.HaveKey("test", db)
	if err != nil {
		panic(err)
	}
	if has {
		// 获取数据
		if data, err := dbapi.GetData("test", db); err != nil {
			panic(err)
		} else {
			log.Println(string(data))
		}
	}

}
```
