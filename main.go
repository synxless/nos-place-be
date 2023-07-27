package main

import (
	"log"
	"nosplace/apis"
	"nosplace/storage"
	placewatcher "nosplace/workers/place-watcher"
	"os"
)

func main() {

	defaultRPC := os.Getenv("DEFAULT_RPC")
	defaultRPCWS := os.Getenv("DEFAULT_RPCWS")
	port := os.Getenv("PORT")
	domain := os.Getenv("DOMAIN")
	dbName := os.Getenv("DB")
	mongoURI := os.Getenv("MONGO_URI")
	mode := os.Getenv("MODE")
	_ = mode
	_ = defaultRPC
	_ = defaultRPCWS
	_ = port
	_ = domain

	storageInst, err := storage.InitStorage(dbName, mongoURI)
	if err != nil {
		log.Println(err)
		return
	}

	err = storageInst.CreateIndex()
	if err != nil {
		log.Println(err)
	}
	indexer := placewatcher.PlaceIndexer{
		Storage:    storageInst,
		DefaultRPC: defaultRPC,
	}
	go indexer.Start()

	newRouter := apis.Router{
		Port:    port,
		Domain:  domain,
		Storage: storageInst,
	}
	newRouter.Start()
}
