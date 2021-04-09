package db

import (
	_ "bufio"
	_ "encoding/json"
	"github.com/dgraph-io/ristretto"
	"log"
	_ "os"
	_ "strings"
	"tc-micro-idp/models"
	. "tc-micro-idp/utils"
	"time"
)

var (
	//Cache In memory cache
	Cache *ristretto.Cache
	Error error
	ok    bool
)

func init() {
	Cache, Error = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e8,
		MaxCost:     1 << 30,
		BufferItems: 1000,
	})
	if Error != nil {
		log.Fatalln(err)
	}
	//SetEntities()
}

func SetClientTable(Table []models.Client) (ok bool) {
	ok = Cache.SetWithTTL(ClientTable, Table, 0, 24*time.Hour)
	return
}

func GetClientsTable() (Table []models.Client, ok bool) {

	var table interface{}
	table, ok = Cache.Get(ClientTable)
	if !ok {
		ok = SetClientTable(getClientsTable())
		GetClientsTable()
	}
	log.Println("Client Table:", table)
	Table = table.([]models.Client)
	return
}

func CanRegister(TcClient string) (bool, int64) {
	table := FindClient(TcClient)
	if TcClient != table.Issuer {
		return false, 0
	}
	return table.CanRegister, table.Id
}

func CanLogin(TcClient string) (ClientTable models.Client, ok bool) {
	ClientTable = FindClient(TcClient)
	if ClientTable.Id == 0 {
		return ClientTable, false
	}
	return ClientTable, ClientTable.CanLogin
}

//func SetEntities() {
//	file, err := os.OpenFile("urls.json", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
//	if err != nil {
//		log.Fatalln(err)
//		return
//	}
//	defer file.Close()
//	var accesses []models.Access
//	reader := bufio.NewReader(file)
//	err = json.NewDecoder(reader).Decode(&accesses)
//	if err != nil {
//		log.Fatalln(err)
//		return
//	}
//	for _, acc := range accesses {
//		Cache.Set(strings.ToLower(acc.Url), acc.Roles, 1)
//		log.Println(acc.Url, " : ", acc.Roles)
//	}
//}

func GetRoles(Url string) (Roles []string, err error) {
	data, found := Cache.Get(Url)
	if !found {
		return nil, err
	}
	return data.([]string), nil
}
