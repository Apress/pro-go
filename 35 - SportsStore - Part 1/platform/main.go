package main

import (
    "platform/services"
    "platform/placeholder"
)

func main() {
    services.RegisterDefaultServices()
    placeholder.Start()
}
