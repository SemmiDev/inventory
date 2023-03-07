package repository

import (
	"inventory/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository() *ItemRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// bikinin table otomatis
	db.AutoMigrate(&entity.Item{})

	return &ItemRepository{db: db}
}

func (r *ItemRepository) Save(item *entity.Item) error {
	return r.db.Create(item).Error
}

func (r *ItemRepository) FindAll() ([]entity.Item, error) {
	var items []entity.Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r *ItemRepository) FindById(id uint) (entity.Item, error) {
	var item entity.Item
	err := r.db.First(&item, id).Error
	return item, err
}

func (r *ItemRepository) Update(item *entity.Item) error {
	return r.db.Save(item).Error
}

func (r *ItemRepository) Delete(item *entity.Item) error {
	return r.db.Delete(item).Error
}

func (r *ItemRepository) FindByStatus(status entity.Status) ([]entity.Item, error) {
	var items []entity.Item
	err := r.db.Where("status = ?", status.String()).Find(&items).Error
	return items, err
}
