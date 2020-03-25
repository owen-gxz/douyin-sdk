package service

import (
	"github.com/owen-gxz/douyin-sdk/corp"
	"github.com/owen-gxz/douyin-sdk/im"
	"github.com/owen-gxz/douyin-sdk/oauth"
	"github.com/owen-gxz/douyin-sdk/resource"
	"github.com/owen-gxz/douyin-sdk/user"
)

// 获取openid的token
func (s *Service) GetAccessToken(openid string) (string, error) {
	return s.tokenService.GetToken(openid)
}

// 设置openid的token
func (s *Service) SetAccessToken(response oauth.TokenResponse) error {
	return s.tokenService.SaveToken(response)
}

// 企业 start
// 获取意向用户列表
func (s *Service) LeadsUserList(openid string, cursor, count int, startTime, endTime int64, leadsLevel, actionType int) (*corp.LeadsUserListResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return corp.LeadsUserList(token, openid, cursor, count, startTime, endTime, leadsLevel, actionType)
}

// 获取意向用户详情
func (s *Service) LeadsUserDetail(openid, userID string) (*corp.LeadsUserDetailResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return corp.LeadsUserDetail(token, openid, userID)
}

// 获取意向用户互动记录
func (s *Service) LeadsUserActionList(openid string, cursor, count int, userID string, actionType int) (*corp.LeadsUserActionListResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return corp.LeadsUserActionList(token, openid, cursor, count, userID, actionType)
}

//获取标签列表
func (s *Service) LeadsTagList(openid string, cursor, count int64) (*corp.LeadsTagListResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return corp.LeadsTagList(token, openid, cursor, count)

}

//获取标签用户列表
func (s *Service) LeadsTagUserList(openid string, cursor, count, tagID int64) (*corp.LeadsTagUserListResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return corp.LeadsTagUserList(token, openid, cursor, count, tagID)

}

//创建标签
func (s *Service) LeadsTagCreate(openid, name string) (*corp.LeadsTagCreateResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return corp.LeadsTagCreate(token, openid, name)
}

//更新标签
func (s *Service) LeadsTagUpdate(openid string, name string, id int64) (*corp.LeadsTagUpdateResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return corp.LeadsTagUpdate(token, openid, name, id)
}

// 删除标签
func (s *Service) LeadsTagDelete(openid string, id int64) (*corp.LeadsTagDeleteResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return corp.LeadsTagDelete(token, openid, id)
}

// 更新用户标签
func (s *Service) LeadsTagUserUpdate(openid string, tag corp.LeadsTagUserUpdateRequest) (*corp.LeadsTagUserUpdateResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return corp.LeadsTagUserUpdate(token, openid, tag)
}

// 企业end

// 素材上传
func (s *Service) MediaUpload(openid string, fileData []byte) (*im.MediaUploadResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return im.MediaUpload(token, openid, fileData)
}

//发送图片消息
func (s *Service) SendImageMessage(openid, toUser string, fileData []byte) (error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return err
	}
	return im.SendImageMessage(token, openid, toUser, fileData)
}

//发送私信消息
func (s *Service) SendMessage(openid string, reply im.MessageReq) (*im.SendMessageResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return im.SendMessage(token, openid, reply)
}

//查看视频评论列表
func (s *Service) GetVideoCommentList(openid string, cursor, count int, itemID string) (*resource.VideoCommentListResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return resource.GetVideoCommentList(token, openid, cursor, count, itemID)
}

// 查看评论回复列表
func (s *Service) GetVideoCommentReplyList(openid string, cursor, count int, itemID, commentID string) (*resource.VideoCommentReplyListResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return resource.GetVideoCommentReplyList(token, openid, cursor, count, itemID, commentID)
}

//回复评论
func (s *Service) VideoCommentReply(openid string, reply resource.ComentReq) (*resource.VideoCommentReplyResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return resource.VideoCommentReply(token, openid, reply)
}

// 置顶视频评论(企业号)
func (s *Service) VideoCommentTop(openid string, reply resource.VideoCommentTopRequest) (*resource.VideoCommentTopResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return resource.VideoCommentTop(token, openid, reply)
}

// 创建图片素材并发布到抖音
func (s *Service) Image2DouYin(openid string, fileData []byte, title string, ats []string) (*resource.ResourceCreateResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return resource.Image2DouYin(token, openid, fileData, title, ats)
}

// 创建视频素材并发布到抖音
func (s *Service) Video2DouYin(openid string, fileData []byte, title string, ats []string) (*resource.ResourceCreateResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return resource.Video2DouYin(token, openid, fileData, title, ats)
}

// 查询视频列表
func (s *Service) GetVideoList(openid string, cursor, count int) (*resource.VideoListResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return resource.GetVideoList(token, openid, cursor, count)
}

// 查询多个视频详情
func (s *Service) GetVideosInfo(openid string, itemIDs []string) (*resource.VideoInfoResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return resource.GetVideosInfo(token, openid, itemIDs)
}

// 删除视频
func (s *Service) RemoveVideo(openid string, itemID string) (*resource.RemoveVideoResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return resource.RemoveVideo(token, openid, itemID)
}

// 分片上传视频
func (s *Service) VideoPart(openid string, fileData []byte) (*resource.VideoPartCompleteResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return resource.VideoPart(token, openid, fileData)
}

// 获取粉丝列表
func (s *Service) GetFans(openid string, cursor, count int) (*user.FansResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return user.GetFans(token, openid, cursor, count)
}

//该接口用于查询用户的粉丝数据，如性别分布，年龄分布，地域分布等。
//注：用户首次授权应用后，需要间隔2天才会产生全部的数据。
func (s *Service) GetFansData(openid string) (*user.FansDataResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return user.GetFansData(token, openid)
}

// 获取粉丝列表
func (s *Service) GetFollowing(openid string, cursor, count int) (*user.FollowingResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return user.GetFollowing(token, openid, cursor, count)
}

// 获取粉丝列表
func (s *Service) GetUserInfo(openid string) (*user.UserInfoResponse, error) {
	token, err := s.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	return user.GetUserInfo(token, openid)
}
