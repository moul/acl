package acl_grpctransport



import (
        "fmt"

	context "golang.org/x/net/context"
        pb "github.com/moul/acl/gen/pb"
        endpoints "github.com/moul/acl/gen/endpoints"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(ctx context.Context, endpoints endpoints.Endpoints) pb.AclServiceServer {
	options := []grpctransport.ServerOption{}
	return &grpcServer{
		
                
                
		hasperm: grpctransport.NewServer(
			ctx,
			endpoints.HaspermEndpoint,
			decodeHaspermRequest,
			encodeHaspermResponse,
			options...,
		),
                
		
                
	}
}

type grpcServer struct {
	
	hasperm grpctransport.Handler
	
}


func (s *grpcServer) Hasperm(ctx context.Context, req *pb.HasPermRequest) (*pb.HasPermResponse, error) {
	_, rep, err := s.hasperm.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HasPermResponse), nil
}

func decodeHaspermRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeHaspermResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.HasPermResponse)
	return resp, nil
}

