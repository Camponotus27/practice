package main

type Product struct {
	Barcode       string
	CategoryID    int32
	SubCategoryID int32
}

func getBarcodeNotFounded(barcodes []string, products []Product) []string {
	return nil
}
