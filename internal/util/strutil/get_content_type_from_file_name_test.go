package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContentTypeFromFileName(t *testing.T) {
	tests := []struct {
		fileName string
		expected string
	}{
		{"documento.pdf", "application/pdf"},
		{"imagen.png", "image/png"},
		{"foto.jpg", "image/jpeg"},
		{"foto.jpeg", "image/jpeg"},
		{"animacion.gif", "image/gif"},
		{"imagen.bmp", "image/bmp"},
		{"datos.json", "application/json"},
		{"tabla.csv", "text/csv"},
		{"config.xml", "application/xml"},
		{"nota.txt", "text/plain"},
		{"pagina.html", "text/html"},
		{"archivo.zip", "application/zip"},
		{"archivo.sql", "application/sql"},
		{"archivo.desconocido", "application/octet-stream"}, // unknown extension
		{"MAYUSCULAS.JPG", "image/jpeg"},                    // upper case
		{"MezclaDeMayusculasYMinusculas.PnG", "image/png"},  // mixed case
	}

	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			result := GetContentTypeFromFileName(tt.fileName)
			assert.Equal(t, tt.expected, result)
		})
	}
}
