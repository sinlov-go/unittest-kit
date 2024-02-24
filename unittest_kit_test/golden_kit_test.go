package unittest_kit_test

import (
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestGoldenKitUse(t *testing.T) {
	// mock GoldenKitUse
	type args struct {
		testGoldenKit unittest_file_kit.TestGoldenKit
	}
	type compArgs struct {
		testBaseFolder string
		testDataFolder string
	}
	tests := []struct {
		name     string
		args     args
		compArgs compArgs
	}{
		{
			name: "sample",
			args: args{
				testGoldenKit: *testGoldenKit,
			},
			compArgs: compArgs{
				testBaseFolder: testBaseFolderPath,
				testDataFolder: filepath.Join(testBaseFolderPath, "testdata"),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do GoldenKitUse

			// verify GoldenKitUse
			assert.Equal(t, tc.compArgs.testBaseFolder, tc.args.testGoldenKit.GetTestBaseFolderFullPath())
			assert.Equal(t, tc.compArgs.testDataFolder, tc.args.testGoldenKit.GetTestDataFolderFullPath())
		})
	}
}
