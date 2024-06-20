package db
import(
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
	"fmt"
)
type(
	DbType interface{
		init()
		
	}
	DbStruct struct{

	}

)

func (dbStruct *DbStruct) Init() *gorm.DB{
	dbUser := os.Getenv("DB_USER")
dbPass := os.Getenv("DB_PASS")
dbHost := os.Getenv("DB_HOST")
dbName := os.Getenv("DB_NAME")
dbPort := os.Getenv("DB_PORT")

dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)

db,err :=gorm.Open(postgres.New(postgres.Config{
	DSN: dsn,
	PreferSimpleProtocol:true,
}))
if(err!=nil){
	panic(err)
} 
return db
}

func(dbStruct *DbStruct) CloseConnection(db *gorm.DB){
dbSQl,err :=db.DB()
if err!=nil{
	panic(err)
}
dbSQl.Close()
}
// Hi