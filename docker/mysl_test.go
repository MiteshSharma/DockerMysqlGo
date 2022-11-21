package docker

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestMySQLContainer(t *testing.T) {

	password := "password"
	port := 23306
	user := "user"
	dbname := "unitest"
	imageName := "mysql:5.7"


	docker := Docker{
		ContainerID:   "mysql-unittest",
		ContainerName: "mysql-unitest",
	}

	mysqlContainer := Container{
		Docker: docker,
		ImageName: imageName,
	}

	mysqlContainer.StartMysqlDocker(user,password,port,dbname)
	defer mysqlContainer.Stop()

	time.Sleep(time.Second * 15)

	// get DB connection
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True", user, password, "127.0.0.1", port, dbname, "utf8")
	log.Printf(" got URI as %s", dbURI)
	Db, err := sql.Open("mysql", dbURI)
	assert.NoError(t,err)

	err = Db.Ping()
	assert.NoError(t,err)

	createTableSQL := "CREATE TABLE test_table ( test_key varchar(50) NOT NULL, test_value varchar(50) NOT NULL);";

	_, err = Db.Exec(createTableSQL)
	assert.NoError(t,err)

	testKey := "test_key_name"
	testData := "test data here"

	var sampleData sql.NullString

	stmt, err := Db.Prepare("INSERT INTO test_table (test_key,test_value) VALUE (?,?) ")
	if err != nil {

		log.Printf("got errors here %s ",err.Error())
		assert.NoError(t,err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(testKey,testData)
	if err != nil {

		log.Printf("got errors here %s ",err.Error())
		assert.NoError(t,err)
	}

	err = Db.QueryRow("SELECT test_value FROM test_table WHERE test_key = ? ",testKey).Scan(&sampleData)
	if err != nil {

		log.Printf("got errors here %s ",err.Error())
		assert.Fail(t,err.Error())
	}

	assert.Equal(t,testData,sampleData.String)

}