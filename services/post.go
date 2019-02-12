package services

import (
	"time"

	"github.com/hbl-duytv/intern-csm/models"
)

func GetPostWithAdminPermission() ([]models.TransformPost, error) {
	var posts []models.TransformPost
	rows, err := DB.Table("post").Select("post.id, user.name, post.title, post.topic, post.description, post.content, post.status, post.created_at, post.updated_at").Joins("join user on user.id = post.creator_id").Rows()
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		var id int
		var name, title, topic, description, content string
		var status int
		var createdAt, updatedAt time.Time
		rows.Scan(&id, &name, &title, &topic, &description, &content, &status, &createdAt, &updatedAt)
		post := models.TransformPost{
			ID:          id,
			Creator:     name,
			Title:       title,
			Topic:       topic,
			Description: description,
			Content:     content,
			Status:      status,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostWithEditorPermission(id int) ([]models.TransformPost, error) {
	var posts []models.TransformPost
	rows, err := DB.Table("post").Select("post.id, user.name, post.title, post.topic, post.description, post.content, post.status, post.created_at, post.updated_at").Joins("join user on user.id = post.creator_id where post.creator_id=?", id).Rows()
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		var id int
		var name, title, topic, description, content string
		var status int
		var createdAt, updatedAt time.Time
		rows.Scan(&id, &name, &title, &topic, &description, &content, &status, &createdAt, &updatedAt)
		post := models.TransformPost{
			ID:          id,
			Creator:     name,
			Title:       title,
			Topic:       topic,
			Description: description,
			Content:     content,
			Status:      status,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}
		posts = append(posts, post)
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
