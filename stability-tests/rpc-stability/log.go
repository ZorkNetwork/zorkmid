package main

import (
	"github.com/ZorkNetwork/zorkmid/infrastructure/logger"
	"github.com/ZorkNetwork/zorkmid/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
