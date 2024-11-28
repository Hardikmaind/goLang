/*

package main

import (
	"fmt"
	"net/http"
	"os/user"

	"github.com/gohugoio/hugo/tpl/data"
)
func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		userId:= r.FormValue("userId")
		password:= r.FormValue("password")

		userId, dbpass := database.getUserData(userId)			//! NOW HERE YOU CAN SEE WE ARE QUERYING THE DATABASE FOR THE USERID AND PASSWORD
		if userId == "" || dbpass != password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	userId:=r.FormValue("userId")
	userName,_:=database.getUserData(userId)	//! ALSO HERE YOU CAN SEE WE ARE QUERYING THE DATABASE FOR THE USERID AND PASSWORD. SO THIS IS TWO TIMES QUERYING THE DATABASE. WE CAN AVOID THIS BY USING CONTEXT
	fmt.Fprintln(w,"Hello %s!",userName)
}

*/

package main

import (
	"context"
	"fmt"
	"net/http"
)

// this is a middleware function
func authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		userId := r.FormValue("userId")
		password := r.FormValue("password")

		userName, dbpass := database.getUserData(userName) //! NOW HERE YOU CAN SEE WE ARE QUERYING THE DATABASE FOR THE USERID AND PASSWORD
		if userName == "" || dbpass != password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
		ctx, cancel := context.WithValue(r.Context(), "userName", userName) //? HERE WE ARE ADDING THE USERID TO THE CONTEXT
	}
}

// this is the main handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("userId")
	//userName,_:=database.getUserData(userId)	//! ALSO HERE YOU CAN SEE WE ARE QUERYING THE DATABASE FOR THE USERID AND PASSWORD. SO THIS IS TWO TIMES QUERYING THE DATABASE. WE CAN AVOID THIS BY USING CONTEXT

	userName := r.Context().Value("userName").(string) //?HERE WE ARE GETTING THE USERID FROM THE CONTEXT, which was send in the authenticate function
	fmt.Fprintln(w, "Hello %s!", userName)
}
func main() {
	//in tihs "/hello" is the endpoint and helloHandler is the handler function and authenticate is the middleware function
	http.HandleFunc("/hello", authenticate(helloHandler))
	http.ListenAndServe(":8080", nil)
}

/*
!Key Components

?Middleware Function (authenticate):
A middleware wraps around a handler function (helloHandler) and performs some processing before or after calling the original handler.
Here, the authenticate middleware checks the user's credentials and, if valid, adds the user's name to the context.

?Request Handler (helloHandler):
The primary handler responds with a personalized greeting, using the user's name retrieved from the context.

?main Function:
Configures the server and registers the /hello endpoint with the authenticate middleware applied to the helloHandler.*/
