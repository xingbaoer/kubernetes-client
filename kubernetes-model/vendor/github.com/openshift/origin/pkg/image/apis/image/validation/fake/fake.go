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
package fake

import (
	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	"github.com/openshift/origin/pkg/image/apis/image/validation/whitelist"
)

type RegistryWhitelister struct{}

func (rw *RegistryWhitelister) AdmitHostname(host string, transport whitelist.WhitelistTransport) error {
	return nil
}
func (rw *RegistryWhitelister) AdmitPullSpec(pullSpec string, transport whitelist.WhitelistTransport) error {
	return nil
}
func (rw *RegistryWhitelister) AdmitDockerImageReference(ref *imageapi.DockerImageReference, transport whitelist.WhitelistTransport) error {
	return nil
}
func (rw *RegistryWhitelister) WhitelistRegistry(hostPortGlob string, transport whitelist.WhitelistTransport) error {
	return nil
}
func (rw *RegistryWhitelister) WhitelistPullSpecs(pullSpec ...string) {}
func (rw *RegistryWhitelister) Copy() whitelist.RegistryWhitelister {
	return &RegistryWhitelister{}
}
