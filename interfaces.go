package figuras

import (
	"fmt"
)

type Figura interface {
	area() float64
	perimetro() float64
}

func Medidas(figura Figura) {
	fmt.Println("medidas: ", figura)
	fmt.Println("area: ", figura.area())
	fmt.Println("perimetro: ", figura.perimetro())
}
