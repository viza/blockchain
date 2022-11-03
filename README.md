# Blockchain
Repository for course "Blockchain and Decentralized Technologies" from Distributed Lab organisation

# Task 1

1. Print possible variants of keys for 8-,16-,32-,64-,128-, 256-, 512- , 1024-, 2048-, 4096- bits sequence

Number of keys equals 2^n, where n - bits (8,16,32 etc)

2. Generate random key in range of 0x00..0 - 0xFF..F (len for 8,16,32 etc bit sequence)

It's not a problem to generate 8-,16,32,64 etc 4096 len key, but starting from 64 bit it's time consuming to brute force it

3. Using brute force method try to find random generated key in key sequence 

Code shows brute force method for 8, 16, 32 bits key
Time to brute force 8,16 key is seconds, 32 bits took about 6 hours for 0xDB... key (Apple M1 Pro chip with 16 RAM)

Example of terminal layout
```
Total possible key values for 8 bits are 256
Total possible key values for 16 bits are 65536
Total possible key values for 32 bits are 4.294967296e+09
Total possible key values for 64 bits are 1.8446744073709552e+19
Total possible key values for 128 bits are 3.402823669209385e+38
Total possible key values for 256 bits are 1.157920892373162e+77
Total possible key values for 512 bits are 1.3407807929942597e+154
Total possible key values for 1024 bits are +Inf
Total possible key values for 2048 bits are +Inf
Total possible key values for 4096 bits are +Inf

Generated 8-bits key =  86
Key 0 = 00
Key 1 = 01
Key 2 = 02
Key 3 = 03
Key 4 = 04
Key 5 = 05
Key 6 = 06
Key 7 = 07
Key 8 = 08
Key 9 = 09
Key 10 = 0A
Key 11 = 0B
Key 12 = 0C
Key 13 = 0D
Key 14 = 0E
Key 15 = 0F
Key 16 = 10
Key 17 = 11
Key 18 = 12
Key 19 = 13
Key 20 = 14
Key 21 = 15
Key 22 = 16
Key 23 = 17
Key 24 = 18
Key 25 = 19
Key 26 = 1A
Key 27 = 1B
Key 28 = 1C
Key 29 = 1D
Key 30 = 1E
Key 31 = 1F
Key 32 = 20
Key 33 = 21
Key 34 = 22
Key 35 = 23
Key 36 = 24
Key 37 = 25
Key 38 = 26
Key 39 = 27
Key 40 = 28
Key 41 = 29
Key 42 = 2A
Key 43 = 2B
Key 44 = 2C
Key 45 = 2D
Key 46 = 2E
Key 47 = 2F
Key 48 = 30
Key 49 = 31
Key 50 = 32
Key 51 = 33
Key 52 = 34
Key 53 = 35
Key 54 = 36
Key 55 = 37
Key 56 = 38
Key 57 = 39
Key 58 = 3A
Key 59 = 3B
Key 60 = 3C
Key 61 = 3D
Key 62 = 3E
Key 63 = 3F
Key 64 = 40
Key 65 = 41
Key 66 = 42
Key 67 = 43
Key 68 = 44
Key 69 = 45
Key 70 = 46
Key 71 = 47
Key 72 = 48
Key 73 = 49
Key 74 = 4A
Key 75 = 4B
Key 76 = 4C
Key 77 = 4D
Key 78 = 4E
Key 79 = 4F
Key 80 = 50
Key 81 = 51
Key 82 = 52
Key 83 = 53
Key 84 = 54
Key 85 = 55
Key 86 = 56
Key 87 = 57
Key 88 = 58
Key 89 = 59
Key 90 = 5A
Key 91 = 5B
Key 92 = 5C
Key 93 = 5D
Key 94 = 5E
Key 95 = 5F
Key 96 = 60
Key 97 = 61
Key 98 = 62
Key 99 = 63
Key 100 = 64
Key 101 = 65
Key 102 = 66
Key 103 = 67
Key 104 = 68
Key 105 = 69
Key 106 = 6A
Key 107 = 6B
Key 108 = 6C
Key 109 = 6D
Key 110 = 6E
Key 111 = 6F
Key 112 = 70
Key 113 = 71
Key 114 = 72
Key 115 = 73
Key 116 = 74
Key 117 = 75
Key 118 = 76
Key 119 = 77
Key 120 = 78
Key 121 = 79
Key 122 = 7A
Key 123 = 7B
Key 124 = 7C
Key 125 = 7D
Key 126 = 7E
Key 127 = 7F
Key 128 = 80
Key 129 = 81
Key 130 = 82
Key 131 = 83
Key 132 = 84
Key 133 = 85
Key 134 = 86
Key 135 = 87
Key 136 = 88
Key 137 = 89
Key 138 = 8A
Key 139 = 8B
Key 140 = 8C
Key 141 = 8D
Key 142 = 8E
Key 143 = 8F
Key 144 = 90
Key 145 = 91
Key 146 = 92
Key 147 = 93
Key 148 = 94
Key 149 = 95
Key 150 = 96
Key 151 = 97
Key 152 = 98
Key 153 = 99
Key 154 = 9A
Key 155 = 9B
Key 156 = 9C
Key 157 = 9D
Key 158 = 9E
Key 159 = 9F
Key 160 = A0
Key 161 = A1
Key 162 = A2
Key 163 = A3
Key 164 = A4
Key 165 = A5
Key 166 = A6
Key 167 = A7
Key 168 = A8
Key 169 = A9
Key 170 = AA
Key 171 = AB
Key 172 = AC
Key 173 = AD
Key 174 = AE
Key 175 = AF
Key 176 = B0
Key 177 = B1
Key 178 = B2
Key 179 = B3
Key 180 = B4
Key 181 = B5
Key 182 = B6
Key 183 = B7
Key 184 = B8
Key 185 = B9
Key 186 = BA
Key 187 = BB
Key 188 = BC
Key 189 = BD
Key 190 = BE
Key 191 = BF
Key 192 = C0
Key 193 = C1
Key 194 = C2
Key 195 = C3
Key 196 = C4
Key 197 = C5
Key 198 = C6
Key 199 = C7
Key 200 = C8
Key 201 = C9
Key 202 = CA
Key 203 = CB
Key 204 = CC
Key 205 = CD
Key 206 = CE
Key 207 = CF
Key 208 = D0
Key 209 = D1
Key 210 = D2
Key 211 = D3
Key 212 = D4
Key 213 = D5
Key 214 = D6
Key 215 = D7
Key 216 = D8
Key 217 = D9
Key 218 = DA
Key 219 = DB
Key 220 = DC
Key 221 = DD
Key 222 = DE
Key 223 = DF
Key 224 = E0
Key 225 = E1
Key 226 = E2
Key 227 = E3
Key 228 = E4
Key 229 = E5
Key 230 = E6
Key 231 = E7
Key 232 = E8
Key 233 = E9
Key 234 = EA
Key 235 = EB
Key 236 = EC
Key 237 = ED
Key 238 = EE
Key 239 = EF
Key 240 = F0
Key 241 = F1
Key 242 = F2
Key 243 = F3
Key 244 = F4
Key 245 = F5
Key 246 = F6
Key 247 = F7
Key 248 = F8
Key 249 = F9
Key 250 = FA
Key 251 = FB
Key 252 = FC
Key 253 = FD
Key 254 = FE
Key 255 = FF
Generated key in array is found 86 = 86
Find the key... took 4.5µs
Generated 16-bits key =  11E1
Generated key in array is found 11E1 = 11E1
Find the key... took 409.166µs
Generated 32-bits key = C2C2BD17
```
# Task 2

