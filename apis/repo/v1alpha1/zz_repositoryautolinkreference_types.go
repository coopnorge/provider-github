/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RepositoryAutolinkReferenceInitParameters struct {

	// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
	// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
	IsAlphanumeric *bool `json:"isAlphanumeric,omitempty" tf:"is_alphanumeric,omitempty"`

	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
	KeyPrefix *string `json:"keyPrefix,omitempty" tf:"key_prefix,omitempty"`

	// The repository of the autolink reference.
	// The repository name
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-github/apis/repo/v1alpha1.Repository
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// The template of the target URL used for the links; must be a valid URL and contain <num> for the reference number
	// The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
	TargetURLTemplate *string `json:"targetUrlTemplate,omitempty" tf:"target_url_template,omitempty"`
}

type RepositoryAutolinkReferenceObservation struct {

	// An etag representing the autolink reference object.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
	// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
	IsAlphanumeric *bool `json:"isAlphanumeric,omitempty" tf:"is_alphanumeric,omitempty"`

	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
	KeyPrefix *string `json:"keyPrefix,omitempty" tf:"key_prefix,omitempty"`

	// The repository of the autolink reference.
	// The repository name
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// The template of the target URL used for the links; must be a valid URL and contain <num> for the reference number
	// The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
	TargetURLTemplate *string `json:"targetUrlTemplate,omitempty" tf:"target_url_template,omitempty"`
}

type RepositoryAutolinkReferenceParameters struct {

	// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters. Default is true.
	// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
	// +kubebuilder:validation:Optional
	IsAlphanumeric *bool `json:"isAlphanumeric,omitempty" tf:"is_alphanumeric,omitempty"`

	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit
	// +kubebuilder:validation:Optional
	KeyPrefix *string `json:"keyPrefix,omitempty" tf:"key_prefix,omitempty"`

	// The repository of the autolink reference.
	// The repository name
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-upjet-github/apis/repo/v1alpha1.Repository
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// The template of the target URL used for the links; must be a valid URL and contain <num> for the reference number
	// The template of the target URL used for the links; must be a valid URL and contain `<num>` for the reference number
	// +kubebuilder:validation:Optional
	TargetURLTemplate *string `json:"targetUrlTemplate,omitempty" tf:"target_url_template,omitempty"`
}

// RepositoryAutolinkReferenceSpec defines the desired state of RepositoryAutolinkReference
type RepositoryAutolinkReferenceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryAutolinkReferenceParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RepositoryAutolinkReferenceInitParameters `json:"initProvider,omitempty"`
}

// RepositoryAutolinkReferenceStatus defines the observed state of RepositoryAutolinkReference.
type RepositoryAutolinkReferenceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryAutolinkReferenceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RepositoryAutolinkReference is the Schema for the RepositoryAutolinkReferences API. Creates and manages autolink references for a single repository
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type RepositoryAutolinkReference struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.keyPrefix) || (has(self.initProvider) && has(self.initProvider.keyPrefix))",message="spec.forProvider.keyPrefix is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.targetUrlTemplate) || (has(self.initProvider) && has(self.initProvider.targetUrlTemplate))",message="spec.forProvider.targetUrlTemplate is a required parameter"
	Spec   RepositoryAutolinkReferenceSpec   `json:"spec"`
	Status RepositoryAutolinkReferenceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryAutolinkReferenceList contains a list of RepositoryAutolinkReferences
type RepositoryAutolinkReferenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RepositoryAutolinkReference `json:"items"`
}

// Repository type metadata.
var (
	RepositoryAutolinkReference_Kind             = "RepositoryAutolinkReference"
	RepositoryAutolinkReference_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RepositoryAutolinkReference_Kind}.String()
	RepositoryAutolinkReference_KindAPIVersion   = RepositoryAutolinkReference_Kind + "." + CRDGroupVersion.String()
	RepositoryAutolinkReference_GroupVersionKind = CRDGroupVersion.WithKind(RepositoryAutolinkReference_Kind)
)

func init() {
	SchemeBuilder.Register(&RepositoryAutolinkReference{}, &RepositoryAutolinkReferenceList{})
}
