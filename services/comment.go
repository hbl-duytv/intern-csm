package services

import (
	"github.com/hbl-duytv/intern-csm/models"
)

func GetCommentsById(idPost int) ([]models.Comment, error) {
	var comments []models.Comment

	if err := DB.Find(&comments, "post_id=?", idPost).Error; err != nil {
		return comments, err
	}
	return comments, nil
}
