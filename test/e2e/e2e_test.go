package e2e

import (
	goflag "flag"
	"os"
	"testing"

	"github.com/TheWeatherCompany/icm-redis-operator/test/e2e/framework"
	"github.com/spf13/pflag"
)

func TestE2E(t *testing.T) {
	RunE2ETests(t)
}

func TestMain(m *testing.M) {

	pflag.StringVar(&framework.FrameworkContext.KubeConfigPath, "kubeconfig", "", "Path to kubeconfig")
	pflag.StringVar(&framework.FrameworkContext.ImageTag, "image-tag", "local", "image tag")
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()
	goflag.CommandLine.Parse([]string{})

	os.Exit(m.Run())
}
