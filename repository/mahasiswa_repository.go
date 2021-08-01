package repository

import (
	"database/sql"

	"github.com/alfiankan/majelisit-rest/entity"
)


type MahasiswaRepository struct {
	MysqlDB		*sql.DB
}


func (repo *MahasiswaRepository) InsertOne(mhs entity.Mahasiswa) (entity.Mahasiswa, error) {
	
	res, err := repo.MysqlDB.Exec("INSERT INTO mahasiswa (nama, alamat) VALUES (?,?)",mhs.Nama, mhs.Alamat)
	if err != nil {
		return mhs, err
	}
	mhs.Nim, _ = res.LastInsertId()
	return mhs, err
}
