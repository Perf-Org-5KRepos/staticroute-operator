// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/iks/v1.StaticRoute":       schema_pkg_apis_iks_v1_StaticRoute(ref),
		"./pkg/apis/iks/v1.StaticRouteSpec":   schema_pkg_apis_iks_v1_StaticRouteSpec(ref),
		"./pkg/apis/iks/v1.StaticRouteStatus": schema_pkg_apis_iks_v1_StaticRouteStatus(ref),
	}
}

func schema_pkg_apis_iks_v1_StaticRoute(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StaticRoute is the Schema for the staticroutes API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/iks/v1.StaticRouteSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/iks/v1.StaticRouteStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/iks/v1.StaticRouteSpec", "./pkg/apis/iks/v1.StaticRouteStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_iks_v1_StaticRouteSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StaticRouteSpec defines the desired state of StaticRoute",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"subnet": {
						SchemaProps: spec.SchemaProps{
							Description: "Subnet defines the required IP subnet in the form of: \"x.x.x.x/x\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"gateway": {
						SchemaProps: spec.SchemaProps{
							Description: "Gateway the gateway the subnet is routed through (optional, discovered if not set)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"subnet"},
			},
		},
	}
}

func schema_pkg_apis_iks_v1_StaticRouteStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StaticRouteStatus defines the observed state of StaticRoute",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodeStatus": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/iks/v1.StaticRouteNodeStatus"),
									},
								},
							},
						},
					},
				},
				Required: []string{"nodeStatus"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/iks/v1.StaticRouteNodeStatus"},
	}
}
