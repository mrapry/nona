package library

import (
	"testing"

	"github.com/Kamva/mgm/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/stretchr/testify/assert"
)

func TestFilter_CalculateOffset(t *testing.T) {
	testName := "Test calculate offset"

	t.Run(testName, func(t *testing.T) {
		filter := Filter{Limit: 1, Page: 1}

		filter.CalculateOffset()

		assert.Equal(t, int32(0), filter.Offset)
	})
}

func TestFilter_DateCondition(t *testing.T) {
	type fields struct {
		Limit   int32
		Page    int32
		Offset  int32
		Search  string
		OrderBy string
		Sort    string
		SortInt int
		ShowAll bool
	}
	type args struct {
		column string
		date   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "positive",
			fields: fields{},
			args: args{
				column: "created_at",
				date:   "2021-09-21",
			},
			want: "created_at BETWEEN '2021-09-21 00:00:00' AND '2021-09-21 23:59:59'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{
				Limit:   tt.fields.Limit,
				Page:    tt.fields.Page,
				Offset:  tt.fields.Offset,
				Search:  tt.fields.Search,
				OrderBy: tt.fields.OrderBy,
				Sort:    tt.fields.Sort,
				SortInt: tt.fields.SortInt,
				ShowAll: tt.fields.ShowAll,
			}
			if got := f.DateCondition(tt.args.column, tt.args.date); got != tt.want {
				t.Errorf("DateCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter_LikeCondition(t *testing.T) {
	type fields struct {
		Limit   int32
		Page    int32
		Offset  int32
		Search  string
		OrderBy string
		Sort    string
		SortInt int
		ShowAll bool
	}
	type args struct {
		searchFields []string
		value        string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "positive",
			fields: fields{Search: "matius"},
			args:   args{searchFields: []string{"name"}, value: "matius"},
			want:   "CAST(lower(name) as text) LIKE '%matius%'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{
				Limit:   tt.fields.Limit,
				Page:    tt.fields.Page,
				Offset:  tt.fields.Offset,
				Search:  tt.fields.Search,
				OrderBy: tt.fields.OrderBy,
				Sort:    tt.fields.Sort,
				SortInt: tt.fields.SortInt,
				ShowAll: tt.fields.ShowAll,
			}
			if got := f.LikeCondition(tt.args.searchFields, tt.args.value); got != tt.want {
				t.Errorf("LikeCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter_SearchCondition(t *testing.T) {
	type fields struct {
		Limit   int32
		Page    int32
		Offset  int32
		Search  string
		OrderBy string
		Sort    string
		SortInt int
		ShowAll bool
	}
	type args struct {
		searchFields []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "positive",
			fields: fields{Search: "matius"},
			args:   args{searchFields: []string{"name"}},
			want:   "CAST(lower(name) as text) LIKE '%matius%'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Filter{
				Limit:   tt.fields.Limit,
				Page:    tt.fields.Page,
				Offset:  tt.fields.Offset,
				Search:  tt.fields.Search,
				OrderBy: tt.fields.OrderBy,
				Sort:    tt.fields.Sort,
				SortInt: tt.fields.SortInt,
				ShowAll: tt.fields.ShowAll,
			}
			if got := f.SearchCondition(tt.args.searchFields); got != tt.want {
				t.Errorf("SearchCondition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter_SetSearch(t *testing.T) {
	testName := "Test set search"

	t.Run(testName, func(t *testing.T) {
		search := gofakeit.Word()
		filter := Filter{Search: search}

		query := []bson.M{}
		fields := []string{"name"}

		// search
		query = filter.SetSearch(query, fields)

		// expected
		expected := []bson.M{
			bson.M{operator.Or: []bson.M{
				bson.M{"name": bson.M{operator.Regex: primitive.Regex{Pattern: `^` + search + `.*`, Options: "i"}}},
			}},
		}

		assert.Equal(t, expected, query)
	})
}

func TestFilter_SetSort(t *testing.T) {
	testCase := map[string]struct {
		sort     string
		expected int
	}{
		"Test set sort asc": {
			sort:     SortAsc,
			expected: 1,
		},
		"Test set sort desc": {
			sort:     SortDesc,
			expected: -1,
		},
	}

	for name, test := range testCase {
		t.Run(name, func(t *testing.T) {
			filter := Filter{Sort: test.sort}

			filter.SetSort()

			assert.Equal(t, test.expected, filter.SortInt)
		})
	}
}
