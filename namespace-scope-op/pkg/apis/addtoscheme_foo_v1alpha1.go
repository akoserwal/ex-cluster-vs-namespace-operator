package apis

import (
	"github.com/akoserwal/ex-cluster-vs-namespace-operatorr/namespace-scope-op/pkg/apis/foo/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
