package tests

import (
	"os"
	"testing"

	"github.com/fudaoji/go-utils"
)

func TestDatToImg(t *testing.T) {
	path, _ := os.Getwd()
	datFile := path + "/img/d3466f2b5cd6be0851f1b5d9de045de3_t.dat"
	outDir := path + "/img/"
	file, err := utils.DatToImg(datFile, outDir)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(file)
}
