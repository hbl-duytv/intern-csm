package services

import (
	"github.com/hbl-duytv/intern-csm/constant"
	"github.com/hbl-duytv/intern-csm/models"
)

func GetPostWithAdminPermission() ([]models.TransformPost, error) {
	var posts []models.TransformPost
	err := DB.Debug().Table("post").Select("post.id as id, user.name as creator, post.title as title, post.topic as topic, post.description as description, post.content as content, post.status as status, post.created_at as created_at, post.updated_at as updated_at").Joins("join user on user.id = post.creator").Find(&posts).Error
	if err != nil {
		return posts, err
	}
	return posts, nil
}

func GetPostWithEditorPermission(id int) ([]models.TransformPost, error) {
	var posts []models.TransformPost
	err := DB.Table("post").Select("post.id as id, user.name as creator, post.title as title, post.topic as topic, post.description as description, post.content as content, post.status as status, post.created_at as created_at, post.updated_at as updated_at").Joins("join user on user.id = post.creator where post.creator=?", id).Find(&posts).Error
	if err != nil {
		return posts, err
	}
	return posts, nil

}
func GetPostById(id int) (models.Post, error) {
	var post models.Post
	if err := DB.Find(&post, "id=?", id).Error; err != nil {
		return post, err
	}
	return post, nil
}
func GetAllPostActive() ([]models.TransformPost, error) {
	var posts []models.Post
	var transformPosts []models.TransformPost
	if err := DB.Where("status=?", constant.ActiveNumber).Find(&posts).Error; err != nil {
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
	if err := DB.Model(&models.Post{}).Where("status=?", constant.ActiveNumber).Count(&count).Error; err != nil {
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
func DeletePost(id int) error {
	if err := DB.Where("id=?", id).Delete(models.Post{}).Error; err != nil {
		return err
	}
	return nil
}
func UpdateContentPost(id int, title, topic, des, content string) error {
	var post models.Post
	if err := DB.Where("id=?", id).Find(&post).Error; err != nil {
		return err
	}
	post.Title = title
	post.Topic = topic
	post.Description = des
	post.Content = content
	if err := DB.Save(&post).Error; err != nil {
		return err
	}
	return nil
}
func ChangeStatusPostWithComment(idPost, idUser, status int, mess string) error {
	var post models.Post
	if err := DB.First(&post, idPost).Error; err != nil {
		return err
	}
	if err := DB.Model(&post).Update("status", status).Error; err != nil {
		return err
	}
	newComment := models.Comment{
		PostID:        idPost,
		CommentatorID: idUser,
		Message:       mess,
	}
	if err := DB.Save(&newComment).Error; err != nil {
		return err
	}
	return nil
}
func CreatePost(post *models.Post) error {
	if err := DB.Save(&post).Error; err != nil {
		return err
	}
	return nil

}
func GetPostActiveLimit(page int) ([]models.TransformPost, int, error) {
	var posts []models.Post
	var transformPosts []models.TransformPost
	if err := DB.Where("status=?", constant.ActiveNumber).Offset((page - 1) * constant.LimitPost).Limit(constant.LimitPost).Find(&posts).Error; err != nil {
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

// idCreator int, title, topic, des, content string
