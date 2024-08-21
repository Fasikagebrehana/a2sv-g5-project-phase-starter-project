// Usecases/blog_usecases.go
package usecases

import (
	"meleket/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogUsecase struct {
	blogRepo domain.BlogRepositoryInterface
}

func NewBlogUsecase(br domain.BlogRepositoryInterface) *BlogUsecase {
	return &BlogUsecase{blogRepo: br}
}

// CreateBlogPost creates a new blog post
func (u *BlogUsecase) CreateBlogPost(blog *domain.BlogPost) (interface{}, error) {
	id,err := u.blogRepo.Save(blog)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// GetAllBlogPosts retrieves all blog posts
func (u *BlogUsecase) GetAllBlogPosts() ([]domain.BlogPost, error) {
	blogs, err := u.blogRepo.GetAllBlog()
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

// GetBlogByID retrieves a blog post by its ID
func (u *BlogUsecase) GetBlogByID(id primitive.ObjectID) (*domain.BlogPost, error) {
	blog, err := u.blogRepo.GetBlogByID(id)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

// UpdateBlogPost updates an existing blog post
func (u *BlogUsecase) UpdateBlogPost(id primitive.ObjectID, blog *domain.BlogPost) (*domain.BlogPost, error) {
	updatedBlog, err := u.blogRepo.Update(blog)
	if err != nil {
		return nil, err
	}
	return updatedBlog, nil
}

// SearchBlogPosts searches for blog posts based on search query
// func (u *BlogUsecase) SearchBlogPosts(query *domain.SearchBlogPost) ([]domain.BlogPost, error) {
// 	blogs, err := u.blogRepo.Search(query.Title)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return blogs, nil
// }

// DeleteBlogPost deletes a blog post by its ID
func (u *BlogUsecase) DeleteBlogPost(id primitive.ObjectID) error {
	err := u.blogRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
