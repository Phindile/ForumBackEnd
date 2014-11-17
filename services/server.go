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
		&rest.Route{"GET", "/topics", GetAllTopics},
		&rest.Route{"POST", "/topics", PostTopic},
		&rest.Route{"GET", "/topics/:TopocID", GetTopic},
		&rest.Route{"DELETE", "/topics/:TopocID", DeleteTopic},
	)
	http.ListenAndServe(":3000", &handler)
}

var mytopic= map[string]*domain.Topic{}

func GetTopic(w rest.ResponseWriter, r *rest.Request) {
	topicid := r.PathParam("TopicID")
	country := mytopic[topicid]
	if country == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&country)
}

func GetAllTopics(w rest.ResponseWriter, r *rest.Request) {
	topics := make([]*domain.Topic, len(mytopic))
	i := 0
	for _, topic := range mytopic {
		topics[i] = topic
		i++
	}
	w.WriteJson(&topics)
}

func PostTopic(w rest.ResponseWriter, r *rest.Request) {
	topic := domain.Topic{}
	err := r.DecodeJsonPayload(&topic)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if topic.TopicID == "" {
		rest.Error(w, "TopicID is required", 400)
		return
	}
	if topic.Title== "" {
		rest.Error(w, "Topic Title is required", 400)
		return
	}
	if topic.Description== "" {
		rest.Error(w, "Description is required", 400)
		return
	}
	if topic.DatePosted== "" {
		rest.Error(w, "DatePosted is required", 400)
		return
	}
	mytopic[topic.TopicID] = &topic
	w.WriteJson(&topic)
}

func DeleteTopic(w rest.ResponseWriter, r *rest.Request) {
	topicid := r.PathParam("TopicID")
	delete(mytopic,topicid)
}
