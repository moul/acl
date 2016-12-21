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
		
                
                
		gettoken: grpctransport.NewServer(
			ctx,
			endpoints.GetTokenEndpoint,
			decodeGetTokenRequest,
			encodeGetTokenResponse,
			options...,
		),
                
		
                
                
                
		addtoken: grpctransport.NewServer(
			ctx,
			endpoints.AddTokenEndpoint,
			decodeAddTokenRequest,
			encodeAddTokenResponse,
			options...,
		),
                
		
                
                
                
		hasperm: grpctransport.NewServer(
			ctx,
			endpoints.HasPermEndpoint,
			decodeHasPermRequest,
			encodeHasPermResponse,
			options...,
		),
                
		
                
	}
}

type grpcServer struct {
	
	gettoken grpctransport.Handler
	
	addtoken grpctransport.Handler
	
	hasperm grpctransport.Handler
	
}


func (s *grpcServer) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	_, rep, err := s.gettoken.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTokenResponse), nil
}

func decodeGetTokenRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeGetTokenResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetTokenResponse)
	return resp, nil
}

func (s *grpcServer) AddToken(ctx context.Context, req *pb.AddTokenRequest) (*pb.AddTokenResponse, error) {
	_, rep, err := s.addtoken.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddTokenResponse), nil
}

func decodeAddTokenRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeAddTokenResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.AddTokenResponse)
	return resp, nil
}

func (s *grpcServer) HasPerm(ctx context.Context, req *pb.HasPermRequest) (*pb.HasPermResponse, error) {
	_, rep, err := s.hasperm.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HasPermResponse), nil
}

func decodeHasPermRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeHasPermResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.HasPermResponse)
	return resp, nil
}

