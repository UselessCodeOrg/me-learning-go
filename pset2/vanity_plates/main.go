package main

import (
	"fmt"
	"helpers"
	"unicode"
	"reflect"
	"runtime"
)

func isInScopeChar(char rune) bool {
	if unicode.IsLetter(char) || unicode.IsDigit(char){
		return true
	}
	return false
}

func numInPlate(number_plate string) bool{
	for _,i := range number_plate{
		if unicode.IsDigit(i){
			return true
		}
	}
	return false
}

func checkScope(number_plate string) bool{
	for _,char := range number_plate{
		if isInScopeChar(char){
			continue
		} else{
			return false
		}
	}
	return true
}

func checkStartLetters(number_plate string) bool{
	result := false
	defer func(){
		if  r:= recover();r!=nil{
		result = false
		}
	}()
	if unicode.IsLetter(rune(number_plate[0])) && unicode.IsLetter(rune(number_plate[1])){
		result = true
	}

	return result
}

func checkNumberOfChar(number_plate string) bool{
	if len(number_plate) <= 6 && len(number_plate) >= 2{
		return true
	}else{
		return false
	}
}

func checkNumberPos(number_plate string) bool{
	f_chardetected := false
	if numInPlate(number_plate){
		for i := range len(number_plate){
			counter := len(number_plate)-i-1
			if unicode.IsLetter(rune(number_plate[counter])) && !f_chardetected{
				f_chardetected = true
			}
			if unicode.IsDigit(rune(number_plate[counter])) && f_chardetected{
				return false
			}
		}
	}
	return true
}

func checkFirstAppearingNum(number_plate string)bool{
	f_digitdetected := false
	if  numInPlate(number_plate){
		for _,c := range number_plate{
			if unicode.IsDigit(rune(c)) && !f_digitdetected{
				f_digitdetected = true
			}
			if f_digitdetected{
				if c == '0'{
					return false
				}
				return true
			}
		}
	}
	return true
}

func checkAll(list []bool) bool{
	for _,v := range list{
		if !v{
			return false
		}
	}
	return true
}

func main(){
	check_funcs := []func(string) bool{
		checkFirstAppearingNum,
		checkNumberPos,
		checkNumberOfChar,
		checkStartLetters,
		checkScope,
	}

	for {
		checks := []bool{}
		number_plate := helpers.Input("Plate: ")
		for _, fn := range check_funcs {
    			fmt.Println(runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(),fn(number_plate))
			checks = append(checks,fn(number_plate))
		}
		if checkAll(checks){
			fmt.Println("Valid")
		}else{
			fmt.Println("Invalid")
		}
	}
}

