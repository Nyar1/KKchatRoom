package tools

import (
	"fmt"
	"go.yaml.in/yaml/v3"
	"os"
	"strconv"
)

const ConfigPath = "base/configResources/serverConfig.yaml"

// Config 保存应用配置。
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	IP string `yaml:"-"`
}

// LoadConfig 从指定路径读取并解析配置。
func LoadConfig(path string) (Config, error) {
	var config Config
	defer func() {
		if config.IP == "" {
			config.IP = "127.0.0.1:8080"
		}
	}()
	path = ConfigPath
	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("读取配置文件失败: %w", err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, fmt.Errorf("解析配置文件失败: %w", err)
	}

	if config.Server.Port < 1 || config.Server.Port > 65535 {
		return config, fmt.Errorf("端口必须在 1～65535 之间")
	}
	config.IP = config.Server.Host + ":" + strconv.Itoa(config.Server.Port)
	return config, nil
}
