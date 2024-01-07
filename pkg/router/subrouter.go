package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Test_Data struct {
	Test string `json:"test"`
}

func subrouter(r chi.Router) {
	r.Get("/api/test", func(w http.ResponseWriter, r *http.Request) {
		t := Test_Data{Test: "test"}
		strJ, _ := json.Marshal(&t)
		//		 json.NewEncoder(w).Encode(t)
		w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(t)
		// w.Write([]byte(string(strJ)))
		_, err := w.Write(strJ)
		if err != nil {
			fmt.Println(err)
		}
	})
}
