package main

import (
	"fmt"
	"helpers"
//	"strings"
	"strconv"
)

func main(){
	user_deposit := 0
	coke_bottle_price := 50
	accepted_denominations := []int{5,10,25}

	for user_deposit  < coke_bottle_price {
	        user_input := helpers.Input("Insert Coin: ")
        	amount_inserted,_ := strconv.Atoi(user_input)
		if contains(accepted_denominations,amount_inserted){
			user_deposit = user_deposit+amount_inserted
		}
		if user_deposit < coke_bottle_price{
			amount_due := coke_bottle_price-user_deposit
			fmt.Println("Amount Due: ",amount_due)
		}
	}
	fmt.Println("Change owed: ",user_deposit-coke_bottle_price)
}

func contains(list []int,element int) bool{
	for _,e := range list{
		if e == element{
			return true
		}
	}
	return false
}
