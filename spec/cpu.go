package spec

type Microarchitecture struct {
	Name       string                `json:"name"`
	From       []string              `json:"from"`
	Vendor     []string              `json:"vendor"`
	Generation int                   `json:"generation"`
	Features   []string              `json:"features"`
	Compilers  map[string][]Compiler `json:"compilers"`
}

type Compiler struct {
	Name     string   `json:"name"`
	Versions string   `json:"versions"`
	Flags    string   `json:"flags"`
	Family   []string `json:"family"`
	Warnings []string `json:"warnings"`
}

var CpuArches = map[string]Microarchitecture{
	"x86": {
		Name:   "x86",
		Vendor: []string{"generic"},
	},
	"i686": {
		Name:   "i686",
		Vendor: []string{"GenuineIntel"},
	},
	"pentium2": {
		Name:     "pentium2",
		From:     []string{"i686"},
		Vendor:   []string{"GenuineIntel"},
		Features: []string{"mmx"},
	},
	"pentium3": {
		Name:     "pentium3",
		From:     []string{"pentium2"},
		Vendor:   []string{"GenuineIntel"},
		Features: []string{"mmx", "sse"},
	},
	"pentium4": {
		Name:     "pentium4",
		From:     []string{"pentium3"},
		Vendor:   []string{"GenuineIntel"},
		Features: []string{"mmx", "sse", "sse2"},
	},
	"prescott": {
		Name:     "prescott",
		From:     []string{"pentium4"},
		Vendor:   []string{"GenuineIntel"},
		Features: []string{"mmx", "sse", "sse2", "sse3"},
	},
	"x86_64": {
		Name:   "x86_64",
		Vendor: []string{"generic"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "gcc",
					Versions: "4.2.0:",
					Flags:    "-march={name} -mtune=generic",
				},
				{
					Name:     "gcc",
					Versions: ":4.1.2",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"apple-clang": {
				{
					Versions: ":",
					Name:     "x86-64",
					Flags:    "-march={name}",
				},
			},
			"clang": {
				{
					Versions: ":",
					Name:     "x86-64",
					Flags:    "-march={name} -mtune=generic",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Name:     "x86-64",
					Flags:    "-march={name} -mtune=generic",
				},
			},
			"intel": {
				{
					Versions: ":",
					Name:     "pentium4",
					Flags:    "-march={name} -mtune=generic",
				},
			},
		},
	},
	"x86_64_v2": {
		Name:   "x86_64_v2",
		From:   []string{"x86_64"},
		Vendor: []string{"generic"},
		Features: []string{"cx16",
			"lahf_lm",
			"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "11.1:",
					Name:     "x86-64-v2",
					Flags:    "-march={name} -mtune=generic",
				},
				{
					Versions: "4.6:11.0",
					Name:     "x86-64",
					Flags:    "-march={name} -mtune=generic -mcx16 -msahf -mpopcnt -msse3 -msse4.1 -msse4.2 -mssse3",
				},
			},
			"clang": {
				{
					Versions: "12.0:",
					Name:     "x86-64-v2",
					Flags:    "-march={name} -mtune=generic",
				},
				{
					Versions: "3.9:11.1",
					Name:     "x86-64",
					Flags:    "-march={name} -mtune=generic -mcx16 -msahf -mpopcnt -msse3 -msse4.1 -msse4.2 -mssse3",
				},
			},
		},
	},
	"x86_64_v3": {
		From:   []string{"x86_64_v2"},
		Vendor: []string{"generic"},
		Features: []string{"cx16",
			"lahf_lm",
			"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"avx",
			"avx2",
			"bmi1",
			"bmi2",
			"f16c",
			"fma",
			"abm",
			"movbe",
			"xsave"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "11.1:",
					Name:     "x86-64-v3",
					Flags:    "-march={name} -mtune=generic",
				},
				{
					Versions: "4.8:11.0",
					Name:     "x86-64",
					Flags:    "-march={name} -mtune=generic -mcx16 -msahf -mpopcnt -msse3 -msse4.1 -msse4.2 -mssse3 -mavx -mavx2 -mbmi -mbmi2 -mf16c -mfma -mlzcnt -mmovbe -mxsave",
				},
			},
			"clang": {
				{
					Versions: "12.0:",
					Name:     "x86-64-v3",
					Flags:    "-march={name} -mtune=generic",
				},
				{
					Versions: "3.9:11.1",
					Name:     "x86-64",
					Flags:    "-march={name} -mtune=generic -mcx16 -msahf -mpopcnt -msse3 -msse4.1 -msse4.2 -mssse3 -mavx -mavx2 -mbmi -mbmi2 -mf16c -mfma -mlzcnt -mmovbe -mxsave",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Name:     "x86-64",
					Flags:    "-march={name} -mtune=generic -mcx16 -msahf -mpopcnt -msse3 -msse4.1 -msse4.2 -mssse3 -mavx -mavx2 -mbmi -mbmi2 -mf16c -mfma -mlzcnt -mmovbe -mxsave",
				},
			},
		},
	},
	"x86_64_v4": {

		From:   []string{"x86_64_v3"},
		Vendor: []string{"generic"},
		Features: []string{"cx16",
			"lahf_lm",
			"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"avx",
			"avx2",
			"bmi1",
			"bmi2",
			"f16c",
			"fma",
			"abm",
			"movbe",
			"xsave",
			"avx512f",
			"avx512bw",
			"avx512cd",
			"avx512dq",
			"avx512vl"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "11.1:",
					Name:     "x86-64-v4",
					Flags:    "-march={name} -mtune=generic",
				},
				{
					Versions: "6.0:11.0",
					Name:     "x86-64",
					Flags:    "-march={name} -mtune=generic -mcx16 -msahf -mpopcnt -msse3 -msse4.1 -msse4.2 -mssse3 -mavx -mavx2 -mbmi -mbmi2 -mf16c -mfma -mlzcnt -mmovbe -mxsave -mavx512f -mavx512bw -mavx512cd -mavx512dq -mavx512vl",
				},
			},
			"clang": {
				{
					Versions: "12.0:",
					Name:     "x86-64-v4",
					Flags:    "-march={name} -mtune=generic",
				},
				{
					Versions: "3.9:11.1",
					Name:     "x86-64",
					Flags:    "-march={name} -mtune=generic -mcx16 -msahf -mpopcnt -msse3 -msse4.1 -msse4.2 -mssse3 -mavx -mavx2 -mbmi -mbmi2 -mf16c -mfma -mlzcnt -mmovbe -mxsave -mavx512f -mavx512bw -mavx512cd -mavx512dq -mavx512vl",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Name:     "x86-64",
					Flags:    "-march={name} -mtune=generic -mcx16 -msahf -mpopcnt -msse3 -msse4.1 -msse4.2 -mssse3 -mavx -mavx2 -mbmi -mbmi2 -mf16c -mfma -mlzcnt -mmovbe -mxsave -mavx512f -mavx512bw -mavx512cd -mavx512dq -mavx512vl",
				},
			},
		},
	},
	"nocona": {
		From:   []string{"x86_64"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"sse3"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.0.4:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Flags:    "-march={name} -mtune=generic",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:",
					Name:     "pentium4",
					Flags:    "-march={name} -mtune=generic",
				},
			},
		},
	},
	"core2": {
		From:   []string{"nocona"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.3.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Flags:    "-march={name} -mtune=generic",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"nehalem": {
		From:   []string{"core2", "x86_64_v2"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.9:",
					Flags:    "-march={name} -mtune={name}",
				},
				{
					Versions: "4.6:4.8.5",
					Name:     "corei7",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Flags:    "-march={name} -mtune=generic",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:",
					Name:     "corei7",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"westmere": {
		From:   []string{"nehalem"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"aes",
			"pclmulqdq"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Flags:    "-march={name} -mtune=generic",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:",
					Name:     "corei7",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"sandybridge": {
		From:   []string{"westmere"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"aes",
			"pclmulqdq",
			"avx"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.9:",
					Flags:    "-march={name} -mtune={name}",
				},
				{
					Versions: "4.6:4.8.5",
					Name:     "corei7-avx",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Flags:    "-march={name} -mtune=generic",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "18.0:",
					Flags:    "-march={name} -mtune={name}",
				},
				{
					Versions: "16.0:17.9.0",
					Name:     "corei7-avx",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"ivybridge": {
		From:   []string{"sandybridge"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"aes",
			"pclmulqdq",
			"avx",
			"rdrand",
			"f16c"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.9:",
					Flags:    "-march={name} -mtune={name}",
				},
				{
					Versions: "4.6:4.8.5",
					Name:     "core-avx-i",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:17.9.0",
					Name:     "core-avx-i",
					Flags:    "-march={name} -mtune={name}",
				},
				{
					Versions: "18.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"haswell": {
		From:   []string{"ivybridge", "x86_64_v3"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"aes",
			"pclmulqdq",
			"avx",
			"rdrand",
			"f16c",
			"movbe",
			"fma",
			"avx2",
			"bmi1",
			"bmi2"},

		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.9:",
					Flags:    "-march={name} -mtune={name}",
				},
				{
					Versions: "4.8:4.8.5",
					Name:     "core-avx2",
					Flags:    "-march={name} -mtune={name}",
				},
			},

			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:17.9.0",
					Name:     "core-avx2",
					Flags:    "-march={name} -mtune={name}",
				},
				{
					Versions: "18.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"broadwell": {
		From:   []string{"haswell"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"aes",
			"pclmulqdq",
			"avx",
			"rdrand",
			"f16c",
			"movbe",
			"fma",
			"avx2",
			"bmi1",
			"bmi2",
			"rdseed",
			"adx"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "18.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"skylake": {
		From:   []string{"broadwell"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"aes",
			"pclmulqdq",
			"avx",
			"rdrand",
			"f16c",
			"movbe",
			"fma",
			"avx2",
			"bmi1",
			"bmi2",
			"rdseed",
			"adx",
			"clflushopt",
			"xsavec",
			"xsaveopt"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "6.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "18.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"mic_knl": {
		From:   []string{"broadwell"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"aes",
			"pclmulqdq",
			"avx",
			"rdrand",
			"f16c",
			"movbe",
			"avx2",
			"fma",
			"avx2",
			"bmi1",
			"bmi2",
			"rdseed",
			"adx",
			"avx512f",
			"avx512pf",
			"avx512er",
			"avx512cd"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "5.1:",
					Name:     "knl",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Name:     "knl",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Name:     "knl",
					Flags:    "-march={name} -mtune=generic",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "18.0:",
					Name:     "knl",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"skylake_avx512": {
		From:   []string{"skylake", "x86_64_v4"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"aes",
			"pclmulqdq",
			"avx",
			"rdrand",
			"f16c",
			"movbe",
			"fma",
			"avx2",
			"bmi1",
			"bmi2",
			"rdseed",
			"adx",
			"clflushopt",
			"xsavec",
			"xsaveopt",
			"avx512f",
			"clwb",
			"avx512vl",
			"avx512bw",
			"avx512dq",
			"avx512cd"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "6.0:",
					Name:     "skylake-avx512",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Name:     "skylake-avx512",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Name:     "skylake-avx512",
					Flags:    "-march={name} -mtune=generic",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "18.0:",
					Name:     "skylake-avx512",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"cannonlake": {
		From:   []string{"skylake"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"aes",
			"pclmulqdq",
			"avx",
			"rdrand",
			"f16c",
			"movbe",
			"fma",
			"avx2",
			"bmi1",
			"bmi2",
			"rdseed",
			"adx",
			"clflushopt",
			"xsavec",
			"xsaveopt",
			"avx512f",
			"avx512vl",
			"avx512bw",
			"avx512dq",
			"avx512cd",
			"avx512vbmi",
			"avx512ifma",
			"sha",
			"umip"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"apple-clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "18.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"cascadelake": {
		From:   []string{"skylake_avx512"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"aes",
			"pclmulqdq",
			"avx",
			"rdrand",
			"f16c",
			"movbe",
			"fma",
			"avx2",
			"bmi1",
			"bmi2",
			"rdseed",
			"adx",
			"clflushopt",
			"xsavec",
			"xsaveopt",
			"avx512f",
			"clwb",
			"avx512vl",
			"avx512bw",
			"avx512dq",
			"avx512cd",
			"avx512_vnni"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "9.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"apple-clang": {
				{
					Versions: "11.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "19.0.1:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"icelake": {
		From:   []string{"cascadelake", "cannonlake"},
		Vendor: []string{"GenuineIntel"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"popcnt",
			"aes",
			"pclmulqdq",
			"avx",
			"rdrand",
			"f16c",
			"movbe",
			"fma",
			"avx2",
			"bmi1",
			"bmi2",
			"rdseed",
			"adx",
			"clflushopt",
			"xsavec",
			"xsaveopt",
			"avx512f",
			"avx512vl",
			"avx512bw",
			"avx512dq",
			"avx512cd",
			"avx512vbmi",
			"avx512ifma",
			"sha_ni",
			"umip",
			"clwb",
			"rdpid",
			"gfni",
			"avx512_vbmi2",
			"avx512_vpopcntdq",
			"avx512_bitalg",
			"avx512_vnni",
			"vpclmulqdq",
			"vaes"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "icelake-client",
					Versions: "8.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "7.0:",
					Name:     "icelake-client",
					Flags:    "-march={name} -mtune={name}",
				},
				{
					Versions: "6.0:6.9",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Name:     "icelake-client",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"apple-clang": {
				{
					Versions: "10.0.1:",
					Name:     "icelake-client",
					Flags:    "-march={name} -mtune={name}",
				},
				{
					Versions: "10.0.0:10.0.99",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "18.0:",
					Name:     "icelake-client",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"k10": {
		From:   []string{"x86_64"},
		Vendor: []string{"AuthenticAMD"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"sse4a",
			"abm",
			"cx16",
			"3dnow",
			"3dnowext"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "amdfam10",
					Versions: "4.3:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Name:     "amdfam10",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Name:     "amdfam10",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:",
					Flags:    "-msse2",
					Warnings: []string{"Intel's compilers may or may not optimize to the same degree for non-Intel microprocessors for optimizations that are not unique to Intel microprocessors"},
				},
			},
		},
	},
	"bulldozer": {
		From:   []string{"x86_64_v2"},
		Vendor: []string{"AuthenticAMD"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"sse4a",
			"abm",
			"avx",
			"xop",
			"fma4",
			"aes",
			"pclmulqdq",
			"cx16",
			"ssse3",
			"sse4_1",
			"sse4_2"},

		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "bdver1",
					Versions: "4.7:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Name:     "bdver1",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Name:     "bdver1",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:",
					Flags:    "-msse3",
					Warnings: []string{"Intel's compilers may or may not optimize to the same degree for non-Intel microprocessors for optimizations that are not unique to Intel microprocessors"},
				},
			},
		},
	},
	"piledriver": {
		From:   []string{"bulldozer"},
		Vendor: []string{"AuthenticAMD"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"sse4a",
			"abm",
			"avx",
			"xop",
			"fma4",
			"aes",
			"pclmulqdq",
			"cx16",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"bmi1",
			"f16c",
			"fma",
			"tbm"},

		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "bdver2",
					Versions: "4.7:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Name:     "bdver2",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Name:     "bdver2",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:",
					Flags:    "-msse3",
					Warnings: []string{"Intel's compilers may or may not optimize to the same degree for non-Intel microprocessors for optimizations that are not unique to Intel microprocessors"},
				},
			},
		},
	},
	"steamroller": {
		From:   []string{"piledriver"},
		Vendor: []string{"AuthenticAMD"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"sse4a",
			"abm",
			"avx",
			"xop",
			"fma4",
			"aes",
			"pclmulqdq",
			"cx16",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"bmi1",
			"f16c",
			"fma",
			"fsgsbase",
			"tbm"},

		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "bdver3",
					Versions: "4.8:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Name:     "bdver3",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Name:     "bdver3",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:",
					Flags:    "-msse4.2",
					Warnings: []string{"Intel's compilers may or may not optimize to the same degree for non-Intel microprocessors for optimizations that are not unique to Intel microprocessors"},
				},
			},
		},
	},
	"excavator": {
		From:   []string{"steamroller", "x86_64_v3"},
		Vendor: []string{"AuthenticAMD"},
		Features: []string{"mmx",
			"sse",
			"sse2",
			"sse4a",
			"abm",
			"avx",
			"xop",
			"fma4",
			"aes",
			"pclmulqdq",
			"cx16",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"bmi1",
			"f16c",
			"fma",
			"fsgsbase",
			"bmi2",
			"avx2",
			"movbe",
			"tbm"},

		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "bdver4",
					Versions: "4.9:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Name:     "bdver4",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Name:     "bdver4",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:",
					Flags:    "-march={name} -mtune={name}",
					Name:     "core-avx2",
					Warnings: []string{"Intel's compilers may or may not optimize to the same degree for non-Intel microprocessors for optimizations that are not unique to Intel microprocessors"},
				},
			},
		},
	},
	"zen": {
		From:   []string{"x86_64_v3"},
		Vendor: []string{"AuthenticAMD"},
		Features: []string{"bmi1",
			"bmi2",
			"f16c",
			"fma",
			"fsgsbase",
			"avx",
			"avx2",
			"rdseed",
			"clzero",
			"aes",
			"pclmulqdq",
			"cx16",
			"movbe",
			"mmx",
			"sse",
			"sse2",
			"sse4a",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"abm",
			"xsavec",
			"xsaveopt",
			"clflushopt",
			"popcnt"},

		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "znver1",
					Versions: "6.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "4.0:",
					Name:     "znver1",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Name:     "znver1",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:",
					Flags:    "-march={name} -mtune={name}",
					Name:     "core-avx2",
					Warnings: []string{"Intel's compilers may or may not optimize to the same degree for non-Intel microprocessors for optimizations that are not unique to Intel microprocessors"},
				},
			},
		},
	},
	"zen2": {
		From:   []string{"zen"},
		Vendor: []string{"AuthenticAMD"},
		Features: []string{"bmi1",
			"bmi2",
			"f16c",
			"fma",
			"fsgsbase",
			"avx",
			"avx2",
			"rdseed",
			"clzero",
			"aes",
			"pclmulqdq",
			"cx16",
			"movbe",
			"mmx",
			"sse",
			"sse2",
			"sse4a",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"abm",
			"xsavec",
			"xsaveopt",
			"clflushopt",
			"popcnt",
			"clwb"},

		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "znver2",
					Versions: "9.0:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "9.0:",
					Name:     "znver2",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "2.2:",
					Name:     "znver2",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"intel": {
				{
					Versions: "16.0:",
					Flags:    "-march={name} -mtune={name}",
					Name:     "core-avx2",
					Warnings: []string{"Intel's compilers may or may not optimize to the same degree for non-Intel microprocessors for optimizations that are not unique to Intel microprocessors"},
				},
			},
		},
	},
	"zen3": {
		From:   []string{"zen2"},
		Vendor: []string{"AuthenticAMD"},
		Features: []string{"bmi1",
			"bmi2",
			"f16c",
			"fma",
			"fsgsbase",
			"avx",
			"avx2",
			"rdseed",
			"clzero",
			"aes",
			"pclmulqdq",
			"cx16",
			"movbe",
			"mmx",
			"sse",
			"sse2",
			"sse4a",
			"ssse3",
			"sse4_1",
			"sse4_2",
			"abm",
			"xsavec",
			"xsaveopt",
			"clflushopt",
			"popcnt",
			"clwb",
			"vaes",
			"vpclmulqdq",
			"pku"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "znver3",
					Versions: "10.3:",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "12.0:",
					Name:     "znver3",
					Flags:    "-march={name} -mtune={name}",
				},
			},
			"aocc": {
				{
					Versions: "3.0:",
					Name:     "znver3",
					Flags:    "-march={name} -mtune={name}",
				},
			},
		},
	},
	"ppc64": {
		From:     []string{},
		Vendor:   []string{"generic"},
		Features: []string{},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "powerpc64",
					Versions: ":",
					Flags:    "-mcpu={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: ":",
					Name:     "znver3",
					Flags:    "-mcpu={name} -mtune={name}",
				},
			},
		},
	},
	"power7": {
		From:       []string{"ppc64"},
		Vendor:     []string{"IBM"},
		Generation: 7,
		Features:   []string{},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.4:",
					Flags:    "-mcpu={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Name:     "znver3",
					Flags:    "-mcpu={name} -mtune={name}",
				},
			},
		},
	},
	"power8": {
		From:       []string{"power7"},
		Vendor:     []string{"IBM"},
		Generation: 8,
		Features:   []string{},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.9:",
					Flags:    "-mcpu={name} -mtune={name}",
				},
				{
					Versions: "4.8:4.8.5",
					Flags:    "-mcpu={name} -mtune={name}",
					Warnings: []string{"Using GCC 4.8 to optimize for Power 8 might not work if you are not on Red Hat Enterprise Linux 7, where a custom backport of the feature has been done. Upstream support from GCC starts in version 4.9"},
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-mcpu={name} -mtune={name}",
				},
			},
		},
	},
	"power9": {
		From:       []string{"power8"},
		Vendor:     []string{"IBM"},
		Generation: 9,
		Features:   []string{},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "6.0:",
					Flags:    "-mcpu={name} -mtune={name}",
				},
				{
					Versions: "4.8:4.8.5",
					Flags:    "-mcpu={name} -mtune={name}",
					Warnings: []string{"Using GCC 4.8 to optimize for Power 8 might not work if you are not on Red Hat Enterprise Linux 7, where a custom backport of the feature has been done. Upstream support from GCC starts in version 4.9"},
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-mcpu={name} -mtune={name}",
				},
			},
		},
	},
	"ppc64le": {
		From:     []string{},
		Vendor:   []string{"generic"},
		Features: []string{},

		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "powerpc64le",
					Versions: "4.8:",
					Flags:    "-mcpu={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: ":",
					Flags:    "-mcpu={name} -mtune={name}",
				},
			},
		},
	},
	"power8le": {
		From:       []string{"ppc64le"},
		Vendor:     []string{"IBM"},
		Generation: 8,
		Features:   []string{},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "power8",
					Versions: "4.9:",
					Flags:    "-mcpu={name} -mtune={name}",
				},
				{
					Name:     "power8",
					Warnings: []string{"Using GCC 4.8 to optimize for Power 8 might not work if you are not on Red Hat Enterprise Linux 7, where a custom backport of the feature has been done. Upstream support from GCC starts in version 4.9"},
					Versions: "4.8:4.8.5",
					Flags:    "-mcpu={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Family:   []string{"ppc64le"},
					Name:     "power8",
					Flags:    "-mcpu={name} -mtune={name}",
				},
			},
		},
	},
	"power9le": {
		From:       []string{"power8le"},
		Vendor:     []string{"IBM"},
		Generation: 9,
		Features:   []string{},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Name:     "power9",
					Versions: "6.0:",
					Flags:    "-mcpu={name} -mtune={name}",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Family:   []string{"ppc64le"},
					Name:     "power9",
					Flags:    "-mcpu={name} -mtune={name}",
				},
			},
		},
	},
	"aarch64": {
		From:     []string{},
		Vendor:   []string{"generic"},
		Features: []string{},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.8.0:",
					Flags:    "-march=armv8-a -mtune=generic",
				},
			},
			"clang": {
				{
					Versions: ":",
					Flags:    "-march=armv8-a -mtune=generic",
				},
			},
			"apple-clang": {
				{
					Versions: ":",
					Flags:    "-march=armv8-a -mtune=generic",
				},
			},
			"arm": {
				{
					Versions: ":",
					Flags:    "-march=armv8-a -mtune=generic",
				},
			},
		},
	},
	"thunderx2": {
		From:   []string{"aarch64"},
		Vendor: []string{"Cavium"},
		Features: []string{"fp",
			"asimd",
			"evtstrm",
			"aes",
			"pmull",
			"sha1",
			"sha2",
			"crc32",
			"atomics",
			"cpuid",
			"asimdrdm"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.8:4.8.9",
					Flags:    "-march=armv8-a",
				},
				{
					Versions: "4.9:5.9",
					Flags:    "-march=armv8-a+crc+crypto",
				},
				{
					Versions: "6:6.9",
					Flags:    "-march=armv8.1-a+crc+crypto",
				},
				{
					Versions: "7:",
					Flags:    "-mcpu=thunderx2t99",
				},
			},
			"clang": {
				{
					Versions: "3.9:4.9",
					Flags:    "-march=armv8.1-a+crc+crypto",
				},
				{
					Versions: "5:",
					Flags:    "-mcpu=thunderx2t99",
				},
			},
		},
	},
	"a64fx": {
		From:   []string{"aarch64"},
		Vendor: []string{"Fujitsu"},
		Features: []string{"fp",
			"asimd",
			"evtstrm",
			"aes",
			"pmull",
			"sha1",
			"sha2",
			"crc32",
			"atomics",
			"cpuid",
			"asimdrdm",
			"fphp",
			"asimdhp",
			"fcma",
			"dcpop",
			"sve"},

		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.8:4.8.9",
					Flags:    "-march=armv8-a",
				},
				{
					Versions: "4.9:5.9",
					Flags:    "-march=armv8-a+crc+crypto",
				},
				{
					Versions: "6:6.9",
					Flags:    "-march=armv8.1-a+crc+crypto",
				},
				{
					Versions: "7:7.9",
					Flags:    "-march=armv8.2-a+crc+crypto+fp16",
				},
				{
					Versions: "8:",
					Flags:    "-march=armv8.2-a+crc+aes+sha2+fp16+sve -msve-vector-bits=512",
				},
			},
			"clang": {
				{
					Versions: "3.9:4.9",
					Flags:    "-march=armv8.2-a+crc+crypto+fp16",
				},
				{
					Versions: "5:",
					Flags:    "-march=armv8.2-a+crc+crypto+fp16+sve",
				},
			},
			"arm": {
				{
					Versions: "20:",
					Flags:    "-march=armv8.2-a+crc+crypto+fp16+sve",
				},
			},
		},
	},
	"graviton": {
		From:   []string{"aarch64"},
		Vendor: []string{"ARM"},
		Features: []string{"fp",
			"asimd",
			"evtstrm",
			"aes",
			"pmull",
			"sha1",
			"sha2",
			"crc32",
			"cpuid"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.8:4.8.9",
					Flags:    "-march=armv8-a",
				},
				{
					Versions: "4.9:5.9",
					Flags:    "-march=armv8-a+crc+crypto",
				},
				{
					Versions: "6:",
					Flags:    "-march=armv8-a+crc+crypto -mtune=cortex-a72",
				},
			},
			"clang": {
				{
					Versions: "3.9:",
					Flags:    "-march=armv8-a+crc+crypto",
				},
			},
			"arm": {
				{
					Versions: "20:",
					Flags:    "-march=armv8.2-a+crc+crypto+fp16+sve",
				},
			},
		},
	},
	"graviton2": {
		From:   []string{"aarch64"},
		Vendor: []string{"ARM"},
		Features: []string{"fp",
			"asimd",
			"evtstrm",
			"aes",
			"pmull",
			"sha1",
			"sha2",
			"crc32",
			"atomics",
			"fphp",
			"asimdhp",
			"cpuid",
			"asimdrdm",
			"lrcpc",
			"dcpop",
			"asimddp",
			"ssbs"},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "4.8:4.8.9",
					Flags:    "-march=armv8-a",
				},
				{
					Versions: "4.9:5.9",
					Flags:    "-march=armv8-a+crc+crypto",
				},
				{
					Versions: "6:6.9",
					Flags:    "-march=armv8.1-a",
				},
				{
					Versions: "7:7.9",
					Flags:    "-march=armv8.2-a+fp16 -mtune=cortex-a72",
				},
				{
					Versions: "8.0:8.0",
					Flags:    "-march=armv8.2-a+fp16+dotprod+crypto -mtune=cortex-a72",
				},
				{
					Versions: "8.1:8.9",
					Flags:    "-march=armv8.2-a+fp16+rcpc+dotprod+crypto -mtune=cortex-a72",
				},
				{
					Versions: "9.0:",
					Flags:    "-march=armv8.2-a+fp16+rcpc+dotprod+crypto -mtune=neoverse-n1",
				},
			},
			"clang": {
				{
					Versions: "3.9:4.9",
					Flags:    "-march=armv8.2-a+fp16+crc+crypto",
				},
				{
					Versions: "5:",
					Flags:    "-march=armv8.2-a+fp16+rcpc+dotprod+crypto",
				},
			},
			"arm": {
				{
					Versions: "20:",
					Flags:    "-march=armv8.2-a+fp16+rcpc+dotprod+crypto",
				},
			},
		},
	},
	"m1": {
		From:     []string{"aarch64"},
		Vendor:   []string{"Apple"},
		Features: []string{},
		Compilers: map[string][]Compiler{
			"gcc": {
				{
					Versions: "8.0:",
					Flags:    "-march=armv8.4-a -mtune=generic",
				},
			},
			"clang": {
				{
					Versions: "9.0:",
					Flags:    "-march=armv8.4-a",
				},
			},
			"apple-clang": {
				{
					Versions: "11.0:",
					Flags:    "-march=armv8.4-a",
				},
			},
		},
	},
	"arm": {
		From:     []string{},
		Vendor:   []string{"generic"},
		Features: []string{},
		Compilers: map[string][]Compiler{
			"clang": {
				{
					Versions: ":",
					Family:   []string{"arm"},
					Flags:    "-march={family} -mcpu=generic",
				},
			},
			"apple-clang": {
				{
					Versions: "11.0:",
					Flags:    "-march=armv8.4-a",
				},
			},
		},
	},
	"ppc": {
		From:     []string{},
		Vendor:   []string{"generic"},
		Features: []string{},
	},
	"ppcle": {
		From:     []string{},
		Vendor:   []string{"generic"},
		Features: []string{},
	},
	"sparc": {
		From:     []string{},
		Vendor:   []string{"generic"},
		Features: []string{},
	},
	"sparc64": {
		From:     []string{},
		Vendor:   []string{"generic"},
		Features: []string{},
	},
}

