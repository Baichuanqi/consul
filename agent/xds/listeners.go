package xds

import (
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
)

type listenerFilterOpts struct {

	// TODO(m1) removing this line eliminates the problem
	httpAuthzFilter *envoyhttp.HttpFilter
}
