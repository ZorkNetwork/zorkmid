package rpcclient

import (
	"github.com/ZorkNetwork/zorkmid/infrastructure/logger"
	"github.com/ZorkNetwork/zorkmid/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
