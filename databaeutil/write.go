package databaeutil

import (
	"fmt"
)

type Record struct {
	Key      string
	Value    string
	DataType DataTypes
}

func store(record Record) {
	fmt.Printf("/nDataType = %d/nKey = %s/nValue = %s", record.DataType, record.Key, record.Value)
}
