package main

import "log"

func main() {

	set := []int{-3, 0, 10, 23, 18, -5}
	smallest := 0

	for _, v := range set {

		if v > smallest {
			continue
		}

		smallest = v
	}

	log.Println("Smallest value is", smallest)

	smallestMin := set[0]
	for _, v := range set {

		smallestMin = min(smallestMin, v)
	}

	log.Println("Smallest Value is", smallestMin)

}
