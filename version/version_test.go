package version

import (
	"testing"
)

func init() {
	appBuild = "Invalid.Name"
}

func TestCheckAppBuild(t *testing.T) {
	appBuild = "correctName"
	checkAppBuild(appBuild)
}

func TestCheckAppBuildInvalidChars(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the code did not panic as expected")
		}
	}()

	appBuild = "Invalid.Name"
	checkAppBuild(appBuild)
}
