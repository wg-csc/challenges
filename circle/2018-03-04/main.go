package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

const (
	horizMultiplier = 2
)

func main() {
	var radius int
	_, err := fmt.Scanf("%d", &radius)
	if err != nil {
		log.Fatal(err)
	}

	for lineNum := 0; lineNum < radius*2; lineNum++ {
		verticalDist := math.Abs(float64(lineNum - radius))
		if verticalDist == 0 {
			middlePadding := strings.Repeat(" ", horizMultiplier*2*radius)
			fmt.Println("*", middlePadding, "*")
			continue
		}
		angle := math.Acos(verticalDist / float64(radius))
		horizDist := math.Abs(verticalDist * math.Tan(angle))
		// fmt.Println(float64(radius) / verticalDist)
		leftPadding := strings.Repeat(" ", int(horizMultiplier*(float64(radius)-horizDist)))
		middlePadding := strings.Repeat(" ", int(horizMultiplier*2*horizDist))
		fmt.Println(leftPadding, "*", middlePadding, "*")
	}
}
