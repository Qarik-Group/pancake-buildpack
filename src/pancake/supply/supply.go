package supply

import (
	"io"
	"path/filepath"
	"os"

	"github.com/cloudfoundry/libbuildpack"
)

type Stager interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/stager.go
	BuildDir() string
	DepDir() string
	DepsIdx() string
	DepsDir() string
	WriteProfileD(string, string) error
}

type Manifest interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/manifest.go
	AllDependencyVersions(string) []string
	DefaultVersion(string) (libbuildpack.Dependency, error)
}

type Installer interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/installer.go
	InstallDependency(libbuildpack.Dependency, string) error
	InstallOnlyVersion(string, string) error
}

type Command interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/command.go
	Execute(string, io.Writer, io.Writer, string, ...string) error
	Output(dir string, program string, args ...string) (string, error)
}

type Supplier struct {
	Manifest  Manifest
	Installer Installer
	Stager    Stager
	Command   Command
	Log       *libbuildpack.Logger
}

func (s *Supplier) Run() error {
	s.Log.BeginStep("Supplying pancake")

	pancake, err := s.Manifest.DefaultVersion("cf-pancake")
	if err != nil {
		return err
	}
	s.Log.Info("Using cf-pancake version %s", pancake.Version)

	if err := s.Installer.InstallDependency(pancake, s.Stager.DepDir()); err != nil {
		return err
	}

	pancakeBin, err := filepath.Glob(filepath.Join(s.Stager.DepDir(), "cf-pancake*"))
	if err != nil {
		return err
	}

	err = os.Rename(pancakeBin[0], filepath.Join(s.Stager.DepDir(), "bin", "cf-pancake"))
	if err != nil {
		return err
	}

	if err := s.Stager.WriteProfileD("finalize_pancake_export.sh", "#!/bin/bash\neval \"$(cf-pancake exports)\""); err != nil {
		s.Log.Error("Unable to write profile.d: %s", err.Error())
		return err
	}
	return nil
}
