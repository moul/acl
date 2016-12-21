package acl_endpoints



import (
	"fmt"

	context "golang.org/x/net/context"
        pb "github.com/moul/acl/gen/pb"
	"github.com/go-kit/kit/endpoint"
)

var _ = fmt.Errorf

type Endpoints struct {
	
	GetTokenEndpoint endpoint.Endpoint
	
	AddTokenEndpoint endpoint.Endpoint
	
	HasPermEndpoint endpoint.Endpoint
	
}


/*{
  "name": "GetToken",
  "input_type": ".acl.GetTokenRequest",
  "output_type": ".acl.GetTokenResponse",
  "options": {}
}*/

func (e *Endpoints)GetToken(ctx context.Context, in *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	out, err := e.GetTokenEndpoint(ctx, in)
	if err != nil {
		return &pb.GetTokenResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.GetTokenResponse), err
}

/*{
  "name": "AddToken",
  "input_type": ".acl.AddTokenRequest",
  "output_type": ".acl.AddTokenResponse",
  "options": {}
}*/

func (e *Endpoints)AddToken(ctx context.Context, in *pb.AddTokenRequest) (*pb.AddTokenResponse, error) {
	out, err := e.AddTokenEndpoint(ctx, in)
	if err != nil {
		return &pb.AddTokenResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.AddTokenResponse), err
}

/*{
  "name": "HasPerm",
  "input_type": ".acl.HasPermRequest",
  "output_type": ".acl.HasPermResponse",
  "options": {}
}*/

func (e *Endpoints)HasPerm(ctx context.Context, in *pb.HasPermRequest) (*pb.HasPermResponse, error) {
	out, err := e.HasPermEndpoint(ctx, in)
	if err != nil {
		return &pb.HasPermResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.HasPermResponse), err
}



func MakeGetTokenEndpoint(svc pb.AclServiceServer) endpoint.Endpoint {
     	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.GetTokenRequest)
		rep, err := svc.GetToken(ctx, req)
		if err != nil {
			return &pb.GetTokenResponse{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}

func MakeAddTokenEndpoint(svc pb.AclServiceServer) endpoint.Endpoint {
     	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.AddTokenRequest)
		rep, err := svc.AddToken(ctx, req)
		if err != nil {
			return &pb.AddTokenResponse{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}

func MakeHasPermEndpoint(svc pb.AclServiceServer) endpoint.Endpoint {
     	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.HasPermRequest)
		rep, err := svc.HasPerm(ctx, req)
		if err != nil {
			return &pb.HasPermResponse{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}


func MakeEndpoints(svc pb.AclServiceServer) Endpoints {
	return Endpoints{
		
		GetTokenEndpoint: MakeGetTokenEndpoint(svc),
		
		AddTokenEndpoint: MakeAddTokenEndpoint(svc),
		
		HasPermEndpoint: MakeHasPermEndpoint(svc),
		
	}
}
