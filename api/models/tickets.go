package models

import (
	"errors"
	"time"

	"github.com/Flix14/soporte-tickets/api/db"
	"github.com/jinzhu/gorm"
)

//Ticket representa
type Ticket struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Titulo      string    `gorm:"size:255;not null;" json:"titulo"`
	Descripcion string    `gorm:"size:255;not null;" json:"descripcion"`
	TemaAyuda   string    `gorm:"size:255;not null;" json:"tema_ayuda"`
	Producto    string    `gorm:"size:255;not null;" json:"producto"`
	Estado      uint8     `gorm:"not null;default:1" json:"estado"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Agente      Agente    `json:"agente"`
	AgenteID    uint32    `gorm:"not null" json:"agente_id"`
	Usuario     Usuario   `json:"usuario"`
	UsuarioID   uint32    `gorm:"not null" json:"usuario_id"`
}

//DB f
var DB *gorm.DB

//SaveTicket s
func (t *Ticket) SaveTicket() (*Ticket, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	err = DB.Debug().Model(&Ticket{}).Create(&t).Error
	if err != nil {
		return &Ticket{}, err
	}
	if t.UsuarioID != 0 {
		err = DB.Debug().Model(&Usuario{}).Where("id = ?", t.UsuarioID).Take(&t.Usuario).Error
		if err != nil {
			return &Ticket{}, err
		}
	}
	if t.AgenteID != 0 {
		err = DB.Debug().Model(&Agente{}).Where("id = ?", t.AgenteID).Take(&t.Agente).Error
		if err != nil {
			return &Ticket{}, err
		}
	}
	return t, nil
}

//FindAllTickets f
func (t *Ticket) FindAllTickets() (*[]Ticket, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	tickets := []Ticket{}
	err = DB.Debug().Model(&Ticket{}).Limit(100).Find(&tickets).Error
	if err != nil {
		return &[]Ticket{}, err
	}
	if len(tickets) > 0 {
		for i := range tickets {
			if tickets[i].UsuarioID != 0 {
				err := DB.Debug().Model(&Usuario{}).Where("id = ?", tickets[i].UsuarioID).Take(&tickets[i].Usuario).Error
				if err != nil {
					return &[]Ticket{}, err
				}
			}
		}
	}
	if len(tickets) > 0 {
		for i := range tickets {
			if tickets[i].AgenteID != 0 {
				err := DB.Debug().Model(&Agente{}).Where("id = ?", tickets[i].AgenteID).Take(&tickets[i].Agente).Error
				if err != nil {
					return &[]Ticket{}, err
				}
			}
		}
	}
	return &tickets, nil
}

//FindTicketByID f
func (t *Ticket) FindTicketByID(uid uint32) (*Ticket, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	err = DB.Debug().Model(Ticket{}).Where("id = ?", uid).Take(&t).Error
	if err != nil {
		return &Ticket{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Ticket{}, errors.New("Ticket Not Found")
	}

	if t.UsuarioID != 0 {
		err := DB.Debug().Model(&Usuario{}).Where("id = ?", t.UsuarioID).Take(&t.Usuario).Error
		if err != nil {
			return &Ticket{}, err
		}
	}
	if t.AgenteID != 0 {
		err := DB.Debug().Model(&Agente{}).Where("id = ?", t.AgenteID).Take(&t.Agente).Error
		if err != nil {
			return &Ticket{}, err
		}
	}
	return t, err
}

//FindTicketsByUsuario f
func (t *Ticket) FindTicketsByUsuario(uid uint32) (*[]Ticket, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	tickets := []Ticket{}
	//err = DB.Debug().Model(Ticket{}).Where("usuario_id = ?", uid).Take(&t).Error
	err = DB.Debug().Model(&Ticket{}).Where("usuario_id = ?", uid).Find(&tickets).Error
	if err != nil {
		return &[]Ticket{}, err
	}
	if len(tickets) > 0 {
		for i := range tickets {
			if tickets[i].UsuarioID != 0 {
				err := DB.Debug().Model(&Usuario{}).Where("id = ?", tickets[i].UsuarioID).Take(&tickets[i].Usuario).Error
				if err != nil {
					return &[]Ticket{}, err
				}
			}
		}
	}
	if len(tickets) > 0 {
		for i := range tickets {
			if tickets[i].AgenteID != 0 {
				err := DB.Debug().Model(&Agente{}).Where("id = ?", tickets[i].AgenteID).Take(&tickets[i].Agente).Error
				if err != nil {
					return &[]Ticket{}, err
				}
			}
		}
	}
	return &tickets, nil
}

//ChangeEstadoTicket c
func (t *Ticket) ChangeEstadoTicket(uid uint32) (*Ticket, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	db := DB.Debug().Model(&Ticket{}).Where("id = ?", uid).Take(&Ticket{}).UpdateColumns(
		map[string]interface{}{
			"estado": t.Estado,
		},
	)
	if db.Error != nil {
		return &Ticket{}, db.Error
	}

	err = DB.Debug().Model(&Ticket{}).Where("id = ?", uid).Take(&t).Error
	if err != nil {
		return &Ticket{}, err
	}
	return t, nil
}
