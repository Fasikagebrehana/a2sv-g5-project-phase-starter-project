package usecase

import (
	domain "AAiT-backend-group-8/Domain"
	infrastructure "AAiT-backend-group-8/Infrastructure"
	repository "AAiT-backend-group-8/Repository"
	"errors"
	"log"
)

type CommentUseCase struct {
	repository     repository.CommentRepository
	infrastructure infrastructure.Infrastructure
	tokenService   domain.ITokenService
}

func NewCommentUseCase(commentRepository repository.CommentRepository, infrastructure infrastructure.Infrastructure, tokenService domain.ITokenService) *CommentUseCase {
	return &CommentUseCase{
		repository:     commentRepository,
		infrastructure: infrastructure,
		tokenService:   tokenService,
	}
}

func (uc *CommentUseCase) CreateComment(comment *domain.Comment, blogID string) error {
	comment.CreatedAt = uc.infrastructure.GetCurrentTime()
	primitiveID, err := uc.infrastructure.ConvertToPrimitiveObjectID(blogID)
	if err != nil {
		return errors.New("invalid blog id")
	}
	comment.BlogID = primitiveID
	err = uc.repository.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CommentUseCase) GetComments(blogID string) ([]domain.Comment, error) {
	primitiveID, err := uc.infrastructure.ConvertToPrimitiveObjectID(blogID)
	if err != nil {
		return nil, errors.New("invalid blog id")
	}
	comments, err := uc.repository.GetComments(primitiveID)
	if err != nil {
		return nil, err
	}
	return comments, nil

}

func (uc *CommentUseCase) DeleteComment(commentID string) error {
	primitiveID, err := uc.infrastructure.ConvertToPrimitiveObjectID(commentID)
	if err != nil {
		return errors.New("invalid comment id")
	}
	err = uc.repository.DeleteComment(primitiveID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CommentUseCase) UpdateComment(comment *domain.Comment, commentID string) error {
	primitive, err := uc.infrastructure.ConvertToPrimitiveObjectID(commentID)
	if err != nil {
		return errors.New("invalid comment id")
	}
	comment.Id = primitive
	err = uc.repository.UpdateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CommentUseCase) DeleteCommentsOfBlog(blogID string) error {
	primitiveID, err := uc.infrastructure.ConvertToPrimitiveObjectID(blogID)
	if err != nil {

		return errors.New("invalid blog id")
	}
	err = uc.repository.DeleteCommentsOfBlog(primitiveID)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}

func (uc *CommentUseCase) DecodeToken(tokenStr string) (string, string, error) {
	// Parse the token
	myMap, err := uc.tokenService.GetClaimsOfToken(tokenStr)
	if err != nil {
		return "", "", err
	}
	if myMap == nil {
		return "", "", errors.New("invalid token - from decode token")
	}

	// Assert the types of the values to string
	id, ok := myMap["id"].(string)
	if !ok {
		return "", "", errors.New("invalid token - id claim is not a string")
	}

	name, ok := myMap["name"].(string)
	if !ok {
		return "", "", errors.New("invalid token - name claim is not a string")
	}

	return id, name, nil
}
