package blogmodel

import (
	"errors"
	"time"

	"github.com/google/uuid"
	er "github.com/group13/blog/domain/errors"
)

const (
	minTitleLength   = 3
	maxTitleLength   = 100
	minContentLength = 50
	maxContentLength = 10000
)

// Blog represents the aggregate user with private fields.
type Blog struct {
	id           uuid.UUID
	title        string
	content      string
	tags         []string
	createdDate  time.Time
	updatedDate  time.Time
	userid       uuid.UUID
	likeCount    int
	disLikeCount int
	commentCount int
}

// Config holds parameters for creating a new Blog.

type Config struct {
	Title   string
	Content string
	Tags    []string
	UserId  uuid.UUID
}

// MapConfig holds parameters for mapping with Blog from Data Base.
type MapConfig struct {
	Id           uuid.UUID
	UserId       uuid.UUID
	Title        string
	Content      string
	Tags         []string
	CreatedDate  time.Time
	UpdatedDate  time.Time
	LikeCount    int
	DisLikeCount int
	CommentCount int
}

// New creates a new Blog with the provided configuration.
func New(config Config) (*Blog, error) {
	if err := validateTitle(config.Title); err != nil {
		return nil, err
	}

	if err := validateContent(config.Content); err != nil {
		return nil, err
	}

	//returns blog with specified fields
	return &Blog{
		id:           uuid.New(),
		userid:       config.UserId,
		title:        config.Title,
		content:      config.Content,
		tags:         config.Tags,
		createdDate:  time.Now(),
		updatedDate:  time.Now(),
		likeCount:    0,
		disLikeCount: 0,
		commentCount: 0,
	}, nil
}

// Map maps a Blog from database and returns Blog.
func Map(mapConfig MapConfig) (*Blog, error) {

	//returns blog with specified fields
	return &Blog{
		id:           mapConfig.Id,
		userid:       mapConfig.UserId,
		title:        mapConfig.Title,
		content:      mapConfig.Content,
		tags:         mapConfig.Tags,
		createdDate:  mapConfig.CreatedDate,
		updatedDate:  mapConfig.UpdatedDate,
		likeCount:    mapConfig.LikeCount,
		disLikeCount: mapConfig.DisLikeCount,
		commentCount: mapConfig.CommentCount,
	}, nil
}

// ID returns Blog Id.
func (b *Blog) ID() uuid.UUID {
	return b.id
}

// UserId returns Blog's Auther Id.
func (b *Blog) UserId() uuid.UUID {
	return b.userid
}

// Title returns Blog's title.
func (b *Blog) Title() string {
	return b.title
}

// Content returns Blog's Content
func (b *Blog) Content() string {
	return b.content
}

// Tags returns Blog's Tags
func (b *Blog) Tags() []string {
	return b.tags
}

// CreatedDate returns Blog's Created Date
func (b *Blog) CreatedDate() time.Time {
	return b.createdDate
}

// UpdatedDate returns Blog's UpdatedDate
func (b *Blog) UpdatedDate() time.Time {
	return b.updatedDate
}

// LikeCount returns Blog's Like Count
func (b *Blog) LikeCount() int {
	return b.likeCount
}

// DislikeCount returns Blog's DisLike Count
func (b *Blog) DisLikeCount() int {
	return b.disLikeCount
}

// Comment Count returns Blog's Comment Count
func (b *Blog) CommentCount() int {
	return b.commentCount
}

func validateTitle(title string) error {
	if len(title) < minTitleLength {
		return er.TitleTooShort
	}
	if len(title) > maxTitleLength {
		return er.TitleTooLong
	}

	return nil
}

func validateContent(content string) error {
	if len(content) < minContentLength {
		return er.ContentTooShort
	}
	if len(content) > maxContentLength {
		return er.ContentTooLong
	}

	return nil
}

// UpdateTitle updates the blog's Title after validating the data.
func (b *Blog) UpdateTitle(config Config) error {
	if err := validateTitle(config.Title); err != nil {
		return err
	}

	b.title = config.Title
	return nil
}

// UpdateContent updates the blog's content  after validating the data.
func (b *Blog) UpdateContent(config Config) error {

	if err := validateContent(config.Content); err != nil {
		return err
	}

	b.content = config.Content
	return nil
}

// UpdateTags updates the blog's tags.
func (b *Blog) UpdateTags(config Config) error {
	b.tags = config.Tags
	return nil
}

func (b *Blog) UpdateCommentCount(inc bool) error {
	if b.commentCount == 0 && inc {
		return errors.New("no comment count")
	}
	if inc {
		b.commentCount++
	} else {
		b.commentCount--
	}
	return nil
}
func (b *Blog) UpdateLikeCount(inc bool) error {
	if b.likeCount == 0 && inc {
		return errors.New("no likes count")
	}
	if inc {
		b.likeCount++
	} else {
		b.likeCount--
	}
	return nil
}
func (b *Blog) UpdateDisLikeCount(inc bool) error {
	if b.disLikeCount == 0 && inc {
		return errors.New("no comment count")
	}
	if inc {
		b.disLikeCount++
	} else {
		b.disLikeCount--
	}
	return nil
}
