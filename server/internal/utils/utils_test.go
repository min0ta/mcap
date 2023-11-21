package utils

import (
	"testing"
)

func TestXxx(t *testing.T) {
	fn := ParseParams("/users/:id/comments/:comment")
	res, err := fn("/users/432/comments/12")
	if err != nil {
		t.Error(err)
	}
	if res["id"] != "432" || res["comment"] != "12" {
		t.Error(res)
	}
}
