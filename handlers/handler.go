package handler

import (
	"context"
	"databaseserverintegration/database"
	"databaseserverintegration/models"

	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.Collection

// insertOneCourse
func InsertoneCourse(course models.Course) primitive.ObjectID {
	inserted, err := collection.InsertOne(context.TODO(), course)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("The value of inserted is ", inserted)
	return inserted.InsertedID.(primitive.ObjectID)

}

// update one course
func UpdateoneCourse(courseId string) int64 {
	id, err := primitive.ObjectIDFromHex(courseId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true, "completed": true}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
	return result.ModifiedCount
}

//delete one

func DeleteoneCourse(courseId string) int64 {
	id, err := primitive.ObjectIDFromHex(courseId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("Course got deleted with the delete count: ", deleteResult)
	return deleteResult.DeletedCount
}

//delete all record from monogdb

func DeleteallCourse() int64 {
	filter := bson.M{}
	deleteResult, err := collection.DeleteMany(context.TODO(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	return deleteResult.DeletedCount

}

//getallcourse

func GetallCourse() []primitive.M {
	courseCur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var courses []primitive.M
	for courseCur.Next(context.Background()) {
		var course bson.M
		err := courseCur.Decode(&course)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, course)
	}
	defer courseCur.Close(context.Background())
	return courses
}

//getOneCourse

func GetoneCourse(courseId string) primitive.M {
	id, err := primitive.ObjectIDFromHex(courseId)
	if err != nil {
		log.Fatal(err)
	}
	filterOneCourse := bson.M{"_id": id}

	oneCourseCur := collection.FindOne(context.TODO(), filterOneCourse)
	var course primitive.M
	err = oneCourseCur.Decode(&course)
	if err != nil {
		log.Fatal(err)
	}

	return course
}
