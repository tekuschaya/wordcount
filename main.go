package main

import (
	//"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	//a := Scan1()
	//fmt.Println(args[0])
	if len(args) > 0 {
		arr := strings.Split(args[0], " ")
		//fmt.Println(arr)
		fmt.Println(len(arr))
	} else {
		fmt.Println('0')
	}
	//fmt.Println("LINE1 : ", Scan1())
	//fmt.Println("LINE2 : ", Scan2())
	//str := readInput()
	//fmt.Println(len(str))
}
/*func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
	  fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
	}
	return in.Text()
  }*/
  
  /*func Scan2() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
	  fmt.Println("Ошибка ввода: ", err)
	}
	return str
  }*/

/*func readInput() []string {
	var str []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		str = append(str, word)
	}
	return str
}*/