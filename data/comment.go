package data

import (
	"context"
	"errors"
	"time"

	"github.com/Smart-Pot/pkg/db"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrCommentNotFound = errors.New("comment not found")
	ErrVoteFailed      = errors.New("vote failed!")
)

type Comment struct {
	ID      string   `json:"id" `
	PostID  string   `json:"postId" validate:"required"`
	UserID  string   `json:"userId" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Like    []string `json:"like"`
	Deleted bool     `json:"deleted"`
	Date    string   `json:"-"`
}

func (c *Comment) Validate() error {
	v := validator.New()
	return v.Struct(c)
}

func findComments(ctx context.Context, filter interface{}, opts *options.FindOptions) ([]*Comment, error) {
	var results []*Comment

	cur, err := db.Collection().Find(ctx, filter, opts)
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

func DeleteUsersComments(ctx context.Context, userID string) error {
	filter := bson.M{"userid": userID}
	updatePost := bson.M{"$set": bson.M{"deleted": true}}
	_, err := db.Collection().UpdateMany(ctx, filter, updatePost)

	if err != nil {
		return err
	}

	return nil
}

// GetComments : Find every comment in database
func GetCommentsByPostID(ctx context.Context, postID string, pageNumber, pageSize int) ([]*Comment, error) {
	skip := int64((pageNumber - 1) * pageSize)
	limit := int64(pageSize)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	filter := bson.M{
		"postid":  postID,
		"deleted": false,
	}
	comments, err := findComments(ctx, filter, &opts)
	return comments, err
}

func GetCommentsByUserID(ctx context.Context, userID string, pageNumber, pageSize int) ([]*Comment, error) {
	skip := int64((pageNumber - 1) * pageSize)
	limit := int64(pageSize)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	filter := bson.M{
		"userid":  userID,
		"deleted": false,
	}
	comments, err := findComments(ctx, filter, &opts)
	return comments, err
}

// AddComment : Add comment to database
func AddComment(ctx context.Context, c Comment) error {
	c.Date = time.Now().UTC().String()
	c.ID = generateID()
	c.Like = []string{}
	_, err := db.Collection().InsertOne(ctx, c)

	return err
}

func Vote(ctx context.Context, userID string, commentID string) error {
	res := db.Collection().FindOne(ctx, bson.M{"id": commentID})
	var c Comment
	if err := res.Decode(&c); err != nil {
		return err
	}
	c.Like = updateLikes(userID, c.Like)
	filter := bson.M{"id": commentID}
	pushToArray := bson.M{"$set": bson.M{"like": c.Like}}
	result, err := db.Collection().UpdateOne(ctx, filter, pushToArray)
	if result.ModifiedCount <= 0 {
		return ErrVoteFailed
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
	filter := bson.M{"id": commentID}

	updatePost := bson.M{"$set": bson.M{"deleted": true}}

	res, err := db.Collection().UpdateOne(ctx, filter, updatePost)
	if err != nil {
		return err
	}

	if res.ModifiedCount <= 0 {
		return ErrCommentNotFound
	}

	return nil
}

func GetCommentByID(ctx context.Context, id string) (*Comment, error) {
	res := db.Collection().FindOne(ctx, bson.M{"id": id})
	var c Comment
	if err := res.Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func generateID() string {
	return uuid.NewString()
}
