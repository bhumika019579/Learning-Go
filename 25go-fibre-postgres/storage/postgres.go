package storage


import(
	""
)


type config struct{
	Host    string
	Port     string
	password   string
	user     string
	DBName    string
	SSLMode    string
}


func NewConnection(config *config)(*gorm.DB,error){

}