package main

import (
	"LearningProject/internal/catalog/db"
	"LearningProject/internal/config"
	"LearningProject/pkg/client/postgresql"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	program := params.ByName("programId")
	_, err := w.Write([]byte(fmt.Sprintf("Hello, %s. Your program ID is %s", name, program)))
	if err != nil {
		log.Fatalf("write error: %v", err)
	}
	log.Printf("Somebody called GET localhost:1234/%s/program/%s", name, program)
}

func main() {

	cfg := config.GetConfig()

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		log.Fatalf("%v, %v", err, postgreSQLClient)
	}

	repository := catalog.NewRepository(postgreSQLClient)

	all, err := repository.FindAll(context.TODO())
	if err != nil {
		log.Fatalf("%v", err)
	}

	//ctlg := catalog2.Catalog{"", "carrot"}
	//if err := repository.Create(context.TODO(), &ctlg); err != nil {
	//	log.Printf("error while update catalog: %v", err)
	//}

	//if err := repository.Delete(context.TODO(), "8"); err != nil {
	//	log.Printf("error while delete from catalog: %v", err)
	//}

	for _, cat := range all {
		log.Printf("%v", cat)
	}

	router := httprouter.New()
	router.GET("/:name/program/:programId", IndexHandler)
	//router.POST("/:product", IndexHandler)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(fmt.Sprintf("listener returned error: %v", err))
	}

	server := &http.Server{
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("http server error: %v", err)
	}

}
