package services

import (
	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/models"
)

func GetAllPostActive() ([]models.TransformPost, error) {
	var posts []models.Post
	var transformPosts []models.TransformPost
	if err := DB.Where("status=?", constant.UserActived).Find(&posts).Error; err != nil {
		return nil, err
	}
	for _, v := range posts {
		var user models.User
		DB.Select("username").Where("id=?", v.Creator).Find(&user)
		transformPosts = append(transformPosts, models.TransformPost{
			v.ID,
			user.Username,
			v.Title,
			v.Topic,
			v.Description,
			v.Content,
			v.Status,
			v.CreatedAt,
			v.UpdatedAt,
		})
	}
	return transformPosts, nil
}
func GetTotalNumberPost() (int, error) {
	count := 0
	if err := DB.Model(&models.Post{}).Where("status=?", constant.UserActived).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}
func GetPostByID(postID string) (models.Post, error) {
	var post models.Post
	if err := DB.Find(&post, postID).Error; err != nil {
		return post, err
	}
	return post, nil
}
func GetPostActiveLimit(page int) ([]models.TransformPost, int, error) {
	var posts []models.Post
	var transformPosts []models.TransformPost
	if err := DB.Where("status=?", constant.UserActived).Offset((page - 1) * constant.LimitPost).Limit(constant.LimitPost).Find(&posts).Error; err != nil {
		return nil, 0, err
	}
	totalPost := len(posts)
	for _, v := range posts {
		var user models.User
		DB.Select("username").Where("id=?", v.Creator).Find(&user)
		transformPosts = append(transformPosts, models.TransformPost{
			v.ID,
			user.Username,
			v.Title,
			v.Topic,
			v.Description,
			v.Content,
			v.Status,
			v.CreatedAt,
			v.UpdatedAt,
		})
	}
	return transformPosts, totalPost, nil
}
