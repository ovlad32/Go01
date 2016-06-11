package SD

type OneOf struct {
	Presentation
	pool *[]interface{}
	dispatcher *SimpleInt64
}

func NewOneOf(pool *[]interface{}) *OneOf {
	oneOf := new(OneOf)
	oneOf.pool = pool;
	oneOf.dispatcher = NewSimpleInt64().
				SetInitial(0).
				SetSequentialStep(1)
	oneOf.dispatcher.Format=""
	oneOf.dispatcher.getLowerBoundFunc = func()  (interface{}) {
		return 0
	}
	oneOf.dispatcher.getUpperBoundFunc = func() (interface{}) {
		return len(*(oneOf.pool))
	}
	return  oneOf
}
func (oneOf *OneOf) SetCyclic(cyclic bool) *OneOf {
	oneOf.dispatcher.SetCyclic(cyclic)
	return oneOf
}

func (oneOf *OneOf) SetRandom(randomType RandomType) *OneOf {
	oneOf.dispatcher.SetRandomType(randomType);
	return oneOf
}
func(oneOf *OneOf) NextValue() (*DataPair) {
	if oneOf.dispatcher.Format != "" {
		panic("Internal index producer has to have empty .Format!")
	}
	internal := oneOf.dispatcher.NextValue();
	if internal.BoundExceeded {
		internal.RawValue = nil
		internal.StringValue = ""
		return internal;
	} else {
		index := internal.RawValue.(int64);
		result := newDataPair((*oneOf.pool)[index], oneOf.Presentation);
		return result;
	}
}
