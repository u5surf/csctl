package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/provision/types"
)

var (
	asps = []types.AutoscalingPolicy{
		{
			Name:           strptr("test1"),
			ID:             types.UUID("1234"),
			MetricsBackend: "prometheus",
			Metric:         strptr("cpu"),
			Policy: &types.ScalingPolicy{
				ScaleUp: &types.ScalingPolicyConfiguration{
					AdjustmentType:     strptr("percent"),
					AdjustmentValue:    int32ptr(10),
					ComparisonOperator: strptr(">="),
					Threshold:          float32ptr(0.5),
				},
				ScaleDown: &types.ScalingPolicyConfiguration{
					AdjustmentType:     strptr("percent"),
					AdjustmentValue:    int32ptr(30),
					ComparisonOperator: strptr("<"),
					Threshold:          float32ptr(0.3),
				},
			},
			PollInterval: int32ptr(15),
			SamplePeriod: int32ptr(600),
		},
		{
			Name:           strptr("test2"),
			ID:             types.UUID("4321"),
			MetricsBackend: "prometheus",
			Metric:         strptr("memory"),
			Policy: &types.ScalingPolicy{
				ScaleUp: &types.ScalingPolicyConfiguration{
					AdjustmentType:     strptr("absolute"),
					AdjustmentValue:    int32ptr(1),
					ComparisonOperator: strptr(">"),
					Threshold:          float32ptr(0.8),
				},
				ScaleDown: &types.ScalingPolicyConfiguration{
					AdjustmentType:     strptr("aboslute"),
					AdjustmentValue:    int32ptr(2),
					ComparisonOperator: strptr("<="),
					Threshold:          float32ptr(0.2),
				},
			},
			PollInterval: int32ptr(20),
			SamplePeriod: int32ptr(800),
		},
		{
			Name:           strptr("test3"),
			ID:             types.UUID("3214"),
			MetricsBackend: "prometheus",
			Metric:         strptr("memory"),
			Policy:         nil,
			PollInterval:   int32ptr(20),
			SamplePeriod:   int32ptr(800),
		},
	}
)

func TestNewAutoscalingPolicies(t *testing.T) {
	a := NewAutoscalingPolicies(nil)
	assert.NotNil(t, a)

	a = NewAutoscalingPolicies(asps)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(asps))

	a = AutoscalingPolicy()
	assert.NotNil(t, a)
}

func TestAutoscalingPoliciesTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewAutoscalingPolicies(asps)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(asps), info.numRows)
}
