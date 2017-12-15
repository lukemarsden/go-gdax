package gdax

type Message struct {
	Type          string           `json:"type"`
	ProductId     string           `json:"product_id"`
	ProductIds    []string         `json:"product_ids"`
	TradeId       int              `json:"trade_id,number"`
	OrderId       string           `json:"order_id"`
	Sequence      int64            `json:"sequence,number"`
	MakerOrderId  string           `json:"maker_order_id"`
	TakerOrderId  string           `json:"taker_order_id"`
	Time          Time             `json:"time,string"`
	RemainingSize string           `json:"remaining_size,string"`
	NewSize       string           `json:"new_size,string"`
	OldSize       string           `json:"old_size,string"`
	Size          string           `json:"size,string"`
	Price         string           `json:"price,string"`
	Side          string           `json:"side"`
	Reason        string           `json:"reason"`
	OrderType     string           `json:"order_type"`
	Funds         string           `json:"funds,string"`
	NewFunds      string           `json:"new_funds,string"`
	OldFunds      string           `json:"old_funds,string"`
	Message       string           `json:"message"`
	Bids          [][]string       `json:"bids,omitempty"`
	Asks          [][]string       `json:"asks,omitempty"`
	Changes       [][]string       `json:"changes,omitempty"`
	LastSize      string           `json:"last_size,string"`
	BestBid       string           `json:"best_bid,string"`
	BestAsk       string           `json:"best_ask,string"`
	Channels      []MessageChannel `json:"channels"`
}

type MessageChannel struct {
	Name       string   `json:"name"`
	ProductIds []string `json:"product_ids"`
}
