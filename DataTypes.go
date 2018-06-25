package databae

//DataTypes is an enum denoting the datatype being sent or stored
type DataTypes int

const (
	Int     DataTypes = 0
	Decimal DataTypes = 1
	String  DataTypes = 2
	Boolean DataTypes = 3
)

type Record struct {
	DataType DataTypes
	Key      string
	Value    string
}
