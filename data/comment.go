package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Comment struct {
	ID      string
	PostID  string
	UserID  string
	Content string
	Like    []string
	Dislike []string
	Date    string
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
	_, err := collection.InsertOne(ctx, c)

	return err
}

func DeleteComment(ctx context.Context, commentID string) error {
	_, err := collection.DeleteOne(ctx, bson.M{"commentid": commentID})

	return err
}
