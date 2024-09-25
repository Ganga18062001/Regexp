package main

import (
	"fmt"
	"regexp"
)

func Gmail(){
	pattern := "^[a-zA-Z0-9]+(?:\\.[a-zA-Z0-9]+)*@gmail\\.com$"
	reg, err := regexp.Compile(pattern)
	if err != nil{
		fmt.Println(err)
		return
	}
	var gmail string
	fmt.Println("Enter your email")
	fmt.Scan(&gmail)
	isMatch := reg.MatchString(gmail)
	if isMatch{
		fmt.Println("Gmail is match successfully")

	}else {
		fmt.Println("Gmail is not match")
	}
}
func MobilenoMatching(){
	var Mobile string
	fmt.Println("Enter your mobile number")
	fmt.Scan(&Mobile)
	pattern := `^[7-9]\d{9}$`
	reg,err := regexp.Compile(pattern)
	if err != nil{
		fmt.Println(err)
	}else{
		isMatch := reg.MatchString(Mobile)
		if isMatch{
			fmt.Println("Successfully match pattern")
		}else{
			fmt.Println("Not match pattren")
		}
	}

}

func IsvalidPancard(panNo string)bool{
	pattern := `^[A-Z]{5}[0-9]{4}[A-Z]{1}$`
	reg , err := regexp.Compile(pattern)
	if err != nil{
		fmt.Println(err)
	}
	isMatch := reg.MatchString(panNo)
	return isMatch
}

func UsingMustcompile(){
	data := "7895-7073-57-3957"
	pattren := `^\d{4}-\d{4}-\d{2}-\d{4}$`
	reg := regexp.MustCompile(pattren)
	isMatch := reg.MatchString(data)
	if isMatch{
		fmt.Println("Match pattern")

	}else{
		fmt.Println("Not match pattern")
	}
}
func Findstring(){
	data := " 4588 Welcome 4308 to pune 569 railay station"
	pattern := `\d+`
	reg :=regexp.MustCompile(pattern)
	isMatch := reg.FindString(data)
	fmt.Println(isMatch)
}
func FindAllString(){
	data := "welcome 143 to the 879 dubai 123"
	pattern := `\d+`
	reg := regexp.MustCompile(pattern)

	find := reg.FindAllString(data,-1)
	fmt.Println(find)
}
func FindStringAllIndex(){
	data := "welcome 1452 2369 to the india "
	pattern := `\d+`
	reg := regexp.MustCompile(pattern)
	find := reg.FindAllIndex([]byte(data),-1)
	fmt.Println(find)
}
func main() {
	FindStringAllIndex()
	//UsingMustcompile()
	//Findstring()
	//FindAllString()
	//Gmail()
	//MobilenoMatching()
// 	var panNo string
// 	fmt.Println("Enter your panNo")
// 	fmt.Scan(&panNo)
// 	if IsvalidPancard(panNo){
// 		fmt.Println("PanNo is valid")
// 	}else{
// 		fmt.Println("PanNo is not valid")
// 	}
 }