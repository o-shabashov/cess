package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "io/ioutil"
    "encoding/json"
    "net/http"
    "net/url"
    "bytes"
    "os"
    "cess/cess"
    "sync"
    log "github.com/sirupsen/logrus"
)

const configFile = "config.json"

func main() {
    config := GetConfig()

    var wg sync.WaitGroup
    wg.Add(len(config.Databases) + len(config.Services))

    for _, db := range config.Databases {
        go TestDatabase(db, &wg)
    }

    for _, service := range config.Services {
        go TestApi(service, &wg)
    }

    wg.Wait()
}

func GetConfig() (config cess.Config) {
    logger := log.WithFields(log.Fields{"config_file": configFile})

    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        logger.Warn("Whoops, config file isn't exist, first run? Whatever, creating sample config for you")
        err := ioutil.WriteFile(configFile, []byte(cess.SampleConfig), 0644)
        CheckError(err, logger)
        log.Info("Done! Please edit config.json and run me again.")
        os.Exit(0)
    }

    config = cess.Config{}
    jsonRaw, er := ioutil.ReadFile(configFile)
    CheckError(er, logger)
    CheckError(json.Unmarshal(jsonRaw, &config), logger)
    log.Info("JSON config file - OK")
    return
}

func GetDBConnection(db cess.Database) (*sql.DB) {

    dbUrl := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", db.Username, db.Password, db.Host, db.Port, db.Name)
    logger := log.WithFields(log.Fields{"db_url": dbUrl})

    dbConnect, err := sql.Open(db.Engine, dbUrl)
    CheckError(err, logger)

    logger.Info("DB connection - OK")
    return dbConnect
}

func TestDatabase(db cess.Database, wg *sync.WaitGroup) {
    logger := log.WithFields(log.Fields{"db_name": db.Name, "db_host": db.Host})

    dbConnect := GetDBConnection(db)
    _, err := dbConnect.Query("SHOW TABLES")
    CheckError(err, logger)
    logger.Info("Test DB query - OK")

    wg.Done()
}

func TestApi(api cess.Api, wg *sync.WaitGroup) {
    apiUrlAction := fmt.Sprintf("%v%v", api.Url, api.Action)

    logger := log.WithFields(log.Fields{"api_name": api.Url, "api_url_action": apiUrlAction})

    data := url.Values{}
    for key, value := range api.Data {
        data.Set(key, value)
    }

    req, err := http.NewRequest(api.Method, apiUrlAction, bytes.NewBufferString(data.Encode()))
    CheckError(err, logger)

    for key, value := range api.Headers {
        req.Header.Set(key, value)
    }

    resp, err := http.DefaultClient.Do(req)
    CheckError(err, logger)

    if resp.StatusCode != 200 {
        body, _ := ioutil.ReadAll(resp.Body)
        logger.Warn(fmt.Sprintf("API returned %v: %v", resp.StatusCode, string(body)))
    } else {
        logger.Info("API returned 200 - OK")
    }
    defer resp.Body.Close()

    wg.Done()
}

func CheckError(err error, logger *log.Entry) {
    if err != nil {
        logger.Error(err.Error())
    }
}
