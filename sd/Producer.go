package SD

import "fmt"

type RandomType int

const (
	NONE      RandomType = iota
	UNIFORM
	GAUSSIAN
)



type DataPair struct{
	RawValue interface{}
	StringValue string
	BoundExceeded bool
}

type Producer interface {
	Reset()
	NextValue() (DataPair)
	CurrentValue(DataPair)
	IsCyclic() (bool)
	IsRandom() (bool)
	doStep() (interface{})
	doRandom() (interface{})
	isBoundExceeded (bool)
}

type Presentation struct{
	NullProbability  int8
	NullPresentation string
	Format string
}

func newDataPair(rawValue interface{}, presentation Presentation) (*DataPair) {
	if rawValue == nil {
		panic("Raw Value is not defined")
	}
	result := new(DataPair)
	result.RawValue = rawValue;
	if str ,ok := result.RawValue.(string); ok {
		result.StringValue = str;
	} else {
		var format string
		if (presentation == nil || presentation.Format == "") {
			format = "%v"
		} else {
			format = presentation.Format
		}
		result.StringValue = format
	}
	return result;
}