package models

import (
	"context"
	"fmt"
	"log"
	"peddie-backend/library/constants"
	"peddie-backend/library/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StoreRecipe(recipe dto.Recipe) error {
	coll := mongoClient.Database(constants.Db).Collection(constants.CollName)

	inserted, err := coll.InsertOne(context.TODO(), recipe)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("inserted a recipe with id: ", inserted.InsertedID)

	return nil
}

func GetOneRecipe(id string) (*dto.Recipe, error) {
	coll := mongoClient.Database(constants.Db).Collection(constants.CollName)
	var recipe dto.Recipe

	recipeId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": recipeId}
	err = coll.FindOne(context.TODO(), filter).Decode(&recipe)
	if err != nil {
		return nil, err
	}

	return &recipe, nil

}

func GetRecipes(id string) (*[]dto.Recipe, error) {
	coll := mongoClient.Database(constants.Db).Collection(constants.CollName)
	var recipe []dto.Recipe

	creatorId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"creator_id": creatorId}
	result, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	result.Decode(&recipe)

	return &recipe, nil

}
