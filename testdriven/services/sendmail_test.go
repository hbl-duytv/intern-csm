package testdriven

import (
	"testing"

	"github.com/hbl-duytv/intern-csm/services"
)

func TestSendMail(t *testing.T) {
	if err := services.SendMail("nguyentrugduc248@gmail.com", "test driven send mail"); err != nil {
		t.Errorf("error send mail: %s", err)
	}
}
