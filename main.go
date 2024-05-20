package main

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

const backlightPath = "/sys/class/backlight/intel_backlight/"

func main() {

	content, err := os.ReadFile(backlightPath + "max_brightness")
	if err != nil {
		fmt.Println("Error reading max brightness: ", err)
	}
	maxBrightness := strconv.ParseInt(string(content))

	content, err = os.ReadFile(backlightPath + "actual_brightness")
	if err != nil {
		fmt.Println("Error reading actual brightness: ", err)
	}
	currentBrightness := string(content)

	fmt.Println("Your brightness is currently set at", currentBrightness)
	fmt.Println("Enter brightness level between 0 -", maxBrightness)

	var input string

	_, err = fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Println("Error reading input", err)
	}

	err = os.WriteFile(backlightPath+"brightness", []byte(input), fs.ModeAppend)
	if err != nil {
		fmt.Println("Error setting brightness", err)
	}

	fmt.Println("Set brightness to", input)
}
