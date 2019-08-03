package pakage

//Global Variables to be initialized on app start
var PackageModelMap map[int64]PackageModel

func Init() {
	PackageModelMap = make(map[int64]PackageModel)

	LoadPakageDataFile()
}

func LoadPakageDataFile() {
	for _, pkg := range PakageData {
		PackageModelMap[pkg.ID] = pkg
	}
}
