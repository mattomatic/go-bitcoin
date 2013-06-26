package common

type CommonOrder struct {
	exchange Exchange
	symbol   Symbol
	volume   Volume
	price    Price
	side     Side
}

func NewOrder(e Exchange, s Symbol, v Volume, p Price, side Side) *CommonOrder {
	return &CommonOrder{e, s, v, p, side}
}

func (o *CommonOrder) GetExchange() Exchange { return o.exchange }
func (o *CommonOrder) GetSymbol() Symbol     { return o.symbol }
func (o *CommonOrder) GetPrice() Price       { return o.price }
func (o *CommonOrder) GetVolume() Volume     { return o.volume }
func (o *CommonOrder) GetSide() Side         { return o.side }
