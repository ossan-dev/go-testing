package archiver

import (
	"fmt"

	"github.com/google/uuid"
)

type InvoiceManager struct {
	Archiver Archiver
}

type Archiver interface {
	Archive() (id string, err error)
}

type Store struct{}

func (s *Store) Archive() (id string, err error) {
	// ... do some logic
	return uuid.NewString(), nil
}

func (i *InvoiceManager) RecordInvoice() (err error) {
	id, err := i.Archiver.Archive()
	if err != nil {
		return err
	}
	fmt.Printf("recorded invoice with id: %s\n", id)
	return nil
}
