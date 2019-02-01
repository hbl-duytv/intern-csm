package testdriven

import (
	"testing"

	"github.com/hbl-duytv/intern-csm/lib_fn"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func TestGetCommentById(t *testing.T) {

	want := []models.Comment{
		{
			ID:            2,
			PostID:        3,
			CommentatorID: 14,
			Message:       "khá ổn",
			CreatedAt:     lib_fn.ParseTime("2019-01-23 14:50:35"),
			UpdatedAt:     lib_fn.ParseTime("2019-01-23 14:50:35"),
		},
		{
			ID:            3,
			PostID:        3,
			CommentatorID: 14,
			Message:       "title chưa ổn",
			CreatedAt:     lib_fn.ParseTime("2019-01-23 14:55:35"),
			UpdatedAt:     lib_fn.ParseTime("2019-01-23 14:55:35"),
		},
	}
	input, _ := services.GetCommentsById(3)

	if !lib_fn.CompareTwoArrComment(want, input) {
		t.Error("get comment by post id failed")
	}

}
