package endpoint

import (
	"testing"

	motan "github.com/Loopring/motan-go/core"
)

func TestGetURL(t *testing.T) {

	url := &motan.URL{}
	grpcEndpoint := &GrpcEndPoint{url: url}
	if grpcEndpoint.GetURL() != url {
		t.Error("GetName Error")
	}
}
