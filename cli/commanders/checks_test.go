package commanders_test

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/xerrors"

	"github.com/greenplum-db/gpupgrade/cli/commanders"
	"github.com/greenplum-db/gpupgrade/idl"
	"github.com/greenplum-db/gpupgrade/idl/mock_idl"
	"github.com/greenplum-db/gpupgrade/utils/disk"

	"github.com/golang/mock/gomock"
)

func TestCheckVersion(t *testing.T) {
	t.Run("it prints out version check is OK and that check version request was processed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock_idl.NewMockCliToHubClient(ctrl)
		client.EXPECT().CheckVersion(
			gomock.Any(),
			&idl.CheckVersionRequest{},
		).Return(&idl.CheckVersionReply{IsVersionCompatible: true}, nil)
		err := commanders.CheckVersion(client)

		if err != nil {
			t.Errorf("No error was expected, but got %#v", err)
		}
	})

	t.Run("it prints out version check failed and that check version request was processed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock_idl.NewMockCliToHubClient(ctrl)
		client.EXPECT().CheckVersion(
			gomock.Any(),
			&idl.CheckVersionRequest{},
		).Return(&idl.CheckVersionReply{IsVersionCompatible: false}, nil)
		err := commanders.CheckVersion(client)

		if err == nil {
			t.Fatalf("An error was expected, but got nil")
		}
		expectedError := "Version Compatibility Check Failed"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error to contain %q but got: %#v", expectedError, err)
		}
	})

	t.Run("it prints out that it was unable to connect to hub", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		client := mock_idl.NewMockCliToHubClient(ctrl)
		client.EXPECT().CheckVersion(
			gomock.Any(),
			&idl.CheckVersionRequest{},
		).Return(&idl.CheckVersionReply{IsVersionCompatible: false}, errors.New("something went wrong"))
		err := commanders.CheckVersion(client)

		if err == nil {
			t.Fatalf("An error was expected, but got nil")
		}
		expectedError := "gRPC call to hub failed"
		if !strings.Contains(err.Error(), expectedError) {
			t.Errorf("Expected error to contain %q but got: %#v", expectedError, err)
		}
	})
}

func TestFormatBytes(t *testing.T) {
	cases := []struct {
		bytes    uint64
		expected string
	}{
		// Testing units
		{42, "42 KiB"},
		{42 * 1024, "42 MiB"},
		{42 * 1024 * 1024, "42 GiB"},
		{42 * 1024 * 1024 * 1024, "42 TiB"},
		{42 * 1024 * 1024 * 1024 * 1024, "42 PiB"},
		{42 * 1024 * 1024 * 1024 * 1024 * 1024, "42 EiB"},
		// Testing rounding
		{0, "0 KiB"},
		{4200, "4.102 MiB"},
		{4291, "4.19 MiB"},
		{4300, "4.199 MiB"},
		{12636, "12.34 MiB"},
		{126362, "123.4 MiB"},
		{1048064, "1024 MiB"},
	}
	for _, c := range cases {
		actual := commanders.FormatBytes(c.bytes)
		if actual != c.expected {
			t.Errorf("FormatBytes(%d)=%q, want %q", c.bytes, actual, c.expected)
		}
	}
}

func TestDiskSpaceError(t *testing.T) {
	err := &commanders.DiskSpaceError{
		Failed: disk.SpaceFailures{
			"sdw1: /": {
				Available: 15,
				Required:  20,
			},
			"sdw1: /proc": {
				Available: 15,
				Required:  20,
			},
			"mdw: /": {
				Available: 1024,
				Required:  2048,
			},
		},
	}

	rows := err.Table()

	expected := [][]string{
		{"Hostname", "Filesystem", "Shortfall", "Available", "Required"},
		{"mdw", "/", commanders.FormatBytes(1024), commanders.FormatBytes(1024), commanders.FormatBytes(2048)},
		{"sdw1", "/", commanders.FormatBytes(5), commanders.FormatBytes(15), commanders.FormatBytes(20)},
		{"sdw1", "/proc", commanders.FormatBytes(5), commanders.FormatBytes(15), commanders.FormatBytes(20)},
	}
	if !reflect.DeepEqual(expected, rows) {
		t.Errorf("got table %q, want %q", rows, expected)
	}
}

func TestDiskSpaceCheck(t *testing.T) {
	cases := []struct {
		name    string
		failed  disk.SpaceFailures
		grpcErr error
	}{
		{"reports completion on success",
			disk.SpaceFailures{},
			nil,
		},
		{"reports failure when hub returns full disks",
			disk.SpaceFailures{
				"mdw": {Required: 300, Available: 1},
			},
			nil,
		},
		{"reports failure on gRPC error",
			disk.SpaceFailures{},
			errors.New("gRPC failure"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// exact value doesn't matter; we simply verify that it's passed
			// through to gRPC as-is
			ratio := float64(0.5)

			client := mock_idl.NewMockCliToHubClient(ctrl)
			client.EXPECT().CheckDiskSpace(
				gomock.Any(),
				&idl.CheckDiskSpaceRequest{Ratio: ratio},
			).Return(&idl.CheckDiskSpaceReply{Failed: c.failed}, c.grpcErr)

			d := bufferStandardDescriptors(t)
			defer d.Close()

			err := commanders.CheckDiskSpace(client, ratio)
			actualOut, _ := d.Collect()

			expectedStatus := idl.Status_FAILED
			switch {
			case c.grpcErr != nil:
				if !xerrors.Is(err, c.grpcErr) {
					t.Errorf("returned error %#v, want %#v", err, c.grpcErr)
				}

			case len(c.failed) != 0:
				var diskSpaceError commanders.DiskSpaceError
				if !xerrors.As(err, &diskSpaceError) {
					t.Errorf("returned error %#v, want a DiskSpaceError", err)
				} else if !reflect.DeepEqual(diskSpaceError.Failed, c.failed) {
					t.Errorf("error contents were %v, want %v", diskSpaceError.Failed, c.failed)
				}

			default:
				expectedStatus = idl.Status_COMPLETE
				if err != nil {
					t.Errorf("returned error %#v, expected no error", err)
				}
			}

			expected := commanders.Format("Checking disk space...", expectedStatus)
			if !strings.Contains(string(actualOut), expected) {
				t.Errorf("Expected string %q to contain %q", actualOut, expected)
			}
		})
	}
}
