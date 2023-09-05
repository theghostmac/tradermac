package core

type Money float64

type Tick struct {
    Date string
    TimeStamp int64
    Price Money
    TradedVolume Money 
    CounterParty1 any
    CounterParty2 any
}

type Order struct {
    Size Money
    Bid bool
    Limit *Limit
    TimeStamp int64
}

type Limit struct {
    Price Money
    Orders Orders
    TotalVolume Money
}

type Orders []*Order
