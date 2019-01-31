package tests

import (
	"testing"

	"github.com/hbl-duytv/intern-csm/services"
)

func TestGetPostActiveLimit(t *testing.T) {
	dataTest := []int{0, 1, 2, 3, 4, 5, 6}
	dataCorrect := []int{5, 5, 5, 5, 5, 5, 1}
	for _, value := range dataTest {
		_, resultTest, err := services.GetPostActiveLimit(value)
		if resultTest != dataCorrect[value] || err != nil {
			t.Errorf("Error result test : %d, data correct : %d", resultTest, dataCorrect[value])
			return
		}
	}
}
func TestGetTotalNumberPost(t *testing.T) {
	dataCorrect := 26
	resultTest, err := services.GetTotalNumberPost()
	if resultTest != dataCorrect || err != nil {
		t.Errorf("Error result test : %d, data correct : %d", resultTest, dataCorrect)
		return
	}
}
func TestGetNumberTestByMonth(t *testing.T) {
	dataTest := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	dataCorrect := []int{14, 2, 1, 1, 1, 5, 2, 1, 1, 1, 1, 2}
	for _, value := range dataTest {
		resultTest, err := services.GetNumberPostByMonth(value)
		if resultTest != dataCorrect[value-1] || err != nil {
			t.Errorf("Error result test : %d, data correct : %d", resultTest, dataCorrect[value-1])
			return
		}
	}
}
