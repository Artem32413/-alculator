package expense  

import (  
	"errors"  
	"testing"  

	app "github.com/Artem32413/Aplication"
)  

func TestCalc(t *testing.T) {  
	testCasesSuccess := []struct {  
		name           string  
		expression     string  
		expectedResult float64  
	}{  
		{  
			name:           "simple",  
			expression:     "1+1",  
			expectedResult: 2,  
		},  
		{  
			name:           "priority",  
			expression:     "(2+2)*2",  
			expectedResult: 8,  
		},  
		{  
			name:           "priority",  
			expression:     "2+2*2",  
			expectedResult: 6,  
		},  
		{  
			name:           "division",  
			expression:     "1/2",  
			expectedResult: 0.5,  
		},  
	}  

	for _, testCase := range testCasesSuccess {  
		t.Run(testCase.name, func(t *testing.T) {  
			val, err := app.Calc(testCase.expression)  
			if err != nil {  
				t.Fatalf("successful case %s returns error: %v", testCase.expression, err)  
			}  
			if val != testCase.expectedResult {  
				t.Fatalf("expected %f but got %f for expression %s", testCase.expectedResult, val, testCase.expression)  
			}  
		})  
	}  

	testCasesFail := []struct {  
		name        string  
		expression  string  
		expectedErr error  
	}{  
		{  
			name:       "trailing operator",  
			expression: "1+1*",  
			expectedErr: errors.New("invalid expression: trailing operator"),  
		},  
		{  
			name:       "mismatched parentheses",  
			expression: "(1+2",  
			expectedErr: errors.New("invalid expression: mismatched parentheses"),  
		},  
		{  
			name:       "div by zero",  
			expression: "1/0",  
			expectedErr: errors.New("division by zero"),  
		},  
		{  
			name:       "invalid character",  
			expression: "1+@",  
			expectedErr: errors.New("invalid expression: contains invalid character"),  
		},  
		{  
			name:       "empty expression",  
			expression: "",  
			expectedErr: errors.New("invalid expression: empty expression"),  
		},  
	}  

	for _, testCase := range testCasesFail {  
		t.Run(testCase.name, func(t *testing.T) {  
			val, err := app.Calc(testCase.expression)  
			if err == nil {  
				t.Fatalf("expected error for case %s but got result: %f", testCase.expression, val)  
			}  
			if err.Error() != testCase.expectedErr.Error() {  
				t.Fatalf("expected error message '%v' but got '%v'", testCase.expectedErr, err)  
			}  
		})  
	}  
}