

package db


import (
   "database/sql"
   "log"
   "time"


   "gorm.io/driver/postgres"
   "gorm.io/gorm"
)


type Config struct {
   URL       string
}


var DB *gorm.DB


// Init the connection to DB
func Init(config *Config) error {
   if DB == nil {
       sqlDB, err := sql.Open("postgres", config.URL)
       if err != nil {
           log.Println("Unable to open postges connection. Err:", err)
           return err
       }

       sqlDB.SetConnMaxLifetime(time.Hour)


       DB, err = gorm.Open(postgres.New(postgres.Config{
           Conn: sqlDB,
       }), &gorm.Config{})
       if err != nil {
           log.Println("Unable to open postges gorm connection. Err:", err)
           return err
       }


       log.Println("Successfully established database connection")
   }


   return nil
}


type DBConn struct {
   *gorm.DB
}


func New() *DBConn {
   return &DBConn{
       DB: DB,
   }
}

