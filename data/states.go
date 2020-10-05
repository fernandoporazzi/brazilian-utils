package data

// State is a model representing a single Brazilian State
type State struct {
	Code      string
	AreaCodes []int
	IeLength  []int
	Name      string
}

// States holds needed informations about all Brazilian states
var States = map[string]State{
	"AC": {"2", []int{68}, []int{13}, "Acre"},
	"AL": {"4", []int{82}, []int{9}, "Alagoas"},
	"AP": {"2", []int{96}, []int{9}, "Amapá"},
	"AM": {"2", []int{92, 97}, []int{9}, "Amazonas"},
	"BA": {"5", []int{71, 73, 74, 75, 77}, []int{8, 9}, "Bahia"},
	"CE": {"3", []int{85, 88}, []int{9}, "Ceará"},
	"DF": {"1", []int{61}, []int{13}, "Distrito Federal"},
	"ES": {"7", []int{27, 28}, []int{9}, "Espírito Santo"},
	"GO": {"1", []int{62, 64}, []int{9}, "Goiás"},
	"MA": {"3", []int{98, 99}, []int{9}, "Maranhão"},
	"MG": {"6", []int{31, 32, 33, 34, 35, 37, 38}, []int{13}, "Minas Gerais"},
	"MT": {"1", []int{65, 66}, []int{11}, "Mato Grosso"},
	"MS": {"1", []int{67}, []int{9}, "Mato Grosso do Sul"},
	"PA": {"2", []int{91, 93, 94}, []int{9}, "Pará"},
	"PB": {"4", []int{83}, []int{9}, "Paraíba"},
	"PE": {"4", []int{81, 87}, []int{9}, "Pernambuco"},
	"PI": {"3", []int{86, 89}, []int{9}, "Piauí"},
	"PR": {"9", []int{41, 42, 43, 44, 45, 46}, []int{10}, "Paraná"},
	"RJ": {"7", []int{21, 22, 24}, []int{8}, "Rio de Janeiro"},
	"RN": {"4", []int{84}, []int{9, 10}, "Rio Grande do Norte"},
	"RO": {"2", []int{69}, []int{14}, "Rondônia"},
	"RS": {"0", []int{51, 53, 54, 55}, []int{10}, "Rio Grande do Sul"},
	"RR": {"2", []int{95}, []int{9}, "Roraima"},
	"SC": {"9", []int{47, 48, 49}, []int{9}, "Santa Catarina"},
	"SE": {"5", []int{79}, []int{9}, "Sergipe"},
	"SP": {"8", []int{11, 12, 13, 14, 15, 16, 17, 18, 19}, []int{12}, "São Paulo"},
	"TO": {"1", []int{63}, []int{9, 11}, "Tocantins"},
}

// AreaCodes holds all area codes of the country
var AreaCodes []int

func init() {
	for _, state := range States {
		for _, code := range state.AreaCodes {
			AreaCodes = append(AreaCodes, code)
		}
	}
}
