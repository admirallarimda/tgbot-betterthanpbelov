package mybot

import "log"
import "github.com/admirallarimda/tgbotbase"

import "github.com/admirallarimda/tgbot-betterthanpbelov/mybot/commandhandler"

func Start(cfg_filename string) error {
	log.Print("Starting my bot")

	fullcfg, err := NewConfig(cfg_filename)
	if err != nil {
		log.Printf("My bot cannot be sarted due to error: %s", err)
		return err
	}

	log.Printf("Starting bot with full config: %+v", fullcfg)

	tgcfg := tgbotbase.Config{TGBot: fullcfg.TGBot,
		Proxy_SOCKS5: fullcfg.Proxy_SOCKS5}
	bot := tgbotbase.NewBot(tgcfg)

	rediscfg := fullcfg.Redis
	redispool := tgbotbase.NewRedisPool(rediscfg)
	propstorage := tgbotbase.NewRedisPropertyStorage(redispool)
	remindstorage := cmd.NewRedisReminderStorage(redispool)

	cron := tgbotbase.NewCron()

	bot.AddHandler(tgbotbase.NewIncomingMessageDealer(cmd.NewPropertyHandler(propstorage)))
<<<<<<< HEAD
	bot.AddHandler(tgbotbase.NewIncomingMessageDealer(cmd.NewWeatherHandler(fullcfg.Weather.Token, redispool, propstorage)))
	bot.AddHandler(tgbotbase.NewIncomingMessageDealer(cmd.NewRemindHandler(cron, remindstorage, propstorage)))
=======
<<<<<<< HEAD
	bot.AddHandler(tgbotbase.NewIncomingMessageDealer(cmd.NewRemindHandler(cron, remindstorage)))
>>>>>>> 8bddb4f2c18ab5ed87e6c866e626614ad169ef20
	bot.AddHandler(tgbotbase.NewBackgroundMessageDealer(cmd.NewKittiesHandler(cron, propstorage)))
=======
	bot.AddHandler(tgbotbase.NewIncomingMessageDealer(cmd.NewWeatherHandler(fullcfg.Weather.Token, redispool, propstorage)))
	bot.AddHandler(tgbotbase.NewIncomingMessageDealer(cmd.NewRemindHandler(cron, remindstorage, propstorage)))
>>>>>>> ac6b42e6a926cc9ee4db23bfe6a69486b5d89a79
	/*     handlers = append(handlers, cmd.NewKittiesHandler(),
	       cmd.NewWeatherHandler(cfg.Weather.Token, opts),
	       cmd.NewDeathHandler(),
	       cmd.NewRemindHandler(notificationChannel))
	*/
	bot.Start()

	log.Print("Stopping my bot")
	return nil
}
