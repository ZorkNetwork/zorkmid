package server

import (
	"context"

	"github.com/ZorkNetwork/zorkmid/cmd/wallet/daemon/pb"
	"github.com/ZorkNetwork/zorkmid/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
