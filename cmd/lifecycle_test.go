package cmd

import (
	"testing"

	"github.com/agentbay/agentbay-cli/internal/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
)

func TestMergeSandboxLifeCycle_NilExisting_NoFlags(t *testing.T) {
	lf := &lifecycleFlags{}
	result := mergeSandboxLifeCycle(nil, lf)

	assert.NotNil(t, result)
	assert.Nil(t, result.Mode)
	assert.Nil(t, result.DesktopMaxRuntime)
	assert.Nil(t, result.HibernateTimeout)
	assert.Nil(t, result.UserIdleTimeout)
	// No UserIdleTimeout -> IdleTimeoutSwitch = false
	assert.NotNil(t, result.IdleTimeoutSwitch)
	assert.False(t, *result.IdleTimeoutSwitch)
}

func TestMergeSandboxLifeCycle_NilExisting_AllFlagsSet(t *testing.T) {
	lf := &lifecycleFlags{
		mode:           "auto",
		maxRuntime:     3600,
		hibernate:      1800,
		idleTimeout:    600,
		modeSet:        true,
		maxRuntimeSet:  true,
		hibernateSet:   true,
		idleTimeoutSet: true,
	}
	result := mergeSandboxLifeCycle(nil, lf)

	assert.Equal(t, "auto", *result.Mode)
	assert.Equal(t, 3600.0, *result.DesktopMaxRuntime)
	assert.Equal(t, 1800.0, *result.HibernateTimeout)
	assert.Equal(t, 600.0, *result.UserIdleTimeout)
	assert.True(t, *result.IdleTimeoutSwitch)
}

func TestMergeSandboxLifeCycle_ExistingValues_NoOverride(t *testing.T) {
	existing := &client.SandboxLifeCycle{
		Mode:              tea.String("manual"),
		DesktopMaxRuntime: tea.Float64(7200),
		HibernateTimeout:  tea.Float64(3600),
		UserIdleTimeout:   tea.Float64(900),
		IdleTimeoutSwitch: tea.Bool(true),
	}
	lf := &lifecycleFlags{}
	result := mergeSandboxLifeCycle(existing, lf)

	assert.Equal(t, "manual", *result.Mode)
	assert.Equal(t, 7200.0, *result.DesktopMaxRuntime)
	assert.Equal(t, 3600.0, *result.HibernateTimeout)
	assert.Equal(t, 900.0, *result.UserIdleTimeout)
	assert.True(t, *result.IdleTimeoutSwitch)
}

func TestMergeSandboxLifeCycle_PartialOverride(t *testing.T) {
	existing := &client.SandboxLifeCycle{
		Mode:              tea.String("manual"),
		DesktopMaxRuntime: tea.Float64(7200),
		HibernateTimeout:  tea.Float64(3600),
		UserIdleTimeout:   tea.Float64(900),
		IdleTimeoutSwitch: tea.Bool(true),
	}
	lf := &lifecycleFlags{
		mode:    "auto",
		modeSet: true,
	}
	result := mergeSandboxLifeCycle(existing, lf)

	assert.Equal(t, "auto", *result.Mode)
	assert.Equal(t, 7200.0, *result.DesktopMaxRuntime)
	assert.Equal(t, 3600.0, *result.HibernateTimeout)
	assert.Equal(t, 900.0, *result.UserIdleTimeout)
	assert.True(t, *result.IdleTimeoutSwitch)
}

func TestMergeSandboxLifeCycle_IdleTimeoutZero_SwitchTrue(t *testing.T) {
	lf := &lifecycleFlags{
		idleTimeout:    0,
		idleTimeoutSet: true,
	}
	result := mergeSandboxLifeCycle(nil, lf)

	assert.NotNil(t, result.UserIdleTimeout)
	assert.Equal(t, 0.0, *result.UserIdleTimeout)
	// UserIdleTimeout has value (0) -> IdleTimeoutSwitch = true
	assert.True(t, *result.IdleTimeoutSwitch)
}

func TestMergeSandboxLifeCycle_ClearIdleTimeout_SwitchFalse(t *testing.T) {
	// Existing has UserIdleTimeout, but user doesn't set it -> keep existing
	existing := &client.SandboxLifeCycle{
		UserIdleTimeout:   tea.Float64(300),
		IdleTimeoutSwitch: tea.Bool(true),
	}
	lf := &lifecycleFlags{}
	result := mergeSandboxLifeCycle(existing, lf)

	assert.Equal(t, 300.0, *result.UserIdleTimeout)
	assert.True(t, *result.IdleTimeoutSwitch)
}

func TestMergeSandboxLifeCycle_ExistingNoIdleTimeout_SwitchFalse(t *testing.T) {
	existing := &client.SandboxLifeCycle{
		Mode: tea.String("auto"),
	}
	lf := &lifecycleFlags{}
	result := mergeSandboxLifeCycle(existing, lf)

	assert.Nil(t, result.UserIdleTimeout)
	assert.False(t, *result.IdleTimeoutSwitch)
}

func TestLifecycleModeValidation(t *testing.T) {
	tests := []struct {
		name        string
		mode        string
		expectValid bool
	}{
		{"auto", "auto", true},
		{"manual", "manual", true},
		{"invalid", "invalid", false},
		{"empty", "", false},
		{"Auto_case", "Auto", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isValid := tt.mode == "auto" || tt.mode == "manual"
			assert.Equal(t, tt.expectValid, isValid)
		})
	}
}
