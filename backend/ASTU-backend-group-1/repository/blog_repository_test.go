package repository

import (
	"astu-backend-g1/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	mongomocks "github.com/sv-tools/mongoifc/mocks/mockery"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var mockBlogs = []domain.Blog{
	{Id: "1", Title: "title 1", Content: "content 1", AuthorId: "author 1", Date: time.Now(), Tags: []string{"tag1"}},
	{Id: "2", Title: "title 2", Content: "content 2", AuthorId: "author 2", Date: time.Now(), Tags: []string{"tag2"}},
	{Id: "3", Title: "title 3", Content: "content 3", AuthorId: "author 3", Date: time.Now(), Tags: []string{"tag3"}},
	{Id: "4", Title: "title 4", Content: "content 4", AuthorId: "author 4", Date: time.Now(), Tags: []string{"tag4"}},
	{Id: "5", Title: "title 5", Content: "content 5", AuthorId: "author 5", Date: time.Now(), Tags: []string{"tag5"}},
	{Id: "6", Title: "title 6", Content: "content 6", AuthorId: "author 6", Date: time.Now(), Tags: []string{"tag6"}},
	{Id: "7", Title: "title 7", Content: "content 7", AuthorId: "author 7", Date: time.Now(), Tags: []string{"tag7"}},
	{Id: "8", Title: "title 8", Content: "content 8", AuthorId: "author 8", Date: time.Now(), Tags: []string{"tag8"}},
	{Id: "9", Title: "title 9", Content: "content 9", AuthorId: "author 9", Date: time.Now(), Tags: []string{"tag9"}},
	{Id: "10", Title: "title 10", Content: "content 10", AuthorId: "author 10", Date: time.Now(), Tags: []string{"tag10"}},
}

type BlogRespositoryTestSuite struct {
	suite.Suite
	coll           *mongomocks.Collection
	BlogRepository domain.BlogRepository
}

func (suite *BlogRespositoryTestSuite) SetupSuite() {
	suite.coll = &mongomocks.Collection{}
	suite.BlogRepository = NewBlogRepository(suite.coll)
}

func (suite *BlogRespositoryTestSuite) TearDownSuite() {
	suite.coll.AssertExpectations(suite.T())
}

func (suite *BlogRespositoryTestSuite) TestCreate() {
	assert := assert.New(suite.T())
	suite.coll.On("InsertOne", mock.Anything, mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{
		InsertedID: primitive.NewObjectID(),
	}, nil)
	result, err := suite.BlogRepository.CreateBlog(mockBlogs[0])
	assert.NoError(err)
	assert.Equal(result, mockBlogs[0])
}

func (suite *BlogRespositoryTestSuite) TestGet() {
	assert := assert.New(suite.T())
	suite.T().Parallel()
	suite.T().Run("Getting all Blogs", func(t *testing.T) {
		cur := &mongomocks.Cursor{}
		for i, Blog := range mockBlogs {
			cur.On("Next", mock.Anything).Return(i < len(mockBlogs)).Once()
			cur.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
				arg := args.Get(0).(*domain.Blog)
				*arg = Blog
			}).Return(nil).Once()
		}
		cur.On("Next", mock.Anything).Return(false).Once()
		suite.coll.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(cur, nil)
		defer cur.AssertExpectations(suite.T())
		result, err := suite.BlogRepository.GetBlog(domain.BlogFilterOption{})
		assert.NoError(err)
		assert.Equal(mockBlogs, result)
	})
	// todo this is for a single result
	suite.T().Run("Getting by Blogname", func(t *testing.T) {
		cur := &mongomocks.Cursor{}
		singleResult := &mongomocks.SingleResult{}
		singleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.Blog)
			*arg = mockBlogs[0]
		}).Return(nil)
		suite.coll.On("FindOne", mock.Anything, mock.Anything, mock.Anything).Return(singleResult)
		defer cur.AssertExpectations(suite.T())
		result, err := suite.BlogRepository.GetBlog(domain.BlogFilterOption{
			Filter: domain.BlogFilters{
				Title: mockBlogs[0].Title,
			},
		})
		assert.NoError(err)
		assert.Equal(mockBlogs, result[0])
	})
	suite.T().Run("Getting by Email", func(t *testing.T) {
		cur := &mongomocks.Cursor{}
		singleResult := &mongomocks.SingleResult{}
		singleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.Blog)
			*arg = mockBlogs[0]
		}).Return(nil)
		suite.coll.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(singleResult)
		defer cur.AssertExpectations(suite.T())
		result, err := suite.BlogRepository.GetBlog(domain.BlogFilterOption{
			Filter: domain.BlogFilters{
                Tags: mockBlogs[0].Tags,
            },
		})
		assert.NoError(err)
		assert.Equal(mockBlogs, result[0])
	})
	suite.T().Run("Getting by Id", func(t *testing.T) {
		cur := &mongomocks.Cursor{}
		singleResult := &mongomocks.SingleResult{}
		singleResult.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
			arg := args.Get(0).(*domain.Blog)
			*arg = mockBlogs[0]
		}).Return(nil)
		suite.coll.On("FindOne", mock.Anything, mock.Anything, mock.Anything).Return(singleResult)
		defer cur.AssertExpectations(suite.T())
		result, err := suite.BlogRepository.GetBlog(domain.BlogFilterOption{
			Filter: domain.BlogFilters{
                AuthorId: mockBlogs[0].AuthorId,
            },
		})
		assert.NoError(err)
		assert.Equal(mockBlogs, result[0])
	})
}

func TestBlogRepository(t *testing.T) {
	suite.Run(t, new(BlogRespositoryTestSuite))
}