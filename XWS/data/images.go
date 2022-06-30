package data

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type Image struct {
	FileID   primitive.ObjectID `bson:"_id,omitempty"`
	Filename string             `json:"filename"`
	postID   primitive.ObjectID
}

type Images []*Image

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
func GetImageByPostIDs(postIDs []primitive.ObjectID) Images {
	fmt.Println("[DEBUG] entered getting images")
	db := Client.Database("xws")
	imagesCollection := db.Collection("images.files")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var results Images
	for i := 0; i < len(postIDs); i++ {
		var result Image
		fmt.Println(postIDs[i])
		filter := bson.D{{"metadata.postID", postIDs[i]}}
		err := imagesCollection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			fmt.Println("[ERROR] couldnt find/decode image")
			log.Fatal(err)
		}
		results = append(results, &result)
	}

	bucket, _ := gridfs.NewBucket(db, options.GridFSBucket().SetName("images"))

	for i := 0; i < len(results); i++ {
		var buf bytes.Buffer
		fmt.Println(results[i].Filename)
		dStream, err := bucket.DownloadToStreamByName(results[i].Filename, &buf)
		if err != nil {
			fmt.Println("[ERROR] downloading image")
			log.Fatal(err)
		}
		fmt.Printf("File size to download: %v \n", dStream)
		ioutil.WriteFile(results[i].Filename, buf.Bytes(), 0644)
	}
	return results
}
