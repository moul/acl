package acl_httptransport



import (
       "log"
	"net/http"
	"encoding/json"
	context "golang.org/x/net/context"

        pb "github.com/moul/acl/gen/pb"
        gokit_endpoint "github.com/go-kit/kit/endpoint"
        httptransport "github.com/go-kit/kit/transport/http"
        endpoints "github.com/moul/acl/gen/endpoints"
)


func MakeHaspermHandler(ctx context.Context, svc pb.AclServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeHaspermRequest,
		encodeHaspermResponse,
                []httptransport.ServerOption{}...,
	)
}

func decodeHaspermRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.HasPermRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeHaspermResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}


func RegisterHandlers(ctx context.Context, svc pb.AclServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {
	
        log.Println("new HTTP endpoint: \"/Hasperm\" (service=Acl)")
	mux.Handle("/Hasperm", MakeHaspermHandler(ctx, svc, endpoints.HaspermEndpoint))
	
	return nil
}
