package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"domain"
)

func main() {

	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}
	handler.SetRoutes(
		&rest.Route{"GET", "/reputation",GetAllReputation},
		&rest.Route{"POST", "/reputation", CreateReputation},
		&rest.Route{"GET", "/reputation/:reputationid", GetReputation},
		&rest.Route{"DELETE", "/reputation/:reputationid", DeleteTopic},
	)
	http.ListenAndServe(":3000", &handler)
}

var myreputation= map[string]*domain.Reputation{}

func GetReputation(w rest.ResponseWriter, r *rest.Request) {
	reputationid := r.PathParam("reputationid")
	reputation:= myreputation[reputationid]
	if reputation== nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&reputation)
}

func GetAllReputation(w rest.ResponseWriter, r *rest.Request) {
	reputations:= make([]*domain.Reputation, len(myreputation))
	i := 0
	for _, reputation := range myreputation{
		reputations[i] = reputation
		i++
	}
	w.WriteJson(&reputations)
}

func CreateReputation(w rest.ResponseWriter, r *rest.Request) {
	reputation := domain.Reputation{}
	err := r.DecodeJsonPayload(&reputation)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if reputation.ReputationID== "" {
		rest.Error(w, "Reputation ID is required", 400)
		return
	}
	if reputation.Owner.UserID== "" {
		rest.Error(w, "UserID is required", 400)
		return
	}
	if reputation.Rate==0{
		rest.Error(w, "Rate is required", 400)
		return
	}

	myreputation[reputation.ReputationID] = &reputation
	w.WriteJson(&reputation)
}

func DeleteTopic(w rest.ResponseWriter, r *rest.Request) {
	reputationid := r.PathParam("reputationid")
	delete(myreputation,reputationid)
}
