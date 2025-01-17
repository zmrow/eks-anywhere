package config_test

import (
	"os"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/aws/eks-anywhere/pkg/config"
)

type testSetup struct {
	*WithT
}

func newTest(t *testing.T) *testSetup {
	return &testSetup{
		WithT: NewWithT(t),
	}
}

func TestGetMaxWaitPerMachineDefault(t *testing.T) {
	tt := newTest(t)

	maxWaitPerMachine := config.GetMaxWaitPerMachine()
	tt.Expect(maxWaitPerMachine).To(Equal(10 * time.Minute))
}

func TestGetMaxWaitPerMachineFromValidEnv(t *testing.T) {
	tt := newTest(t)

	oldEnv := os.Getenv(config.EksaReplicasReadyTimeoutEnv)
	os.Setenv(config.EksaReplicasReadyTimeoutEnv, "15m")
	defer os.Setenv(config.EksaReplicasReadyTimeoutEnv, oldEnv)

	maxWaitPerMachine := config.GetMaxWaitPerMachine()
	tt.Expect(maxWaitPerMachine).To(Equal(15 * time.Minute))
}

func TestGetMaxWaitPerMachineFromInvalidEnv(t *testing.T) {
	tt := newTest(t)

	oldEnv := os.Getenv(config.EksaReplicasReadyTimeoutEnv)
	os.Setenv(config.EksaReplicasReadyTimeoutEnv, "15x")
	defer os.Setenv(config.EksaReplicasReadyTimeoutEnv, oldEnv)

	maxWaitPerMachine := config.GetMaxWaitPerMachine()
	tt.Expect(maxWaitPerMachine).To(Equal(10 * time.Minute))
}

func TestGetExternalEtcdTimeoutDefault(t *testing.T) {
	tt := newTest(t)

	externalEtcdTimeout := config.GetExternalEtcdTimeout()
	tt.Expect(externalEtcdTimeout).To(Equal("60m"))
}

func TestGetExternalEtcdTimeoutFromValidEnv(t *testing.T) {
	tt := newTest(t)

	oldEnv := os.Getenv(config.ExternalEtcdTimeoutEnv)
	os.Setenv(config.ExternalEtcdTimeoutEnv, "15m")
	defer os.Setenv(config.ExternalEtcdTimeoutEnv, oldEnv)

	externalEtcdTimeout := config.GetExternalEtcdTimeout()
	tt.Expect(externalEtcdTimeout).To(Equal("15m"))
}

func TestGetExternalEtcdTimeoutFromInvalidEnv(t *testing.T) {
	tt := newTest(t)

	oldEnv := os.Getenv(config.ExternalEtcdTimeoutEnv)
	os.Setenv(config.ExternalEtcdTimeoutEnv, "15x")
	defer os.Setenv(config.ExternalEtcdTimeoutEnv, oldEnv)

	externalEtcdTimeout := config.GetExternalEtcdTimeout()
	tt.Expect(externalEtcdTimeout).To(Equal("60m"))
}
