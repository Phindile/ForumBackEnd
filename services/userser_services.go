
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
&rest.Route{"GET", "/users", GetAllUsers},
&rest.Route{"POST", "/users", CreateUser},
&rest.Route{"GET", "/users/:userID", GetUser},
&rest.Route{"DELETE", "/users/:userID", DeleteUser},
)
http.ListenAndServe(":3000", &handler)
}

var myusers= map[string]*domain.User{}

func GetUser(w rest.ResponseWriter, r *rest.Request) {
	userid := r.PathParam("UserID")
	user:= myusers[userid]
	if user == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&user)
}

func GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
	users:= make([]*domain.User, len(myusers))
	i := 0
	for _, user:= range myusers {
		users[i] = user
		i++
	}
	w.WriteJson(&users)
}

func CreateUser(w rest.ResponseWriter, r *rest.Request) {
	user:= domain.User{}
    err := r.DecodeJsonPayload(&user)
   if err != nil {
    rest.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
if user.UserID== "" {
    rest.Error(w, "UserID is required", 400)
    return
}
if user.Avatar== "" {
rest.Error(w, " PLease Upload a Picture ", 400)
return
}
if user.Contacts.Cell== "" {
rest.Error(w, "Cell NUmber is required", 400)
return
}
if user.Contacts.Email== "" {
rest.Error(w, "Email Address is required", 400)
return
}
	if user.Contacts.HomePhone== "" {
		rest.Error(w, "HomePhone is required", 400)
		return
	}
	if user.Contacts.Work== "" {
		rest.Error(w, "Work Number is required", 400)
		return
	}
myusers[user.UserID] = &user
w.WriteJson(&user)
}

func DeleteUser(w rest.ResponseWriter, r *rest.Request) {
	userid := r.PathParam("userD")
	delete(myusers,userid)
}
