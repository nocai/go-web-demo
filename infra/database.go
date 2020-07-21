package infra
//
//import (
//	"fmt"
//	"github.com/golang/protobuf/ptypes/timestamp"
//	"github.com/nocai/chat/api"
//	"time"
//
//	"github.com/golang/glog"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"
//
//	"github.com/go-sql-driver/mysql"
//	"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/mysql"
//)
//
//type DBClient struct {
//	*gorm.DB
//}
//
//func NewDB() *DBClient {
//	args := `%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local`
//	args = fmt.Sprintf(args,
//		Config.DataBase.Username,
//		Config.DataBase.Password,
//		Config.DataBase.Host,
//		Config.DataBase.Port,
//		Config.DataBase.Dbname,
//	)
//
//	db, err := gorm.Open("mysql", args)
//	if err != nil {
//		glog.Error(err)
//		panic(err)
//	}
//
//	if err := db.DB().Ping(); err != nil {
//		glog.Error(err)
//		panic(err)
//	}
//
//	db.DB().SetMaxIdleConns(Config.DataBase.MaxIdleConns)
//	db.DB().SetMaxOpenConns(Config.DataBase.MaxOpenConns)
//
//	db.LogMode(Config.DataBase.LogMode)
//
//	return &DBClient{db}
//}
//
//type Model struct {
//	ID        uint64    `gorm:"primary_key" json:"id"`
//	CreatedAt time.Time `json:"createdAt"`
//	UpdatedAt time.Time `json:"updatedAt"`
//}
//
//func (m *Model) Proto() *api.Model {
//	return &api.Model{
//		Id:        m.ID,
//		CreatedAt: &timestamp.Timestamp{Seconds: m.CreatedAt.Unix()},
//		UpdatedAt: &timestamp.Timestamp{Seconds: m.UpdatedAt.Unix()},
//	}
//}
//
//func (m *Model) BeforeCreate() (err error) {
//	if m.ID == 0 {
//		m.ID, err = NextUint64()
//		return err
//	}
//	return nil
//}
//
//func (db *DBClient) IsDuplicateEntry(err error) bool {
//	if err == nil {
//		return false
//	}
//
//	if mysqlError, ok := err.(*mysql.MySQLError); ok {
//		if mysqlError.Number == 1062 {
//			return true
//		}
//	}
//	return false
//}
//
//func (db *DBClient) Wrap(err error) error {
//	if db.IsDuplicateEntry(err) {
//		return status.Error(codes.AlreadyExists, err.Error())
//	}
//	if gorm.IsRecordNotFoundError(err) {
//		return status.Error(codes.NotFound, err.Error())
//	}
//	return status.Error(codes.Internal, err.Error())
//}
