package SD

import (
	"math/rand"
	"fmt"
	"time"
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
type Presentation struct{
	NullProbability  int8
	NullPresentation string
	Format string
}

type Simple struct {
	Presentation
	lowerBound interface{}
	upperBound interface{}
	initial interface{}
	randomTypeValue RandomType
	sequentialStep interface{}
	currentValue interface{}
	cyclic bool
	//
	doStepHandler func()
	doRandomHandler func()
	isBoundExceededHandler func() bool

	getLowerBoundFunc func() (interface{})
	getUpperBoundFunc func() (interface{})
}


func (simple *Simple) Reset() {
	simple.currentValue = nil;
}
func(simple Simple) IsRandom() (bool) {
	return simple.randomTypeValue != NONE;
}

func(simple Simple) IsCyclic() (bool) {
	return simple.cyclic;
}

func (simple *Simple) doStep()  {
	if simple.doStepHandler == nil {
		panic("doStep handler has been defined")
	} else {
		simple.doStepHandler()
	}
}
func (simple *Simple) doRandom() {

	if simple.doRandomHandler == nil {
		panic("doRandom handler has been defined")
	} else {
		simple.doRandomHandler()
	}
}
func (simple Simple) isBoundExceeded() (bool){
	if simple.isBoundExceededHandler == nil {
		panic("isBoundExceeded handler has been defined")
	} else {
		return  simple.isBoundExceededHandler()
	}
}
func (simple Simple) GetLowerBound() interface {} {
	if simple.getLowerBoundFunc == nil {
		return  simple.lowerBound
	} else {
		return  simple.getLowerBoundFunc()
	}
}
func (simple Simple) GetUpperBound() interface {} {
	if simple.getUpperBoundFunc == nil {
		return  simple.upperBound
	} else {
		return  simple.getUpperBoundFunc()
	}
}


func (simple *Simple) NextValue() (DataPair) {
	var result DataPair;
	if simple.IsRandom() {
		simple.doRandom();
	} else {
		if simple.currentValue == nil {
			simple.currentValue = simple.GetLowerBound()
		} else {
			simple.doStep();
		}


		if  !simple.isBoundExceeded() {
			result.RawValue = simple.currentValue;
		} else {
			if simple.IsCyclic() {
				simple.Reset()
				result = simple.NextValue();
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
			if(simple.Format == "") {simple.Format = "%v"}
			result.StringValue = fmt.Sprintf(simple.Format, result.RawValue)
		}
	}
	return result;
}

func (simple *Simple) SetCyclic(value bool) *Simple {
	simple.sequentialStep = value
	return simple
}

func (simple *Simple) SetRandomType(value RandomType) *Simple {
	simple.randomTypeValue = value
	return simple
}


type SimpleInt64 struct {
	Simple
}

func (simpleInt64 *SimpleInt64) doStep()  {
	current := simpleInt64.currentValue.(int64);
	step := simpleInt64.sequentialStep.(int64);
	simpleInt64.currentValue = current + step;
}

func (simpleInt64 *SimpleInt64) doRandom()   {
	lb := simpleInt64.GetLowerBound().(int64)
	ub := simpleInt64.GetUpperBound().(int64)
	simpleInt64.currentValue = lb + rand.Int63n(ub - lb)
}

func (simpleInt64 *SimpleInt64) isBoundExceeded() (bool)  {
	step := simpleInt64.sequentialStep.(int64)
	upperBound := simpleInt64.GetUpperBound().(int64)
	current := simpleInt64.currentValue.(int64)

	return (step > 0 && current > upperBound ) ||
		(step < 0 && current > upperBound);
}


func NewSimpleInt64() (*SimpleInt64)  {
	result := new(SimpleInt64);
	result.lowerBound = 1
	result.upperBound = 10
	result.initial = 1;
	result.sequentialStep = 1;
	result.randomTypeValue = NONE
	result.isBoundExceededHandler = result.isBoundExceeded
	result.doStepHandler = result.doStep
	result.doRandomHandler = result.doRandom
	return result;
}
func (simpleInt64 *SimpleInt64) SetLowerBound(value int64) *SimpleInt64 {
	simpleInt64.lowerBound = value
	return simpleInt64
}

func (simpleInt64 *SimpleInt64) SetUpperBound(value int64) *SimpleInt64 {
	simpleInt64.upperBound = value
	return simpleInt64
}

func (simpleInt64 *SimpleInt64) SetInitial(value int64) *SimpleInt64 {
	simpleInt64.initial = value
	return simpleInt64
}
func (simpleInt64 *SimpleInt64) SetSequentialStep(value int64) *SimpleInt64 {
	simpleInt64.sequentialStep = value
	return simpleInt64
}





type SimpleFloat64 struct {
	Simple
}

func (simpleFloat64 *SimpleFloat64) doStep()  {
	current := simpleFloat64.currentValue.(float64);
	step := simpleFloat64.sequentialStep.(float64);
	simpleFloat64.currentValue = current + step;
}

func (simpleFloat64 *SimpleFloat64) doRandom()   {
	lb := simpleFloat64.GetLowerBound().(float64)
	ub := simpleFloat64.GetUpperBound().(float64)
	simpleFloat64.currentValue = lb + rand.Float64()*(ub - lb)
}

func (simpleFloat64 *SimpleFloat64) isBoundExceeded() (bool)  {
	step := simpleFloat64.sequentialStep.(float64)
	upperBound := simpleFloat64.GetUpperBound().(float64)
	current := simpleFloat64.currentValue.(float64)

	return (step > 0 && current > upperBound ) ||
	(step < 0 && current > upperBound);
}


func NewSimpleFloat64() (*SimpleFloat64)  {
	result := new(SimpleFloat64);
	result.lowerBound = 1.0
	result.upperBound = 10.0
	result.initial = 1.0;
	result.sequentialStep = 1.0;
	result.randomTypeValue = NONE
	result.isBoundExceededHandler = result.isBoundExceeded
	result.doStepHandler = result.doStep
	result.doRandomHandler = result.doRandom
	return result;
}
func (simpleFloat64 *SimpleFloat64) SetLowerBound(value float64) *SimpleFloat64 {
	simpleFloat64.lowerBound = value
	return simpleFloat64
}

func (simpleFloat64 *SimpleFloat64) SetUpperBound(value float64) *SimpleFloat64 {
	simpleFloat64.upperBound = value
	return simpleFloat64
}

func (simpleFloat64 *SimpleFloat64) SetInitial(value float64) *SimpleFloat64 {
	simpleFloat64.initial = value
	return simpleFloat64
}
func (simpleFloat64 *SimpleFloat64) SetSequentialStep(value float64) *SimpleFloat64 {
	simpleFloat64.sequentialStep = value
	return simpleFloat64
}




type SimpleTime struct {
	Simple
}

func (simpleTime *SimpleTime) doStep()  {
	current := simpleTime.currentValue.(time.Time);
	step := simpleTime.sequentialStep.(time.Duration);
	simpleTime.currentValue = current.Add(step);
}

func (simpleTime *SimpleTime) doRandom()   {
	lb := simpleTime.GetLowerBound().(time.Time)
	ub := simpleTime.GetUpperBound().(time.Time)
	simpleTime.currentValue = lb.Add(time.Duration(rand.Int63n(int64(ub.Sub(lb)))))
}

func (simpleTime *SimpleTime) isBoundExceeded() (bool)  {
	step := simpleTime.sequentialStep.(time.Duration)
	upperBound := simpleTime.GetUpperBound().(time.Time)
	current := simpleTime.currentValue.(time.Time)

	return (step > 0 && current.After(upperBound) ) ||
	(step < 0 && current.Before(upperBound));
}


func NewSimpleTime() (*SimpleTime)  {
	result := new(SimpleTime);
	result.lowerBound = time.Date(1900,01,01,00,00,00,00,time.UTC)
	result.upperBound = time.Date(2099,12,31,23,59,59,99,time.UTC)
	result.initial = time.Date(1900,01,01,00,00,00,00,time.UTC)
	result.sequentialStep = time.Duration(time.Second)*60*60*24
	result.randomTypeValue = NONE
	result.isBoundExceededHandler = result.isBoundExceeded
	result.doStepHandler = result.doStep
	result.doRandomHandler = result.doRandom
	return result;
}
func (simpleTime *SimpleTime) SetLowerBound(value time.Time) *SimpleTime {
	simpleTime.lowerBound = value
	return simpleTime
}

func (simpleTime *SimpleTime) SetUpperBound(value time.Time) *SimpleTime {
	simpleTime.upperBound = value
	return simpleTime
}

func (simpleTime *SimpleTime) SetInitial(value time.Time) *SimpleTime {
	simpleTime.initial = value
	return simpleTime
}
func (simpleTime *SimpleTime) SetSequentialStep(value time.Time) *SimpleTime {
	simpleTime.sequentialStep = value
	return simpleTime
}



