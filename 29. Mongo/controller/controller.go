package controller

import (
	"encoding/json"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/Hycient195/Mongo_Exercise/models"
	"fmt"
	"log"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const MONGO_URL = "mongodb://localhost/demo-blog";
const dbName = "blog-reader";
const collectionName = "blogs";

var collection *mongo.Collection;

func init() {
	clientOption := options.Client().ApplyURI(MONGO_URL);

	client, err := mongo.Connect(context.TODO(), clientOption);
	HandleFatalError(err);

	fmt.Println("Sucessfully connected to database");

	collection = client.Database(dbName).Collection(collectionName);

	fmt.Println("Collection instance created");
}


/* ================================== */
/* MongoDB Handler To Insert One Post */
/* ================================== */
func insertOnePost(post model.BlogPost) *mongo.InsertOneResult {
	inserted, err := collection.InsertOne(context.Background(), post);
	HandleFatalError(err);

	fmt.Printf("Inserted one post %v\n", inserted);
	return inserted;
}


/* =============================== */
/* MongoDB Handler To Get One Post */
/* =============================== */
func getOnePost(postId string) *mongo.SingleResult {
	mongo_id, _ := primitive.ObjectIDFromHex(postId)
	filter := bson.M{"_id": mongo_id};
	result := collection.FindOne(context.Background(), filter);
	fmt.Println(postId);
	return result;
}


/* ======================================= */
/* MongoDB Handler To Update Only One Post */
/* ======================================= */
func updateOnePost(postId string, postData model.BlogPost) *mongo.UpdateResult {
	// We obtain the Id of the post to be updated as string, and then converti it to MongoDB ObjectId BSON form;
	id, _ := primitive.ObjectIDFromHex(postId);

	// The MongoDB Update method accepts the crieteria to find what to update, and the value to update as parameters
	// These are not defined belod;

	filter := bson.M{"_id": id};
	// update := bson.M{"$set": bson.M{ "isRead": true }};
	update := bson.M{"$set": postData};

	result, err := collection.UpdateOne(context.Background(), filter, update);
	HandleFatalError(err);

	fmt.Printf("Update result is %v\n", result);
	return result;
}

func deleteOnePost(postId string) *mongo.DeleteResult {
	id, _ := primitive.ObjectIDFromHex(postId);
	filter := bson.M{"_id": id};

	deleteResult, err := collection.DeleteOne(context.Background(), filter);
	HandleFatalError(err);
	return deleteResult;
}


func deleteManyPosts() *mongo.DeleteResult {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}});
	HandleFatalError(err);
	fmt.Printf("Delete result is %v\n", deleteResult);
	return deleteResult;
}


func getAllPosts() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{});
	HandleFatalError(err);

	var posts []primitive.M

	for cursor.Next(context.Background()) {
		var post bson.M;
		err := cursor.Decode(&post);
		HandleFatalError(err);
		posts = append(posts, post);
	}

	defer cursor.Close(context.Background());
	return posts;
}




/* ============================= */
/* Defining The Control Hanlders */
/* ============================= */
func HandleInsertOnePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode");
	var blogPostInstance model.BlogPost;
	err := json.NewDecoder(r.Body).Decode(&blogPostInstance);
	HandlePanicError(err);
	response := insertOnePost(blogPostInstance)
	json.NewEncoder(w).Encode(response)
}

func HandleGetOnePost(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r);
	foundPost := getOnePost(postId["postId"]);
	json.NewEncoder(w).Encode(foundPost);
}

func HandleUpdateOnePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode");
	w. Header () . Set ("Allow-Control-Allow-Methods", "POST" );
	postId := mux.Vars(r);
	var postUpdate model.BlogPost;
	json.NewDecoder(r.Body).Decode(&postUpdate);
	updateResponse := updateOnePost(postId["postId"], postUpdate);
	er := json.NewEncoder(w).Encode(updateResponse);
	HandlePanicError(er);
}

func HandleDeleteOnePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode");
	postId := mux.Vars(r);
	result := deleteOnePost(postId["postId"]);
	json.NewEncoder(w).Encode(result);
}

func HandleDeleteManyPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode");
	result := deleteManyPosts();
	json.NewEncoder(w).Encode(result);
}

func HandleGetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode");
	post := getAllPosts();
	marshalledData, err := json.MarshalIndent(post, "", "\t")
	HandlePanicError(err);
	w.Write([]byte(marshalledData));
}



func HandleFatalError(e error) {
	if e != nil {
		log.Fatal(e);
	}
}

func HandlePanicError(e error) {
	if e != nil {
		panic(e);
	}
}