package acl_httptransport

import (
	"encoding/json"
	context "golang.org/x/net/context"
	"log"
	"net/http"

	gokit_endpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	endpoints "github.com/moul/acl/gen/endpoints"
	pb "github.com/moul/acl/gen/pb"
)

func MakeGetTokenHandler(ctx context.Context, svc pb.AclServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeGetTokenRequest,
		encodeGetTokenResponse,
		[]httptransport.ServerOption{}...,
	)
}

func decodeGetTokenRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.GetTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeGetTokenResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

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

	log.Println("new HTTP endpoint: \"/GetToken\" (service=Acl)")
	mux.Handle("/GetToken", MakeGetTokenHandler(ctx, svc, endpoints.GetTokenEndpoint))

	log.Println("new HTTP endpoint: \"/AddToken\" (service=Acl)")
	mux.Handle("/AddToken", MakeAddTokenHandler(ctx, svc, endpoints.AddTokenEndpoint))

	log.Println("new HTTP endpoint: \"/HasPerm\" (service=Acl)")
	mux.Handle("/HasPerm", MakeHasPermHandler(ctx, svc, endpoints.HasPermEndpoint))

	return nil
}
