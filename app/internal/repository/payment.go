package repository

import (
	"github.com/Nukie90/my-fluffy/app/domain/entity"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type PaymentRepo struct {
	DB *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) *PaymentRepo {
	return &PaymentRepo{DB: db}
}

func (pr *PaymentRepo) CreatePayment(payment *entity.Payment) error {
	result := pr.DB.Create(payment)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *PaymentRepo) GetPaymentFromSpecificUser(userID ulid.ULID) ([]entity.Payment, error) {
	var payments []entity.Payment
	result := pr.DB.Where("receiver_id = ? or user_id = ?", userID, userID).Find(&payments)
	if result.Error != nil {
		return nil, result.Error
	}

	return payments, nil
}
