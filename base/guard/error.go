package guard

const (
	ErrRefProviderNil   = "reference provider is nil"
	ErrRefValueNotFound = "reference value not found"
	ErrRefValueNotEQ    = "evaluated value not equal to reference"
	ErrRefValueNotNE    = "evaluated value equal to reference"
	ErrRefValueNotGT    = "evaluated value not greater than reference"
	ErrRefValueNotGE    = "evaluated value not greater-equal than reference"
	ErrRefValueNotLT    = "evaluated value not less than reference"
	ErrRefValueNotLE    = "evaluated value not less-equal than reference"
	ErrValueOutOfRange  = "value out of range"
	ErrValueNotMatch    = "value not match"
)
