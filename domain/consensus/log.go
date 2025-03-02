package consensus

import (
	"github.com/ZorkNetwork/zorkmid/infrastructure/logger"
	"github.com/ZorkNetwork/zorkmid/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
