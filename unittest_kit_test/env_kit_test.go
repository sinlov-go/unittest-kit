package unittest_kit_test

import (
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/sinlov-go/unittest-kit/unittest_env_kit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMustHasEnvSetByArray(t *testing.T) {
	// mock MustHasEnvSetByArray

	unittest_env_kit.SetEnvBool(t, keyEnvDebug, true)

	unittest_env_kit.SetEnvInt64(t, keyEnvCiNum, 2)

	unittest_env_kit.SetEnvStr(t, keyEnvCiKey, "foo")

	type args struct {
		envs []string
	}
	tests := []struct {
		name    string
		args    args
		wantRes bool
	}{
		{
			name: "has",
			args: args{
				envs: []string{
					keyEnvDebug,
					keyEnvCiNum,
				},
			},
			wantRes: false,
		},
		{
			name: "has",
			args: args{
				envs: []string{
					"foo",
					"bar",
				},
			},
			wantRes: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do MustHasEnvSetByArray

			gotResult := env_kit.MustHasEnvSetByArray(t, tc.args.envs)

			// verify MustHasEnvSetByArray
			assert.Equal(t, tc.wantRes, gotResult)
		})
	}
}
