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
		&rest.Route{"GET", "/votes", GetAllVotes},
		&rest.Route{"POST", "/votes", CreateVotes},
		&rest.Route{"GET", "/votes/:voteid", GetVote},
		&rest.Route{"DELETE", "/votes/:voteid",DeleteVote},
	)
	http.ListenAndServe(":3000", &handler)
}

var myvotes= map[string]*domain.Vote{}

func GetVote(w rest.ResponseWriter, r *rest.Request) {
	votesid := r.PathParam("voteid")
	vote:= myvotes[votesid]
	if vote == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&vote)
}

func GetAllVotes(w rest.ResponseWriter, r *rest.Request) {
	votes := make([]*domain.Vote, len(myvotes))
	i := 0
	for _, vote := range myvotes{
		votes[i] = vote
		i++
	}
	w.WriteJson(&votes)
}

func CreateVotes(w rest.ResponseWriter, r *rest.Request) {
	vote:= domain.Vote{}
	err := r.DecodeJsonPayload(&vote)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if vote.VoteID== "" {
		rest.Error(w, "VoteID is required", 400)
		return
	}
	if vote.Comment.CommentID== "" {
		rest.Error(w, "CommentID is required", 400)
		return
	}
	if vote.VoteState== "" {
		rest.Error(w, "YES OR NO is required", 400)
		return
	}

	myvotes[vote.VoteID] = &vote
	w.WriteJson(&vote)
}

func DeleteVote(w rest.ResponseWriter, r *rest.Request) {
	voteid := r.PathParam("voteID")
	delete(mytopic,voteid)
}
