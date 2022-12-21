package main

import "testing"

func Test_isAdjacent(t *testing.T) {
	type args struct {
		rowDifference int
		colDifference int
	}
	tests := []struct {
		name string
		args args
	}{{

		/*
			_ h
			_ t
		*/
		args: args{
			rowDifference: 1,
			colDifference: 0,
		},
	},
		{
			/*

				t h
			*/
			args: args{
				rowDifference: 0,
				colDifference: 1,
			},
		},
		{
			/*
				_ t
				_ h
			*/
			args: args{
				rowDifference: -1,
				colDifference: 0,
			},
		},
		{
			args: args{
				/*
					h t

				*/
				rowDifference: 0,
				colDifference: -1,
			},
		},

		{
			/*
				_ h
				t _
			*/
			args: args{
				rowDifference: 1,
				colDifference: 1,
			},
		},
		{
			/*
				h _
				_ t
			*/
			args: args{
				rowDifference: 1,
				colDifference: -1,
			},
		},
		{
			/*
				t _
				_ h
			*/
			args: args{
				rowDifference: -1,
				colDifference: 1,
			},
		},
		{
			/*
				_ t
				h _
			*/
			args: args{
				rowDifference: -1,
				colDifference: -1,
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdjacent(tt.args.rowDifference, tt.args.colDifference); got != true {
				t.Errorf("isAdjacent() = %v, want true for rowdiff=%d, coldiff=%d", got, tt.args.rowDifference, tt.args.colDifference)
			}
		})
	}
}
