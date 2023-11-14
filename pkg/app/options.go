package app

import (
	"log"

	"github.com/khulnasoft-lab/go-application-framework/internal/utils"
	"github.com/khulnasoft-lab/go-application-framework/pkg/configuration"
	"github.com/khulnasoft-lab/go-application-framework/pkg/workflow"
	"github.com/rs/zerolog"
)

type Opts func(engine workflow.Engine)

func WithLogger(logger *log.Logger) Opts {
	return func(engine workflow.Engine) {
		console := &zerolog.ConsoleWriter{
			Out:        &utils.ToLog{Logger: logger},
			NoColor:    true,
			PartsOrder: []string{zerolog.MessageFieldName},
		}
		log := zerolog.New(console)
		engine.SetLogger(&log)
	}
}

func WithConfiguration(config configuration.Configuration) Opts {
	return func(engine workflow.Engine) {
		engine.SetConfiguration(config)
	}
}

func WithZeroLogger(logger *zerolog.Logger) Opts {
	return func(engine workflow.Engine) {
		engine.SetLogger(logger)
	}
}

func WithInitializers(initializers ...workflow.ExtensionInit) Opts {
	return func(engine workflow.Engine) {
		for _, i := range initializers {
			engine.AddExtensionInitializer(i)
		}
	}
}
