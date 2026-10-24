package controller

import (
	"context"
	"fmt"
	"log"

	model "github.com/farhaan-shamsee/golang-practise/03-web/08-mongo-api/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://farrowaman:<pwd>@cluster0.kaqhouk.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT

var collection *mongo.Collection

// connect with mongo DB

func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MOngoDB connection successful")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection instance is ready")

}

// MONGODB helpers

// insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in DB with ID: ", inserted.InsertedID)
}

// update one record
func updateOneMovie(movieId string) {
	id, err := bson.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete 1 record

func deleteOneMovie(movieId string) {
	id, err := bson.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete count: ", deleteCount)
}

// delete all records from MongoDB

func deleteAllMovie() int64 {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete count: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database
func getAllMovies() []bson.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []bson.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}
