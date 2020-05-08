package models

import (
	"github.com/jinzhu/gorm"
)

// Item is a struct containing Item data
type Item struct {
	gorm.Model
	Name string `gorm:"index:name"` // название изделия
	Slug string  `gorm:"index;unique;not null"`// строка адреса
	Category Category `gorm:"foreignkey:CategoryID"` // категория техники
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
	Tabel Tabel `gorm:"foreignkey:TabelID"`// нормы табелизации
	Image Image `gorm:"foreignkey:ImageRefer"`// фотография
}

// нормы табелизации
type Tabel struct {
	gorm.Model
	Escadrile uint // отдельная эскадрилься
	Polk uint // авиационный полк
	Division uint // ав. дивизия
	Army uint // А ВВС
	Center uint // учебно авиационный центр
	CY uint // ЦУ вида/рода
	Polygon uint // межвидовой полигон
	Komend uint // ав. комендатура
	School uint // учебное заведение
	// внешний ключ
	ItemID uint
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
	ImageRefer uint
}