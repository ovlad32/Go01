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


type Nulling struct{
	probability float32
	presentation string
}
func NewNullingExt(probability float32, presentation string) (Nulling){
	//TODO: check parameters
	return Nulling{
		probability: probability,
		presentation: presentation,
	}
}
func NewNulling(probability float32) (Nulling){
	return NewNullingExt(probability,"");
}

type Formatting struct {
	format string
	locale string
}
func NewFormattingExt(format, locale string) (Formatting){
	//TODO: check parameters
	return Formatting{
		format: format,
		locale: locale,
	}
}
func NewFormatting(format string) (Formatting){
	return NewFormattingExt(format,"")
}

type Randomizing struct {
	RandomTypeValue RandomType
}

type bound struct {
	Nulling Nulling
	Formatting Formatting
	randomizing Randomizing
	cyclic bool
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
func(bound bound) IsRandom() (bool) {
	return bound.randomizing.RandomTypeValue != NONE;
}
func(bound bound) IsCyclic() (bool) {
	return bound.cyclic;
}
func (bound bound) doStep() (interface{}) {
	panic("Abstract doStep has been called")
}
func (bound bound) doRandom() (interface{}) {
	panic("Abstract doRandom has been called")
}
func (bound bound) isBoundExceeded() (bool){
	panic("Abstract isBoundExceeded has been called")
}

func(bound *bound) NextValue() (DataPair) {
	var result DataPair;
	if bound.IsRandom() {
		bound.current = bound.doRandom();
	} else {
		if bound == nil {
			bound.current = bound.lowerBound
		} else {
			bound.current = bound.doStep();
		}

		if  bound.isBoundExceeded() {
			if bound.IsCyclic() {
				bound.Reset()
				result = bound.NextValue();
			} else {
				result.BoundExceeded = true;
				result.RawValue = nil
			}
		}
	}

	if result.RawValue != nil {
		if str ,ok := result.RawValue.(string); ok {
			result.StringValue = str;
		} else {
			result.StringValue = fmt.Sprintf(bound.Formatting.format,result.RawValue)
		}
	}
	return result;
}
type BoundInt64 struct{
	bound bound;
}
type BoundInt64I interface {
	Producer
}

func (boundInt64 BoundInt64) doStep() (interface{})  {
	current := boundInt64.bound.current.(int64);
	step := boundInt64.bound.step.(int64);
	return current + step;
}
func (boundInt64 BoundInt64) doRandom() (interface{})  {
	return boundInt64.bound.lowerBound + rand.Int63n(boundInt64.bound.upperBound - boundInt64.bound.lowerBound)
}
func (boundInt64 BoundInt64) isBoundExceeded() (bool)  {
	return (boundInt64.bound.step > 0 && boundInt64.bound.current >boundInt64.bound.upperBound) ||
		(boundInt64.bound.step < 0 && boundInt64.bound.current >boundInt64.bound.upperBound);
}

func NewBoundInt64Sequential(lowerBound, upperBound, step, initialValue int64, cyclic bool) (error,*BoundInt64)  {
	//TODO: check parameters
	return &BoundInt64{
		bound.lowerBound : lowerBound,
		bound.upperBound : upperBound,
		bound.initial : initialValue,
		bound.step : step,
		bound.Formatting : NewFormatting("%v"),
		bound.Nulling : NewNulling(0),
		bound.cyclic : cyclic,
		bound.randomizing.RandomTypeValue : NONE,
	}

}

func NewBoundInt64Random(lowerBound, upperBound int64, randomizing Randomizing) (error, *BoundInt64)  {
	//TODO: check parameters
	return &BoundInt64{
		bound.lowerBound : lowerBound,
		bound.upperBound : upperBound,
		bound.initial : nil,
		bound.step : nil,
		bound.Formatting : NewFormatting("%v"),
		bound.Nulling : NewNulling(0),
		bound.cyclic : false,
		bound.randomizing : randomizing,
	}

}