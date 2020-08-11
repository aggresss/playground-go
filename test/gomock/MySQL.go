package MySQL

import db "github.com/aggresss/playground-go/test/gomock/infra"

type MySQL struct {
	DB db.Repository
}

func NewMySQL(db db.Repository) *MySQL {
	return &MySQL{DB: db}
}

func (mysql *MySQL) CreateData(key string, value []byte) error {
	return mysql.DB.Create(key, value)
}

func (mysql *MySQL) GetData(key string) ([]byte, error) {
	return mysql.DB.Retrieve(key)
}

func (mysql *MySQL) DeleteData(key string) error {
	return mysql.DB.Delete(key)
}

func (mysql *MySQL) UpdateData(key string, value []byte) error {
	return mysql.DB.Update(key, value)
}
