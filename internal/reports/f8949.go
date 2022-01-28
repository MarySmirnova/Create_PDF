package reports

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/desertbit/fillpdf"

	"github.com/MarySmirnova/create_pdf/internal/form"
)

type Report8949Generator struct {
	templatesPath string
	report8949    string
}

func NewReport8949Generator(templatesPath, report8949 string) *Report8949Generator {
	return &Report8949Generator{
		templatesPath: templatesPath,
		report8949:    report8949,
	}
}

func (g *Report8949Generator) Generate(report form.Report) ([]byte, error) {
	filename := generateTemporaryFilename(".pdf")
	err := fillpdf.Fill(form.ConvertReportToPDFForm(report), filepath.Join(g.templatesPath, g.report8949), filename)
	if err != nil {
		return nil, err
	}
	defer os.Remove(filename)

	// TODO: prepare and send the response
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// generateTemporaryFilename generates a temporary filename for use in testing or whatever
func generateTemporaryFilename(suffix string) string {
	randBytes := make([]byte, 16)
	_, _ = rand.Read(randBytes)
	return filepath.Join(os.TempDir(), hex.EncodeToString(randBytes)+suffix)
}
