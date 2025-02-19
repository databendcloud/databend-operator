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

// FileLogSpecApplyConfiguration represents a declarative configuration of the FileLogSpec type for use
// with apply.
type FileLogSpecApplyConfiguration struct {
	Enabled *bool   `json:"enabled,omitempty"`
	Format  *string `json:"format,omitempty"`
	Level   *string `json:"level,omitempty"`
	Dir     *string `json:"dir,omitempty"`
}

// FileLogSpecApplyConfiguration constructs a declarative configuration of the FileLogSpec type for use with
// apply.
func FileLogSpec() *FileLogSpecApplyConfiguration {
	return &FileLogSpecApplyConfiguration{}
}

// WithEnabled sets the Enabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enabled field is set to the value of the last call.
func (b *FileLogSpecApplyConfiguration) WithEnabled(value bool) *FileLogSpecApplyConfiguration {
	b.Enabled = &value
	return b
}

// WithFormat sets the Format field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Format field is set to the value of the last call.
func (b *FileLogSpecApplyConfiguration) WithFormat(value string) *FileLogSpecApplyConfiguration {
	b.Format = &value
	return b
}

// WithLevel sets the Level field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Level field is set to the value of the last call.
func (b *FileLogSpecApplyConfiguration) WithLevel(value string) *FileLogSpecApplyConfiguration {
	b.Level = &value
	return b
}

// WithDir sets the Dir field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Dir field is set to the value of the last call.
func (b *FileLogSpecApplyConfiguration) WithDir(value string) *FileLogSpecApplyConfiguration {
	b.Dir = &value
	return b
}
