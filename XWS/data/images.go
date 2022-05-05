package data

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func StoreImageToDB(filename string, postID primitive.ObjectID) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	bucket, err := gridfs.NewBucket(Client.Database("xws"), options.GridFSBucket().SetName("images"))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	opts := options.GridFSUpload()
	opts.SetMetadata(bsonx.Doc{{Key: "postID", Value: bsonx.ObjectID(postID)}})
	uploadStream, err := bucket.OpenUploadStream(filename, opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer uploadStream.Close()
	fileSize, err := uploadStream.Write(data)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)
}
