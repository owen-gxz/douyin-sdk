package oauth

import "strings"

const (
	// 用户想关
	ScopeUserInfo      = "user_info"      // 获取用户公开信息
	ScopeFansList      = "fans.list"      //粉丝列表
	ScopeFollowingList = "following.list" //关注列表
	ScopeFansData      = "fans.data"      //查询创作者粉丝数据
	// 视频相关
	ScopeVideoCreate = "video.create" //上传视频到文件服务器 - 创建抖音视频 -上传图片到文件服务器 - 发布图片
	ScopeVideoList   = "video.list"   //列出已发布的视频
	ScopeVideoData   = "video.data"   //查询指定视频数据
	ScopeAwemeshare  = "aweme.share"  //抖音分享id机制
	ScopeVideoDelete = "video.delete" //删除抖音视频
	ScopeHotsearch   = "hotsearch"    //获取实时热点词 --获取热点词聚合的视频
	// 互动
	ScopeVideoComment = "video.comment" //评论列表 ---评论回复列表 ---回复视频评论 ---置顶视频评论(企业号)
	ScopeIm           = "im"            //给抖音用户发送消息  --- 上传素材
)

func GetUserScope() string {
	scopes := []string{ScopeUserInfo, ScopeFansList, ScopeFollowingList, ScopeFansData}
	return strings.Join(scopes, ",")
}

func GetVideoScope() string {
	scopes := []string{ScopeVideoCreate, ScopeVideoList, ScopeVideoData, ScopeAwemeshare, ScopeVideoDelete, ScopeHotsearch}
	return strings.Join(scopes, ",")
}

func GetInteractScope() string {
	scopes := []string{ScopeVideoComment, ScopeIm}
	return strings.Join(scopes, ",")
}

func GetAllScope() string {
	scopes := []string{GetInteractScope(), GetVideoScope(), GetUserScope()}
	return strings.Join(scopes, ",")
}
