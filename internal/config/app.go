package config

type Application struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"DEBUG"`

	ReportTemplatesPath string `env:"REPORT_TEMPLATES_PATH" envDefault:"forms"`
	Report8949          string `env:"REPORT_8949" envDefault:"f8949.pdf"`

	REST REST
}
