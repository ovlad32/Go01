package SD

import "fmt"

type RandomType int

const (
	NONE      RandomType = iota
	UNIFORM
	GAUSSIAN
)



type DataPair struct{
	Presentation
	formatted bool
	RawValue interface{}
	stringValue string
	BoundExceeded bool
}

type Producer interface {
	Reset()
	NextValue() (*DataPair)
	CurrentValue(*DataPair)
	IsCyclic() (bool)
	IsRandom() (bool)
	doStep() (interface{})
	doRandom() (interface{})
	isBoundExceeded (bool)
}

type Presentation struct{
	NullPresentation string
	Format string
}

func newDataPair(rawValue interface{}) (*DataPair) {
	/*if rawValue == nil {
		panic("Raw Value is not defined")
	}*/
	result := new(DataPair)
	result.RawValue = rawValue;
	return result;
}

func NewBoundExceeded() (*DataPair) {
	result := new(DataPair)
	result.BoundExceeded = true;
	return result
}

func(dataPair *DataPair) SetPresentation(presentation Presentation) (*DataPair){
	dataPair.Presentation.Format = presentation.Format
	dataPair.NullPresentation = presentation.NullPresentation
	return dataPair;
}

func (dataPair *DataPair) String()  string {
	if !dataPair.formatted {
		if sValue ,ok := dataPair.RawValue.(string); ok {
			dataPair.stringValue = sValue;
		} else if dataPair.Format != "" {
			if dataPair.RawValue == nil {
				dataPair.stringValue = dataPair.NullPresentation
			} else {
				dataPair.stringValue = fmt.Sprintf(
					dataPair.Format,
					dataPair.RawValue,
				)
			}
		}
		dataPair.formatted = true;
	}
	return dataPair.stringValue;
}

