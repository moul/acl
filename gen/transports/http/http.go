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


func MakeAddTokenHandler(ctx context.Context, svc pb.AclServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeAddTokenRequest,
		encodeAddTokenResponse,
                []httptransport.ServerOption{}...,
	)
}

func decodeAddTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.AddTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeAddTokenResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func MakeHasPermHandler(ctx context.Context, svc pb.AclServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeHasPermRequest,
		encodeHasPermResponse,
                []httptransport.ServerOption{}...,
	)
}

func decodeHasPermRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.HasPermRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeHasPermResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}


func RegisterHandlers(ctx context.Context, svc pb.AclServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {
	
        log.Println("new HTTP endpoint: \"/AddToken\" (service=Acl)")
	mux.Handle("/AddToken", MakeAddTokenHandler(ctx, svc, endpoints.AddTokenEndpoint))
	
        log.Println("new HTTP endpoint: \"/HasPerm\" (service=Acl)")
	mux.Handle("/HasPerm", MakeHasPermHandler(ctx, svc, endpoints.HasPermEndpoint))
	
	return nil
}
