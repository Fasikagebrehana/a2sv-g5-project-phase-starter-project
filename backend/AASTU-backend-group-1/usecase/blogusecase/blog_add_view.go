package blogusecase

import (
	"blogs/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddView implements domain.BlogUsecase.
func (b *BlogUsecase) AddView(view []primitive.ObjectID, claim domain.LoginClaims) error {

	var views []*domain.View

	for _, v := range view {
		temp := domain.View{
			BlogID: v,
			User:   claim.Username,
		}
		views = append(views, &temp)

		err := b.BlogRepo.IncrmentBlogViews(v.Hex())
		if err != nil {
			return err
		}
	}

	return b.BlogRepo.AddView(views)

}
