// Copyright (c) 2017 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v2

import (
	"github.com/projectcalico/libcalico-go/lib/numorstring"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindHostEndpoint     = "HostEndpoint"
	KindHostEndpointList = "HostEndpointList"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HostEndpoint contains information about a HostEndpoint resource that represents a “bare-metal”
// interface attached to the host that is running Calico’s agent, Felix. By default, Calico doesn’t
// apply any policy to such interfaces.
type HostEndpoint struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the HostEndpoint.
	Spec HostEndpointSpec `json:"spec,omitempty"`
}

// HostEndpointSpec contains the specification for a HostEndpoint resource.
type HostEndpointSpec struct {
	// The node name identifying the Calico node instance.
	Node string `json:"node,omitempty" validate:"omitempty,name"`
	// The name of the linux interface to apply policy to; for example “eth0”.
	// If "InterfaceName" is not present then at least one expected IP must be specified.
	InterfaceName string `json:"interfaceName,omitempty" validate:"omitempty,interface"`
	// The expected IP addresses (IPv4 and IPv6) of the endpoint.
	// If "InterfaceName" is not present, Calico will look for an interface matching any
	// of the IPs in the list and apply policy to that.
	// Note:
	// 	When using the selector|tag match criteria in an ingress or egress security Policy
	// 	or Profile, Calico converts the selector into a set of IP addresses. For host
	// 	endpoints, the ExpectedIPs field is used for that purpose. (If only the interface
	// 	name is specified, Calico does not learn the IPs of the interface for use in match
	// 	criteria.)
	ExpectedIPs []string `json:"expectedIPs,omitempty" validate:"omitempty,dive,ip"`
	// A list of identifiers of security Profile objects that apply to this endpoint. Each
	// profile is applied in the order that they appear in this list.  Profile rules are applied
	// after the selector-based security policy.
	Profiles []string `json:"profiles,omitempty" validate:"omitempty,dive,namespacedname"`
	// Ports contains the endpoint's named ports, which may be referenced in security policy rules.
	Ports []EndpointPort `json:"ports,omitempty" validate:"dive"`
}

type EndpointPort struct {
	Name     string               `json:"name" validate:"name"`
	Protocol numorstring.Protocol `json:"protocol"`
	Port     uint16               `json:"port" validate:"gt=0"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HostEndpointList contains a list of HostEndpoint resources.
type HostEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []HostEndpoint `json:"items"`
}

// NewHostEndpoint creates a new (zeroed) HostEndpoint struct with the TypeMetadata initialised to the current
// version.
func NewHostEndpoint() *HostEndpoint {
	return &HostEndpoint{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindHostEndpoint,
			APIVersion: GroupVersionCurrent,
		},
	}
}

// NewHostEndpointList creates a new (zeroed) HostEndpointList struct with the TypeMetadata initialised to the current
// version.
func NewHostEndpointList() *HostEndpointList {
	return &HostEndpointList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindHostEndpointList,
			APIVersion: GroupVersionCurrent,
		},
	}
}
