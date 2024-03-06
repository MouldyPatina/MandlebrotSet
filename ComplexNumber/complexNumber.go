package ComplexNumber

import (
	"strconv"
	"math"
)

func (start *ComplexNum) BlowUp(bound float64, numTries int) int {
	z := ComplexNum{};
	for i := 0; i < numTries ; i++ {
		z.Multiply(&z).Add(start);
		if z.Abs() > bound {
			return i;
		}
	}
	return -1;
}

type ComplexNum struct {
	Real float64
	Imagine float64
}

func (num *ComplexNum) Add(num2 *ComplexNum) *ComplexNum {
	num.Real += num2.Real;
	num.Imagine += num2.Imagine;
	return num;
}

func (num *ComplexNum) Multiply(num2 *ComplexNum) *ComplexNum {
	RealPart := num.Real * num2.Real - num.Imagine * num2.Imagine;
	ImaginePart := num.Imagine * num2.Real + num.Real * num2.Imagine;
	num.Real = RealPart;
	num.Imagine = ImaginePart;
	return num;
}

func (num *ComplexNum) toString() string {
	str := strconv.FormatFloat(num.Real, 'g', 5, 64);
	if num.Imagine < 0 {
		str += " - " + strconv.FormatFloat(-1 * num.Imagine, 'g', 5, 64);
	} else {
		str += " + " + strconv.FormatFloat(num.Imagine, 'g', 5, 64);
	}
	
	return str + "i";
}

func (num *ComplexNum) Abs() float64 {
	return math.Sqrt(num.Real * num.Real + num.Imagine * num.Imagine);
}
