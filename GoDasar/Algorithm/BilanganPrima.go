package main

import "fmt"

func BilanganPrima(n int)  {
	for i:=1; i <= n; i++ {
		l :=0;
		for j:=1;j <=n; j++ {
			if i%j == 0 {
				l++;
			}
		}
		if (l == 2) && (i != 1){
			fmt.Println(i)
		}
	}

}

func main()  {
	BilanganPrima(20)
}