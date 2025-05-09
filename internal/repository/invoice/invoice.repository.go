package repository

import (
	"database/sql"

	"github.com/marcelochb/training-go-payment-gateway/internal/domain"
)

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

// Save salva uma fatura no banco de dados
func (repository *InvoiceRepository) Save(invoice *domain.InvoiceEntity) error {
	_, err := repository.db.Exec(
		"INSERT INTO invoices (id, amount, status, created_at) VALUES ($1, $2, $3, $4)",
		invoice.ID,
		invoice.Amount,
		invoice.Status,
		invoice.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repository *InvoiceRepository) FindByID(id string) (*domain.InvoiceEntity, error) {
	var invoice domain.InvoiceEntity
	err := repository.db.QueryRow(
		"SELECT id, account_id, description, payment_type, card_last_digits, amount, status, created_at FROM invoices WHERE id = $1",
		id,
	).Scan(
		&invoice.ID,
		&invoice.AccountID,
		&invoice.Description,
		&invoice.Amount,
		&invoice.PaymentType,
		&invoice.CardLastDigits,
		&invoice.Status,
		&invoice.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrInvoiceNotFound
	}
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

func (repository *InvoiceRepository) FindByAccountID(accountID string) ([]*domain.InvoiceEntity, error) {
	rows, err := repository.db.Query(
		"SELECT id, amount, status, created_at FROM invoices WHERE account_id = $1",
		accountID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []*domain.InvoiceEntity
	for rows.Next() {
		var invoice domain.InvoiceEntity
		err := rows.Scan(&invoice.ID, &invoice.Amount, &invoice.Status, &invoice.CreatedAt)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}

func (repository *InvoiceRepository) UpdateStatus(invoice *domain.InvoiceEntity) error {
	result, err := repository.db.Exec(
		"UPDATE invoices SET status = $1, updated_at = $2 WHERE id = $3",
		invoice.Status,
		invoice.UpdatedAt,
		invoice.ID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrInvoiceNotFound
	}

	return nil
}
