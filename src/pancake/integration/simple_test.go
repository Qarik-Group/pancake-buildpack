package integration_test

import (
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack/cutlass"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Simple Integration Test", func() {
	var app *cutlass.App
	AfterEach(func() {
		if app != nil {
			// app.Destroy()
		}
		app = nil
	})

	It("app deploys", func() {
		app = cutlass.New(filepath.Join(bpDir, "fixtures", "simple_test"))
		PushAppAndConfirm(app)
		Expect(app.GetBody("/")).To(ContainSubstring("Something on your website"))
	})
})
