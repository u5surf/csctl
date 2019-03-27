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
			ScalingPolicy: &types.ScalingPolicy{
				ScaleUp: &types.ScalingPolicyConfiguration{
					AdjustmentType:     strptr("percent"),
					AdjustmentValue:    float32ptr(10),
					ComparisonOperator: strptr(">="),
					Threshold:          float32ptr(0.5),
				},
				ScaleDown: &types.ScalingPolicyConfiguration{
					AdjustmentType:     strptr("percent"),
					AdjustmentValue:    float32ptr(30),
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
			ScalingPolicy: &types.ScalingPolicy{
				ScaleUp: &types.ScalingPolicyConfiguration{
					AdjustmentType:     strptr("absolute"),
					AdjustmentValue:    float32ptr(1),
					ComparisonOperator: strptr(">"),
					Threshold:          float32ptr(0.8),
				},
				ScaleDown: &types.ScalingPolicyConfiguration{
					AdjustmentType:     strptr("aboslute"),
					AdjustmentValue:    float32ptr(2),
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
			ScalingPolicy:  nil,
			PollInterval:   int32ptr(20),
			SamplePeriod:   int32ptr(800),
		},
	}
	aspsSingle = []types.AutoscalingPolicy{
		{
			Name:           strptr("test4"),
			ID:             types.UUID("1234"),
			MetricsBackend: "prometheus",
			Metric:         strptr("cpu"),
			ScalingPolicy: &types.ScalingPolicy{
				ScaleUp: &types.ScalingPolicyConfiguration{
					AdjustmentType:     strptr("percent"),
					AdjustmentValue:    float32ptr(10),
					ComparisonOperator: strptr(">="),
					Threshold:          float32ptr(0.5),
				},
				ScaleDown: &types.ScalingPolicyConfiguration{
					AdjustmentType:     strptr("percent"),
					AdjustmentValue:    float32ptr(30),
					ComparisonOperator: strptr("<"),
					Threshold:          float32ptr(0.3),
				},
			},
			PollInterval: int32ptr(15),
			SamplePeriod: int32ptr(600),
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

func TestAutoscalingPoliciesDisableListView(t *testing.T) {
	a := NewAutoscalingPolicies(aspsSingle)
	assert.NotNil(t, a)
	a.resource.DisableListView()
	assert.Equal(t, a.resource.listView, false)
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

func TestAutoscalingPoliciesJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewAutoscalingPolicies(aspsSingle)
	err := a.JSON(buf)
	assert.Nil(t, err)
	a.resource.DisableListView()
	err = a.JSON(buf)
	assert.Nil(t, err)
}

func TestAutoscalingPoliciesYAML(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewAutoscalingPolicies(aspsSingle)
	err := a.YAML(buf)
	assert.Nil(t, err)
	a.resource.DisableListView()
	err = a.YAML(buf)
	assert.Nil(t, err)
}
