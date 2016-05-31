package SD

import (
	"math/rand"
	"fmt"
)

type RandomType int

const (
	NONE      RandomType = iota
	UNIFORM
	GAUSSIAN
)


type nulling struct{
	NullProbability int
	NullPresentation string
}
type formatting struct {
	Locale string
	format string
}

type randomizing struct {
	RandomTypeValue RandomType
}
type cycling struct  {
	cycle bool
}

type bound struct {
	nulling nulling
	formatting formatting
	randomizing randomizing
	cycling
	lowerBound interface{}
	upperBound interface{}
	initial interface{}
	step interface{}
	current interface{}
}

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
func(bound *bound) Reset() {
	bound.current = nil;
}
func(bound bound) isCyclic() (bool) {
	return bound.cycle == true;
}
func(bound bound) isRandom() (bool) {
	return bound.randomizing.RandomTypeValue == true;
}
func (bound bound) doStep() (interface{}) {
	panic("Abstract doStep has been called")
}
func (bound bound) doRandom() (interface{}) {
	panic("Abstract doRandom has been called")
}
func (bound bound) isBoundExceeded() {
	panic("Abstract isBoundExceeded has been called")
}

func(bound *bound) NextValue() (DataPair) {
	var result DataPair;
	if bound.isRandom() {
		bound.current = bound.doRandom();
	} else {
		if bound == nil {
			bound.current = bound.lowerBound
		} else {
			bound.current = bound.doStep();
		}

		if  bound.isBoundExceeded (){
			if bound.isCyclic() {
				bound.Reset()
				result = bound.NextValue();
			} else {
				result.BoundExceeded = true;
				result.RawValue = nil
			}
		}
	}

	if result.RawValue != nil {
		if _,ok := result.RawValue.(string); ok {
			result.StringValue = result.RawValue
		} else {
			result.StringValue = fmt.Sprintf(bound.formatting.format,result.RawValue)
		}
	}
	return result;
}
type BoundInt64 struct{
	bound bound;
}

func (boundInt64 BoundInt64) doStep() (interface{})  {
	return boundInt64.bound.current + boundInt64.bound.step
}
func (boundInt64 BoundInt64) doRandom() (interface{})  {
	return boundInt64.bound.lowerBound + rand.Int63n(boundInt64.bound.upperBound - boundInt64.bound.lowerBound)
}
func (boundInt64 BoundInt64) isBoundExceeded() (bool)  {
	return (boundInt64.bound.step > 0 && boundInt64.bound.current >boundInt64.bound.upperBound) ||
		(boundInt64.bound.step < 0 && boundInt64.bound.current >boundInt64.bound.upperBound);
}

