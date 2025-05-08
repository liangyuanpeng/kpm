package client

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"kcl-lang.io/kpm/pkg/downloader"
	pkg "kcl-lang.io/kpm/pkg/package"
	"kcl-lang.io/kpm/pkg/reporter"
)

// go test -timeout 30s -run ^TestPush$ kcl-lang.io/kpm/pkg/client -v
func TestPush(t *testing.T) {
	testFunc := func(t *testing.T, kpmcli *KpmClient) {
		// if runtime.GOOS == "windows" {
		// 	t.Skip("Skipping test on Windows")
		// }
		// err := mock.StartDockerRegistry()
		// if err != nil {
		// 	t.Errorf("Error starting docker registry: %v", err)
		// }

		// defer func() {
		// 	err = mock.CleanTestEnv()
		// 	if err != nil {
		// 		t.Errorf("Error stopping docker registry: %v", err)
		// 	}
		// }()

		kpmcli.SetInsecureSkipTLSverify(true)
		kpmcli.SetPlainHttp(true)
		err := kpmcli.LoginOci("192.168.3.187:5001", "test", "1234")
		if err != nil {
			t.Errorf("Error logging in to docker registry: %v", err)
		}

		var buf bytes.Buffer
		kpmcli.SetLogWriter(&buf)

		testDir := getTestDir("test_push")
		pushedModPath := filepath.Join(testDir, "push_0")

		err = kpmcli.Push(
			WithPushModPath(pushedModPath),
			WithPushSource(
				downloader.Source{
					Oci: &downloader.Oci{
						Reg:  "192.168.3.187:5001",
						Repo: "test/push_0",
					},
				},
			),
		)

		if err != (*reporter.KpmEvent)(nil) {
			t.Errorf("Error pushing kcl package: %v", err)
		}

		assert.Contains(t, buf.String(), "package 'push_0' will be pushed")
		assert.Contains(t, buf.String(), "pushed [registry] 192.168.3.187:5001/test/push_0")
		assert.Contains(t, buf.String(), "digest: sha256:")

		testPushModPath := filepath.Join(testDir, "test_pushed_mod")

		err = kpmcli.Init(
			WithInitModPath(testPushModPath),
		)
		if err != nil {
			t.Errorf("Error initializing kcl package: %v", err)
		}

		testMod, err := pkg.LoadKclPkgWithOpts(
			pkg.WithPath(testPushModPath),
		)

		err = kpmcli.Add(
			WithAddKclPkg(testMod),
			WithAddSourceUrl("oci://192.168.3.187:5001/test/push_0"),
			WithAddModSpec(&downloader.ModSpec{
				Name:    "push_0",
				Version: "0.0.1",
			}),
		)

		if err != nil {
			t.Errorf("Error adding kcl package: %v", err)
		}

	}
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestPush", TestFunc: testFunc}})
}
