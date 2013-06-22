package bitstamp

import (
	"encoding/json"
)

func (o *Order) UnmarshalJSON(bytes []byte) error {
	values := &[2]string{}
	err := json.Unmarshal(bytes, values)

	if err != nil {
		panic(err.Error())
	}

	o.Price = values[0]
	o.Volume = values[1]

	return err
}
