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
		&rest.Route{"GET", "/discussion ", GetAllDiscussions},
		&rest.Route{"POST", "/discussion", CreateDiscussion },
		&rest.Route{"GET", "/discussion /:discussionid ", GetDiscussion},
		&rest.Route{"DELETE", "/discussion/:discussionid", DeleteDiscussion},
	)
	http.ListenAndServe(":3000", &handler)
}
var mydiscuss= map[string]*domain.Discussion{}

func GetDiscussion(w rest.ResponseWriter, r *rest.Request) {
	discussid:= r.PathParam("discussionid")
	discussion := mydiscuss=[discussid]
	if discussion  == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&discussion)
}

func GetAllDiscussions(w rest.ResponseWriter, r *rest.Request) {
	discussions := make([]*domain.Comment, len(mycomments))
	i := 0
	for _,discuss := range mydiscuss{
		discussions [i] = discuss
		i++
	}
	w.WriteJson(&discussions)
}

func CreateDiscussion(w rest.ResponseWriter, r *rest.Request) {
	discussion := domain.Discussion{}
	err := r.DecodeJsonPayload(&discussion)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if discussion.DiscussionID== "" {
		rest.Error(w, "DiscussionID is required", 400)
		return
	}
	if discussion.Topic.TopicID== "" {
		rest.Error(w, "TopicID is required", 400)
		return
	}
	if discussion.Discription== "" {
		rest.Error(w, "Description is required", 400)
		return
	}


	mydiscuss[discussion.DiscussionID] = &discussion
	w.WriteJson(&discussion)
}

func DeleteDiscussion(w rest.ResponseWriter, r *rest.Request) {
	discussionid:= r.PathParam("discussionid")
	delete(mycomments,discussionid)
}
