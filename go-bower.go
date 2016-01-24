package main

import (
	"github.com/cyberlight/go-bower/configuration"
	"github.com/spf13/viper"
)

func main() {
	bowerRc := viper.New()
	bowerRcConfig := configuration.NewBowerConfigReader(bowerRc)
	bowerRcConfig.Read("")
}
