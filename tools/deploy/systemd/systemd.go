package systemd

import (
	"fmt"

	"github.com/jakewright/home-automation/tools/deploy/config"
)

// Deploy deploys the given service
func Deploy(service *config.Service) error {
	if service == nil {
		return fmt.Errorf("service is nil")
	}
	if service.System != config.SysSystemd {
		return fmt.Errorf("system is not %s", config.SysSystemd)
	}

	switch service.Language {
	case config.LangGo:
		return deployGo(service)
	case config.LangJavaScript:
		return deployJavaScript(service)
	default:
		return fmt.Errorf("unsupported language '%s'", service.Language)
	}
}

func deployJavaScript(service *config.Service) error {
	return nil
}
