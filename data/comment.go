package data

import (
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Comment struct {
	ID      string   `json:"id" `
	PostID  string   `json:"postId" validate:"required"`
	UserID  string   `json:"userId" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Like    []string `json:"like"`
	Date    string   `json:"-"`
}

func (c *Comment) Validate() error {
	v := validator.New()
	return v.Struct(c)
}

func findComments(ctx context.Context, key, value string) ([]*Comment, error) {
	var results []*Comment
	cur, err := collection.Find(ctx, bson.D{{key, value}})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var cmnt Comment
		err := cur.Decode(&cmnt)
		if err != nil {
			return nil, err
		}

		results = append(results, &cmnt)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())

	return results, err
}

// GetComments : Find every comment in database
func GetCommentsByPostID(ctx context.Context, postID string) ([]*Comment, error) {
	comments, err := findComments(ctx, "postid", postID)
	return comments, err
}

func GetCommentsByUserID(ctx context.Context, userID string) ([]*Comment, error) {
	comments, err := findComments(ctx, "userid", userID)
	return comments, err
}

// AddComment : Add comment to database
func AddComment(ctx context.Context, c Comment) error {
	c.Date = time.Now().UTC().String()
	c.ID = generateID()
	c.Like = []string{}
	_, err := collection.InsertOne(ctx, c)

	return err
}

func Vote(ctx context.Context, userID string, commentID string) error {
	res := collection.FindOne(ctx, bson.M{"id": commentID})
	var c Comment
	if err := res.Decode(&c); err != nil {
		return err
	}
	c.Like = updateLikes(userID, c.Like)
	filter := bson.M{"id": commentID}
	pushToArray := bson.M{"$set": bson.M{"like": c.Like}}
	result, err := collection.UpdateOne(ctx, filter, pushToArray)
	if result.ModifiedCount <= 0 {
		return errors.New("vote failed!")
	}
	return err
}

func updateLikes(userID string, likes []string) []string {
	for i, v := range likes {
		if v == userID {
			return append(likes[:i], likes[i+1:]...)
		}
	}
	return append(likes, userID)
}

func DeleteComment(ctx context.Context, commentID string) error {
	_, err := collection.DeleteOne(ctx, bson.M{"id": commentID})
	return err
}

func GetCommentByID(ctx context.Context, id string) (*Comment, error) {
	res := collection.FindOne(ctx, bson.M{"id": id})
	var c Comment
	if err := res.Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func generateID() string {
	return uuid.NewString()
}
