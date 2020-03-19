package docker_registry

import (
	"io/ioutil"

	"github.com/google/go-containerregistry/pkg/logs"

	"github.com/flant/logboek"
)

var generic *api

func Init(insecureRegistry, skipTlsVerifyRegistry bool) error {
	if logboek.Debug.IsAccepted() {
		logs.Progress.SetOutput(logboek.GetOutStream())
		logs.Warn.SetOutput(logboek.GetErrStream())
		logs.Debug.SetOutput(logboek.GetOutStream())
	} else {
		logs.Progress.SetOutput(ioutil.Discard)
		logs.Warn.SetOutput(ioutil.Discard)
		logs.Debug.SetOutput(ioutil.Discard)
	}

	generic = newAPI(apiOptions{
		InsecureRegistry:      insecureRegistry,
		SkipTlsVerifyRegistry: skipTlsVerifyRegistry,
	})

	return nil
}

func API() *api {
	return generic
}
