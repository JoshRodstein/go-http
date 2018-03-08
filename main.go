/*
* Sample API with GET and POST endpoint.
* POST data is converted to string and saved in internal memory.
* GET endpoint returns all strings in an array.
 */
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

)

type User struct {
		Username string
}

var (
	// flagPort is the open port the application listens on
	flagPort = flag.String("port", "9000", "Port to listen on")
)

func main() {


	mux := http.NewServeMux()
	mux.HandleFunc("/get", GetHandler)
	mux.HandleFunc("/post", PostHandler)

	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))
}



// GetHandler handles the index route
func GetHandler(w http.ResponseWriter, r *http.Request) {
		user := User{}

		// decode json request data into User struct
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
				http.Error(w, "json body empty", 200)
				return
		}

		// marchal data from user struct back into json
		// for writing
		userJson, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "Error converting results to json", 400)
			return
		}

		fmt.Fprint(w, "jason data parsed and recieved\n:")
		w.Write(userJson)
		fmt.Fprint(w, "\n")
}

// PostHandler converts post request body to string
func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var u User
		err := decoder.Decode(&u)
		if err != nil {
			http.Error(w, "Unable to parse JSON body", 400)
		}
		defer r.Body.Close()
		fmt.Fprint(w, u.Username)
		fmt.Fprint(w, "\n")
		return
	} else {
		http.Error(w, "Ivalid request method", http.StatusMethodNotAllowed)
	}
}



func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	flag.Parse()
}
