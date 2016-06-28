package SD

type Multibit struct {
	Presentation
	NullProbability  int8
	bits[] *Producer
	temp[] *DataPair
	current[] *DataPair

}



func( multibit *Multibit) Reset() {
	for _,bit := range multibit.bits {
		(*bit).Reset();
	}
	multibit.temp = nil;

}

func(multibit Multibit) IsRandom() (bool) {
	for _,bit := range multibit.bits {
		if  (*bit).IsRandom() {
			return true;
		}
	}
	return false;
}

func(multibit Multibit) IsCyclic() (bool) {
	result := true;
	for _, bit := range multibit.bits {
		if  result = result && (*bit).IsCyclic(); !result {
			break;
		}
	}
	return result;
}
func(multibit *Multibit) NextValue() (result DataPair) {
  return DataPair{};
}

func(multibit *Multibit) getBitValue(index int) (result[] *DataPair) {
	var bitLen int = len(multibit.bits)

	 if index == -1 {
		 if multibit.temp == nil {
			 multibit.temp = make([]*DataPair, bitLen)
			 multibit.temp[0] = NewBoundExceeded()
		 }
	 } else {
		 if multibit.temp == nil {
			 multibit.temp = make([]*DataPair, bitLen )

			for bitIndex, _ := range multibit.bits {
				reverseIndex := bitLen - 1 - bitIndex;
				multibit.temp[reverseIndex] =
					(*multibit.bits[reverseIndex]).NextValue();
			}
		} else {
			currentProducer :=  (*multibit.bits[index]);
			multibit.temp[index] = currentProducer.NextValue() ;
			if multibit.temp[index].BoundExceeded {
				currentProducer.Reset()
				multibit.temp[index] = currentProducer.NextValue();
				result = multibit.getBitValue(index - 1)
			}
		}
 	}
	copy(result, multibit.temp);
	return result;
}

