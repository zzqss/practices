package bignumber

type BigNumber struct {
	positive bool
	number   []byte
}

func NewBigNumber(number int) *BigNumber {
	bigNumber := &BigNumber{}
	bigNumber.number = make([]byte, 0)
	bigNumber.positive = true
	for {
		if number == 0 {
			break
		}
		leftNumber := number % 10
		number = number / 10
		bigNumber.number = append(bigNumber.number, byte(leftNumber))
	}
	return bigNumber
}
func NewBigNumberBySlice(positive bool, slice []byte) *BigNumber {
	newSlice := make([]byte, len(slice))
	for index, num := range slice {
		newSlice[index] = num
	}
	bigNumber := &BigNumber{positive, newSlice}
	bigNumber.OptimizeLength()
	return bigNumber
}

func (this *BigNumber) Add(number *BigNumber) {
	if this.positive == number.positive {
		this.SliceAdd(number.number)
	} else {
		if !this.SliceBiggerThan(number.number) {
			this.positive = number.positive
		}
		this.SliceSubtract(number.number)
	}
}
func (this *BigNumber) Subtract(number *BigNumber) {
	if this.positive != number.positive {
		this.SliceAdd(number.number)
	} else {
		if this.SliceLesserThan(number.number) {
			this.positive = !number.positive
		}
		this.SliceSubtract(number.number)
	}
}
func (this *BigNumber) Multiply(number *BigNumber) {
	this.positive = this.positive == number.positive
	this.SliceMultiply(number.number)
}
func (this *BigNumber) Divide(number *BigNumber) {
	this.positive = this.positive == number.positive
	this.SliceDivide(number.number)
}

func (this *BigNumber) Pow(number int) {
	tempThis := this.Copy()
	for i := 1; i < number; i++ {
		this.Multiply(tempThis)
	}
}

func (this *BigNumber) SliceDivide(number []byte) {
	if !this.SliceBiggerThan(number) {
		this.number = []byte{0}
		return
	}
	// result must bigger than 1
	// base divide step
	//
	resultByteSlice := make([]byte, 0)
	tempThis := this.Copy()
	tempThis.positive = true
	tempNumber := NewBigNumberBySlice(true, number)
	itLength := len(tempNumber.number)
	for {
		originLength := len(tempThis.number)
		if originLength < itLength {
			break
		}
		copyLength := itLength
		copySlice := CopySlice(tempThis.number, copyLength, false)
		tempForSubtract := NewBigNumberBySlice(true, copySlice)
		// tempForSubtract should not top by 0
		if tempForSubtract.SliceLesserThan(tempNumber.number) {
			resultByteSlice = append(resultByteSlice, byte(0))
			copyLength += 1
			if copyLength > len(tempNumber.number) {
				break
			}
			copySlice = CopySlice(tempThis.number, copyLength, false)
			tempForSubtract = NewBigNumberBySlice(true, copySlice)
		}
		subtractNum := byte(0)
		for {
			tempForSubtract.Subtract(tempNumber)
			subtractNum += byte(1)
			if len(tempForSubtract.number) == 1 && tempForSubtract.number[0] == byte(0) {
				break
			}
			if !tempForSubtract.positive {
				subtractNum -= byte(1)
				tempForSubtract.Add(tempNumber)
				break
			}
		}

		tempForSubtractLength := len(tempForSubtract.number)
		leftLength := originLength - copyLength + tempForSubtractLength
		tempThis.number = tempThis.number[0:leftLength]
		for k := 0; k < tempForSubtractLength; k++ {
			tempThis.number[leftLength+k-tempForSubtractLength] = tempForSubtract.number[k]
		}
		resultByteSlice = append(resultByteSlice, subtractNum)
		supposeEnd := originLength + itLength - copyLength - 1
		tempThis.OptimizeLength()
		for i := len(tempThis.number); i < supposeEnd; i++ {
			resultByteSlice = append(resultByteSlice, byte(0))
		}
	}
	resultLength := len(resultByteSlice)
	this.number = make([]byte, resultLength)
	for i := 0; i < resultLength; i++ {
		this.number[i] = resultByteSlice[resultLength-i-1]
	}
	this.OptimizeLength()
}

