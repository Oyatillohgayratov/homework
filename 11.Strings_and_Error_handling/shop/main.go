package main

import (
	"fmt"
	"os"
	"errors"
)
type Basket struct {
	p_basket map[string]int
	price float64
}


func (b *Basket) buy_products(p map[string]float64){
	var a string
	fmt.Print(":::maxsulotlar:::\n")
	for k, v := range p {
	   fmt.Printf("%s summasi %.1f\n", k, v)
	}

	for{
		fmt.Printf("\nsavatga nima solasiz(chiqish uchun 'exit'): ")
		fmt.Scanln(&a)
		if a == "exit"{
			break
		}	
		for k, v := range p {
			if k == a{
				fmt.Printf("%s savatga solindi",a)
				b.p_basket[a]++
				b.price += v
			}
		}
	}


}

func (b *Basket) return_products(p map[string]float64){
	var a string
	fmt.Printf("savatda bor maxsulotlar\n")
    for k := range b.p_basket {
       fmt.Printf("%s\n", k)
    }

    for{
        fmt.Printf("\nqayrtaradigan maxsulotning nomimni kiriting(chiqish uchun 'exit'): ")
        fmt.Scanln(&a)
        if a == "exit"{
            break
        }    
        for k, v := range p {
            if k == a{
				fmt.Printf("%s savatdan olib tashlandi", a)
				b.p_basket[k]--
                if b.p_basket[a] == 0{
					delete(b.p_basket,a)
				}
                b.price -= v
            }
        }
    }
}

func (b *Basket)check(p map[string]float64){
	fmt.Printf("\n    :::check:::   \n")
	i := 0
	for k, v := range p {
		for s := range b.p_basket{
			if k == s{
				i++
                fmt.Printf("%d ta %s summasi %.1f\n",i, k, v)
            }
		}
		i = 0
	}
	fmt.Printf("summa  %.1f\n", b.price)
}

func main() {
	var products = map[string]float64{"banana":19.100,"apple":15.200,"water":3.200}
	basket := Basket{p_basket: make(map[string]int)}
	var a int

	
	
	for {
		fmt.Printf("\n\n1.Maxsulot olish\n2.Maxsulotni atmen qilish\n3.chek\n0.chiqib ketish\n")
		fmt.Scanln(&a)
		switch a {
			case 0:
				os.Exit(0)
			case 1:
				basket.buy_products(products)
			case 2:
				basket.return_products(products)
			case 3:
				basket.check(products)
				os.Exit(0)
			default:
				fmt.Print(errors.New("Invalid Input!"))
		}
		
	}
	
	
}
