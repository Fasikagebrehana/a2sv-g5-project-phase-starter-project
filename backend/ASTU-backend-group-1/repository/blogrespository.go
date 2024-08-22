package repository

import (
	"astu-backend-g1/domain"
	"context"
	"errors"
	"time"

	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoBlogRepository struct {
	collection mongoifc.Collection
}

func NewBlogRepository(collection mongoifc.Collection) domain.BlogRepository {
	return &MongoBlogRepository{
		collection: collection,
	}
}

func CreateBlogQuery(b domain.Blog) bson.M {
	query := bson.M{}
	if b.Title != "" {
		query["title"] = b.Title
	}
	if b.Content != "" {
		query["content"] = b.Content
	}
	if b.Id != "" {
		id, err := IsValidObjectID(b.Id)
		if err == nil {
			query["_id"] = id
		}
	}
	if b.AuthorId != "" {
		id, err := IsValidObjectID(b.AuthorId)
		if err == nil {
			query["author_id"] = id
		}
	}
	query["date"] = b.Date
	query["tags"] = b.Tags
	query["likes"] = []string{}
	query["dislikes"] = []string{}
	query["comments"] = []string{}
	query["views"] = []string{}

	return query
}

func (r *MongoBlogRepository) CreateBlog(b domain.Blog) (domain.Blog, error) {
	b.Id = primitive.NewObjectID().Hex()
	b.Date = time.Now()

	create := CreateBlogQuery(b)

	_, err := r.collection.InsertOne(context.Background(), create)
	if err != nil {
		return domain.Blog{}, err
	}
	return b, nil
}

func (r *MongoBlogRepository) FindPopularBlog() ([]domain.Blog, error) {
	pipeline := mongo.Pipeline{
		{
			{Key: "$addFields", Value: bson.D{
				{Key: "likesCount", Value: bson.D{{"$size", bson.D{{"$ifNull", bson.A{"$likes", bson.A{}}}}}}},
				{Key: "dislikesCount", Value: bson.D{{"$size", bson.D{{"$ifNull", bson.A{"$dislikes", bson.A{}}}}}}},
				{Key: "viewsCount", Value: bson.D{{"$size", bson.D{{"$ifNull", bson.A{"$views", bson.A{}}}}}}},
				{Key: "commentsCount", Value: bson.D{{"$size", bson.D{{"$ifNull", bson.A{"$comments", bson.A{}}}}}}},
			}},
		},
		{
			{Key: "$sort", Value: bson.D{
				{Key: "likesCount", Value: -1},
				{Key: "viewsCount", Value: -1},
				{Key: "commentsCount", Value: -1},
				{Key: "dislikesCount", Value: 1},
			}},
		},
	}

	cursor, err := r.collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var blogs []domain.Blog
	for cursor.Next(context.TODO()) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}
func BuildBlogQueryAndOptions(filterOption domain.BlogFilterOption) bson.M {
	filter := bson.M{}
	findOptions := options.Find()

	if filterOption.Filter.BlogId != "" {
		id, err := IsValidObjectID(filterOption.Filter.BlogId)
		if err == nil {
			filter["_id"] = id
		}
	}

	if filterOption.Filter.Title != "" {
		filter["title"] = filterOption.Filter.Title
	}

	if filterOption.Filter.AuthorId != "" {
		id, err := IsValidObjectID(filterOption.Filter.AuthorId)
		if err == nil {
			filter["author_id"] = id
		}
	}

	if !filterOption.Filter.Date.IsZero() {
		filter["date"] = filterOption.Filter.Date
	}

	if len(filterOption.Filter.Tags) > 0 {
		filter["tags"] = bson.M{"$in": filterOption.Filter.Tags}
	}

	if filterOption.Pagination.PageSize > 0 {
		findOptions.SetLimit(int64(filterOption.Pagination.PageSize))
	}
	if filterOption.Pagination.Page > 0 {
		findOptions.SetSkip(int64((filterOption.Pagination.Page - 1) * filterOption.Pagination.PageSize))
	}

	return filter
}

func (r *MongoBlogRepository) GetBlog(opts domain.BlogFilterOption) ([]domain.Blog, error) {
	filter := BuildBlogQueryAndOptions(opts)

	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var blogs []domain.Blog
	for cursor.Next(context.Background()) {
		var blog domain.Blog
		if err := cursor.Decode(&blog); err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func UpdateBlogQuery(b domain.Blog) bson.M {
	update := bson.M{}
	if b.Title != "" {
		update["title"] = b.Title
	}
	if b.Content != "" {
		update["content"] = b.Content
	}
	if b.AuthorId != "" {

		id, err := IsValidObjectID(b.AuthorId)
		if err != nil {
		} else {
			update["author_id"] = id
		}
	}
	if len(b.Tags) > 0 {

		update["tags"] = b.Tags
	}
	if len(b.Views) > 0 {

		update["views"] = b.Views
	}
	if len(b.Likes) > 0 {

		update["likes"] = b.Likes
	}
	if len(b.Comments) > 0 {

		update["comments"] = b.Comments
	}
	if len(b.Dislikes) > 0 {

		update["dislikes"] = b.Dislikes
	}
	return update
}

func (r *MongoBlogRepository) UpdateBlog(strBlogId string, updateData domain.Blog) (domain.Blog, error) {
	blogId, err := IsValidObjectID(strBlogId)
	if err != nil {
		return domain.Blog{}, err
	}
	filter := bson.M{"_id": blogId}
	update := bson.M{"$set": UpdateBlogQuery(updateData)}

	result, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil || result.MatchedCount == 0 {
		return domain.Blog{}, errors.New("Failed to delete blog with ID" + strBlogId + ":" + err.Error())
	}

	return updateData, nil
}

func (r *MongoBlogRepository) DeleteBlog(blogId, authorId string) error {
	id, err := IsValidObjectID(blogId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	var blog domain.Blog
	err = r.collection.FindOne(context.Background(), filter).Decode(&blog)
	if err != nil {
		return err
	}
	if blog.AuthorId != authorId {
		return errors.New("unauthorized to delete this blog")
	}
	result, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil || result.DeletedCount == 0 {
		return errors.New("Failed to delete blog with ID" + blogId + ":" + err.Error())
	}
	return nil
}

func (r *MongoBlogRepository) LikeOrDislikeBlog(blogId, userId string, like int) error {

	id, err := IsValidObjectID(blogId)
	if err != nil {
		return err
	}
	uid, err := IsValidObjectID(userId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{}
	if like == 1 {
		result := bson.M{}
		dislikeFinder := bson.M{"_id": id, "dislikes": uid}
		err := r.collection.FindOne(context.TODO(), dislikeFinder).Decode(&result)
		if err == nil {
			_, err = r.collection.UpdateOne(context.TODO(), filter, bson.M{
				"$pull": bson.M{
					"dislikes": uid,
				},
			})
			if err != nil {
				return err
			}
		}
		update["$addToSet"] = bson.M{"likes": uid, "view": uid}
	} else if like == -1 {
		result := bson.M{}
		likeFinder := bson.M{"_id": id, "likes": uid}
		err := r.collection.FindOne(context.TODO(), likeFinder).Decode(&result)
		if err == nil {
			_, err = r.collection.UpdateOne(context.TODO(), filter, bson.M{
				"$pull": bson.M{
					"likes": uid,
				},
			})
			if err != nil {
				return err
			}
		}
		update["$addToSet"] = bson.M{"dislikes": uid, "view": uid}
	} else {
		update["$addToSet"] = bson.M{"view": uid}
	}

	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
func (r *MongoBlogRepository) GetBlogById(blogId string) (domain.Blog, error) {
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return domain.Blog{}, err
	}
	filter := bson.M{"_id": bid}
	var result domain.Blog
	err = r.collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return domain.Blog{}, err
	}
	return result, nil

}

func (r *MongoBlogRepository) GetCommentById(blogId, commentId string) (domain.Comment, error) {
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return domain.Comment{}, err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return domain.Comment{}, err
	}

	filter := bson.M{
		"_id": bid,
	}

	projection := bson.M{
		"comments": bson.M{"$elemMatch": bson.M{"comment_id": cid}},
	}

	var result struct {
		Comments []domain.Comment `bson:"comments"`
	}

	err = r.collection.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return domain.Comment{}, err
	}

	if len(result.Comments) == 0 {
	}

	return result.Comments[0], nil
}

func (r *MongoBlogRepository) LikeOrDislikeComment(blogId, commentId, userId string, like int) error {
	id, err := IsValidObjectID(blogId)
	if err != nil {
		return err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return err
	}
	uid, err := IsValidObjectID(userId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id, "comments.comment_id": cid}
	update := bson.M{}
	if like == 1 {
		result := bson.M{}
		likeFinder := bson.M{"_id": id, "comments.comment_id": cid, "comments.dislikes": uid}
		err := r.collection.FindOne(context.TODO(), likeFinder).Decode(&result)
		if err == nil {
			_, err = r.collection.UpdateOne(context.TODO(), filter, bson.M{
				"$pull": bson.M{
					"comments.$.dislikes": uid,
				},
			})
			if err != nil {
				return err
			}
		}
		update["$addToSet"] = bson.M{"comments.$.likes": uid, "comments.$.views": uid}
	} else if like == -1 {
		result := bson.M{}
		likeFinder := bson.M{"_id": id, "comments.comment_id": cid, "comments.likes": uid}
		err := r.collection.FindOne(context.TODO(), likeFinder).Decode(&result)
		if err == nil {
			_, err = r.collection.UpdateOne(context.TODO(), filter, bson.M{
				"$pull": bson.M{
					"comments.$.likes": uid,
				},
			})
			if err != nil {
				return err
			}
		}
		update["$addToSet"] = bson.M{"comments.$.dislikes": uid, "comments.$.views": uid}
	} else {
		update["$addToSet"] = bson.M{"comments.$.views": uid}
	}
	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func CreateCommentQuery(r domain.Comment) bson.M {

	query := bson.M{}
	query["comment_id"] = primitive.NewObjectID()
	if r.Content != "" {
		query["content"] = r.Content
	}

	if r.AuthorId != "" {
		id, err := IsValidObjectID(r.AuthorId)
		if err == nil {
			query["author_id"] = id
		}
	}

	query["likes"] = []string{}
	query["dislikes"] = []string{}
	query["replies"] = []string{}
	query["views"] = []string{}
	return query
}
func (r *MongoBlogRepository) AddComment(blogId string, comment domain.Comment) error {
	id, err := IsValidObjectID(blogId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	comm := CreateCommentQuery(comment)

	update := bson.M{"$push": bson.M{"comments": comm}}

	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func CreateReplyQuery(c domain.Reply) bson.M {

	query := bson.M{}
	query["reply_id"] = primitive.NewObjectID()
	if c.Content != "" {
		query["content"] = c.Content
	}

	if c.AuthorId != "" {
		id, err := IsValidObjectID(c.AuthorId)
		if err == nil {
			query["author_id"] = id
		}
	}

	query["likes"] = []string{}
	query["dislikes"] = []string{}
	query["replies"] = []string{}
	query["views"] = []string{}
	return query
}

func (r *MongoBlogRepository) GetAllComments(blogId string) ([]domain.Comment, error) {
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return []domain.Comment{}, err
	}
	filter := bson.M{"_id": bid}
	projection := bson.M{"comments": 1, "_id": 0}
	var result struct {
		Comments []domain.Comment `bson:"comments"`
	}
	err = r.collection.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return []domain.Comment{}, err
	}
	return result.Comments, nil

}

func (r *MongoBlogRepository) LikeOrDislikeReply(blogId, commentId, replyId, userId string, like int) error {
	id, err := IsValidObjectID(blogId)
	if err != nil {
		return err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return err
	}
	rid, err := IsValidObjectID(replyId)
	if err != nil {
		return err
	}
	uid, err := IsValidObjectID(userId)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id":                       id,
		"comments.comment_id":       cid,
		"comments.replies.reply_id": rid,
	}

	update := bson.M{
		"$addToSet": bson.M{"view": uid},
		"$pull":     bson.M{},
	}

	if like == 1 {
		update["$pull"] = bson.M{"comments.$[].replies.$[reply].dislikes": uid}
		update["$addToSet"] = bson.M{"comments.$[].replies.$[reply].likes": uid}
	} else if like == -1 {
		update["$pull"] = bson.M{"comments.$[].replies.$[reply].likes": uid}
		update["$addToSet"] = bson.M{"comments.$[].replies.$[reply].dislikes": uid}
	} else {
		update["$addToSet"] = bson.M{"comments.$[].replies.$[reply].views": uid}
	}

	updateOptions := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.M{"reply.reply_id": rid}},
	})

	_, err = r.collection.UpdateOne(context.Background(), filter, update, updateOptions)
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoBlogRepository) GetAllRepliesForComment(blogId, commentId string) ([]domain.Reply, error) {
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return nil, err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": bid, "comments.comment_id": cid}
	projection := bson.M{
		"comments.$": 1,
	}

	var result struct {
		Comments []struct {
			Replies []domain.Reply `bson:"replies"`
		} `bson:"comments"`
	}

	err = r.collection.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}

	if len(result.Comments) > 0 {
		return result.Comments[0].Replies, nil
	}

	return nil, errors.New("no comments found")
}
func (r *MongoBlogRepository) GetReplyById(blogId, commentId, replyId string) (domain.Reply, error) {
	bid, err := IsValidObjectID(blogId)
	if err != nil {
		return domain.Reply{}, err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return domain.Reply{}, err
	}

	filter := bson.M{
		"_id":                 bid,
		"comments.comment_id": cid,
	}

	projection := bson.M{
		"comments.$": 1,
	}

	var result struct {
		Comments []struct {
			CommentID string         `bson:"comment_id"`
			Replies   []domain.Reply `bson:"replies"`
		} `bson:"comments"`
	}

	err = r.collection.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return domain.Reply{}, err
	}

	if len(result.Comments) > 0 {
		for _, reply := range result.Comments[0].Replies {
			if reply.ReplyId == replyId {
				return reply, nil
			}
		}
	}

	return domain.Reply{}, errors.New("reply not found")

}
func (r *MongoBlogRepository) ReplyToComment(blogId, commentId string, reply domain.Reply) error {
	id, err := IsValidObjectID(blogId)
	if err != nil {
		return err
	}
	cid, err := IsValidObjectID(commentId)
	if err != nil {
		return err
	}
	comm := CreateReplyQuery(reply)
	filter := bson.M{"_id": id, "comments.comment_id": cid}
	update := bson.M{"$push": bson.M{"comments.$.replies": comm}}

	_, err = r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
func IsValidObjectID(id string) (primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return oid, nil
}
