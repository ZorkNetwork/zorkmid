package templatemanager

import (
	"github.com/ZorkNetwork/zorkmid/app/appmessage"
	"github.com/ZorkNetwork/zorkmid/domain/consensus/model/externalapi"
	"github.com/ZorkNetwork/zorkmid/domain/consensus/utils/pow"
	"sync"
)

var currentTemplate *externalapi.DomainBlock
var currentState *pow.State
var isSynced bool
var lock = &sync.Mutex{}

// Get returns the template to work on
func Get() (*externalapi.DomainBlock, *pow.State, bool) {
	lock.Lock()
	defer lock.Unlock()
	// Shallow copy the block so when the user replaces the header it won't affect the template here.
	if currentTemplate == nil {
		return nil, nil, false
	}
	block := *currentTemplate
	state := *currentState
	return &block, &state, isSynced
}

// Set sets the current template to work on
func Set(template *appmessage.GetBlockTemplateResponseMessage) error {
	block, err := appmessage.RPCBlockToDomainBlock(template.Block)
	if err != nil {
		return err
	}
	lock.Lock()
	defer lock.Unlock()
	currentTemplate = block
	currentState = pow.NewState(block.Header.ToMutable())
	isSynced = template.IsSynced
	return nil
}
