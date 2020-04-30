package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"
)

func main() {
	filePattern := "/sys/devices/platform/system76/leds/system76::kbd_backlight/color*"
	matches, err := filepath.Glob(filePattern)
	if err != nil {
		fmt.Println(err)
	}

	freq := 0.3

	for i := 0; i < 64; i++ {
		R := math.Sin(freq*float64(i)+0)*127 + 128
		G := math.Sin(freq*float64(i)+2*math.Pi/3)*127 + 128
		B := math.Sin(freq*float64(i)+4*math.Pi/3)*127 + 128

		output := fmt.Sprintf("%s%s%s", fmt.Sprintf("%X", int(R)), fmt.Sprintf("%X", int(G)), fmt.Sprintf("%X", int(B)))

		for _, match := range matches {
			file, err := os.Create(match)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			file.WriteString(output)
		}
		fmt.Println(output)
		time.Sleep(10 * time.Millisecond)

		if i == 63 {
			i = 0
		}
	}

}
