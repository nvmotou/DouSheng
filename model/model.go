package model

import (
	"time"

	"gorm.io/gorm"
)

// Model 数据库关系实体的基类
type Model struct {
	ID       uint64    `gorm:"comment:自增主键"`
	CreateAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
	UpdateAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
}

// Video 视频：数据库实体
type Video struct {
	Model
	gorm.DeletedAt
	VideoID       int64  `gorm:"type:BIGINT;not null;UNIQUE"`
	VideoName     string `gorm:"type:varchar(100);not null"`
	UserID        int64  `gorm:"type:BIGINT;not null;index:idx_author_id"`
	FavoriteCount int32  `gorm:"type:INT;not null;default:0"`
	CommentCount  int32  `gorm:"type:INT;not null;default:0"`
	PlayURL       string `gorm:"type:varchar(200);not null"`
	CoverURL      string `gorm:"type:varchar(200);not null"`
}

// User 用户:数据库实体
type User struct {
	Model
	UserID        int64  `gorm:"type:bigint;unsigned;not null;unique;uniqueIndex:idx_user_id" json:"user_id"`
	UserName      string `gorm:"type:varchar(50);not null;unique;uniqueIndex:idx_user_name" json:"name" validate:"min=6,max=32"`
	PassWord      string `gorm:"type:varchar(50);not null" json:"password" validate:"min=6,max=32"`
	FollowCount   int64  `gorm:"type:bigint;unsigned;not null;default:0" json:"follow_count"`
	FollowerCount int64  `gorm:"type:bigint;unsigned;not null;default:0" json:"follower_count"`
}

// Comment 评论：数据库实体
type Comment struct {
	Model
	UserID  int64  `gorm:"type:BIGINT;not null;index:idx_user_id;评论用户ID" json:"user_id"`
	VideoID int64  `gorm:"type:BIGINT;not null;index:idx_video_id;comment:被评论视频ID" json:"video_id"`
	Content string `gorm:"type:varchar(300);not null;comment:评论内容" json:"content"`
}

// Favourite 点赞：数据库实体
type Favourite struct {
	Model
	UserID  int64 `gorm:"type:BIGINT;not nul;index:idx_user_id;comment:点赞用户ID"`
	VideoID int64 `gorm:"type:BIGINT;not null;index:idx_video_id;comment:被点赞视频ID"`
}

// Follow 关注：数据库实体
type Follow struct {
	Model
	FromUserID int64 `gorm:"type:BIGINT;not null;index:idx_user_id;comment:粉丝用户ID"`
	ToUserID   int64 `gorm:"type:BIGINT;not null;index:idx_to_user_id;comment:被关注用户ID"`
}
