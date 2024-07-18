// This is implemented based on idea from Temporal
// https://github.com/temporalio/temporal/blob/master/common/config/loader.go

package config

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const (
	baseFile             = "config.yml"
	envDevelopment       = "development"
	defaultConfigRootDir = "etc"
	configDir            = "config"
)

// LoadConfig loads and validates configurations
// The hierarchy is as follows from lowest to highest
//
//	base.yaml
//	    env.yaml   -- environment is one of the input params ex: development
func LoadConfig(ctx context.Context, rootDir string) (*Config, error) {
	log.Println(ctx, "env", envDevelopment, "root dir", rootDir)
	if len(rootDir) == 0 {
		rootDir = defaultConfigRootDir
	}
	files, err := getConfigFiles(ctx, "config", path(rootDir, configDir))
	if err != nil {
		return nil, err
	}

	config := Config{}
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			log.Panic(ctx, "cannot read file", f)
			return nil, err
		}

		if err := yaml.Unmarshal(data, &config); err != nil {
			log.Panic(ctx, "cannot decode file", f)
			return nil, err
		}
	}

	return &config, nil
}

// getConfigFiles returns the list of config files to process in the hierarchy order
func getConfigFiles(ctx context.Context, env string, configDir string) ([]string, error) {
	candidates := []string{
		path(configDir, baseFile),
		path(configDir, file(env, "yml")),
	}

	result := make([]string, 0, len(candidates))
	for _, c := range candidates {
		if _, err := os.Stat(c); err != nil {
			log.Panic(ctx, "file not exist", c)
		}

		result = append(result, c)
	}

	if len(result) == 0 {
		err := fmt.Errorf("no config files found within %v", configDir)
		log.Panic(ctx, err.Error())
		return nil, err
	}

	return result, nil
}

func file(name string, suffix string) string {
	return name + "." + suffix
}

func path(dir string, file string) string {
	return dir + "/" + file
}
