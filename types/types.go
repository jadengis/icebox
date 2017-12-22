package types

// SQLType is an abstract representation of a SQL column type. A given dialect
// implementation will need to map these types to the concrete types.
//
// IceboxType is embedded in the SQLType as every SQLType corresponds to an
// abstract icebox, except with extra argument information.
//
// Args allows one to supply named argument information to the SQL type.
// For example, one may supply the size, or number of decimals as a type.
type SQLType struct {
	IceboxType
	Args map[ArgType]string
}

// NewSQLType constructs a new SQLType object with the given information.
func NewSQLType(iceboxType IceboxType, args map[ArgType]string) *SQLType {
	return &SQLType{
		IceboxType: iceboxType,
		Args:       args,
	}
}

// IceboxType is the abstract data specification of a SQLType.
type IceboxType int

const (
	// Text types
	Char IceboxType = iota
	VarChar
	Text
	MediumText
	LongText
	Blob
	MediumBlob
	LongBlob

	// Numeric types
	TinyInt
	SmallInt
	MediumInt
	Int
	BigInt
	Float
	Double
	Decimal

	// Date Types
	Date
	DateTime
	TimeStamp
	Time
	Year
)

// ArgType is an enumeration of the argument types that can be specified
// when building SQLType.
//
// Size allows one to specify the size of a type (text or numeric)
//
// Decimals allows one to specify the number of decimal places to use.
type ArgType int

const (
	Size ArgType = iota
	Decimals
)
