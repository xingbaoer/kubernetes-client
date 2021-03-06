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
	"strings"
	"testing"

	appsgraph "github.com/openshift/origin/pkg/oc/graph/appsgraph"
	osgraph "github.com/openshift/origin/pkg/oc/graph/genericgraph"
	osgraphtest "github.com/openshift/origin/pkg/oc/graph/genericgraph/test"
	"github.com/openshift/origin/pkg/oc/graph/kubegraph"
)

func TestHPAMissingCPUTargetError(t *testing.T) {
	g, _, err := osgraphtest.BuildGraph("./../../../graph/genericgraph/test/hpa-missing-cpu-target.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	markers := FindHPASpecsMissingCPUTargets(g, osgraph.DefaultNamer)
	if len(markers) != 1 {
		t.Fatalf("expected to find one HPA spec missing a CPU target, got %d", len(markers))
	}

	if actual, expected := markers[0].Severity, osgraph.ErrorSeverity; actual != expected {
		t.Errorf("expected HPA missing CPU target to be %v, got %v", expected, actual)
	}

	if actual, expected := markers[0].Key, HPAMissingCPUTargetError; actual != expected {
		t.Errorf("expected marker type %v, got %v", expected, actual)
	}

	patchString := `-p '{"spec":{"targetCPUUtilizationPercentage": 80}}'`
	if !strings.HasSuffix(string(markers[0].Suggestion), patchString) {
		t.Errorf("expected suggestion to end with patch JSON path, got %q", markers[0].Suggestion)
	}
}

func TestHPAMissingScaleRefError(t *testing.T) {
	g, _, err := osgraphtest.BuildGraph("./../../../graph/genericgraph/test/hpa-missing-scale-ref.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	markers := FindHPASpecsMissingScaleRefs(g, osgraph.DefaultNamer)
	if len(markers) != 1 {
		t.Fatalf("expected to find one HPA spec missing a scale ref, got %d", len(markers))
	}

	if actual, expected := markers[0].Severity, osgraph.ErrorSeverity; actual != expected {
		t.Errorf("expected HPA missing scale ref to be %v, got %v", expected, actual)
	}

	if actual, expected := markers[0].Key, HPAMissingScaleRefError; actual != expected {
		t.Errorf("expected marker type %v, got %v", expected, actual)
	}
}

func TestOverlappingHPAsWarning(t *testing.T) {
	g, _, err := osgraphtest.BuildGraph("./../../../graph/genericgraph/test/overlapping-hpas.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	kubegraph.AddHPAScaleRefEdges(g)
	appsgraph.AddAllDeploymentConfigsDeploymentEdges(g)

	markers := FindOverlappingHPAs(g, osgraph.DefaultNamer)
	if len(markers) != 8 {
		t.Fatalf("expected to find eight overlapping HPA markers, got %d", len(markers))
	}

	for _, marker := range markers {
		if actual, expected := marker.Severity, osgraph.WarningSeverity; actual != expected {
			t.Errorf("expected overlapping HPAs to be %v, got %v", expected, actual)
		}

		if actual, expected := marker.Key, HPAOverlappingScaleRefWarning; actual != expected {
			t.Errorf("expected marker type %v, got %v", expected, actual)
		}
	}
}

func TestOverlappingLegacyHPAsWarning(t *testing.T) {
	g, _, err := osgraphtest.BuildGraph("./../../../graph/genericgraph/test/overlapping-hpas-legacy.yaml")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	kubegraph.AddHPAScaleRefEdges(g)
	appsgraph.AddAllDeploymentConfigsDeploymentEdges(g)

	markers := FindOverlappingHPAs(g, osgraph.DefaultNamer)
	if len(markers) != 8 {
		t.Fatalf("expected to find eight overlapping HPA markers, got %d", len(markers))
	}

	for _, marker := range markers {
		if actual, expected := marker.Severity, osgraph.WarningSeverity; actual != expected {
			t.Errorf("expected overlapping HPAs to be %v, got %v", expected, actual)
		}

		if actual, expected := marker.Key, HPAOverlappingScaleRefWarning; actual != expected {
			t.Errorf("expected marker type %v, got %v", expected, actual)
		}
	}
}
