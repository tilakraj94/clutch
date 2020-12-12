package k8s

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	lyftk8sapiv1 "github.com/lyft/clutch/backend/api/k8s/v1"
)

func (a *k8sAPI) DeleteFacet(ctx context.Context, req *lyftk8sapiv1.DeleteFacetRequest) (*lyftk8sapiv1.DeleteFacetResponse, error) {
	switch req.Type {
	case lyftk8sapiv1.DeleteFacetRequest_SERVICE:
		return a.deleteServiceFacet(ctx, req)
	default:
		return nil, fmt.Errorf("unsupported facet type %v", req.Type)
	}
}

func (a *k8sAPI) deleteServiceFacet(ctx context.Context, req *lyftk8sapiv1.DeleteFacetRequest) (*lyftk8sapiv1.DeleteFacetResponse, error) {
	// delete the deployment object first and return an error if it fails
	err := a.k8s.DeleteDeployment(ctx, req.Clientset, req.Cluster, req.Namespace, req.Name)
	if err != nil {
		return nil, err
	}

	// delete HPA and service object in parallel after deployment has been deleted and
	// log errors in case of failure. We do not return an error in this case to account
	// for edge cases where an HPA and service object might not be present in the namespace.
	errs, _ := errgroup.WithContext(ctx)
	errs.Go(func() error {
		return a.k8s.DeleteHPA(ctx, req.Clientset, req.Cluster, req.Namespace, req.Name)
	})
	errs.Go(func() error {
		return a.k8s.DeleteService(ctx, req.Clientset, req.Cluster, req.Namespace, req.Name)
	})

	/*
		if err := errs.Wait(); err != nil {
			// log the errors if delete HPA or delete Service fail
			a.logger.Error("delete service facet error:", zap.Error(err))
		}
	*/
	return &lyftk8sapiv1.DeleteFacetResponse{}, nil
}
