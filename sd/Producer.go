package SD

import (
	"fmt"
	"math/rand"
)

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
	GetCurrentValue() (interface{})
	setCurrentValue( interface{})
	IsCyclic() (bool)
	IsRandom() (bool)
	initializeCurrentValue()
	doStep()
	doRandom()
	isBoundExceeded() (bool)
	getPresentation() (Presentation)
	getNullProbability() int8
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

func getNullOccurance(nullProbability int8) bool {
	if !(nullProbability>=0 && nullProbability<=100) {
		panic(fmt.Sprintf("Null probability is out of range %v. Has to be in range 0..100!",nullProbability));
	}
	if nullProbability == 0 {
		return false;
	} else if nullProbability == 100 {
		return true;
	} else if level := int8(rand.Int31n(100)); level < nullProbability {
		return true;
	}
	return true
}

func nextValue(prod Producer) (*DataPair) {
	var result *DataPair;
	var makeValue bool

	makeValue = !getNullOccurance(prod.getNullProbability())

	if prod.IsRandom() {
		if makeValue {
			prod.doRandom()
		} else {
			prod.setCurrentValue(nil)
		}
		result = newDataPair(
			prod.GetCurrentValue(),
		).SetPresentation(
			prod.getPresentation(),
		)
	} else {
		if prod.GetCurrentValue() == nil {
			prod.initializeCurrentValue()
		} else {
			prod.doStep();
		}
		if  !prod.isBoundExceeded() {
			if !makeValue {
				prod.setCurrentValue(nil)
			}
			result = newDataPair(
				prod.GetCurrentValue(),
			).SetPresentation(
				prod.getPresentation(),
			)
		} else {
			if prod.IsCyclic() {
				prod.Reset()
				result = prod.NextValue();
			} else {
				result = NewBoundExceeded();
			}
		}
	}
	return result;
}
