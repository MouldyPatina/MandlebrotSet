package main

import (
	"fmt"
	"os"
	"image"
	"image/color"
	"image/png"
	"math"
	"strconv"
)

func main () {
	Picture(2560, 1440, 500);
	fmt.Println("see the image");
}

func Picture(widthSteps int, heightSteps int, accuracy int) {
	upLeft := image.Point{0, 0};
	lowRight := image.Point{widthSteps, heightSteps};
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight});
	var height float64 =  1.0;
	var width float64 =  2.1;
	for y := 0; y <= heightSteps; y++ {
		for x := 0; x <= widthSteps; x++ {
			z := complexNum{real: -width + (width * float64(x) * 3) / (2 * float64(widthSteps)), imagine: height - (height * 2 * float64(y)) / float64(heightSteps)}; 
			if blowUp := z.BlowUp(float64(accuracy), 5 * accuracy); blowUp == -1 {
				img.Set(x, y, color.Black);	
			} else {
				scale := (uint8) (100 - 100 / blowUp );
				pixelColor := color.RGBA{ 2*scale, scale, scale, 0xff};
				img.Set(x, y, pixelColor);
			}
		}
	}
	f, _ := os.Create("MandleBrot_Set.png");
	png.Encode(f, img);
}

func (start *complexNum) BlowUp(bound float64, numTries int) int {
	z := complexNum{};
	for i := 0; i < numTries ; i++ {
		z.Multiply(&z).Add(start);
		if z.Abs() > bound {
			return i;
		}
	}
	return -1;
}

type complexNum struct {
	real float64
	imagine float64
}

func (num *complexNum) Add(num2 *complexNum) *complexNum {
	num.real += num2.real;
	num.imagine += num2.imagine;
	return num;
}

func (num *complexNum) Multiply(num2 *complexNum) *complexNum {
	realPart := num.real * num2.real - num.imagine * num2.imagine;
	imaginePart := num.imagine * num2.real + num.real * num2.imagine;
	num.real = realPart;
	num.imagine = imaginePart;
	return num;
}

func (num *complexNum) toString() string {
	str := strconv.FormatFloat(num.real, 'g', 5, 64);
	if num.imagine < 0 {
		str += " - " + strconv.FormatFloat(-1 * num.imagine, 'g', 5, 64);
	} else {
		str += " + " + strconv.FormatFloat(num.imagine, 'g', 5, 64);
	}
	
	return str + "i";
}

func (num *complexNum) Abs() float64 {
	return math.Sqrt(num.real * num.real + num.imagine * num.imagine);
}
