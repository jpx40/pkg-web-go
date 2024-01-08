package main

import (
	"fmt"
	"time"

	_ "github.com/IBM/fp-go/function"
	"github.com/jpx40/pkg-web-go/pkg/router"

	_ "github.com/gorilla/schema"
	_ "github.com/markbates/goth"
	_ "github.com/phakornkiong/go-pattern-match/pattern"
	_ "github.com/sourcegraph/conc"
)

func main() {
	startTime := int64(time.Now().UnixNano())
	router.Router()

	endTime := int64(time.Now().UnixNano())

	time := startTime - endTime

	fmt.Println("Time taken to run: ", time)
}
