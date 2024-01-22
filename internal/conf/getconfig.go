package conf

import (
	"github.com/spf13/viper"

	"github.com/tnvmadhav/ExerciseDiary/internal/check"
	"github.com/tnvmadhav/ExerciseDiary/internal/models"
)

// Get - read config from file or env
func Get(path string) models.Conf {
	var config models.Conf

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8851")
	viper.SetDefault("THEME", "grass")
	viper.SetDefault("COLOR", "light")
	viper.SetDefault("HEATCOLOR", "#03a70c")
	viper.SetDefault("PAGESTEP", 10)

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)
	config.Color, _ = viper.Get("COLOR").(string)
	config.HeatColor, _ = viper.Get("HEATCOLOR").(string)
	config.PageStep = viper.GetInt("PAGESTEP")

	return config
}

// Write - write config to file
func Write(config models.Conf) {

	viper.SetConfigFile(config.ConfPath)
	viper.SetConfigType("yaml")

	viper.Set("host", config.Host)
	viper.Set("port", config.Port)
	viper.Set("theme", config.Theme)
	viper.Set("color", config.Color)
	viper.Set("heatcolor", config.HeatColor)
	viper.Set("pagestep", config.PageStep)

	err := viper.WriteConfig()
	check.IfError(err)
}
