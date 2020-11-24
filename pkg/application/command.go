package application

import (
	"dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/api/launcher"
	idcreator "dev-gitlab.wanxingrowth.com/wanxin-go-micro/base/utils/idcreator/snowflake"
	"github.com/spf13/viper"

	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/config"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/address"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/protos"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/rpc/user"
	"dev-gitlab.wanxingrowth.com/fanli/user/pkg/utils/log"
)

func Start() {

	app := launcher.NewApplication(
		launcher.SetApplicationDescription(
			&launcher.ApplicationDescription{
				ShortDescription: "users service",
				LongDescription:  "support user data management function.",
			},
		),
		launcher.SetApplicationLogger(log.GetLogger()),
		launcher.SetApplicationEvents(
			launcher.NewApplicationEvents(
				launcher.SetOnInitEvent(func(app *launcher.Application) {

					unmarshalConfiguration()

					registerUserRPCRouter(app)

					idcreator.InitCreator(app.GetServiceId())
				}),
				launcher.SetOnStartEvent(func(app *launcher.Application) {

					autoMigration()
				}),
			),
		),
	)

	app.Launch()
}

func registerUserRPCRouter(app *launcher.Application) {

	rpcService := app.GetRPCService()
	if rpcService == nil {

		log.GetLogger().WithField("stage", "onInit").Error("get rpc service is nil")
		return
	}

	protos.RegisterUserControllerServer(rpcService.GetRPCConnection(), &user.Controller{})
	protos.RegisterAddressControllerServer(rpcService.GetRPCConnection(), &address.Controller{})
}

func unmarshalConfiguration() {
	err := viper.Unmarshal(config.Config)
	if err != nil {

		log.GetLogger().WithError(err).Error("unmarshal config error")
	}
}
