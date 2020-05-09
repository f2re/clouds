package models

import (
	"github.com/jinzhu/gorm"
)

// Item is a struct containing Item data
type Item struct {
	gorm.Model
	Name string `gorm:"index:name"` // название изделия
	Slug string  `gorm:"index;unique;not null"`// строка адреса
	Category Category `gorm:"foreignkey:ID"` // категория техники
	Index string `gorm:"index:indx"` // индекс иделия
	Snabjenie string // принят на снабжение
	KVT string `gorm:"index:kvt"` // код КВТ
	Nomenklature string `gorm:"index:nomenklature"` // номенклатурный номер
	Dovorgan string   // довольствующий орган
	Reqorgan string  // заказывающий орган
	Explorgan string  // эксплуатирующий орган
	Creator string  // изготовитель
	Description string  // описание
	Destination string  // назначение
	Composition string  // состав
	TTH string  // ТТХ
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

// Category категория техники
type Category struct {
	gorm.Model
	Name string `gorm:"index;unique;not null"`
	Slug string `gorm:"unique;not null"`
}

// Images фотографии
type Image struct {
	gorm.Model
	Name string 
	Path string 
	// внешний ключ
	// ImageRefer uint
}