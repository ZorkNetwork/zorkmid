package prefixmanager

import (
	"github.com/ZorkNetwork/zorkmid/infrastructure/logger"
	"github.com/ZorkNetwork/zorkmid/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
