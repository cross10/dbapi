package dbapi

//BatchDataList 批处理数据
type BatchDataList struct {
	Method string //操作方法    put写入 del删除
	Key    string
	Value  []byte
}
