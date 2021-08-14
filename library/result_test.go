package library

import (
	"reflect"
	"testing"
)

func TestMeta_CalculatePages(t *testing.T) {
	type fields struct {
		Page         int64
		Limit        int64
		TotalRecords int64
		TotalPages   int64
		OrderBy      string
	}
	tests := []struct {
		name            string
		fields          fields
		wantTotalRecord int64
		wantError       bool
	}{
		{
			name: "positif case",
			fields: fields{
				Page:         1,
				Limit:        10,
				TotalRecords: 20,
				OrderBy:      "",
			},
			wantTotalRecord: 2,
			wantError:       false,
		},
		{
			name: "negatif case",
			fields: fields{
				Page:         1,
				Limit:        10,
				TotalRecords: 20,
				OrderBy:      "",
			},
			wantTotalRecord: 10,
			wantError:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Meta{
				Page:         tt.fields.Page,
				Limit:        tt.fields.Limit,
				TotalRecords: tt.fields.TotalRecords,
				TotalPages:   tt.fields.TotalPages,
				OrderBy:      tt.fields.OrderBy,
			}

			m.CalculatePages()
			if tt.wantError {
				if tt.wantTotalRecord == m.TotalPages {
					t.Errorf("Data totalPage same with result")
				}
			} else {
				if tt.wantTotalRecord != m.TotalPages {
					t.Errorf("Data totalPage not same with result")
				}
			}
		})
	}
}

func TestMeta_ToResolver(t *testing.T) {
	type fields struct {
		Page         int64
		Limit        int64
		TotalRecords int64
		TotalPages   int64
		OrderBy      string
	}
	tests := []struct {
		name   string
		fields fields
		want   *MetaResolver
	}{
		{
			name: "positif case",
			fields: fields{
				Page:         1,
				Limit:        10,
				TotalRecords: 2,
				TotalPages:   1,
			},
			want: &MetaResolver{
				Page:         1,
				Limit:        10,
				TotalRecords: 2,
				TotalPages:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Meta{
				Page:         tt.fields.Page,
				Limit:        tt.fields.Limit,
				TotalRecords: tt.fields.TotalRecords,
				TotalPages:   tt.fields.TotalPages,
				OrderBy:      tt.fields.OrderBy,
			}
			if got := m.ToResolver(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToResolver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMeta(t *testing.T) {
	type args struct {
		page         int64
		limit        int64
		totalRecords int64
	}
	tests := []struct {
		name string
		args args
		want *Meta
	}{
		{
			name: "positif",
			args: args{
				page:         1,
				limit:        10,
				totalRecords: 8,
			},
			want: &Meta{
				Page:         1,
				Limit:        10,
				TotalRecords: 8,
				TotalPages:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMeta(tt.args.page, tt.args.limit, tt.args.totalRecords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}
