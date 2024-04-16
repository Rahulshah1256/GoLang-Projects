package main

import (
	"context"
	"log"

	//pdf "pdf_generate_grpc/proto" // Import your generated proto package

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pdf.NewPDFServiceClient(conn)

	// Create a PDF request
	req := &pdf.PDFRequest{
		Content: "Hello, World!",
	}

	// Call the GeneratePDF service
	res, err := client.GeneratePDF(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to generate PDF: %v", err)
	}

	// Use the generated PDF
	pdfData := res.PdfData
	// You can save the PDF data to a file or use it as needed

	log.Println("PDF generated successfully!")
}
