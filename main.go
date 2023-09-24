package main

import (
	"log"
	"time"

	"github.com/SimonWaldherr/ws2812/pixarray"
)

func main() {
	var leds pixarray.LEDStrip
	var err error
	pixels := 149
	ws281xPin0 := 18
	order := pixarray.GRB
	var ws281xFreq uint = 800000
	ws281xDma := 10

	leds, err = pixarray.NewWS281x(pixels, 3, order, ws281xFreq, ws281xDma, []int{ws281xPin0})
	if err != nil {
		log.Fatalf("failed creating WS281x: %v", err)
	}

	// clear old values
	pa := pixarray.NewPixArray(pixels, 3, leds)
	pa.SetAll(pixarray.Pixel{})
	pa.Write()

	colorIndex := 0
	for {
		for i := 0; i < pixels; i++ {
			colorIndex++
			colorIndex = colorIndex % 3
			log.Print(colorIndex)

			red := 0
			green := 0
			blue := 0
			if colorIndex == 0 {
				red = 50
			}
			if colorIndex == 1 {
				green = 50
			}
			if colorIndex == 2 {
				blue = 50
			}

			pa.SetOne(i, pixarray.Pixel{R: red, G: green, B: blue})

		}
		err = pa.Write()
		if err != nil {
			log.Fatalf("failed to write: %v", err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
