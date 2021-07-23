package config

import (
	"bytes"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"strings"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/config/medadata/report"
	"dubbo.apache.org/dubbo-go/v3/config/metric"
	"dubbo.apache.org/dubbo-go/v3/config/registry"
)

// RootConfig is the root config
type RootConfig struct {
	ConfigCenter *CenterConfig `yaml:"config-center" json:"config-center,omitempty"`

	// since 1.5.0 version
	//Remotes              map[string]*config.RemoteConfig `yaml:"remote" json:"remote,omitempty" property:"remote"`

	ServiceDiscoveries map[string]*ServiceDiscoveryConfig `yaml:"service-discovery" json:"service-discovery,omitempty" property:"service_discovery"`

	MetadataReportConfig *report.Config `yaml:"metadata_report" json:"metadata-report,omitempty" property:"metadata-report"`

	// Application application config
	Application *ApplicationConfig `yaml:"application" json:"application,omitempty" property:"application"`

	// Registries registry config
	Registries map[string]*registry.Config `yaml:"registries" json:"registries" property:"registries"`

	Protocols map[string]*ProtocolConfig `yaml:"protocols" json:"protocols" property:"protocols"`

	// prefix              string
	fatherConfig        interface{}
	EventDispatcherType string               `default:"direct" yaml:"event_dispatcher_type" json:"event_dispatcher_type,omitempty"`
	MetricConfig        *metric.MetricConfig `yaml:"metrics" json:"metrics,omitempty"`
	fileStream          *bytes.Buffer

	// validate
	//Validate *validator.Validate
	// cache file used to store the current used configurations.
	CacheFile string `yaml:"cache_file" json:"cache_file,omitempty" property:"cache_file"`
}

// Prefix dubbo
func (RootConfig) Prefix() string {
	return constant.DUBBO
}

func verification(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Error())
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}
