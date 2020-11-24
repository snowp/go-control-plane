package ttl

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/golang/protobuf/ptypes/any"
)

// Helper functions for interacting with TTL resources for xDS V2. Since TTL resources are not supported for V2, these are
// essentially noops.

func MaybeCreateTtlResourceIfSupported(resource types.ResourceWithTtl, name string, resourceTypeUrl string, heartbeat bool) (types.Resource, string, error) {
	return resource.Resource, resourceTypeUrl, nil
}

func IsTTLResource(resource *any.Any) bool {
	return true
}
