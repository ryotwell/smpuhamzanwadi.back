package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Fullname  string    `json:"fullname" gorm:"type:varchar(255);"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegister struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Session struct {
	gorm.Model
	Token    string    `json:"token"`
	Username string    `json:"username"`
	Expiry   time.Time `json:"expiry"`
}

// type Student struct {
// 	gorm.Model
// 	Name    string `json:"name"`
// 	Address string `json:"address"`
// 	ClassId int    `json:"class_id"`
// }

// type Class struct {
// 	ID         int    `gorm:"primaryKey"`
// 	Name       string `json:"name"`
// 	Professor  string `json:"professor"`
// 	RoomNumber int    `json:"room_number"`
// }

// type StudentClass struct {
// 	Name       string `json:"name"`
// 	Address    string `json:"address"`
// 	ClassName  string `json:"class_name"`
// 	Professor  string `json:"professor"`
// 	RoomNumber int    `json:"room_number"`
// }

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

// type ErrorResponse struct {
// 	Error string `json:"error"`
// }

// type SuccessResponse struct {
// 	Username string `json:"username"`
// 	Message  string `json:"message"`
// }

type SuccessResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Success bool              `json:"success"`
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

// Parent Model
type Parent struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	FatherName      *string `json:"father_name"`
	FatherEducation *string `json:"father_education"`
	FatherJob       *string `json:"father_job"`
	FatherIncome    *string `json:"father_income"`

	MotherName      *string `json:"mother_name"`
	MotherEducation *string `json:"mother_education"`
	MotherJob       *string `json:"mother_job"`
	MotherIncome    *string `json:"mother_income"`

	ParentEmail *string `json:"parent_email"`

	WaliName       *string `json:"wali_name"`
	AlamatOrtuWali *string `json:"alamat_ortu_wali"`
	NoHpOrtuWali   *string `json:"no_hp_ortu_wali"`

	Student *Student `json:"student"`
}

// Student Model
type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
)

type BloodType string

const (
	BloodA       BloodType = "A"
	BloodB       BloodType = "B"
	BloodAB      BloodType = "AB"
	BloodO       BloodType = "O"
	BloodUnknown BloodType = "UNKNOWN"
)

type TinggalBersama string

const (
	OrangTua       TinggalBersama = "ORANG_TUA"
	KakekNenek     TinggalBersama = "KAKEK_NENEK"
	PamanBibi      TinggalBersama = "PAMAN_BIBI"
	SaudaraKandung TinggalBersama = "SAUDARA_KANDUNG"
	Kerabat        TinggalBersama = "KERABAT"
	PantiPontRen   TinggalBersama = "PANTI_PONTREN"
	Lainnya        TinggalBersama = "LAINNYA"
)

type StatusKeluarga string

const (
	AnakKandung StatusKeluarga = "ANAK_KANDUNG"
	AnakTiri    StatusKeluarga = "ANAK_TIRI"
	AnakAngkat  StatusKeluarga = "ANAK_ANGKAT"
)

type KeadaanOrtu string

const (
	Lengkap    KeadaanOrtu = "LENGKAP"
	Yatim      KeadaanOrtu = "YATIM"
	Piatu      KeadaanOrtu = "PIATU"
	YatimPiatu KeadaanOrtu = "YATIM_PIATU"
)

type Religion string

const (
	Islam     Religion = "ISLAM"
	Christian Religion = "CHRISTIAN"
	Catholic  Religion = "CATHOLIC"
	Hindu     Religion = "HINDU"
	Buddha    Religion = "BUDDHA"
	Konghucu  Religion = "KONGHUCU"
	OtherRel  Religion = "OTHER"
)

type Student struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	FullName              string          `json:"full_name"`
	Nisn                  *string         `gorm:"unique" json:"nisn"`
	Nik                   *string         `gorm:"uniqueIndex" json:"nik"`
	AsalSekolah           *string         `json:"asal_sekolah"`
	Gender                Gender          `json:"gender"`
	TempatLahir           *string         `json:"tempat_lahir"`
	TanggalLahir          *string         `json:"tanggal_lahir"`
	Agama                 *Religion       `json:"agama"`
	KeadaanOrtu           *KeadaanOrtu    `json:"keadaan_ortu"`
	StatusKeluarga        *StatusKeluarga `json:"status_keluarga"`
	AnakKe                *int            `json:"anak_ke"`
	DariBersaudara        *int            `json:"dari_bersaudara"`
	TinggalBersama        *TinggalBersama `json:"tinggal_bersama"`
	TinggalBersamaLainnya *string         `json:"tinggal_bersama_lainnya"`
	Kewarganegaraan       *string         `json:"kewarganegaraan"`
	AlamatJalan           *string         `json:"alamat_jalan"`
	Rt                    *string         `json:"rt"`
	Rw                    *string         `json:"rw"`
	DesaKel               *string         `json:"desa_kel"`
	Kecamatan             *string         `json:"kecamatan"`
	Kabupaten             *string         `json:"kabupaten"`
	Provinsi              *string         `json:"provinsi"`
	KodePos               *string         `json:"kode_pos"`
	Phone                 *string         `json:"phone"`
	Email                 *string         `json:"email"`

	BloodType       *BloodType `json:"blood_type"`
	BeratKg         *int       `json:"berat_kg"`
	TinggiCm        *int       `json:"tinggi_cm"`
	RiwayatPenyakit *string    `json:"riwayat_penyakit"`

	ParentId *int    `gorm:"unique" json:"parent_id"`
	Parent   *Parent `json:"parent"`
}
