package http_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func TestServerBadDecode(t *testing.T) {
	handler := httptransport.NewServer(
		context.Background(),
		func(context.Context, interface{}) (interface{}, error) { return struct{}{}, nil },
		func(context.Context, *http.Request) (interface{}, error) { return struct{}{}, errors.New("dang") },
		func(context.Context, http.ResponseWriter, interface{}) error { return nil },
	)
	server := httptest.NewServer(handler)
	defer server.Close()
	resp, _ := http.Get(server.URL)
	if want, have := http.StatusInternalServerError, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
}

func TestServerBadEndpoint(t *testing.T) {
	handler := httptransport.NewServer(
		context.Background(),
		func(context.Context, interface{}) (interface{}, error) { return struct{}{}, errors.New("dang") },
		func(context.Context, *http.Request) (interface{}, error) { return struct{}{}, nil },
		func(context.Context, http.ResponseWriter, interface{}) error { return nil },
	)
	server := httptest.NewServer(handler)
	defer server.Close()
	resp, _ := http.Get(server.URL)
	if want, have := http.StatusInternalServerError, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
}

func TestServerBadEncode(t *testing.T) {
	handler := httptransport.NewServer(
		context.Background(),
		func(context.Context, interface{}) (interface{}, error) { return struct{}{}, nil },
		func(context.Context, *http.Request) (interface{}, error) { return struct{}{}, nil },
		func(context.Context, http.ResponseWriter, interface{}) error { return errors.New("dang") },
	)
	server := httptest.NewServer(handler)
	defer server.Close()
	resp, _ := http.Get(server.URL)
	if want, have := http.StatusInternalServerError, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
}

func TestServerErrorEncoder(t *testing.T) {
	errTeapot := errors.New("teapot")
	code := func(err error) int {
		if err == errTeapot {
			return http.StatusTeapot
		}
		return http.StatusInternalServerError
	}
	handler := httptransport.NewServer(
		context.Background(),
		func(context.Context, interface{}) (interface{}, error) { return struct{}{}, errTeapot },
		func(context.Context, *http.Request) (interface{}, error) { return struct{}{}, nil },
		func(context.Context, http.ResponseWriter, interface{}) error { return nil },
		httptransport.ServerErrorEncoder(func(_ context.Context, err error, w http.ResponseWriter) { w.WriteHeader(code(err)) }),
	)
	server := httptest.NewServer(handler)
	defer server.Close()
	resp, _ := http.Get(server.URL)
	if want, have := http.StatusTeapot, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
}

func TestServerHappyPath(t *testing.T) {
	_, step, response := testServer(t)
	step()
	resp := <-response
	defer resp.Body.Close()
	buf, _ := ioutil.ReadAll(resp.Body)
	if want, have := http.StatusOK, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d (%s)", want, have, buf)
	}
}

func TestServerFinalizer(t *testing.T) {
	c := make(chan int)
	handler := httptransport.NewServer(
		context.Background(),
		endpoint.Nop,
		func(context.Context, *http.Request) (interface{}, error) {
			return struct{}{}, nil
		},
		func(_ context.Context, w http.ResponseWriter, _ interface{}) error {
			w.WriteHeader(<-c)
			return nil
		},
		httptransport.ServerFinalizer(func(_ context.Context, code int, _ *http.Request) {
			c <- code
		}),
	)

	server := httptest.NewServer(handler)
	defer server.Close()
	go http.Get(server.URL)

	want := http.StatusTeapot
	c <- want   // give status code to response encoder
	have := <-c // take status code from finalizer

	if want != have {
		t.Errorf("want %d, have %d", want, have)
	}
}

type enhancedResponse struct {
	Foo string `json:"foo"`
}

func (e enhancedResponse) StatusCode() int      { return http.StatusPaymentRequired }
func (e enhancedResponse) Headers() http.Header { return http.Header{"X-Edward": []string{"Snowden"}} }

func TestEncodeJSONResponse(t *testing.T) {
	handler := httptransport.NewServer(
		context.Background(),
		func(context.Context, interface{}) (interface{}, error) { return enhancedResponse{Foo: "bar"}, nil },
		func(context.Context, *http.Request) (interface{}, error) { return struct{}{}, nil },
		httptransport.EncodeJSONResponse,
	)

	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if want, have := http.StatusPaymentRequired, resp.StatusCode; want != have {
		t.Errorf("StatusCode: want %d, have %d", want, have)
	}
	if want, have := "Snowden", resp.Header.Get("X-Edward"); want != have {
		t.Errorf("X-Edward: want %q, have %q", want, have)
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	if want, have := `{"foo":"bar"}`, strings.TrimSpace(string(buf)); want != have {
		t.Errorf("Body: want %s, have %s", want, have)
	}
}

type enhancedError struct{}

func (e enhancedError) Error() string                { return "enhanced error" }
func (e enhancedError) StatusCode() int              { return http.StatusTeapot }
func (e enhancedError) MarshalJSON() ([]byte, error) { return []byte(`{"err":"enhanced"}`), nil }
func (e enhancedError) Headers() http.Header         { return http.Header{"X-Enhanced": []string{"1"}} }

func TestEnhancedError(t *testing.T) {
	handler := httptransport.NewServer(
		context.Background(),
		func(context.Context, interface{}) (interface{}, error) { return nil, enhancedError{} },
		func(context.Context, *http.Request) (interface{}, error) { return struct{}{}, nil },
		func(_ context.Context, w http.ResponseWriter, _ interface{}) error { return nil },
	)

	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if want, have := http.StatusTeapot, resp.StatusCode; want != have {
		t.Errorf("StatusCode: want %d, have %d", want, have)
	}
	if want, have := "1", resp.Header.Get("X-Enhanced"); want != have {
		t.Errorf("X-Enhanced: want %q, have %q", want, have)
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	if want, have := `{"err":"enhanced"}`, strings.TrimSpace(string(buf)); want != have {
		t.Errorf("Body: want %s, have %s", want, have)
	}
}

func testServer(t *testing.T) (cancel, step func(), resp <-chan *http.Response) {
	var (
		ctx, cancelfn = context.WithCancel(context.Background())
		stepch        = make(chan bool)
		endpoint      = func(context.Context, interface{}) (interface{}, error) { <-stepch; return struct{}{}, nil }
		response      = make(chan *http.Response)
		handler       = httptransport.NewServer(
			ctx,
			endpoint,
			func(context.Context, *http.Request) (interface{}, error) { return struct{}{}, nil },
			func(context.Context, http.ResponseWriter, interface{}) error { return nil },
			httptransport.ServerBefore(func(ctx context.Context, r *http.Request) context.Context { return ctx }),
			httptransport.ServerAfter(func(ctx context.Context, w http.ResponseWriter) context.Context { return ctx }),
		)
	)
	go func() {
		server := httptest.NewServer(handler)
		defer server.Close()
		resp, err := http.Get(server.URL)
		if err != nil {
			t.Error(err)
			return
		}
		response <- resp
	}()
	return cancelfn, func() { stepch <- true }, response
}
