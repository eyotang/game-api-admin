package config

type GamePlugin struct {
	Base string `mapstructure:"base" json:"base" yaml:"base"` // 本地文件路径
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径
}