type FeatureAlias struct {
	Reason   string   `json:"reason"`
	AnyOf    []string `json:"any_of"`
	Families []string `json:"families"`
}

var FeatureAliases = map[string]FeatureAlias{
	"sse": {
		Reason: "ssse3 is a superset of sse3 and might be the only one listed",
		AnyOf:  []string{"ssse3"},
	},

	"avx512": {
		Reason: "avx512 indicates generic support for any of the avx512 instruction sets",
		AnyOf: []string{
			"avx512f",
			"avx512vl",
			"avx512bw",
			"avx512dq",
			"avx512cd"},
	},
	"altivec": {
		Reason: "altivec is supported by Power PC architectures, but might not be listed in features",
		Families: []string{
			"ppc64le",
			"ppc64"},
	},
	"vsx": {
		Reason:   "VSX alitvec extensions are supported by PowerISA from v2.06 (Power7+), but might not be listed in features",
		Families: []string{"ppc64le", "ppc64"},
	},
	"fma": {
		Reason: "FMA has been supported by PowerISA since Power1, but might not be listed in features",
		Families: []string{"ppc64le",
			"ppc64"},
	},
	"sse4.1": {
		Reason: "permits to refer to sse4_1 also as sse4.1",
		AnyOf:  []string{"sse4_1"},
	},
	"sse4.2": {
		Reason: "permits to refer to sse4_2 also as sse4.2",
		AnyOf:  []string{"sse4_2"},
	},
	"neon": {
		Reason:   "NEON is required in all standard ARMv8 implementations",
		Families: []string{"aarch64"},
	},
}

// Conversions that map some platform specific values to canonical values
type Conversion map[string]string

var Conversions = map[string]Conversion{
	"arm_vendors": {
		"0x41": "ARM",
		"0x42": "Broadcom",
		"0x43": "Cavium",
		"0x44": "DEC",
		"0x46": "Fujitsu",
		"0x48": "HiSilicon",
		"0x49": "Infineon Technologies AG",
		"0x4d": "Motorola",
		"0x4e": "Nvidia",
		"0x50": "APM",
		"0x51": "Qualcomm",
		"0x53": "Samsung",
		"0x56": "Marvell",
		"0x61": "Apple",
		"0x66": "Faraday",
		"0x68": "HXT",
		"0x69": "Intel",
	},
	"darwin_flags": {
		"sse4.1":  "sse4_1",
		"sse4.2":  "sse4_2",
		"avx1.0":  "avx",
		"clfsopt": "clflushopt",
		"xsave":   "xsavec xsaveopt",
	},
}
