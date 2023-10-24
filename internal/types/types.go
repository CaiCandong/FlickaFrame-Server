// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	NickName  string `json:"nickName"`  // 昵称
	AvatarUrl string `json:"avatarUrl"` // 头像地址
}

type Video struct {
	Id         int64  `json:"id"`         // 视频ID
	PlayUrl    string `json:"playUrl"`    // 视频播放地址
	ThumbUrl   string `json:"thumbUrl"`   // 视频封面地址
	Context    string `json:"context"`    // 视频描述
	FavNum     int64  `json:"favNum"`     // 点赞数
	CommentNum int64  `json:"commentNum"` // 评论数
	ShareNum   int64  `json:"shareNum"`   // 分享数
	IsFav      bool   `json:"isFav"`      // 当前用户是否已点赞
	IsFollow   bool   `json:"isFollow"`   // 当前用户是否已关注该用户
	Author     User   `json:"author"`     // 作者信息
}

type FeedReq struct {
	LatestTime int    `json:"latest_time" binding:"option"` // 最新视频时间
	Token      string `json:"token" binding:"option"`       // 用户token
}

type FeedResp struct {
	VideoList []Video `json:"videoList"`
}
