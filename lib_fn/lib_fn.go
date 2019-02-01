package lib_fn

import (
	"time"

	"github.com/hbl-duytv/intern-csm/models"
)

func ParseTime(src string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", src)
	return t
}
func CompareTwoArrComment(arr1, arr2 []models.Comment) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v.ID != arr2[i].ID || v.PostID != arr2[i].PostID ||
			v.CommentatorID != arr2[i].CommentatorID || v.Message != arr2[i].Message || !CompareTime(v.CreatedAt, arr2[i].CreatedAt) || !CompareTime(v.UpdatedAt, arr2[i].UpdatedAt) {
			return false
		}
	}
	return true

}
func CompareTime(t1, t2 time.Time) bool {
	if t1.Year() != t2.Year() || t1.Month() != t2.Month() || t1.Day() != t2.Day() ||
		t1.Hour() != t2.Hour() || t1.Minute() != t2.Minute() || t1.Second() != t2.Second() {
		return false
	}
	return true
}
func CompareTwoArrPost(arr1, arr2 []models.TransformPost) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v.ID != arr2[i].ID || v.Creator != arr2[i].Creator ||
			v.Title != arr2[i].Title || v.Topic != arr2[i].Topic ||
			v.Description != arr2[i].Description || v.Content != arr2[i].Content || v.Status != arr2[i].Status ||
			!CompareTime(v.CreatedAt, arr2[i].CreatedAt) || !CompareTime(v.UpdatedAt, arr2[i].UpdatedAt) {
			return false
		}
	}
	return true

}
func CompareTwoPost(post1, post2 models.Post) bool {
	if post1.ID != post2.ID || post1.Creator != post2.Creator ||
		post1.Title != post2.Title || post1.Topic != post2.Topic ||
		post1.Description != post2.Description || post1.Content != post2.Content || post1.Status != post2.Status ||
		!CompareTime(post1.CreatedAt, post2.CreatedAt) || !CompareTime(post1.UpdatedAt, post2.UpdatedAt) {
		return false
	}
	return true
}
func CompareTwoArrUser(arr1, arr2 []models.User) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i, v := range arr1 {
		if v.ID != arr2[i].ID || v.Username != arr2[i].Username || v.Password != arr2[i].Password ||
			v.Email != arr2[i].Email || v.Name != arr2[i].Name || v.Gender != arr2[i].Gender ||
			v.PhoneNumber != arr2[i].PhoneNumber || v.Type != arr2[i].Type ||
			v.Status != arr2[i].Status || v.Token != arr2[i].Token ||
			v.Confirm != arr2[i].Confirm || !CompareTime(v.CreatedAt, arr2[i].CreatedAt) ||
			!CompareTime(v.UpdatedAt, arr2[i].UpdatedAt) {
			return false
		}
	}
	return true

}
func CompareTwoUser(user1, user2 models.User) bool {
	if user1.ID != user2.ID || user1.Username != user2.Username || user1.Password != user2.Password ||
		user1.Email != user2.Email || user1.Name != user2.Name || user1.Gender != user2.Gender ||
		user1.PhoneNumber != user2.PhoneNumber || user1.Type != user2.Type ||
		user1.Status != user2.Status || user1.Token != user2.Token ||
		user1.Confirm != user2.Confirm || !CompareTime(user1.CreatedAt, user2.CreatedAt) ||
		!CompareTime(user1.UpdatedAt, user2.UpdatedAt) {
		return false
	}
	return true
}
