package txscript

import (
	"github.com/ZorkNetwork/zorkmid/infrastructure/logger"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// set log level to trace, so that logClosures passed to log.Tracef are covered
	log.SetLevel(logger.LevelTrace)
	logger.InitLogStdout(logger.LevelTrace)

	os.Exit(m.Run())
}
