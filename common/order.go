package common

type CommonOrder struct {
	exchange Exchange
	symbol   Symbol
	volume   Volume
	price    Price
	fee      Fee
	side     Side
}

func NewOrder(e Exchange, s Symbol, v Volume, p Price, f Fee, side Side) *CommonOrder {
	return &CommonOrder{e, s, v, p, f, side}
}

func (o *CommonOrder) GetExchange() Exchange { return o.exchange }
func (o *CommonOrder) GetSymbol() Symbol     { return o.symbol }
func (o *CommonOrder) GetPrice() Price       { return o.price }
func (o *CommonOrder) GetFee() Fee           { return o.fee }
func (o *CommonOrder) GetVolume() Volume     { return o.volume }
func (o *CommonOrder) GetSide() Side         { return o.side }
