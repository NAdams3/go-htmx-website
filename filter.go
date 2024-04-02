package main

import (
    "fmt"
)

type receipt struct {
    UserID int
    Merchant string
}

func FilterTest() {
    receipts := []receipt{
        {UserID: 0, Merchant: "0"},
        {UserID: 1, Merchant: "1"},
        {UserID: 2, Merchant: "2"},
        {UserID: 0, Merchant: "3"},
        {UserID: 4, Merchant: "4"},
        {UserID: 5, Merchant: "5"},
        {UserID: 0, Merchant: "6"},
        {UserID: 7, Merchant: "7"},
        {UserID: 8, Merchant: "8"},
    }
    processReceipts(receipts)

}

func processReceipts( receipts []receipt ) {
    filteredReceipts := processReceipt(0, receipts, 0)
    fmt.Println(filteredReceipts)
}

func processReceipt(index int, receipts []receipt, userID int) []receipt {
    empty := make([]receipt, 0)

    if index >= len(receipts) {
        return empty 
    }

    filtered := processReceipt(index+1, receipts, userID)
    receipt := receipts[index]
    if receipt.UserID == userID {
        return append(filtered, receipt)
    }

    return empty 

}
