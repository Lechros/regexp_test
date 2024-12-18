package main

import "testing"
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

func BenchmarkStandardFindGroups(b *testing.B) {
	for _, pattern := range groupPatterns {
		b.Run(pattern, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StandardFindGroups(pattern)
			}
		})
	}
}

func BenchmarkRuReFindGroups(b *testing.B) {
	for _, pattern := range groupPatterns {
		b.Run(pattern, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RuReFindGroups(pattern)
			}
		})
	}
}
