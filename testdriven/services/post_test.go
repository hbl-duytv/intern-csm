package testdriven

import (
	"testing"

	"github.com/hbl-duytv/intern-csm/libfn"
	"github.com/hbl-duytv/intern-csm/models"
	"github.com/hbl-duytv/intern-csm/services"
)

func TestGetPostWithAdminPermission(t *testing.T) {
	want := []models.TransformPost{
		{
			ID:          2,
			CreatorID:   "Tuấn Anh",
			Title:       "I believe every human has a finite number of heartbeats. I don't intend to waste any of mine.",
			Topic:       "We predict too much for the next year and yet far too little for the next ten.",
			Description: "We predict too much for the next year and yet far too little for the next ten.",
			Content:     "Never in all their history have men been able truly to conceive of the world as one: a single sphere, a globe, having the qualities of a globe, a round earth in which all the directions eventually meet, in which there is no center because every point, or none, is center — an equal earth which all men occupy as equals. The airman's earth, if free men make it, will be truly round: a globe in practice, not in theory.",
			Status:      1,
			CreatedAt:   libfn.ParseTime("2019-01-28 00:00:00"),
			UpdatedAt:   libfn.ParseTime("2019-01-28 00:00:00"),
			Tag:         "",
		},
		{
			ID:          3,
			CreatorID:   "vip1",
			Title:       "I believe every human has a finite number of heartbeats",
			Topic:       "We predict too much for the next year and yet far too little for the next ten.",
			Description: "We predict too much for the next year and yet far too little for the next ten.",
			Content:     "Never in all their history have men been able truly to conceive of the world as one: a single sphere, a globe, having the qualities of a globe, a round earth in which all the directions eventually meet, in which there is no center because every point, or none, is center — an equal earth which all men occupy as equals. The airman's earth, if free men make it, will be truly round: a globe in practice, not in theory.",
			Status:      1,
			CreatedAt:   libfn.ParseTime("2019-01-28 00:00:00"),
			UpdatedAt:   libfn.ParseTime("2019-01-28 00:00:00"),
			Tag:         "",
		},
	}
	input, _ := services.GetPostWithAdminPermission()
	if !libfn.CompareTwoArrPost(want, input) {
		t.Error("get post admin permission failed")
	}

}
func TestGetPostWithEditorPermission(t *testing.T) {
	want := []models.TransformPost{
		{
			ID:          3,
			CreatorID:   "vip1",
			Title:       "I believe every human has a finite number of heartbeats",
			Topic:       "We predict too much for the next year and yet far too little for the next ten.",
			Description: "We predict too much for the next year and yet far too little for the next ten.",
			Content:     "Never in all their history have men been able truly to conceive of the world as one: a single sphere, a globe, having the qualities of a globe, a round earth in which all the directions eventually meet, in which there is no center because every point, or none, is center — an equal earth which all men occupy as equals. The airman's earth, if free men make it, will be truly round: a globe in practice, not in theory.",
			Status:      1,
			CreatedAt:   libfn.ParseTime("2019-01-28 00:00:00"),
			UpdatedAt:   libfn.ParseTime("2019-01-28 00:00:00"),
			Tag:         "",
		},
	}
	input, _ := services.GetPostWithEditorPermission(40)
	if !libfn.CompareTwoArrPost(want, input) {
		t.Error("get post editor permission failed")
	}

}
func TestGetPostById(t *testing.T) {
	want := models.Post{
		ID:          3,
		CreatorID:   40,
		Title:       "I believe every human has a finite number of heartbeats",
		Topic:       "We predict too much for the next year and yet far too little for the next ten.",
		Description: "We predict too much for the next year and yet far too little for the next ten.",
		Content:     "Never in all their history have men been able truly to conceive of the world as one: a single sphere, a globe, having the qualities of a globe, a round earth in which all the directions eventually meet, in which there is no center because every point, or none, is center — an equal earth which all men occupy as equals. The airman's earth, if free men make it, will be truly round: a globe in practice, not in theory.",
		Status:      1,
		CreatedAt:   libfn.ParseTime("2019-01-28 00:00:00"),
		UpdatedAt:   libfn.ParseTime("2019-01-28 00:00:00"),
		Tag:         "",
	}
	input, _ := services.GetPostById(3)
	if !libfn.CompareTwoPost(want, input) {
		t.Error("get post by id failed")
	}

}
func TestCreatePost(t *testing.T) {
	input := services.CreatePost(40, "ahihi", "ahehe", "ahuhu", "hixhix", "a1,a2,a3,a4,a5")
	if input != nil {
		t.Error("create post failed")
	}
}
func TestDeletePost(t *testing.T) {
	input := services.DeletePost(17)
	if input != nil {
		t.Error("delete post failed")
	}
}
func TestUpdateContentPost(t *testing.T) {
	input := services.UpdateContentPost(16, "ahihihihihi", "ahehehehe", "ahuhuhuh", "hixixix", "a1,a2,a3,a4,a5")
	if input != nil {
		t.Error("update content post failed")
	}
}
func TestChangeStatusPostWithComment(t *testing.T) {
	input := services.ChangeStatusPostWithComment(16, 44, 1, "test change status")
	if input != nil {
		t.Error("change status post failed")
	}

}
