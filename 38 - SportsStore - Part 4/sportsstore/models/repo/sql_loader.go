package repo

import (
    "os"
    "database/sql"
    "reflect"
    "platform/config"
    "platform/logging"
    _ "modernc.org/sqlite"
)

func openDB(config config.Configuration, logger logging.Logger) (db *sql.DB, 
        commands *SqlCommands, needInit bool) {
    driver := config.GetStringDefault("sql:driver_name", "sqlite")
    connectionStr, found := config.GetString("sql:connection_str")
    if !found {
        logger.Panic("Cannot read SQL connection string from config")
        return
    } 
    if _, err := os.Stat(connectionStr); os.IsNotExist(err) {
        needInit = true
    }
    var err error
    if db, err = sql.Open(driver, connectionStr); err == nil {
        commands = loadCommands(db, config, logger)
    } else {
        logger.Panic(err.Error())
    }
    return   
}

func loadCommands(db *sql.DB, config config.Configuration, 
        logger logging.Logger) (commands *SqlCommands)  {
    commands = &SqlCommands {}
    commandVal := reflect.ValueOf(commands).Elem()
    commandType := reflect.TypeOf(commands).Elem()
    for i := 0; i < commandType.NumField(); i++ {
        commandName := commandType.Field(i).Name
        logger.Debugf("Loading SQL command: %v", commandName)
        stmt := prepareCommand(db, commandName, config, logger)
        commandVal.Field(i).Set(reflect.ValueOf(stmt))
    }
    return commands
}

func prepareCommand(db *sql.DB, command string, config config.Configuration, 
        logger logging.Logger) *sql.Stmt {
    filename, found := config.GetString("sql:commands:" + command)
    if !found {
        logger.Panicf("Config does not contain location for SQL command: %v", 
            command)
    }
    data, err := os.ReadFile(filename)
    if err != nil {
        logger.Panicf("Cannot read SQL command file: %v", filename)
    }
    statement, err := db.Prepare(string(data))
    if (err != nil) {
        logger.Panicf(err.Error())
    }
    return statement
}
