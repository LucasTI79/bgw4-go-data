package product_test

import (
	"testing"

	"github.com/lucasti79/bgw4-put-patch-delete/internal/domain"
	"github.com/lucasti79/bgw4-put-patch-delete/internal/product"
	"github.com/lucasti79/bgw4-put-patch-delete/tests/utils"
	"github.com/stretchr/testify/assert"
)

func Test_MySqlRepositoryWithTxDb_Store_Mock(t *testing.T) {
	var p domain.Product

	db, err := utils.InitTxDbDatabase(t)
	defer db.Close()
	assert.NoError(t, err)

	repository := product.NewRepository(db)

	err = repository.UpdateOrCreate(&p)
	assert.NoError(t, err)

	assert.True(t, p.Id > 0)
}

func Test_MySqlRepositoryWithTxDb_Delete_Mock(t *testing.T) {
	var p domain.Product

	db, err := utils.InitTxDbDatabase(t)
	assert.NoError(t, err)

	repository := product.NewRepository(db)

	err = repository.UpdateOrCreate(&p)
	assert.NoError(t, err)

	assert.True(t, p.Id > 0)

	err = repository.Delete(p.Id)
	assert.NoError(t, err)
}
