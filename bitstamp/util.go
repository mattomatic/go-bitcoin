package bitstamp

import (
	"github.com/mattomatic/go-bitcoin/common"
	"strconv"
)

func getPrice(price string) common.Price {
	return common.Price(getFloat(price))
}

func getVolume(volume string) common.Volume {
	return common.Volume(getFloat(volume))
}

func getInt(amount string) int {
	x, err := strconv.Atoi(amount)

	if err != nil {
		panic(err.Error())
	}

	return x
}

func getFloat(amount string) float64 {
	x, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		panic(err.Error())
	}

	return x
}
