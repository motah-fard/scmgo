package inventory

import (
	"testing"
)

func TestMinMaxLevels(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   MinMaxInput
		want    MinMaxResult
		wantErr error
	}{
		{
			name: "valid input",
			input: MinMaxInput{
				ReorderPoint:  300,
				OrderQuantity: 200,
			},
			want: MinMaxResult{
				Min: 300,
				Max: 500,
			},
		},
		{
			name: "zero order quantity",
			input: MinMaxInput{
				ReorderPoint:  300,
				OrderQuantity: 0,
			},
			want: MinMaxResult{
				Min: 300,
				Max: 300,
			},
		},
		{
			name: "negative reorder point",
			input: MinMaxInput{
				ReorderPoint:  -1,
				OrderQuantity: 200,
			},
			wantErr: ErrNegativeReorderPoint,
		},
		{
			name: "negative order quantity",
			input: MinMaxInput{
				ReorderPoint:  300,
				OrderQuantity: -1,
			},
			wantErr: ErrNegativeOrderQuantity,
		},
		{
			name: "zero reorder point with positive order quantity",
			input: MinMaxInput{
				ReorderPoint:  0,
				OrderQuantity: 500,
			},
			want: MinMaxResult{
				Min: 0,
				Max: 500,
			},
		},
		{
			name: "very large values",
			input: MinMaxInput{
				ReorderPoint:  10_000_000,
				OrderQuantity: 5_000_000,
			},
			want: MinMaxResult{
				Min: 10_000_000,
				Max: 15_000_000,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := MinMaxLevels(tt.input)
			if err != tt.wantErr {
				t.Fatalf("MinMaxLevels() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if got != tt.want {
				t.Fatalf("MinMaxLevels() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
