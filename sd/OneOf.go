package SD

type OneOf struct {
	Presentation
	pool *[]interface{}
	dispatcher Simple
}

func NewOneOf(pool *[]interface{}) OneOf {
	oneOf := new(OneOf)
	oneOf.pool = pool;
	oneOf.dispatcher.getLowerBoundFunc = func() {
		return 0;
	}
	oneOf.dispatcher.getUpperBoundFunc = func() {
		return len(&oneOf.pool)
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
	index := oneOf.dispatcher.NextValue().RawValue.(int64);
	result := newDataPair(oneOf.pool[index],oneOf.Presentation);
	return result;
}
