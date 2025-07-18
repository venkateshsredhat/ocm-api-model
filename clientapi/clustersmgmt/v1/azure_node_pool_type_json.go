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

package v1 // github.com/openshift-online/ocm-api-model/clientapi/clustersmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-api-model/clientapi/helpers"
)

// MarshalAzureNodePool writes a value of the 'azure_node_pool' type to the given writer.
func MarshalAzureNodePool(object *AzureNodePool, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	WriteAzureNodePool(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// WriteAzureNodePool writes a value of the 'azure_node_pool' type to the given stream.
func WriteAzureNodePool(object *AzureNodePool, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	var present_ bool
	present_ = object.bitmap_&1 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("os_disk_size_gibibytes")
		stream.WriteInt(object.osDiskSizeGibibytes)
		count++
	}
	present_ = object.bitmap_&2 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("os_disk_storage_account_type")
		stream.WriteString(object.osDiskStorageAccountType)
		count++
	}
	present_ = object.bitmap_&4 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("vm_size")
		stream.WriteString(object.vmSize)
		count++
	}
	present_ = object.bitmap_&8 != 0 && object.encryptionAtHost != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("encryption_at_host")
		WriteAzureNodePoolEncryptionAtHost(object.encryptionAtHost, stream)
		count++
	}
	present_ = object.bitmap_&16 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("ephemeral_os_disk_enabled")
		stream.WriteBool(object.ephemeralOSDiskEnabled)
		count++
	}
	present_ = object.bitmap_&32 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("os_disk_sse_encryption_set_resource_id")
		stream.WriteString(object.osDiskSseEncryptionSetResourceId)
		count++
	}
	present_ = object.bitmap_&64 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("resource_name")
		stream.WriteString(object.resourceName)
	}
	stream.WriteObjectEnd()
}

// UnmarshalAzureNodePool reads a value of the 'azure_node_pool' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalAzureNodePool(source interface{}) (object *AzureNodePool, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = ReadAzureNodePool(iterator)
	err = iterator.Error
	return
}

// ReadAzureNodePool reads a value of the 'azure_node_pool' type from the given iterator.
func ReadAzureNodePool(iterator *jsoniter.Iterator) *AzureNodePool {
	object := &AzureNodePool{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "os_disk_size_gibibytes":
			value := iterator.ReadInt()
			object.osDiskSizeGibibytes = value
			object.bitmap_ |= 1
		case "os_disk_storage_account_type":
			value := iterator.ReadString()
			object.osDiskStorageAccountType = value
			object.bitmap_ |= 2
		case "vm_size":
			value := iterator.ReadString()
			object.vmSize = value
			object.bitmap_ |= 4
		case "encryption_at_host":
			value := ReadAzureNodePoolEncryptionAtHost(iterator)
			object.encryptionAtHost = value
			object.bitmap_ |= 8
		case "ephemeral_os_disk_enabled":
			value := iterator.ReadBool()
			object.ephemeralOSDiskEnabled = value
			object.bitmap_ |= 16
		case "os_disk_sse_encryption_set_resource_id":
			value := iterator.ReadString()
			object.osDiskSseEncryptionSetResourceId = value
			object.bitmap_ |= 32
		case "resource_name":
			value := iterator.ReadString()
			object.resourceName = value
			object.bitmap_ |= 64
		default:
			iterator.ReadAny()
		}
	}
	return object
}
