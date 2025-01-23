package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/johannesboyne/gofakes3"
	"github.com/johannesboyne/gofakes3/backend/s3mem"
)

func main() {
	addr := flag.String("addr", "localhost:9000", "Address to listen on")
	bucket := flag.String("bucket", "um", "bucket name to create")
	flag.Parse()
	backend := s3mem.New()
	backend.CreateBucket(*bucket)
	faker := gofakes3.New(backend)
	if err := http.ListenAndServe(*addr, faker.Server()); err != nil {
		log.Fatalln("cannot run s3 mock", err)
	}
}
