package main

import (
	"reflect"
	"testing"
)

func Test_getBarcodeNotFounded(t *testing.T) {

	// los campos CategoryID, SubCategoryID son irrelevantes
	type args struct {
		barcodes []string
		products []Product
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "cuando no está 1 de los 3 codigos de barra en los productos entonces la funcion revuelve el codigo de barra que no está",
			args: args{
				barcodes: []string{"codigo_1", "codigo_2", "codigo_3"},
				products: []Product{
					{
						Barcode:       "codigo_1",
						CategoryID:    1,
						SubCategoryID: 1,
					},
					{
						Barcode:       "codigo_3",
						CategoryID:    1,
						SubCategoryID: 1,
					},
				},
			},
			want: []string{"codigo_2"},
		},
		{
			name: "cuando no está ningun codigo de los 3 codigos de barra entonces la funcion revuelve todos los codigos de productos",
			args: args{
				barcodes: []string{"codigo_1", "codigo_2", "codigo_3"},
				products: []Product{
					{
						Barcode:       "nuevo_codigo_1",
						CategoryID:    1,
						SubCategoryID: 1,
					},
					{
						Barcode:       "nuevo_codigo_2",
						CategoryID:    1,
						SubCategoryID: 1,
					},
					{
						Barcode:       "nuevo_codigo_3",
						CategoryID:    1,
						SubCategoryID: 1,
					},
				},
			},
			want: []string{"codigo_1", "codigo_2", "codigo_3"},
		},
		{
			name: "cuando estan todos los codigo barra en los productos entonces la funcion revuelve nil",
			args: args{
				barcodes: []string{"codigo_1", "codigo_2", "codigo_3"},
				products: []Product{
					{
						Barcode:       "codigo_1",
						CategoryID:    1,
						SubCategoryID: 1,
					},
					{
						Barcode:       "codigo_2",
						CategoryID:    1,
						SubCategoryID: 1,
					},
					{
						Barcode:       "codigo_3",
						CategoryID:    1,
						SubCategoryID: 1,
					},
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBarcodeNotFounded(tt.args.barcodes, tt.args.products); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBarcodeNotFounded() = %v, want %v", got, tt.want)
			}
		})
	}
}
