package models

import (
	"github.com/jinzhu/gorm"
)

// Item is a struct containing Item data
type Item struct {
	gorm.Model
	Name string `gorm:"index:name"` // название изделия
	Slug string  `gorm:"index;unique;not null"`// строка адреса
	Category uint64 // `gorm:"foreignkey:ID"` // категория техники
	Index string `gorm:"index:indx"` // индекс иделия
	Snabjenie string `gorm:"type:text"`// принят на снабжение
	KVT string `gorm:"index:kvt"` // код КВТ
	Nomenklature string `gorm:"index:nomenklature"` // номенклатурный номер
	Dovorgan string  `gorm:"type:text"` // довольствующий орган
	Reqorgan string `gorm:"type:text"`// заказывающий орган
	Explorgan string `gorm:"type:text"` // эксплуатирующий орган
	Creator string  `gorm:"type:text"` // изготовитель
	Description string `gorm:"type:text"` // описание
	Destination string `gorm:"type:text"` // назначение
	Composition string `gorm:"type:text"` // состав
	TTH string `gorm:"type:text"` // ТТХ
	Tabel Tabel `gorm:"foreignkey:ID"`// нормы табелизации
	Image Image `gorm:"foreignkey:ID"`// фотография
}


// нормы табелизации
type Tabel struct {
	gorm.Model
	Escadrile uint `json:",string"`// отдельная эскадрилься
	Polk uint `json:",string"`// авиационный полк
	Division uint `json:",string"`// ав. дивизия
	Army uint `json:",string"`// А ВВС
	Center uint `json:",string"`// учебно авиационный центр
	CY uint `json:",string"`// ЦУ вида/рода
	Polygon uint `json:",string"`// межвидовой полигон
	Komend uint `json:",string"`// ав. комендатура
	School uint `json:",string"`// учебное заведение
	// внешний ключ
	// ItemID uint
}

// Images фотографии
type Image struct {
	gorm.Model
	Name string 
	Path string 
	// внешний ключ
	// ImageRefer uint
}

type UserItems struct {
	gorm.Model
	DateStart string `gorm:"default:'2004-07-23'"` // дата ввода в эксплуатацию
	DayHours float32 `gorm:"default:24"` // Работа в день
	Srok float32 `gorm:"default:15"`// Время наработки на отказ, лет
	Inventory string //Инвентарный номер
	Primechanie string `gorm:"type:text"`// Примечание
	ItemID uint
	Item Item `gorm:"foreignkey:ItemID"`// Изделие
}