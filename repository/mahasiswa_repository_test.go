package repository

import (
	"strconv"
	"testing"
	"time"

	"github.com/alfiankan/majelisit-rest/entity"
	"github.com/alfiankan/majelisit-rest/helper"
)

func konkurenInsert(mhs entity.Mahasiswa)  {
	
}

func TestMahasiswaRepository_InsertOne(t *testing.T) {
	database := helper.InitMysqlDB()
	repo := MahasiswaRepository{
		MysqlDB: database,
	}



	for i := 0; i < 10000; i++ {
		mhs1 := entity.Mahasiswa{
			Nama: "Budi " + strconv.Itoa(i),
			Alamat: "Solo",
		}
		go repo.InsertOne(mhs1)
	}

	time.Sleep(time.Hour)
	

}
