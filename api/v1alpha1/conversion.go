/*
SPDX-FileCopyrightText: 2025 SAP SE or an SAP affiliate company and valkey-operator-cop contributors
SPDX-License-Identifier: Apache-2.0
*/

package v1alpha1

// Uncomment the following block if conversion is used, and this api version is the conversion hub;
// see https://book.kubebuilder.io/multiversion-tutorial/conversion.html to learn about the concept of hubs and spokes.
/*
import "sigs.k8s.io/controller-runtime/pkg/conversion"

var _ conversion.Hub = &ValkeyOperator{}

func (c *ValkeyOperator) Hub() {}
*/

// Uncomment the following block if conversion is used, and this api version is a conversion spoke,
// and replace _HUB_API_VERSION_ with the api version of the conversion hub;
// see https://book.kubebuilder.io/multiversion-tutorial/conversion.html to learn about the concept of hubs and spokes.
/*
import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/sap/valkey-operator-cop/api/_HUB_API_VERSION_"
)

var _ conversion.Convertible = &ValkeyOperator{}

func (src *ValkeyOperator) ConvertTo(dstHub conversion.Hub) error {
	dst := dstHub.(*_HUB_API_VERSION_.ValkeyOperator)
	dst.ObjectMeta = src.ObjectMeta
	// Add logic here to convert src.Spec into dst.Spec and src.Status into dst.Status.
	return nil
}

func (dst *ValkeyOperator) ConvertFrom(srcHub conversion.Hub) error {
	src := srcHub.(*_HUB_API_VERSION_.ValkeyOperator)
	dst.ObjectMeta = src.ObjectMeta
	// Add logic here to convert src.Spec into dst.Spec and src.Status into dst.Status.
	return nil
}
*/
