package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    

)

var DB *gorm.DB

func InitDatabase() {
    db_uri :="postgres://avnadmin:AVNS_MZREtKn09d-6m9zYmuo@nstasdb2-folajimiopeyemisax13-3385.j.aivencloud.com:26285/defaultdb?sslmode=require"
    db, err := gorm.Open(postgres.Open(db_uri), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB = db
}
