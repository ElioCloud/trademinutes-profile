package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"trademinutes-profile/config"
	"trademinutes-profile/middleware"
	"github.com/ElioCloud/shared-models/models"
)

func UpdateProfileInfoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered UpdateProfileInfoHandler")

	email, ok := r.Context().Value(middleware.EmailKey).(string)
	if !ok {
		log.Println("Failed to get email from context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	log.Printf("Email from context: %s\n", email)

	var req models.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	log.Printf("Decoded request: %+v\n", req)

	collection := config.GetDB().Collection("MyClusterCol")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"program":              req.Program,
		"location":             req.Location,
		"interests":            req.Interests,
		"languages":            req.Languages,
		"joinedDate":           req.JoinedDate,
		"expectedGraduationDate": req.ExpectedGradDate,
	}
	log.Printf("Update document: %+v\n", update)

	result, err := collection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": update})
	if err != nil {
		log.Printf("Failed to update profile: %v\n", err)
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}
	log.Printf("Update result: %+v\n", result)

	w.Write([]byte("Profile information updated successfully"))
}
