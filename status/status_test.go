// Copyright (c) 2018 Palantir Technologies. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package status

import (
	"context"
	"testing"

	"github.com/palantir/witchcraft-go-server/conjure/sls/spec/health"
	"github.com/stretchr/testify/assert"
)

type testHealthCheckSource struct {
	healthStatus health.HealthStatus
}

func (t *testHealthCheckSource) HealthStatus(_ context.Context) health.HealthStatus {
	return t.healthStatus
}

func TestCombinedHealthCheckSource(t *testing.T) {
	sourceA := &testHealthCheckSource{
		healthStatus: health.HealthStatus{
			Checks: map[health.CheckType]health.HealthCheckResult{
				"a": {
					State: health.HealthStateHealthy,
				},
				"b": {
					State: health.HealthStateHealthy,
				},
			},
		},
	}
	sourceB := &testHealthCheckSource{
		healthStatus: health.HealthStatus{
			Checks: map[health.CheckType]health.HealthCheckResult{
				"a": {
					State: health.HealthStateError,
				},
				"c": {
					State: health.HealthStateHealthy,
				},
			},
		},
	}
	combined := NewCombinedHealthCheckSource(sourceA, sourceB)
	actual := combined.HealthStatus(context.Background())
	assert.Equal(t, health.HealthStatus{
		Checks: map[health.CheckType]health.HealthCheckResult{
			"a": {
				State: health.HealthStateError,
			},
			"b": {
				State: health.HealthStateHealthy,
			},
			"c": {
				State: health.HealthStateHealthy,
			},
		},
	}, actual)
}
