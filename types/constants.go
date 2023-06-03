package types

const (

	// Boolean - Represents a boolean (true/false)
	Boolean ArgTypes = 2

	// Flag - Represents a boolean flag (present - true, absent - false)
	Flag ArgTypes = 4

	// Float - Represents a 64-bit floating-point number
	Float ArgTypes = 3

	// Integer - Represents a 64-bit integer
	Integer ArgTypes = 1

	// String - Represents a string
	String ArgTypes = 0
)

const (

	// eMsgTypeCheckBoolean - type-check failed error
	eMsgTypeCheckBoolean = "type-check failed (expected Boolean)"

	// eMsgTypeCheckFlag - type-check failed error
	eMsgTypeCheckFlag = "type-check failed (expected Flag)"

	// eMsgTypeCheckFloat - type-check failed error
	eMsgTypeCheckFloat = "type-check failed (expected Float)"

	// eMsgTypeCheckInteger - type-check failed error
	eMsgTypeCheckInteger = "type-check failed (expected Integer)"

	// eMsgTypeCheckString - type-check failed error
	eMsgTypeCheckString = "type-check failed (expected String)"

	// eMsgTypeCheckUnknown - type-check failed error
	eMsgTypeCheckUnknown = "type-check failed (unknown type)"

	errInvalidArgType = "invalid argument type %s"
)
