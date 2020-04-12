package conf

import (
	_ "despair/app"
	"testing"
)

func TestInitConfig(t *testing.T) {
	InitConfig()
	if Conf == nil {
		t.Errorf("init config failed")
	}

	if Conf.Log.LogDir == "" {
		t.Errorf("logdir is empty")
	}

}
