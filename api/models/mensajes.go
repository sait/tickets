package models

import (
	"github.com/sait/tickets/api/db"
)

//Mensaje representa
type Mensaje struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Descripcion string `gorm:"size:255;not null;" json:"descripcion"`
	Emisor      uint8  `gorm:"not null" json:"emisor"`
	Ticket      Ticket `json:"ticket"`
	TicketID    int    `gorm:"not null" json:"ticket_id"`
}

//SaveMsg s
func (m *Mensaje) SaveMsg() (*Mensaje, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	err = DB.Debug().Model(&Mensaje{}).Create(&m).Error
	if err != nil {
		return &Mensaje{}, err
	}
	if m.TicketID != 0 {
		err = DB.Debug().Model(&Ticket{}).Where("id = ?", m.TicketID).Take(&m.Ticket).Error
		if err != nil {
			return &Mensaje{}, err
		}
	}
	return m, nil
}

//FindAllMsgByTicket f
func (m *Mensaje) FindAllMsgByTicket(uid uint32) (*[]Mensaje, error) {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		return nil, err
	}
	mensajes := []Mensaje{}
	err = DB.Debug().Model(&Mensaje{}).Where("ticket_id = ?", uid).Find(&mensajes).Error
	if err != nil {
		return &[]Mensaje{}, err
	}
	if len(mensajes) > 0 {
		for i := range mensajes {
			if mensajes[i].TicketID != 0 {
				err := DB.Debug().Model(&Ticket{}).Where("id = ?", mensajes[i].TicketID).Take(&mensajes[i].Ticket).Error
				if err != nil {
					return &[]Mensaje{}, err
				}
			}
		}
	}
	return &mensajes, nil
}
