package reg

import (
	"github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"strconv"
)
type Reg struct {
	Host string
	Port int
	client consul.Client
}

func MakeR (host string , port int) (*Reg , error) {
	config := api.DefaultConfig()
	config.Address = host + ":" + strconv.Itoa(port)
	apiclient , err := api.NewClient(config)
	if err != nil {
		return nil ,err
	}
	client := consul.NewClient(apiclient)
	return &Reg{Port: port , Host: host , client: client} , nil
}

func (r Reg) Register (host string ,port int ) bool {
	config := api.AgentServiceRegistration{
		Port: port ,
		Address: host ,
		Name: "testname" ,
		ID: "testid",
		Check: &api.AgentServiceCheck{
			Interval: "15s" ,
			HTTP: "http://" + host + ":" + strconv.Itoa(port) + "/health" ,
			DeregisterCriticalServiceAfter: "30s",
		},
	}
	if err := r.client.Register(&config) ; err != nil {
		return false
	}
	return true
}

func (r Reg) Deregister (id string) bool {
	config := api.AgentServiceRegistration{ID: id}
	if err := r.client.Deregister(&config) ; err != nil {
		return false
	} else {
		return true
	}
}
