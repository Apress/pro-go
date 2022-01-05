package config

func (c *DefaultConfig) GetStringDefault(name, val string) (result string) {
    result, ok := c.GetString(name)
    if !ok {
        result = val
    }
    return
}

func (c *DefaultConfig) GetIntDefault(name string, val int) (result int) {
    result, ok := c.GetInt(name)
    if !ok {
        result = val
    }
    return
}

func (c *DefaultConfig) GetBoolDefault(name string, val bool) (result bool) {
    result, ok := c.GetBool(name)
    if !ok {
        result = val
    }
    return 
}

func (c *DefaultConfig) GetFloatDefault(name string, val float64) (result float64) {
    result, ok := c.GetFloat(name)
    if !ok {
        result = val
    }
    return
}
