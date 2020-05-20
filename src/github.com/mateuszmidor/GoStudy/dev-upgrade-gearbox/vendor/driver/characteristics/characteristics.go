package characteristics

// Characteristics represents engine params
var characteristics = [...]interface{}{2000.0, 1000.0, 1000.0, 0.5, 2500.0, 4500.0, 1500.0, 0.5, 5000.0, 0.7, 5000.0, 5000.0, 1500.0, 2000.0, 3000.0, 6500.0, 14.0}

/*
rivate Object[] characteristics = new Object[]{
	2000d, // 0
	1000d, // 1
	1000d, // 2
	0.5d,  // 3
	2500d, // 4
	4500d, // 5
	1500d, // 6
	0.5d,  // 7
	5000d, // 8
	0.7d,  // 9
	5000d, // 10
	5000d, // 11
	1500d,
	2000d,
	3000d,
	6500d,
	14d
};

index:
0 - tryb ECO - rpm czy podbić bieg przy przyspieszaniu
1 - tryb ECO - rpm czy redukować bieg przy przyspieszaniu
2 - tryb COMFORT - rpm czy redukować bieg przy przyspieszaniu
3 - tryb COMFORT - threshold naciśnięcia pedału gazu, żeby jeszcze to nie był kickdown
4 - tryb COMFORT - rpm czy podbić bieg przy przyspieszaniu
5 - tryb COMFORT - rpm czy zrzucić bieg w kickdown
6 - tryb SPORT -  rpm czy zrzucić bieg przy przyspieszaniu
7 - tryb SPORT -  threshold naciśnięcia pedału gazu, żeby czy lekko przyspieszamy
8 - tryb SPORT -  rpm czy zwiekszamy bieg w lekkim przyspieszeniu
9 - tryb SPORT -  threshold naciśnięcia pedału gazu, żeby czy lekki kickdown
10 - tryb SPORT -  rpm czy redukcja w lekkim kickdown
11 - tryb SPORT -  rpm czy zrzucić 2 biegi w MOCNYM kickdown - zapier...
12 - tryb ECO - rpm zrzucić bieg przy hamowaniu
13 - tryb COMFORT - rpm zrzucić bieg przy hamowaniu
14 - tryb SPORT - rpm zrzucić bieg przy hamowaniu
15 - ???
17 - tryb HIDDEN MODE - kiedy podbić bieg przy przyspieszaniu
18 - tryb HIDDEN MODE - kiedy redukować bieg przy przyspieszaniu w hidden mode
19 - tryb HIDDEN MODE - kiedy redukować bieg przy hamowaniu w hidden mode (chyba)
*/

// GetRPMForSpeedingUpEco returns (min, max) RPM
func GetRPMForSpeedingUpEco() (float64, float64) {
	return characteristics[1].(float64), characteristics[0].(float64)
}

// GetRPMForSpeedingUpSport returns (min, max) RPM
func GetRPMForSpeedingUpSport() (float64, float64) {
	return characteristics[6].(float64), characteristics[8].(float64)
}

// GetGasKickDown1Sport getter
func GetGasKickDown1Sport() float64 {
	return characteristics[7].(float64)
}

// GetGasKickDown2Sport getter
func GetGasKickDown2Sport() float64 {
	return characteristics[9].(float64)
}
