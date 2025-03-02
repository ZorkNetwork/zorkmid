package rpc

import (
	"github.com/ZorkNetwork/zorkmid/infrastructure/logger"
	"github.com/ZorkNetwork/zorkmid/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
