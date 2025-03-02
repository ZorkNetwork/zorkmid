package consensusstatestore

import (
	"github.com/ZorkNetwork/zorkmid/domain/consensus/database/serialization"
	"github.com/ZorkNetwork/zorkmid/domain/consensus/model/externalapi"
	"google.golang.org/protobuf/proto"
)

func serializeOutpoint(outpoint *externalapi.DomainOutpoint) ([]byte, error) {
	return proto.Marshal(serialization.DomainOutpointToDbOutpoint(outpoint))
}

func serializeUTXOEntry(entry externalapi.UTXOEntry) ([]byte, error) {
	return proto.Marshal(serialization.UTXOEntryToDBUTXOEntry(entry))
}

func deserializeOutpoint(outpointBytes []byte) (*externalapi.DomainOutpoint, error) {
	dbOutpoint := &serialization.DbOutpoint{}
	err := proto.Unmarshal(outpointBytes, dbOutpoint)
	if err != nil {
		return nil, err
	}

	return serialization.DbOutpointToDomainOutpoint(dbOutpoint)
}

func deserializeUTXOEntry(entryBytes []byte) (externalapi.UTXOEntry, error) {
	dbEntry := &serialization.DbUtxoEntry{}
	err := proto.Unmarshal(entryBytes, dbEntry)
	if err != nil {
		return nil, err
	}
	return serialization.DBUTXOEntryToUTXOEntry(dbEntry)
}
