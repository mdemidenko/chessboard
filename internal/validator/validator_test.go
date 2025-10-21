package validator

import (
	"testing"
)

func TestValidateSize(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"корректный ввод", "5", 5, false},
		{"ноль", "0", 0, true},
		{"отрицательное число", "-1", 0, true},
		{"не число", "abc", 0, true},
		{"дробное число", "3.14", 0, true},
		{"пустая строка", "", 0, true},
		{"пробелы", "   ", 0, true},
		{"с пробелами", " 8 ", 8, false},
		{"очень большое число", "999999", 999999, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateSize(tt.input)

			// Наличие ошибки
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Корректность возвращаемого значения
			if !tt.wantErr && got != tt.want {
				t.Errorf("ValidateSize() = %v, want %v", got, tt.want)
			}

			// Если ошибка ожидается, убедиться что значение по умолчанию
			if tt.wantErr && got != 0 {
				t.Errorf("ValidateSize() should return 0 on error, got %v", got)
			}
		})
	}
}
