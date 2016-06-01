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

type bound struct {
	Nulling
	Formatting
	Randomizing
	cyclic bool
	lowerBound interface{}
	upperBound interface{}
	initial interface{}
	step interface{}
	current interface{}
}

func(bound *bound) Reset() {
	bound.current = nil;
}
func(bound bound) IsRandom() (bool) {
	return bound.RandomTypeValue != NONE;
}
func(bound bound) IsCyclic() (bool) {
	return bound.cyclic;
}
func (bound *bound) doStep() (interface{}) {
	panic("Abstract doStep has been called")
}
func (bound *bound) doRandom() (interface{}) {
	panic("Abstract doRandom has been called")
}
func (bound *bound) isBoundExceeded() (bool){
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
	bound
}


func (boundInt64 *BoundInt64) doStep() (interface{})  {
	current := boundInt64.current.(int64);
	step := boundInt64.step.(int64);
	return current + step;
}

func (boundInt64 *BoundInt64) doRandom() (interface{})  {
	lb := boundInt64.lowerBound.(int64)
	ub := boundInt64.upperBound.(int64)

	return lb + rand.Int63n(ub - lb)
}

func (boundInt64 *BoundInt64) isBoundExceeded() (bool)  {
	step := boundInt64.step.(int64)
	upperBound := boundInt64.upperBound.(int64)
	current := boundInt64.current.(int64)

	return (step > 0 && current > upperBound ) ||
		(step < 0 && current > upperBound);
}

func NewBoundInt64Sequential(lowerBound, upperBound, step, initial int64, cyclic bool) (error,*BoundInt64)  {
	//TODO: check parameters

	result := new(BoundInt64);
	result.lowerBound = lowerBound
	result.upperBound = upperBound;
	result.initial = initial;
	result.step = step;
	result.Formatting = NewFormatting("%v");
	result.Nulling = NewNulling(0);
	result.cyclic = cyclic;
	result.RandomTypeValue = NONE

	return nil,result;

}
/*

func NewBoundInt64Random(lowerBound, upperBound int64, randomizing Randomizing) (error, *BoundInt64)  {
	//TODO: check parameters
	return &BoundInt64{
		lowerBound : lowerBound,
		upperBound : upperBound,
		initial : nil,
		step : nil,
		Formatting : NewFormatting("%v"),
		Nulling : NewNulling(0),
		cyclic : false,
		RandomTypeValue: randomizing.RandomTypeValue,
	}

}*/