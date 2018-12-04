package resource

import (
	"fmt"
	"io"

	"github.com/containership/csctl/cloud/provision/types"
	"github.com/containership/csctl/resource/table"
)

// AutoscalingPolicies is a list of the associated cloud resource with additional functionality
type AutoscalingPolicies struct {
	resource
	items []types.AutoscalingPolicy
}

// NewAutoscalingPolicies constructs a new AutoscalingPolicies wrapping the given cloud type
func NewAutoscalingPolicies(items []types.AutoscalingPolicy) *AutoscalingPolicies {
	return &AutoscalingPolicies{
		resource: resource{
			name:    "autoscaling-policy",
			plural:  "autoscaling-policies",
			aliases: []string{"asp", "asps"},
		},
		items: items,
	}
}

// AutoscalingPolicy constructs a new AutoscalingPolicies with no underlying items, useful for
// interacting with the metadata itself.
func AutoscalingPolicy() *AutoscalingPolicies {
	return NewAutoscalingPolicies(nil)
}

func (p *AutoscalingPolicies) columns() []string {
	return []string{
		"Name",
		"ID",
		// TODO add scale up/down policies
		"Metrics Backend",
		"Metric",
		"Scale Up",
		"Scale Down",
		"Poll Interval",
		"Sample Period",
	}
}

// Table outputs the table representation to the given writer
func (p *AutoscalingPolicies) Table(w io.Writer) error {
	table := table.New(w, p.columns())

	for _, asp := range p.items {
		var upPercent string
		up := asp.Policy.ScaleUp
		if *up.AdjustmentType == "percent" {
			upPercent = "%"
		}
		scaleUpPolicy := fmt.Sprintf("%s %.2f: +%d%s",
			*up.ComparisonOperator, *up.Threshold, *up.AdjustmentValue, upPercent)

		var downPercent string
		down := asp.Policy.ScaleDown
		if *down.AdjustmentType == "percent" {
			downPercent = "%"
		}
		scaleDownPolicy := fmt.Sprintf("%s %.2f: -%d%s",
			*down.ComparisonOperator, *down.Threshold, *down.AdjustmentValue, downPercent)

		table.Append([]string{
			*asp.Name,
			string(asp.ID),
			asp.MetricsBackend,
			*asp.Metric,
			scaleUpPolicy,
			scaleDownPolicy,
			fmt.Sprintf("%ds", *asp.PollInterval),
			fmt.Sprintf("%ds", *asp.SamplePeriod),
		})
	}

	table.Render()

	return nil
}

// JSON outputs the JSON representation to the given writer
func (p *AutoscalingPolicies) JSON(w io.Writer) error {
	return displayJSON(w, p.items)
}

// YAML outputs the YAML representation to the given writer
func (p *AutoscalingPolicies) YAML(w io.Writer) error {
	return displayYAML(w, p.items)
}

// JSONPath outputs the executed JSONPath template to the given writer
func (p *AutoscalingPolicies) JSONPath(w io.Writer, template string) error {
	return displayJSONPath(w, template, p.items)
}
