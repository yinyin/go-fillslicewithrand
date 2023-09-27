package gofillslicewithrand

func fillIntInt32N(s []int, n int32) {
	if (n & (n - 1)) == 0 {
		rndsrc := ReaderInt32Masked{
			Mask: (n - 1),
		}
		for idx := range s {
			v := int(rndsrc.Int32Masked())
			s[idx] = v
		}
	} else {
		rndsrc := NewReaderInt32N(n)
		for idx := range s {
			v := int(rndsrc.Int32N())
			s[idx] = v
		}
	}
}

func fillIntInt64N(s []int, n int64) {
	if (n & (n - 1)) == 0 {
		rndsrc := ReaderInt64Masked{
			Mask: (n - 1),
		}
		for idx := range s {
			v := int(rndsrc.Int64Masked())
			s[idx] = v
		}
	} else {
		rndsrc := NewReaderInt64N(n)
		for idx := range s {
			v := int(rndsrc.Int64N())
			s[idx] = v
		}
	}
}

func FillIntN(s []int, n int) {
	if n <= ((1 << 31) - 1) {
		fillIntInt32N(s, int32(n))
	} else {
		fillIntInt64N(s, int64(n))
	}
}
