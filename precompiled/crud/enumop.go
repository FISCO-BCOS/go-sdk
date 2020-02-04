package crud

type EnumOP string

const (
	EQ EnumOP = "eq"
	NE EnumOP = "ne"
	GT EnumOP = "gt"
	GE EnumOP = "ge"
	LT EnumOP = "lt"
	LE EnumOP = "le"
	Limit EnumOP = "limit"
)