1 - to convert hex string to little endian
2 - to convert hex string to big endian
3 - to convert little endian to hex string
4 - to convert big endian to hex string

Some console layout:
```
viza@Vitaliys-MacBook-Pro 2 % go run ./hex.go
Please choose what do you want to convert...
Enter 1 - to convert hex string to little endian
Enter 2 - to convert hex string to big endian
Enter 3 - to convert little endian to hex string
Enter 4 - to convert big endian to hex string
1
1 - to convert hex string to little endian
Enter hex number as string please(without 0x). Allowed symbols are: 0123456789ABCDEF
ffcc
Hex = 0xffcc
Binary = [11111111 11001100]
Number of needed bytes to store data is 2
 - - - - - 
Little endian:  52479
 - - - - - 

 - - TEST VECTORS - - 
Hex = 0xff00000000000000000000000000000000000000000000000000000000000000
Binary = [11111111 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
Number of needed bytes to store data is 32
Byte data: [255 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], little endian = 255

Hex = 0xaaaa000000000000000000000000000000000000000000000000000000000000
Binary = [10101010 10101010 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
Number of needed bytes to store data is 32
Byte data: [170 170 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], little endian = 43690

Hex = 0xFFFFFFFF
Binary = [11111111 11111111 11111111 11111111]
Number of needed bytes to store data is 4
Byte data: [255 255 255 255], little endian = 4294967295

Hex = 0xF000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
Binary = [11110000 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
Number of needed bytes to store data is 512
Byte data: [240 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], little endian = 240

viza@Vitaliys-MacBook-Pro 2 % 

Please choose what do you want to convert...
Enter 1 - to convert hex string to little endian
Enter 2 - to convert hex string to big endian
Enter 3 - to convert little endian to hex string
Enter 4 - to convert big endian to hex string
3
3 - to convert little endian to hex string
Enter number please. Allowed symbols are: 0123456789
240
Hex string: f000000000000000
```

# Task 3

SHA1 algorithm 


