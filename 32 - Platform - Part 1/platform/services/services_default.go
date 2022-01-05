package services

import (
    "platform/logging"
    "platform/config"
)

func RegisterDefaultServices() {

    err := AddSingleton(func() (c config.Configuration) {
        c, loadErr :=  config.Load("config.json")
        if (loadErr != nil) {
            panic(loadErr)
        }
        return
    })

    err = AddSingleton(func(appconfig config.Configuration) logging.Logger {
        return logging.NewDefaultLogger(appconfig)
    })
    if (err != nil) {
        panic(err)
    }
}
