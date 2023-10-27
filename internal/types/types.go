// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	ID        int64  `json:"id"`
	Phone     string `json:"phone"`
	NickName  string `json:"nick_name"`
	Sex       int64  `json:"sex"`
	AvatarUrl string `json:"avatar_url"`
	Info      string `json:"info"`
}

type RegisterReq struct {
	Phone    string `json:"phone" validate:"e164,required"`
	Password string `json:"password" validate:"required"`
	NickName string `json:"nick_name,option"`
}

type RegisterResp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo User `json:"user_info"`
}

type VideoUserInfo struct {
	NickName  string `json:"nickName"`  // 昵称
	AvatarUrl string `json:"avatarUrl"` // 头像地址
}

type Video struct {
	ID            int64         `json:"id"`         // 视频ID
	Title         string        `json:"title"`      // 视频标题
	PlayUrl       string        `json:"playUrl"`    // 视频播放地址
	ThumbUrl      string        `json:"thumbUrl"`   // 视频封面地址
	FavoriteCount int64         `json:"favNum"`     // 点赞数
	CommentCount  int64         `json:"commentNum"` // 评论数
	ShareNum      int64         `json:"shareNum"`   // 分享数
	CreatedAt     string        `json:"createdAt"`  // 视频创建时间(毫秒时间戳)
	IsFav         bool          `json:"isFav"`      // 当前用户是否已点赞
	IsFollow      bool          `json:"isFollow"`   // 当前用户是否已关注该用户
	Tags          []string      `json:"tags"`       // 视频标签
	Author        VideoUserInfo `json:"author"`     // 作者信息
}

type FeedReq struct {
	LatestTime int64  `json:"latest_time,optional" form:"latestTime,optional"`  // 最新视频时间(毫秒时间戳)
	Limit      int    `json:"limit,optional" form:"limit,optional"`             // 请求数量
	Token      string `json:"token,optional" form:"token,optional"`             // 登录token
	AuthorID   uint   `json:"author_id,optional" form:"authorID,optional"`      // 作者ID
	Tag        string `json:"tag,optional" form:"tag,optional"`                 // 标签
	CategoryID uint   `json:"category_id,optional" form:"category_id,optional"` // 分类
}

type FeedResp struct {
	VideoList []*Video `json:"video_list"`
	NextTime  int64    `json:"next_time"` // 下次请求时间(毫秒时间戳)
	Length    int      `json:"length"`    // 视频列表长度
}

type CategoryReq struct {
}

type Category struct {
	ID   uint   `json:"id"`   // 分类ID
	Name string `json:"name"` // 分类名称
}

type CategoryResp struct {
	CategoryList []*Category `json:"category_list"`
}

type FollowUser struct {
	UserId   string `json:"user_id" desc:"用户id"`
	UserName string `json:"user_name" desc:"用户名"`
	Avatar   string `json:"avatar" desc:"头像"`
}

type FollowReq struct {
	ContextUserId uint `json:"user_id" path:"user_id" desc:"关注用户id" validate:"required"`
}

type FollowResp struct {
}

type UnFollowReq struct {
	ContextUserId uint `json:"user_id" path:"user_id" desc:"关注用户id" validate:"required"`
}

type UnFollowResp struct {
}

type CheckMyFollowingReq struct {
	ContextUserId uint `json:"user_id" path:"user_id" desc:"关注用户id" validate:"required"`
}

type CheckMyFollowingResp struct {
	Status bool `json:"status" desc:"是否关注"`
}

type ListMyFollowersReq struct {
	Page  int `json:"page" desc:"页码" validate:"required"`
	Limit int `json:"limit" desc:"每页数量" validate:"required"`
}

type ListMyFollowersResp struct {
	FollowUser []*FollowUser `json:"followers" desc:"用户id"`
}

type ListMyFollowingReq struct {
	Page  int `json:"page" desc:"页码" validate:"required"`
	Limit int `json:"limit" desc:"每页数量" validate:"required"`
}

type ListMyFollowingResp struct {
	FollowUser []*FollowUser `json:"followers" desc:"用户id"`
}

type ListFollowersReq struct {
	ContextUserId uint `json:"user_id" path:"user_id" desc:"用户id" validate:"required"`
	Page          int  `json:"page" desc:"页码" validate:"required"`
	Limit         int  `json:"limit" desc:"每页数量" validate:"required"`
}

type ListFollowersResp struct {
	FollowUser []*FollowUser `json:"followers" desc:"用户id"`
}

type ListFollowingReq struct {
	ContextUserId uint `json:"user_id" path:"user_id" desc:"用户id" validate:"required"`
	Page          int  `json:"page" desc:"页码" validate:"required"`
	Limit         int  `json:"limit" desc:"每页数量" validate:"required"`
}

type ListFollowingResp struct {
	FollowUser []*FollowUser `json:"followers" desc:"用户id"`
}

type CheckFollowingReq struct {
	DoerUserId    uint `json:"doer_user_id" path:"doer_user_id" desc:"用户id" validate:"required"`
	ContextUserId uint `json:"context_user_id" path:"doer_user_id" desc:"用户id" validate:"required"`
}

type CheckFollowingResp struct {
	Status bool `json:"status" desc:"是否关注"`
}

type CountFollowReq struct {
	ContextUserId uint `json:"user_id" path:"user_id" desc:"用户id" validate:"required"`
}

type CountFollowResp struct {
	FollowingCount int64 `json:"following_count" desc:"关注数量"`
	FollowersCount int64 `json:"follower_count" desc:"粉丝数量"`
}
