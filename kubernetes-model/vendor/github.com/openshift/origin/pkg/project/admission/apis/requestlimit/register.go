/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package requestlimit

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	GroupName       = "requestlimit.project.openshift.io"
	LegacyGroupName = ""
)

var (
	SchemeGroupVersion       = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
	LegacySchemeGroupVersion = schema.GroupVersion{Group: LegacyGroupName, Version: runtime.APIVersionInternal}

	LegacySchemeBuilder    = runtime.NewSchemeBuilder(addLegacyKnownTypes)
	AddToSchemeInCoreGroup = LegacySchemeBuilder.AddToScheme

	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addLegacyKnownTypes(scheme *runtime.Scheme) error {
	types := []runtime.Object{
		&ProjectRequestLimitConfig{},
	}
	scheme.AddKnownTypes(LegacySchemeGroupVersion, types...)
	return nil
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&ProjectRequestLimitConfig{},
	)
	return nil
}
