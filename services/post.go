package services

import (
	"fmt"
	"time"

	"github.com/hbl-duytv/intern-csm/models"
)

func GetPostWithAdminPermission() ([]models.TransformPost, error) {
	var posts []models.Post
	var transformPosts []models.TransformPost
	if err := DB.Find(&posts).Error; err != nil {
		return transformPosts, err
	}
	for _, v := range posts {
		if user, err := GetUserByID((v.Creator)); err == nil {
			transformPosts = append(transformPosts, models.TransformPost{
				v.ID,
				user.Name,
				v.Title,
				v.Topic,
				v.Description,
				v.Content,
				v.Status,
				v.CreatedAt,
				v.UpdatedAt,
			})
		}
	}
	return transformPosts, nil
}
func GetPostWithEditorPermission(id int) ([]models.TransformPost, error) {
	var tranformPosts []models.TransformPost
	var posts []models.Post
	if err := DB.Find(&posts, "creator=?", id).Error; err != nil {
		fmt.Println(err)
		return tranformPosts, err
	}
	for _, v := range posts {
		if user, err := GetUserByID((v.Creator)); err == nil {
			tranformPosts = append(tranformPosts, models.TransformPost{
				v.ID,
				user.Name,
				v.Title,
				v.Topic,
				v.Description,
				v.Content,
				v.Status,
				v.CreatedAt,
				v.UpdatedAt,
			})
		}
	}
	return tranformPosts, nil
}
func GetPostById(id int) (models.Post, error) {
	var post models.Post
	if err := DB.Find(&post, "id=?", id).Error; err != nil {
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
func CreatePost(idCreator int, title, topic, des, content string) error {
	newPost := models.Post{
		Creator:     idCreator,
		Title:       title,
		Topic:       topic,
		Description: des,
		Content:     content,
		Status:      0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := DB.Save(&newPost).Error; err != nil {
		return err
	}
	return nil
}
