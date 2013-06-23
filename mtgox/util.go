package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
	"strconv"
)

const (
	Divisor = 1.0e5 // mtgox display divisor
)

func getPrice(price string) common.Price {
	return common.Price(getFloat(price))
}

func getVolume(volume string) common.Volume {
	return common.Volume(getFloat(volume))
}

func getFloat(amount string) float64 {
	x, err := strconv.Atoi(amount)

	if err != nil {
		panic(err.Error())
	}

	return float64(x) / Divisor
}
