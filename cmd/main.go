// Необходимо спроектировать архитектуру ресторана, которая позволит вести учет смен
//персонала. В структуре "Ресторан" должны быть методы, позволяющие назначить на смену
//как повара, так и официанта. Однако, следует обратить внимание, что официанта
//нельзя назначить на роль повара, так как их обязанности существенно различаются,
// но повара можно назначать на обе роли.Каждый раз, когда сотрудник выходит на смену
//или принимает заказ, в специальный журнал ресторана записывается информация о событии,
//указывая, кто вышел на смену, кто принял заказ, а также, при необходимости,
//кто будет готовить соответствующее блюдо.

package main

import (
	"fmt"

	"github.com/EvgenTrue/leto2023/internal/restorant"
	"github.com/EvgenTrue/leto2023/internal/staff"
)


 

 
func main(){
	chief:=&staff.Chief{Name: "IVANOV"}
	restorant:=restorant.Restorant{}
	restorant.AddWaiter(chief)
	restorant.AddChief(chief)
	zakaz:=restorant.MakeOrder() 
	fmt.Println(zakaz)
	waiter:=&staff.Waiter{Name: "ARBUZOV"}
	restorant.AddChief(chief)
	restorant.AddWaiter(waiter)
	zakaz1:=restorant.MakeOrder()
	fmt.Println(zakaz1)
}

