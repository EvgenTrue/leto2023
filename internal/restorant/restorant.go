package restorant

import "fmt" 

type waiterinterface interface {
	Order() string
	GetName() string
}
type chiefinterface interface{
   Cook() string
   GetName() string
} 

type Restorant struct {
	Chief chiefinterface 
	Waiter waiterinterface
}

func(r *Restorant)AddChief(chief chiefinterface){
	r.Chief=chief
}
func(r *Restorant)AddWaiter(waiter waiterinterface){
   r.Waiter=waiter
}
func(r *Restorant)MakeOrder()string{
   return fmt.Sprintf("Waiter: %s, Chief: %s", r.Waiter.GetName(), r.Chief.GetName())
}
