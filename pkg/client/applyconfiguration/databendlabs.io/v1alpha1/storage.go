/*
Copyright 2024.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// StorageApplyConfiguration represents a declarative configuration of the Storage type for use
// with apply.
type StorageApplyConfiguration struct {
	S3 *S3StorageApplyConfiguration `json:"s3,omitempty"`
}

// StorageApplyConfiguration constructs a declarative configuration of the Storage type for use with
// apply.
func Storage() *StorageApplyConfiguration {
	return &StorageApplyConfiguration{}
}

// WithS3 sets the S3 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the S3 field is set to the value of the last call.
func (b *StorageApplyConfiguration) WithS3(value *S3StorageApplyConfiguration) *StorageApplyConfiguration {
	b.S3 = value
	return b
}
