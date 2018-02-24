package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"envd/pkg/core"
)

func RunServer(core *core.CoreInstance)  {
	handlerGetKey := func (w http.ResponseWriter, r *http.Request) {
		val, err := core.GetKey(r.URL.Path[len("/"):], nil)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			w.Write([]byte("null"))
		} else {
			res, _ := json.Marshal(val)

			w.Write(res)
		}
	}

	handlerGetKeyDescriptor := func (w http.ResponseWriter, r *http.Request) {
		val, err := core.GetKeyDescriptor(r.URL.Path[len("/descriptor/"):])

		if val == nil || err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			w.Write([]byte("null"))
		} else {
			res, _ := json.Marshal(val.GetDescriptor())

			w.Write(res)
		}
	}

	http.HandleFunc("/", handlerGetKey)
	http.HandleFunc("/descriptor/", handlerGetKeyDescriptor)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
