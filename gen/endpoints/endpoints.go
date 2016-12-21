package acl_endpoints



import (
	"fmt"

	context "golang.org/x/net/context"
        pb "github.com/moul/acl/gen/pb"
	"github.com/go-kit/kit/endpoint"
)

var _ = fmt.Errorf

type Endpoints struct {
	
	HaspermEndpoint endpoint.Endpoint
	
}


/*{
  "name": "Hasperm",
  "input_type": ".acl.HasPermRequest",
  "output_type": ".acl.HasPermResponse",
  "options": {}
}*/

func (e *Endpoints)Hasperm(ctx context.Context, in *pb.HasPermRequest) (*pb.HasPermResponse, error) {
	out, err := e.HaspermEndpoint(ctx, in)
	if err != nil {
		return &pb.HasPermResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.HasPermResponse), err
}



func MakeHaspermEndpoint(svc pb.AclServiceServer) endpoint.Endpoint {
     	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.HasPermRequest)
		rep, err := svc.Hasperm(ctx, req)
		if err != nil {
			return &pb.HasPermResponse{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}


func MakeEndpoints(svc pb.AclServiceServer) Endpoints {
	return Endpoints{
		
		HaspermEndpoint: MakeHaspermEndpoint(svc),
		
	}
}
