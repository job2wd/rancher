// +build !k8s

package k8s

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rancher/norman/pkg/remotedialer"
)

func NewK3sConfig(ctx context.Context, dataDir string, authorizer remotedialer.Authorizer) (context.Context, interface{}, http.Handler, error) {
	return ctx, nil, nil, fmt.Errorf("k3s not compiled in, build with -tags k8s")
}
