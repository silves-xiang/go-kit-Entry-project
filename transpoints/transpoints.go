package transpoints

import (
	"context"
	"encoding/json"
	"errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"baction/endopints"
)

func MakeMux (endpointsz endopints.Bendpoints ) http.Handler {
	r := mux.NewRouter()
	r.Methods("POST").Path("/con/{msgo}/{msgt}").Handler(kithttp.NewServer(endpointsz.ConEndopint ,DecodeCon , EncodeCon))
	r.Methods("POST").Path("/diff/{msgo}/{msgt}").Handler(kithttp.NewServer(endpointsz.DifEndpoint , DecodeDif , EncodDif))
	r.Methods("GET").Path("/health").Handler(kithttp.NewServer(endpointsz.HealEndpoint , DecodHeal , EncodHeal))
	return r
}

func DecodeCon (_ context.Context , r *http.Request) (interface{} , error) {
	vars := mux.Vars(r)
	msgo , ok:= vars["msgo"]
	if !ok {
		return nil , errors.New("params has fault")
	}
	msgt ,ok := vars["msgt"]
	if !ok {
		return nil , errors.New("param has fault")
	}
	return endopints.ConRequest{Msgo: msgo , Msgt: msgt} , nil
}

func EncodeCon (ctx context.Context , w http.ResponseWriter , response interface{} ) error {
	w.Header().Set("Content-Type" , "application/json")
	return json.NewEncoder(w).Encode(response)
}

func DecodeDif (_ context.Context , request *http.Request ) (interface{} , error) {
	vars := mux.Vars(request)
	var DifR  endopints.DifRequest
	if msgo , ok :=vars["msgo"] ; ok {
		DifR.Msgo = msgo
	} else {
		return nil , errors.New("params has fault")
	}
	if msgt , ok := vars["msgt"] ; ok {
		DifR.Msgt = msgt
	} else {
		return nil , errors.New("params has fault")
	}
	return DifR , nil
}

func EncodDif (_ context.Context , w http.ResponseWriter , response interface{}) error {
	w.Header().Set("Content-Type" , "application/json")
	return json.NewEncoder(w).Encode(response)
}

func DecodHeal (_ context.Context , r *http.Request) (response interface{} , err error) {
	return endopints.HealRequest{} , nil
}

func EncodHeal (_ context.Context , w http.ResponseWriter , response interface{}) error {
	w.Header().Set("Content-Type" , "application/json")
	return json.NewEncoder(w).Encode(response)
}