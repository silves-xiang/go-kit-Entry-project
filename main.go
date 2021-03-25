package main
import (
	"baction/reg"
	"baction/services"
	"baction/endopints"
	"baction/transpoints"
	"fmt"
	"net/http"
)


func main () {
	var svc services.Services = services.ServiceA{}
	fmt.Println("o")
	endpoints := endopints.Bendpoints{HealEndpoint: endopints.HealEndpoint(svc) , ConEndopint: endopints.ConEndopint(svc) , DifEndpoint: endopints.DifEndpoint(svc)}
	r := transpoints.MakeMux(endpoints)
	reg ,err := reg.MakeR("127.0.0.1" , 8500)
	if err != nil {
		fmt.Println(err)
		return
	}
	reg.Register("127.0.0.1" , 8081)
	err = http.ListenAndServe(":8081" , r)
	if err != nil {
		fmt.Println(err)
	}
}