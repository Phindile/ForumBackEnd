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
		&rest.Route{"GET", "/topics", GetAllReplys},
		&rest.Route{"POST", "/topics", PostReply},
		&rest.Route{"GET", "/topics/:TopocID", GetReply},
		&rest.Route{"DELETE", "/topics/:TopocID", DeleteReply},
	)
	http.ListenAndServe(":3000", &handler)
}

var myreply= map[string]*domain.Reply{}

func GetReply(w rest.ResponseWriter, r *rest.Request) {
	replyid:= r.PathParam("replyid")
	reply := mytopic[replyid]
	if reply == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&reply)
}

func GetAllReplys(w rest.ResponseWriter, r *rest.Request) {
	replys := make([]*domain.Reply, len(myreply))
	i := 0
	for _, reply:= range myreply {
		replys[i] = reply
		i++
	}
	w.WriteJson(&replys)
}

func PostReply(w rest.ResponseWriter, r *rest.Request) {
	reply := domain.Reply{}
	err := r.DecodeJsonPayload(&reply)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if reply.ReplyID== "" {
		rest.Error(w, "ReplyID is required", 400)
		return
	}
	if reply.Comment.CommentID== "" {
		rest.Error(w, "CommentID is required", 400)
		return
	}
	if reply.Discussion.DiscussionID== "" {
		rest.Error(w, "DiscussionID is required", 400)
		return
	}
	if reply.Reply== "" {
		rest.Error(w, "REPLY is required", 400)
		return
	}
	myreply[reply.ReplyID] = &reply
	w.WriteJson(&reply)
}

func DeleteReply(w rest.ResponseWriter, r *rest.Request) {
	replyid := r.PathParam("replyid")
	delete(myreply,replyid)
}
