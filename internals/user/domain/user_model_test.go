package domain

import (
	"nona/pkg/common/base"
	"testing"
)

func TestUser_CollectionName(t *testing.T) {
	type fields struct {
		ModelMongo base.ModelMongo
		Name       string
		Username   string
		Password   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "positive",
			fields: fields{},
			want:   collectionUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &User{
				ModelMongo: tt.fields.ModelMongo,
				Name:       tt.fields.Name,
				Username:   tt.fields.Username,
				Password:   tt.fields.Password,
			}
			if got := m.CollectionName(); got != tt.want {
				t.Errorf("CollectionName() = %v, want %v", got, tt.want)
			}
		})
	}
}
