
## Test Cases

```
    testcases := []TestCase{
		{
			Identifier: "Test Case 1",
			PlainText:  "karet pedurenan 81",
			Type:       "Address",
		},
		{
			Identifier: "Test Case 2",
			PlainText:  "pedurenan 81",
			Type:       "Address",
		},
		{
			Identifier: "Test Case 3",
			PlainText:  "6281234566789",
			Type:       "Phone Number",
		},
		{
			Identifier: "Test Case 4",
			PlainText:  "6281234566788",
			Type:       "Phone Number",
		},
		{
			Identifier: "Test Case 5",
			PlainText:  "Jl. Jend. Sudirman Kav. 44-46, Jakarta 10210",
			Type:       "Address",
		},
		{
			Identifier: "Test Case 6",
			PlainText:  "Jl. Jend. Sudirman Kav. 1, Jakarta 10220",
			Type:       "Address",
		},
		{
			Identifier: "Test Case 7",
			PlainText:  "Sudirman",
			Type:       "Address",
		},
	}

	for _, tc := range testcases {
		fmt.Printf("%s\nPlainText:%s\nCipher Text:%s\nType:%s\n==========\n\n", tc.Identifier, tc.PlainText, transform.Transform(tc.PlainText), tc.Type)
	}
```


## Results

```
Test Case 1
PlainText:karet pedurenan 81
Cipher Text:nSqU�oeSgq[qW�Lk
Type:Address
==========

Test Case 2
PlainText:pedurenan 81
Cipher Text:oeSgq[qW�Lk
Type:Address
==========

Test Case 3
PlainText:6281234566789
Cipher Text:h^kcccccdccc
Type:Phone Number
==========

Test Case 4
PlainText:6281234566788
Cipher Text:h^kcccccdccd
Type:Phone Number
==========

Test Case 5
PlainText:Jl. Jend. Sudirman Kav. 44-46, Jakarta 10210
Cipher Text:B�r:I[n�r1Bu_[ipW�9NO�rPdk]bnp:MZnSbw�Sebee
Type:Address
==========

Test Case 6
PlainText:Jl. Jend. Sudirman Kav. 1, Jakarta 10220
Cipher Text:B�r:I[n�r1Bu_[ipW�9NO�rSip:MZnSbw�Sebdf
Type:Address
==========

Test Case 7
PlainText:Sudirman
Cipher Text:Bu_[ipW
Type:Address
==========
```