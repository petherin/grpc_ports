package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"portsvc/proto"
)

type mockClientsPort struct {
	list       map[string]*proto.Port
	forceError error
}

func NewMockClientsPort(list map[string]*proto.Port, forceError error) mockClientsPort {
	return mockClientsPort{list: list, forceError: forceError}
}

func (m mockClientsPort) Save(ctx context.Context, in *proto.SavePortRequest, opts ...grpc.CallOption) (*proto.SavePortResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockClientsPort) GetPorts(ctx context.Context, in *proto.GetPortsRequest, opts ...grpc.CallOption) (*proto.PortList, error) {
	list := &proto.PortList{
		Ports: m.list,
	}

	return list, m.forceError
}

func TestHandleGetPorts(t *testing.T) {
	mockPortList := map[string]*proto.Port{
		"testcode": {
			Name:        "testname",
			City:        "testcity",
			Country:     "testcountry",
			Alias:       nil,
			Regions:     nil,
			Coordinates: nil,
			Timezone:    "testtimezone",
			Unlocs:      nil,
			Code:        "testcode",
		},
	}

	tt := []struct {
		name           string
		addr           string
		portsClient    proto.PortsClient
		method         string
		expectedStatus int
		expected       string
	}{
		{
			name:           "1. Happy path, ports returned",
			addr:           ":8080",
			portsClient:    NewMockClientsPort(mockPortList, nil),
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expected:       "{\"testcode\":{\"name\":\"testname\",\"city\":\"testcity\",\"country\":\"testcountry\",\"timezone\":\"testtimezone\",\"code\":\"testcode\"}}",
		},
		{
			name:           "2. Invalid method used, error returned",
			addr:           ":8080",
			portsClient:    NewMockClientsPort(nil, nil),
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
			expected:       "only GET allowed\n",
		},
		{
			name:           "3. Error returned by port client which is returned in a 500 response",
			addr:           ":8080",
			portsClient:    NewMockClientsPort(nil, fmt.Errorf("port client error")),
			method:         http.MethodGet,
			expectedStatus: http.StatusInternalServerError,
			expected:       "port client error\n",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/ports", nil)

			svr := New(tc.addr, tc.portsClient)
			svr.Router.HandleFunc("/ports", svr.handleGetPorts())

			svr.HTTPServer = &http.Server{Handler: &svr.Router}

			// Record the response
			responseRecorder := httptest.NewRecorder()

			// Send request
			svr.Router.ServeHTTP(responseRecorder, request)

			assert.Equal(t, tc.expectedStatus, responseRecorder.Code)
			actual := responseRecorder.Body.String()
			assert.Equal(t, tc.expected, actual)
		})
	}

}
