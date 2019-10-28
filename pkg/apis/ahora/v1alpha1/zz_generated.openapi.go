// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/ahora/v1alpha1.SPA":       schema_pkg_apis_ahora_v1alpha1_SPA(ref),
		"./pkg/apis/ahora/v1alpha1.SPASpec":   schema_pkg_apis_ahora_v1alpha1_SPASpec(ref),
		"./pkg/apis/ahora/v1alpha1.SPAStatus": schema_pkg_apis_ahora_v1alpha1_SPAStatus(ref),
	}
}

func schema_pkg_apis_ahora_v1alpha1_SPA(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SPA is the Schema for the spas API",
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
							Ref: ref("./pkg/apis/ahora/v1alpha1.SPASpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/ahora/v1alpha1.SPAStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/ahora/v1alpha1.SPASpec", "./pkg/apis/ahora/v1alpha1.SPAStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_ahora_v1alpha1_SPASpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SPASpec defines the desired state of SPA",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_ahora_v1alpha1_SPAStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SPAStatus defines the observed state of SPA",
				Type:        []string{"object"},
			},
		},
	}
}