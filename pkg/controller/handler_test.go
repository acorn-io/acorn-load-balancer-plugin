package controller

import (
	"testing"

	"github.com/acorn-io/baaah/pkg/router/tester"
	"github.com/stretchr/testify/assert"
	"github.com/tylerslaton/acorn-load-balancer-plugin/pkg/scheme"
	"k8s.io/client-go/kubernetes/fake"
)

func TestHandler_AddAnnotations(t *testing.T) {
	cases := []struct {
		name         string
		testdata     string
		shouldChange bool
	}{
		{
			name:         "valid",
			testdata:     "testdata/service-lb",
			shouldChange: true,
		},
		{
			name:         "non-lb",
			testdata:     "testdata/service-non-lb",
			shouldChange: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			harness, input, err := tester.FromDir(scheme.Scheme, tt.testdata)
			if err != nil {
				t.Fatal(err)
			}

			existingAnnotations := input.GetAnnotations()

			expected := existingAnnotations
			expected["foo"] = "bar"

			h := Handler{
				client:      fake.NewSimpleClientset(input),
				annotations: map[string]string{"foo": "bar"},
			}

			req := tester.NewRequest(t, harness.Scheme, input, harness.Existing...)

			if err := h.AddAnnotations(req, nil); err != nil {
				t.Fatal(err)
			}

			if tt.shouldChange {
				assert.Equal(t, expected, input.GetAnnotations())
			} else {
				assert.Equal(t, existingAnnotations, input.GetAnnotations())
			}
		})
	}
}
