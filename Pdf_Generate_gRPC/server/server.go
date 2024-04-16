package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pdf "pdf_generate_grpc/proto" // Import your generated proto package

	"github.com/jung-kurt/gofpdf"
)

const (
	port = ":50051"
)

type server struct {
	pdf.UnimplementedPDFServiceServer
}

func (s *server) GeneratePDF(ctx context.Context, req *pdf.PDFRequest) (*pdf.PDFResponse, error) {
	// Create a new PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set the content from the request
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, req.Content)

	// Generate PDF output
	output := pdf.Output(nil)

	// Return the PDF response
	return &pdf.PDFResponse{
		PdfData: output,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pdf.RegisterPDFServiceServer(s, &server{})
	log.Printf("Server listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
