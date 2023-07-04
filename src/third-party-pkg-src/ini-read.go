package main

import (
  "fmt"
  "log"
  "gopkg.in/ini.v1"
)

func main() {
  cfg, err := ini.Load("ini-read.ini")
  if err != nil {
    log.Fatal("Fail to read ini file: ", err)
  }

  fmt.Println("App Name:", cfg.Section("").Key("app_name").String())
  fmt.Println("Log Level:", cfg.Section("").Key("log_level").String())

  fmt.Println("Nginx IP:", cfg.Section("nginx").Key("ip").String())
  nginxPort, err := cfg.Section("nginx").Key("port").Int()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Nginx Port:", nginxPort)

  fmt.Println("MySQL IP:", cfg.Section("mysql").Key("ip").String())
  mysqlPort, err := cfg.Section("mysql").Key("port").Int()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("MySQL Port:", mysqlPort)
  fmt.Println("MySQL User:", cfg.Section("mysql").Key("user").String())
  fmt.Println("MySQL Password:", cfg.Section("mysql").Key("password").String())
  fmt.Println("MySQL Database:", cfg.Section("mysql").Key("database").String())
}
