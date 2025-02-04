package platform_test

import (
	"io/ioutil"
	"testing"

	"bls-ftcosi/cothority/log"
	"bls-ftcosi/cothority/simul/platform"
)

var testfile = `Machines = 8
App = "sign"

Ppm, Rounds
2, 30
4, 30`

func TestReadRunfile(t *testing.T) {
	tplat := &TPlat{}

	tmpfile, err := ioutil.TempFile("", "testrun.toml")
	log.ErrFatal(err)
	_, err = tmpfile.Write([]byte(testfile))
	if err != nil {
		log.Fatal("Couldn't write to tmp-file:", err)
	}
	tmpfile.Close()

	tests := platform.ReadRunFile(tplat, tmpfile.Name())
	log.Lvl2(tplat)
	log.Lvlf2("%+v\n", tests[0])
	if tplat.App != "sign" {
		log.Fatal("App should be 'sign'")
	}
	if len(tests) != 2 {
		log.Fatal("There should be 2 tests")
	}
	if tests[0].Get("machines") != "8" {
		log.Fatal("Machines = 8 has not been copied into RunConfig")
	}
}

type TPlat struct {
	App      string
	Machines int
}

func (t *TPlat) Configure(pc *platform.Config)       {}
func (t *TPlat) Build(s string, arg ...string) error { return nil }
func (t *TPlat) Deploy(rc platform.RunConfig) error  { return nil }
func (t *TPlat) Start(...string) error               { return nil }
func (t *TPlat) Stop() error                         { return nil }
func (t *TPlat) Cleanup() error                      { return nil }
func (t *TPlat) Wait() error                         { return nil }
