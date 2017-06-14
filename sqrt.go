package main

import "fmt"

var h int

func mod(p1, p2 float64) float64{
	ans:=p1-p2
	if ans>=0 {
		return ans
	}
	return ans*(-1)
}

func sqrt(x float64, z float64) float64{
	z1:=(z*z+x)/(2*z)
	fmt.Println(h, z, z1, z-z1)
	if mod(z1,z)>0.00000001{
		z1=sqrt(x, z1)
	}
	return z1

}

func main(){
	h=1.0
	fmt.Println(sqrt(2, 1.0))
}
