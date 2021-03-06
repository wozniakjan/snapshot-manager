package v1alpha1

import (
	controller "github.com/cisco-sso/snapshot-manager/pkg/apis/snapshotmanager"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: controller.GroupName, Version: "v1alpha1"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	// Variables referenced in generation
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&ValidationStrategy{},
		&ValidationRun{},
		&SnapshotRevert{},
	)
	scheme.AddKnownTypeWithName(SchemeGroupVersion.WithKind("ValidationStrategyList"), &ValidationStrategyList{})
	scheme.AddKnownTypeWithName(SchemeGroupVersion.WithKind("ValidationRunList"), &ValidationRunList{})
	scheme.AddKnownTypeWithName(SchemeGroupVersion.WithKind("SnapshotRevertList"), &SnapshotRevertList{})
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
