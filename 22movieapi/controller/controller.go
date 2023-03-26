package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ankit85/movieapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://falconn:farzidesiboyzz@cluster0.ysngla4.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Nextflix"
const colName = "Wachlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Db connection is successful")
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Client is successfull handshake")
}

// insert method
func insertOneMovieInDb(movie model.Nextflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insert value in db:-", inserted)
	fmt.Println("Insert value in db with id:-", inserted.InsertedID)
}

// update
func updateMovieInDbByID(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", result)
	fmt.Println("result:", result.ModifiedCount)
}

// delete One
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleteCount:", deleteCount)
}

// delete All
func deleteAllMovie() {

	deleteCount, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleteCount:", deleteCount.DeletedCount)
}

// get all movies
func getAllMovies() []primitive.M {

	cur, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		if err := cur.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies
}

// actual controller
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movies model.Nextflix
	err := json.NewDecoder(r.Body).Decode(&movies)
	if err != nil {
		log.Fatal(err)
	}

	insertOneMovieInDb(movies)
	json.NewEncoder(w).Encode(movies)
}

func MarkedWatchedMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateMovieInDbByID(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	id := mux.Vars(r)["id"]
	deleteOneMovie(id)
	json.NewEncoder(w).Encode("Deleted Hurray gm!!!")
}
func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteAllMovie()
	json.NewEncoder(w).Encode("Bhai pura delete ho gya abh kya gm")
}
