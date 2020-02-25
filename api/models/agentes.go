package models

import (
	"errors"
	"time"

	"github.com/Flix14/soporte-tickets/api/db"
	"github.com/jinzhu/gorm"
)

//Agente sdlksd
type Agente struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nombre    string    `gorm:"size:255;not null;" json:"nombre"`
	Email     string    `gorm:"size:100;not null;" json:"email"` //unique
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Estado    uint8     `gorm:"not null;default:1" json:"estado"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//SaveAgente s
func (u *Agente) SaveAgente() (*Agente, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	err = DB.Debug().Create(&u).Error
	if err != nil {
		return &Agente{}, err
	}
	return u, nil
}

//FindAllAgentes f
func (u *Agente) FindAllAgentes() (*[]Agente, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	agentes := []Agente{}
	err = DB.Debug().Model(&Agente{}).Limit(100).Find(&agentes).Error
	if err != nil {
		return &[]Agente{}, err
	}
	return &agentes, err
}

//FindAgenteByID f
func (u *Agente) FindAgenteByID(uid uint32) (*Agente, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	err = DB.Debug().Model(Agente{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Agente{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Agente{}, errors.New("Agente Not Found")
	}
	return u, err
}

//UpdateAgente u
func (u *Agente) UpdateAgente(uid uint32) (*Agente, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	db := DB.Debug().Model(&Agente{}).Where("id = ?", uid).Take(&Agente{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"nombre":     u.Nombre,
			"email":      u.Email,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Agente{}, db.Error
	}

	err = DB.Debug().Model(&Agente{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Agente{}, err
	}
	return u, nil
}

//DeleteAgente d
func (u *Agente) DeleteAgente(uid uint32) (*Agente, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	db := DB.Debug().Model(&Agente{}).Where("id = ?", uid).Take(&Agente{}).UpdateColumns(
		map[string]interface{}{
			"estado": 0,
		},
	)
	if db.Error != nil {
		return &Agente{}, db.Error
	}

	err = DB.Debug().Model(&Agente{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Agente{}, err
	}
	return u, nil
}

//ValidateCredentials v
func (u *Agente) ValidateCredentials(email, password string) (*Agente, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return u, err
	}

	err = DB.Debug().Model(Agente{}).Where("email = ? and password = ?", email, password).Take(&u).Error
	if err != nil {
		return u, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return u, errors.New("Agente Not Found")
	}
	return u, nil
}
