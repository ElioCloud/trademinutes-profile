package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID                primitive.ObjectID `bson:"_id,omitempty"`
    Name              string             `bson:"name"`
    Email             string             `bson:"email"`
    Password          string             `bson:"password,omitempty"`
    Program           string             `bson:"program,omitempty"`
    Location          string             `bson:"location,omitempty"`
    Interests         string             `bson:"interests,omitempty"`
    Languages         []string           `bson:"languages,omitempty"`
    JoinedDate        string             `bson:"joinedDate,omitempty"`
    ExpectedGradDate  string             `bson:"expectedGradDate,omitempty"`
    ProfilePictureURL string             `bson:"profilePictureURL,omitempty"`
}
