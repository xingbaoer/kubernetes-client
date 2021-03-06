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
package analysis

import (
	"testing"

	appsedges "github.com/openshift/origin/pkg/oc/graph/appsgraph"
	buildedges "github.com/openshift/origin/pkg/oc/graph/buildgraph"
	osgraph "github.com/openshift/origin/pkg/oc/graph/genericgraph"
	osgraphtest "github.com/openshift/origin/pkg/oc/graph/genericgraph/test"
	imageedges "github.com/openshift/origin/pkg/oc/graph/imagegraph"
)

func TestMissingImageStreamTag(t *testing.T) {
	g, _, err := osgraphtest.BuildGraph("../../../graph/genericgraph/test/missing-istag.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	buildedges.AddAllInputOutputEdges(g)
	appsedges.AddAllTriggerDeploymentConfigsEdges(g)
	imageedges.AddAllImageStreamRefEdges(g)
	imageedges.AddAllImageStreamImageRefEdges(g)

	markers := FindDeploymentConfigTriggerErrors(g, osgraph.DefaultNamer)
	if e, a := 1, len(markers); e != a {
		t.Fatalf("expected %v, got %v", e, a)
	}

	if got, expected := markers[0].Key, MissingImageStreamTagWarning; got != expected {
		t.Fatalf("expected marker key %q, got %q", expected, got)
	}
}

func TestMissingImageStream(t *testing.T) {
	g, _, err := osgraphtest.BuildGraph("../../../graph/genericgraph/test/unpushable-build-2.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	buildedges.AddAllInputOutputEdges(g)
	appsedges.AddAllTriggerDeploymentConfigsEdges(g)
	imageedges.AddAllImageStreamRefEdges(g)
	imageedges.AddAllImageStreamImageRefEdges(g)

	markers := FindDeploymentConfigTriggerErrors(g, osgraph.DefaultNamer)
	if e, a := 1, len(markers); e != a {
		t.Fatalf("expected %v, got %v", e, a)
	}

	if got, expected := markers[0].Key, MissingImageStreamErr; got != expected {
		t.Fatalf("expected marker key %q, got %q", expected, got)
	}
}

func TestMissingReadinessProbe(t *testing.T) {
	g, _, err := osgraphtest.BuildGraph("../../../graph/genericgraph/test/unpushable-build-2.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	buildedges.AddAllInputOutputEdges(g)
	appsedges.AddAllTriggerDeploymentConfigsEdges(g)
	imageedges.AddAllImageStreamRefEdges(g)

	markers := FindDeploymentConfigReadinessWarnings(g, osgraph.DefaultNamer, "command probe")
	if e, a := 1, len(markers); e != a {
		t.Fatalf("expected %v, got %v", e, a)
	}

	if got, expected := markers[0].Key, MissingReadinessProbeWarning; got != expected {
		t.Fatalf("expected marker key %q, got %q", expected, got)
	}
}

func TestSingleHostVolumeError(t *testing.T) {
	g, _, err := osgraphtest.BuildGraph("../../../graph/genericgraph/test/dc-with-claim.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	appsedges.AddAllVolumeClaimEdges(g)

	markers := FindPersistentVolumeClaimWarnings(g, osgraph.DefaultNamer)
	if e, a := 1, len(markers); e != a {
		t.Fatalf("expected %v, got %v", e, a)
	}

	if got, expected := markers[0].Key, SingleHostVolumeWarning; got != expected {
		t.Fatalf("expected marker key %q, got %q", expected, got)
	}
}
