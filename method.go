package dbapi

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// SetData 写入数据
func SetData(key string, value []byte, db *leveldb.DB) (err error) {
	err = db.Put([]byte(key), value, nil)
	return
}

// DelData 删除数据
func DelData(key string, db *leveldb.DB) (err error) {
	err = db.Delete([]byte(key), nil)
	return
}

// GetData 获取数据
func GetData(key string, db *leveldb.DB) (data []byte, err error) {
	data, err = db.Get([]byte(key), nil)
	return
}

// RangeData 遍历数据
func IterateData(db *leveldb.DB) (data []map[string]string, err error) {

	iter := db.NewIterator(nil, nil)

	for iter.Next() {
		data = append(data, map[string]string{
			string(iter.Key()): string(iter.Value()),
		})
	}
	iter.Release()
	err = iter.Error()
	return
}

// SeekIterate 搜索迭代
func SeekIterate(key string, db *leveldb.DB) (data []map[string]string, err error) {
	iter := db.NewIterator(nil, nil)
	for ok := iter.Seek([]byte(key)); ok; ok = iter.Next() {
		data = append(data, map[string]string{
			string(iter.Key()): string(iter.Value()),
		})
	}
	iter.Release()
	err = iter.Error()
	return
}

// HaveKey 判断是否存在指定key
func HaveKey(key string, db *leveldb.DB) (have bool, err error) {
	iter := db.NewIterator(nil, nil)
	for ok := iter.Seek([]byte(key)); ok; ok = iter.Next() {
		have = true
		break
	}
	iter.Release()
	err = iter.Error()
	return
}

// PrefixIterate 特定前缀遍历
func PrefixIterate(pre string, db *leveldb.DB) (data []map[string]string, err error) {
	iter := db.NewIterator(util.BytesPrefix([]byte(pre)), nil)
	for iter.Next() {
		data = append(data, map[string]string{
			string(iter.Key()): string(iter.Value()),
		})
	}
	iter.Release()
	err = iter.Error()

	return
}

// BatchData 批处理数据
func BatchData(datalist []BatchDataList, db *leveldb.DB) (err error) {
	batch := new(leveldb.Batch)

	for i := 0; i < len(datalist); i++ {
		switch datalist[i].Method {
		case "put":
			batch.Put([]byte(datalist[i].Key), datalist[i].Value)
		case "del":
			batch.Delete([]byte(datalist[i].Key))
		}
	}

	err = db.Write(batch, nil)
	return
}

// SliceData 根据key切片 keya起始 keyb结束
func SliceData(keya, keyb string, db *leveldb.DB) (data []map[string]string, err error) {

	iter := db.NewIterator(&util.Range{Start: []byte(keya), Limit: []byte(keyb)}, nil)
	for iter.Next() {
		data = append(data, map[string]string{
			string(iter.Key()): string(iter.Value()),
		})
	}
	iter.Release()
	err = iter.Error()
	return
}