// just multiply byte slice
func (this *BigNumber) SliceMultiply(number []byte) {
	itLength := len(number)
	soloMultiplySlice := make([][]byte, itLength, itLength)
	myLength := len(this.number)
	for i := 0; i < itLength; i++ {
		soloMultiply := make([]byte, i+myLength, i+myLength+2)
		soloMultiplySlice[i] = soloMultiply
		j := 0
		for ; j < i; j++ {
			soloMultiply[j] = 0
		}
		var tempAdd byte = 0
		byteForMultiply := number[i]
		k := 0
		for ; k < myLength; k++ {
			multiplyResult := this.number[k]*byteForMultiply + tempAdd
			if multiplyResult > 10 {
				tempAdd = multiplyResult / 10
				multiplyResult = multiplyResult % 10
			} else {
				tempAdd = 0
			}
			soloMultiply[j+k] = multiplyResult
		}
		if tempAdd > 0 {
			soloMultiply[j+k] = tempAdd
		}
	}
	this.number = soloMultiplySlice[0]
	for i := 1; i < itLength; i++ {
		this.SliceAdd(soloMultiplySlice[i])
	}
}

// just add byte slice
func (this *BigNumber) SliceAdd(number []byte) {
	addedLength := len(number)
	myLength := len(this.number)
	if addedLength > myLength {
		this.number = append(this.number, make([]byte, addedLength-myLength)...)
	}
	maxLength := myLength
	if addedLength > maxLength {
		maxLength = addedLength
	}
	var tempAdd byte = 0
	i := 0
	for ; i < maxLength; i++ {
		addResult := tempAdd
		if i < myLength {
			addResult += this.number[i]
		}
		if i < addedLength {
			addResult += number[i]
		}
		if addResult > 10 {
			tempAdd = 1
			addResult -= 10
		} else {
			tempAdd = 0
		}
		this.number[i] = addResult
	}
	if tempAdd > 0 {
		this.number = append(this.number, tempAdd)
	}
}

// just subtract byte slice
func (this *BigNumber) SliceSubtract(number []byte) {
	var largeNumber []byte
	var lowerNumber []byte
	if this.SliceBiggerThan(number) {
		largeNumber = this.number
		lowerNumber = number
	} else {
		largeNumber = make([]byte, len(number))
		for index, num := range number {
			largeNumber[index] = num
		}
		lowerNumber = this.number
		this.number = largeNumber
	}
	largeLength := len(largeNumber)
	lowerLength := len(lowerNumber)
	var tempSubtract byte = 0
	for i := 0; i < largeLength; i++ {
		subtractResult := byte(0)
		subtractNumber := tempSubtract
		if i < lowerLength {
			subtractNumber += lowerNumber[i]
		} else if subtractNumber == 0 {
			break
		}
		if largeNumber[i] < subtractNumber {
			tempSubtract = 1
			subtractResult = byte(10) + largeNumber[i] - subtractNumber
		} else {
			tempSubtract = 0
			subtractResult = largeNumber[i] - subtractNumber
		}
		largeNumber[i] = subtractResult
	}
	this.OptimizeLength()
}

// if top number of slice is 0,reduce the length
func (this *BigNumber) OptimizeLength() int {
	resultLength := len(this.number)
	optimizeLength := 0
	for i := resultLength - 1; i > 0; i-- {
		if this.number[i] > 0 {
			break
		}
		optimizeLength += 1
		this.number = this.number[0:i]
	}
	return optimizeLength
}

// just compare slice
func (this *BigNumber) SliceBiggerThan(number []byte) bool {
	myLength := len(this.number)
	itLength := len(number)
	if myLength > itLength {
		return true
	} else if myLength < itLength {
		return false
	} else {
		for i := myLength - 1; i >= 0; i-- {
			if this.number[i] == number[i] {
				continue
			}
			if this.number[i] > number[i] {
				return true
			} else {
				return false
			}
		}
	}
	return false
}

// just compare slice
func (this *BigNumber) SliceLesserThan(number []byte) bool {
	myLength := len(this.number)
	itLength := len(number)
	if myLength > itLength {
		return false
	} else if myLength < itLength {
		return true
	} else {
		for i := myLength - 1; i >= 0; i-- {
			if this.number[i] == number[i] {
				continue
			}
			if this.number[i] > number[i] {
				return false
			} else {
				return true
			}
		}
	}
	return false
}

func (this *BigNumber) Copy() *BigNumber {
	bigNumber := &BigNumber{}
	bigNumber.positive = this.positive
	bigNumber.number = make([]byte, len(this.number))
	for index, num := range this.number {
		bigNumber.number[index] = num
	}
	return bigNumber
}

func CopySlice(from []byte, num int, positive bool) []byte {
	result := make([]byte, num)
	if positive {
		for i := 0; i < num; i++ {
			result[i] = from[i]
		}
	} else {
		fromLength := len(from)
		for i := 0; i < num; i++ {
			result[num-1-i] = from[fromLength-1-i]
		}
	}
	return result
}
