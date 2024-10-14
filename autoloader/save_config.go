package autoloader

import (
	"github.com/kyaxcorp/go-helper/file"
	// cassandraConfig "github.com/kyaxcorp/go-db/driver/cassandra/config"

	cfgData "github.com/kyaxcorp/go-config/data"
	"github.com/kyaxcorp/go-helper/errors2"
	"github.com/kyaxcorp/go-helper/hash"

	"github.com/spf13/viper"
)

func SaveConfigFromMemory(cfg Config) error {
	c := viper.New()

	var _err error

	// Setting the main config
	c.Set("main", cfgData.MainConfig)
	// Setting the custom config
	c.Set("custom", cfg.CustomConfig)

	// TODO: save config only by comparing if it's different!
	// If it's diff, then overwrite it!
	configPath := GetConfigFilePath()
	if configPath == "" {
		return errors2.New(0, "config path is empty")
	}

	configTmpPath := GetConfigTmpFilePath()
	if configTmpPath == "" {
		return errors2.New(0, "config tmp path is empty")
	}
	// Save the temporary config file
	_err = c.WriteConfigAs(configTmpPath)
	if _err != nil {
		// log.Println("Failed to generate config!")
		return _err
	}

	// Compare the 2 configs
	tmpConfigHash, _err := hash.FileSha256(configTmpPath)
	if _err != nil {
		return _err
	}
	// Delete the tmp config
	_, _err = file.Delete(configTmpPath)
	if _err != nil {
		return _err
	}

	// CHeck if config exists
	isConfigExists := IsConfigExists()
	if isConfigExists {
		realConfigHash, _err := hash.FileSha256(configPath)
		if _err != nil {
			return _err
		}
		// log.Println(realConfigHash, tmpConfigHash)

		// Compare the 2 configs
		if tmpConfigHash == realConfigHash {
			// It's the same configuration!
			// log.Println("Same config!!! skipping save...")
			return nil
		}
	}

	// Save the real config file
	_err = c.WriteConfigAs(configPath)
	//_err = c.SafeWriteConfigAs(configPath)
	if _err != nil {
		// log.Println("Failed to generate config!")
		return _err
	}
	return nil
}
