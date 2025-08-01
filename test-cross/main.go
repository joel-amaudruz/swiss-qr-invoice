package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	const scale = 4.838 // Échelle principale du <g> (approximative)

	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	// Couleur noire (remplissage du polygone principal)
	pdf.SetFillColor(0, 0, 0)
	pdf.RectFromUpperLeftWithStyle(
		0*scale, 0*scale,
		19.8*scale, 19.8*scale,
		"F",
	)

	// Couleur blanche pour les deux rectangles "vides"
	pdf.SetFillColor(255, 255, 255)

	// Rectangle vertical (x=11.034758, y=5.2231627, w=4.4874606, h=14.958201)
	pdf.RectFromUpperLeftWithStyle(
		11.034758*scale,
		5.2231627*scale,
		4.4874606*scale,
		14.958201*scale,
		"F",
	)

	// Rectangle horizontal (x=5.7313952, y=10.526525, w=14.958201, h=4.4874606)
	pdf.RectFromUpperLeftWithStyle(
		5.7313952*scale,
		10.526525*scale,
		14.958201*scale,
		4.4874606*scale,
		"F",
	)

	// Contour blanc (polygone de bordure)
	pdf.SetStrokeColor(255, 255, 255)
	pdf.SetLineWidth(1.43) // ~1.05601561 * 1.3598 * scale
	pdf.SetLineType("solid")

	pdf.Line(1.6, 0.7, 0.7, 0.7)
	pdf.Line(0.7, 0.7, 0.7, 1.6)
	pdf.Line(0.7, 1.6, 0.7, 18.3)
	pdf.Line(0.7, 18.3, 0.7, 19.1)
	pdf.Line(0.7, 19.1, 1.6, 19.1)
	pdf.Line(1.6, 19.1, 18.3, 19.1)
	pdf.Line(18.3, 19.1, 19.1, 19.1)
	pdf.Line(19.1, 19.1, 19.1, 18.3)
	pdf.Line(19.1, 18.3, 19.1, 1.6)
	pdf.Line(19.1, 1.6, 19.1, 0.7)
	pdf.Line(19.1, 0.7, 18.3, 0.7)
	pdf.Line(18.3, 0.7, 1.6, 0.7)

	// Multiplier toutes les coordonnées du polygone
	// par scale * 1.3598 pour avoir le rendu exact (ou ajuster manuellement)

	if err := pdf.WritePdf("output.pdf"); err != nil {
		log.Fatal(err)
	}
}
