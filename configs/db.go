package configs

import (
	"context"
	"fmt"
	"imdb/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func ConnectDB() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return &DB{client: client}
}

func colHelper(db *DB, collectionName string) *mongo.Collection {
	return db.client.Database("hollywood").Collection(collectionName)
}

func (db *DB) CreateMovie(input *model.NewMovie) (*model.Movie, error) {
	collection := colHelper(db, "movie")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		return nil, err
	}

	movie := &model.Movie{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		ActorID:     input.ActorID,
		Name:        input.Name,
		Description: input.Description,
		Status:      model.StatusNotStarted,
	}

	return movie, err
}

func (db *DB) CreateActor(input *model.NewActor) (*model.Actor, error) {
	collection := colHelper(db, "actor")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		return nil, err
	}

	actor := &model.Actor{
		ID:    res.InsertedID.(primitive.ObjectID).Hex(),
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	return actor, err
}

func (db *DB) GetActors() ([]*model.Actor, error) {
	collection := colHelper(db, "actor")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var actors []*model.Actor
	defer cancel()

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleActor *model.Actor
		if err = res.Decode(&singleActor); err != nil {
			log.Fatal(err)
		}
		actors = append(actors, singleActor)
	}

	return actors, err
}

func (db *DB) GetMovies() ([]*model.Movie, error) {
	collection := colHelper(db, "movie")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var movies []*model.Movie
	defer cancel()

	res, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer res.Close(ctx)
	for res.Next(ctx) {
		var singleMovie *model.Movie
		if err = res.Decode(&singleMovie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, singleMovie)
	}

	return movies, err
}

func (db *DB) SingleActor(ID string) (*model.Actor, error) {
	collection := colHelper(db, "actor")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var actor *model.Actor
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&actor)

	return actor, err
}

func (db *DB) SingleMovie(ID string) (*model.Movie, error) {
	collection := colHelper(db, "movie")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var movie *model.Movie
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(ID)

	err := collection.FindOne(ctx, bson.M{"_id": objId}).Decode(&movie)

	return movie, err
}
