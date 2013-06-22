package mtgox

import (
    "github.com/mattomatic/go-heap/heap"
)

type OrderBook struct {
    bids *heap.Heap
    asks *heap.Heap
}

func NewOrderBook() *OrderBook {
    return &OrderBook{
        bids: heap.NewMaxHeap(),
        asks: heap.NewMinHeap(),
    }
}


    
