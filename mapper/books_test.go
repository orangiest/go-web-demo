package mapper

import (
	"cwm.wiki/web/models"
	"testing"
)

func TestSelectBooks(t *testing.T) {
	models.ConnectDataBase()


	t.Log(SelectBooks())
}
