package utils

import "testing"

func TestGetDateTime(t *testing.T) {
	dt := GetDateTime()
	t.Logf("Current date and time: %s", dt)
}
