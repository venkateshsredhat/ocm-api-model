/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1alpha1 // github.com/openshift-online/ocm-api-model/clientapi/arohcp/v1alpha1

// AzureNodePoolBuilder contains the data and logic needed to build 'azure_node_pool' objects.
//
// Representation of azure node pool specific parameters.
type AzureNodePoolBuilder struct {
	bitmap_                          uint32
	osDiskSizeGibibytes              int
	osDiskStorageAccountType         string
	vmSize                           string
	encryptionAtHost                 *AzureNodePoolEncryptionAtHostBuilder
	osDiskSseEncryptionSetResourceId string
	resourceName                     string
	ephemeralOSDiskEnabled           bool
}

// NewAzureNodePool creates a new builder of 'azure_node_pool' objects.
func NewAzureNodePool() *AzureNodePoolBuilder {
	return &AzureNodePoolBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AzureNodePoolBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// OSDiskSizeGibibytes sets the value of the 'OS_disk_size_gibibytes' attribute to the given value.
func (b *AzureNodePoolBuilder) OSDiskSizeGibibytes(value int) *AzureNodePoolBuilder {
	b.osDiskSizeGibibytes = value
	b.bitmap_ |= 1
	return b
}

// OSDiskStorageAccountType sets the value of the 'OS_disk_storage_account_type' attribute to the given value.
func (b *AzureNodePoolBuilder) OSDiskStorageAccountType(value string) *AzureNodePoolBuilder {
	b.osDiskStorageAccountType = value
	b.bitmap_ |= 2
	return b
}

// VMSize sets the value of the 'VM_size' attribute to the given value.
func (b *AzureNodePoolBuilder) VMSize(value string) *AzureNodePoolBuilder {
	b.vmSize = value
	b.bitmap_ |= 4
	return b
}

// EncryptionAtHost sets the value of the 'encryption_at_host' attribute to the given value.
//
// AzureNodePoolEncryptionAtHost defines the encryption setting for Encryption At Host.
// If not specified, Encryption at Host is not enabled.
func (b *AzureNodePoolBuilder) EncryptionAtHost(value *AzureNodePoolEncryptionAtHostBuilder) *AzureNodePoolBuilder {
	b.encryptionAtHost = value
	if value != nil {
		b.bitmap_ |= 8
	} else {
		b.bitmap_ &^= 8
	}
	return b
}

// EphemeralOSDiskEnabled sets the value of the 'ephemeral_OS_disk_enabled' attribute to the given value.
func (b *AzureNodePoolBuilder) EphemeralOSDiskEnabled(value bool) *AzureNodePoolBuilder {
	b.ephemeralOSDiskEnabled = value
	b.bitmap_ |= 16
	return b
}

// OsDiskSseEncryptionSetResourceId sets the value of the 'os_disk_sse_encryption_set_resource_id' attribute to the given value.
func (b *AzureNodePoolBuilder) OsDiskSseEncryptionSetResourceId(value string) *AzureNodePoolBuilder {
	b.osDiskSseEncryptionSetResourceId = value
	b.bitmap_ |= 32
	return b
}

// ResourceName sets the value of the 'resource_name' attribute to the given value.
func (b *AzureNodePoolBuilder) ResourceName(value string) *AzureNodePoolBuilder {
	b.resourceName = value
	b.bitmap_ |= 64
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AzureNodePoolBuilder) Copy(object *AzureNodePool) *AzureNodePoolBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.osDiskSizeGibibytes = object.osDiskSizeGibibytes
	b.osDiskStorageAccountType = object.osDiskStorageAccountType
	b.vmSize = object.vmSize
	if object.encryptionAtHost != nil {
		b.encryptionAtHost = NewAzureNodePoolEncryptionAtHost().Copy(object.encryptionAtHost)
	} else {
		b.encryptionAtHost = nil
	}
	b.ephemeralOSDiskEnabled = object.ephemeralOSDiskEnabled
	b.osDiskSseEncryptionSetResourceId = object.osDiskSseEncryptionSetResourceId
	b.resourceName = object.resourceName
	return b
}

// Build creates a 'azure_node_pool' object using the configuration stored in the builder.
func (b *AzureNodePoolBuilder) Build() (object *AzureNodePool, err error) {
	object = new(AzureNodePool)
	object.bitmap_ = b.bitmap_
	object.osDiskSizeGibibytes = b.osDiskSizeGibibytes
	object.osDiskStorageAccountType = b.osDiskStorageAccountType
	object.vmSize = b.vmSize
	if b.encryptionAtHost != nil {
		object.encryptionAtHost, err = b.encryptionAtHost.Build()
		if err != nil {
			return
		}
	}
	object.ephemeralOSDiskEnabled = b.ephemeralOSDiskEnabled
	object.osDiskSseEncryptionSetResourceId = b.osDiskSseEncryptionSetResourceId
	object.resourceName = b.resourceName
	return
}
