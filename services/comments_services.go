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
		&rest.Route{"GET", "/comments ", GetAllComments},
		&rest.Route{"POST", "/comments ", CreateComment },
		&rest.Route{"GET", "/comments /:commentid ", GetComment},
		&rest.Route{"DELETE", "/comments/:commentid", DeleteComment},
	)
	http.ListenAndServe(":3000", &handler)
}
var mycomments= map[string]*domain.Comment{}

func GetComment(w rest.ResponseWriter, r *rest.Request) {
	commentid:= r.PathParam("commentid")
	comments := mycomments[commentid]
	if comments  == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&comments)
}

func GetAllComments(w rest.ResponseWriter, r *rest.Request) {
	comments := make([]*domain.Comment, len(mycomments))
	i := 0
	for _,comment := range mycomments{
		comments [i] = comment
		i++
	}
	w.WriteJson(&comments )
}

func CreateComment(w rest.ResponseWriter, r *rest.Request) {
	comment := domain.Comment{}
	err := r.DecodeJsonPayload(&comment)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if comment.CommentID== "" {
		rest.Error(w, "CommentID is required", 400)
		return
	}
	if comment.Discussion.DiscussionID== "" {
		rest.Error(w, "DiscussionID is required", 400)
		return
	}
	if comment.Discription== "" {
		rest.Error(w, "Description is required", 400)
		return
	}

	mycomments[comment.commentID] = &comment
	w.WriteJson(&comment)
}

func DeleteComment(w rest.ResponseWriter, r *rest.Request) {
	commentid := r.PathParam("commentid")
	delete(mycomments,commentid)
}
