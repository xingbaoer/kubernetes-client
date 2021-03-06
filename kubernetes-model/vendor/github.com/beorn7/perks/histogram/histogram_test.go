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
package histogram

import (
	"math/rand"
	"testing"
)

func TestHistogram(t *testing.T) {
	const numPoints = 1e6
	const maxBins = 3

	h := New(maxBins)
	for i := 0; i < numPoints; i++ {
		f := rand.ExpFloat64()
		h.Insert(f)
	}

	bins := h.Bins()
	if g := len(bins); g > maxBins {
		t.Fatalf("got %d bins, wanted <= %d", g, maxBins)
	}

	for _, b := range bins {
		t.Logf("%+v", b)
	}

	if g := count(h.Bins()); g != numPoints {
		t.Fatalf("binned %d points, wanted %d", g, numPoints)
	}
}

func count(bins Bins) int {
	binCounts := 0
	for _, b := range bins {
		binCounts += b.Count
	}
	return binCounts
}
