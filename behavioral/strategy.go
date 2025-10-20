package main
import "fmt"

type ShippingStrategy interface {
    Calculate(weight float64) float64
}

type AirShipping struct{}
func (AirShipping) Calculate(w float64) float64 { return 10 + w*2 }
type SeaShipping struct{}
func (SeaShipping) Calculate(w float64) float64 { return 5 + w*0.5 }

type Order struct { strategy ShippingStrategy }
func (o *Order) SetStrategy(s ShippingStrategy) { o.strategy = s }
func (o *Order) ShippingCost(w float64) float64 { return o.strategy.Calculate(w) }

func main(){
    o := &Order{strategy: AirShipping{}}
    fmt.Println(o.ShippingCost(3.5))
    o.SetStrategy(SeaShipping{})
    fmt.Println(o.ShippingCost(3.5))
}
