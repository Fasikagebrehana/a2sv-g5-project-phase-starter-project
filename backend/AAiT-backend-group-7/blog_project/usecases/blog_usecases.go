package usecases

import (
	"context"
	"errors"

	"blog_project/domain"
)

type BlogUsecases struct {
	BlogRepo    domain.IBlogRepository
	UserUsecase domain.IUserUsecase
}

func NewBlogUsecase(blogrepo domain.IBlogRepository, userusecase domain.IUserUsecase) domain.IBlogUsecase {
	return &BlogUsecases{
		BlogRepo:    blogrepo,
		UserUsecase: userusecase,
	}
}

func (u *BlogUsecases) GetAllBlogs(ctx context.Context) ([]domain.Blog, error) {
	return u.BlogRepo.GetAllBlogs(ctx)
}

func (u *BlogUsecases) GetBlogByID(ctx context.Context, id int) (domain.Blog, error) {
	return u.BlogRepo.GetBlogByID(ctx, id)
}

func (u *BlogUsecases) CreateBlog(ctx context.Context, blog domain.Blog) (domain.Blog, error) {
	// Generate a new unique ID based on the current time in nanoseconds
	blog.ID = generateUniqueID() // Convert nanoseconds to microseconds

	_, err := u.UserUsecase.AddBlog(ctx, blog.AuthorID, blog)
	if err != nil {
		return domain.Blog{}, err
	}

	return u.BlogRepo.CreateBlog(ctx, blog)
}

func (u *BlogUsecases) UpdateBlog(ctx context.Context, id int, updatedBlog domain.Blog) (domain.Blog, error) {
	// Retrieve the existing blog from the repository
	existingBlog, err := u.BlogRepo.GetBlogByID(ctx, id)
	if err != nil {
		return domain.Blog{}, err
	}

	// Update only the fields that are not empty or zero
	if updatedBlog.Title != "" {
		existingBlog.Title = updatedBlog.Title
	}
	if updatedBlog.AuthorID != 0 {
		existingBlog.AuthorID = updatedBlog.AuthorID
	}
	if updatedBlog.Content != "" {
		existingBlog.Content = updatedBlog.Content
	}
	if len(updatedBlog.Comments) > 0 {
		existingBlog.Comments = updatedBlog.Comments
	}
	if len(updatedBlog.Likes) > 0 {
		existingBlog.Likes = updatedBlog.Likes
	}
	if len(updatedBlog.Dislikes) > 0 {
		existingBlog.Dislikes = updatedBlog.Dislikes
	}
	if updatedBlog.Date != "" {
		existingBlog.Date = updatedBlog.Date
	}
	if len(updatedBlog.Tags) > 0 {
		existingBlog.Tags = updatedBlog.Tags
	}
	if updatedBlog.Views != 0 {
		existingBlog.Views = updatedBlog.Views
	}

	// Save the updated blog back to the repository
	return u.BlogRepo.UpdateBlog(ctx, id, existingBlog)
}

func (u *BlogUsecases) DeleteBlog(ctx context.Context, id int) error {
	blog, err := u.BlogRepo.GetBlogByID(ctx, id)
	if err != nil {
		return err
	}

	author := blog.AuthorID

	u.UserUsecase.DeleteBlog(ctx, author, id)

	return u.BlogRepo.DeleteBlog(ctx, id)
}

func (u *BlogUsecases) Search(ctx context.Context, author string, tags []string, title string) ([]domain.Blog, error) {
	var results []domain.Blog
	var tempResults []domain.Blog
	var err error

	// Use a map to track blog occurrences by ID for intersection logic
	blogMap := make(map[int]int)

	// Search by author
	if author != "" {
		tempResults, err = u.BlogRepo.SearchByAuthor(ctx, author)
		if err != nil {
			return nil, err
		}
		for _, blog := range tempResults {
			blogMap[blog.ID]++
		}
	}

	// Search by tags
	if len(tags) > 0 {
		tempResults, err = u.BlogRepo.SearchByTags(ctx, tags)
		if err != nil {
			return nil, err
		}
		for _, blog := range tempResults {
			blogMap[blog.ID]++
		}
	}

	// Search by title
	if title != "" {
		tempResults, err = u.BlogRepo.SearchByTitle(ctx, title)
		if err != nil {
			return nil, err
		}
		for _, blog := range tempResults {
			blogMap[blog.ID]++
		}
	}

	// Collect results that match all criteria
	criteriaCount := 0
	if author != "" {
		criteriaCount++
	}
	if len(tags) > 0 {
		criteriaCount++
	}
	if title != "" {
		criteriaCount++
	}

	for blogID, count := range blogMap {
		if count == criteriaCount {
			blog, err := u.BlogRepo.GetBlogByID(ctx, blogID)
			if err != nil {
				return nil, err
			}
			results = append(results, blog)
		}
	}

	return results, nil
}

func (u *BlogUsecases) LikeBlog(ctx context.Context, blogID int, authorID int) (domain.Blog, error) {
	blog, err := u.BlogRepo.GetBlogByID(ctx, blogID)
	if err != nil {
		return domain.Blog{}, err
	}

	// Check if the user has already liked the blog
	for _, like := range blog.Likes {
		if like.UserID == authorID {
			return domain.Blog{}, errors.New("user already liked this blog")
		}
	}

	// Add like
	newLike := domain.Like{
		ID:     len(blog.Likes) + 1,
		UserID: authorID,
		Date:   "current date",
	}
	blog.Likes = append(blog.Likes, newLike)

	_, err = u.BlogRepo.UpdateBlog(ctx, blogID, blog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (u *BlogUsecases) DislikeBlog(ctx context.Context, blogID int, authorID int) (domain.Blog, error) {
	blog, err := u.BlogRepo.GetBlogByID(ctx, blogID)
	if err != nil {
		return domain.Blog{}, err
	}

	// Check if the user has already disliked the blog
	for _, dislike := range blog.Dislikes {
		if dislike.UserID == authorID {
			return domain.Blog{}, errors.New("user already disliked this blog")
		}
	}

	// Add dislike
	newDislike := domain.Dislike{
		ID:     len(blog.Dislikes) + 1,
		UserID: authorID,
		Date:   "current date",
	}
	blog.Dislikes = append(blog.Dislikes, newDislike)

	_, err = u.BlogRepo.UpdateBlog(ctx, blogID, blog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}

func (u *BlogUsecases) AddComent(ctx context.Context, blogID int, authorID int, content string) (domain.Blog, error) {
	blog, err := u.BlogRepo.GetBlogByID(ctx, blogID)
	if err != nil {
		return domain.Blog{}, err
	}

	// Add comment
	newComment := domain.Comment{
		ID:      len(blog.Comments) + 1,
		UserID:  authorID,
		Content: content,
		Date:    "current date",
	}
	blog.Comments = append(blog.Comments, newComment)

	_, err = u.BlogRepo.UpdateBlog(ctx, blogID, blog)
	if err != nil {
		return domain.Blog{}, err
	}

	return blog, nil
}
