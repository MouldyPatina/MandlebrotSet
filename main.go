package main

import (
	"fmt"
	"os"
	"image"
	"image/color"
	"image/png"
	cn "MandlebrotSet/ComplexNumber"
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
			z := cn.ComplexNum{Real: -width + (width * float64(x) * 3) / (2 * float64(widthSteps)), Imagine: height - (height * 2 * float64(y)) / float64(heightSteps)}; 
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
