package archiver_test

import (
	"testing"

	"github.com/ossan-dev/gotesting/internal/payment/archiver"
	"github.com/ossan-dev/gotesting/mocks"
	"github.com/stretchr/testify/assert"
)

// mockery --output ./mocks --all
func TestRecordInvoice(t *testing.T) {
	// Arrange
	store := mocks.NewArchiver(t)
	store.On("Archive").Return("5ec3fea2-e43f-4d09-8d30-e6ad9757bb6f", nil).Once()
	invoiceManager := &archiver.InvoiceManager{
		Archiver: store,
	}
	// Act
	err := invoiceManager.RecordInvoice()
	// Assert
	assert.NoError(t, err)
	store.AssertExpectations(t)
}
