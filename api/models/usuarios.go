package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sait/tickets/api/db"
)

//Usuario sd
type Usuario struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nombre    string    `gorm:"not null;size:255" json:"nombre"`
	Email     string    `gorm:"not null;unique;size:100" json:"email"` //unique
	Password  string    `gorm:"not null;size:100;default:12345" json:"password"`
	Telefono  string    `gorm:"not null;size:10" json:"telefono"`
	Ext       string    `gorm:"not null;size:100" json:"extension"`
	Estado    uint8     `gorm:"not null;default:1" json:"estado"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//SaveUsuario s
func (u *Usuario) SaveUsuario() (*Usuario, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	err = DB.Debug().Create(&u).Error
	if err != nil {
		return &Usuario{}, err
	}
	return u, nil
}

//FindAllUsuarios f
func (u *Usuario) FindAllUsuarios() (*[]Usuario, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	usuarios := []Usuario{}
	err = DB.Debug().Model(&Usuario{}).Limit(100).Find(&usuarios).Error
	if err != nil {
		return &[]Usuario{}, err
	}
	return &usuarios, err
}

//FindUsuarioByID f
func (u *Usuario) FindUsuarioByID(uid uint32) (*Usuario, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	err = DB.Debug().Model(Usuario{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Usuario{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Usuario{}, errors.New("Usuario Not Found")
	}
	return u, err
}

//FindUsuarioByEmail f
func (u *Usuario) FindUsuarioByEmail(email string) (*Usuario, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	err = DB.Debug().Model(Usuario{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Usuario Not Found")
	}
	return u, err
}

//UpdateUsuario u
func (u *Usuario) UpdateUsuario(uid uint32) (*Usuario, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	db := DB.Debug().Model(&Usuario{}).Where("id = ?", uid).Take(&Usuario{}).UpdateColumns(
		map[string]interface{}{
			"nombre":     u.Nombre,
			"email":      u.Email,
			"telefono":   u.Telefono,
			"extension":  u.Ext,
			"updated_at": time.Now(),
			"estado":     u.Estado,
		},
	)
	if db.Error != nil {
		return &Usuario{}, db.Error
	}

	err = DB.Debug().Model(&Usuario{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Usuario{}, err
	}
	return u, nil
}

//DeleteUsuario d
func (u *Usuario) DeleteUsuario(uid uint32) (*Usuario, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	db := DB.Debug().Model(&Usuario{}).Where("id = ?", uid).Take(&Usuario{}).UpdateColumns(
		map[string]interface{}{
			"estado": 0,
		},
	)
	if db.Error != nil {
		return &Usuario{}, db.Error
	}

	err = DB.Debug().Model(&Usuario{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Usuario{}, err
	}
	return u, nil
}

//ValidateCredentials v
func (u *Usuario) ValidateCredentials(email, password string) (*Usuario, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}

	err = DB.Debug().Model(Usuario{}).Where("email = ? and password = ?", email, password).Take(&u).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Usuario Not Found")
	}
	return u, nil
}
