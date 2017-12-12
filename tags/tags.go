package tags

const (
	// Icebox is the go tag for scoping all icebox subtags.
	// Any StructField with the Icebox tag will be considered a persisted field.
	Icebox string = "icebox"
	// Sep is the icebox subtag. Subtags of the icebox tag will be separated by
	// this value.
	Sep string = ","
	// Name is the subtag for specifying the column name for a field.
	Name       string = "name"
	PrimaryKey string = "primaryKey"
)
