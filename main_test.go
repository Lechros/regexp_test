package main

import (
	"reflect"
	"slices"
	"testing"
)
import "github.com/Lechros/hangul_regexp"

var searches = []string{
	"a",
	"ㅇ",
	"ㅇ ㅅ ㅇ",
	"이",
	"마깃안",
	"앱솔",
	"젠",
	"제네",
	"에광",
	"ㅇㅋㅇㅅㅇㄷ ㅇ",
	"아케인셰이드 아처",
}

func TestConcatMatchAll(t *testing.T) {
	for _, search := range searches {
		pattern, _ := hangul_regexp.GetPattern(search, false, true, true)
		t.Run(search, func(t *testing.T) {
			expected := StandardMatchAll(pattern)
			actual := StandardConcatMatchAll(pattern)
			slices.Sort(expected)
			slices.Sort(actual)
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
			}
		})
	}
}

func BenchmarkStandardMatchAll(b *testing.B) {
	for _, search := range searches {
		pattern, _ := hangul_regexp.GetPattern(search, false, true, true)
		b.Run(search, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StandardMatchAll(pattern)
			}
		})
	}
}

func BenchmarkStandardConcatMatchAll(b *testing.B) {
	for _, search := range searches {
		pattern, _ := hangul_regexp.GetPattern(search, false, true, true)
		b.Run(search, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StandardConcatMatchAll(pattern)
			}
		})
	}
}

func BenchmarkRuReMatchAll(b *testing.B) {
	for _, search := range searches {
		pattern, _ := hangul_regexp.GetPattern(search, false, true, true)
		b.Run(search, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RuReMatchAll(pattern)
			}
		})
	}
}

func BenchmarkRuReConcatMatchAll(b *testing.B) {
	for _, search := range searches {
		pattern, _ := hangul_regexp.GetPattern(search, false, true, true)
		b.Run(search, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RuReConcatMatchAll(pattern)
			}
		})
	}
}

func BenchmarkPcreMatchAll(b *testing.B) {
	for _, search := range searches {
		pattern, _ := hangul_regexp.GetPattern(search, false, true, true)
		b.Run(search, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PcreMatchAll(pattern)
			}
		})
	}
}

func BenchmarkRe2MatchAll(b *testing.B) {
	for _, search := range searches {
		pattern, _ := hangul_regexp.GetPattern(search, false, true, true)
		b.Run(search, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Re2MatchAll(pattern)
			}
		})
	}
}

//func BenchmarkHyperScanMatchAll(b *testing.B) {
//	for _, search := range searches {
//		pattern, _ := hangul_regexp.GetPattern(search, false, true, true)
//		b.Run(search, func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				HyperScanMatchAll(pattern)
//			}
//		})
//	}
//}

var groupPatterns = []string{
	"(ㅇ|[아-잏]).*?(ㅋ|[카-킿]).*?(ㅇ|[아-잏]).*?(ㅅ|[사-싷]).*?(ㅇ|[아-잏]).*?(ㄷ|[다-딯]).*?_.*?(ㅇ|[아-잏])",
	"(에).*?(?:(광)|(과).*?(ㅇ|[아-잏]))",
}

//func BenchmarkStandardFindGroups(b *testing.B) {
//	for _, pattern := range groupPatterns {
//		b.Run(pattern, func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				StandardFindGroups(pattern)
//			}
//		})
//	}
//}

//func BenchmarkRuReFindGroups(b *testing.B) {
//	for _, pattern := range groupPatterns {
//		b.Run(pattern, func(b *testing.B) {
//			for i := 0; i < b.N; i++ {
//				RuReFindGroups(pattern)
//			}
//		})
//	}
//}
