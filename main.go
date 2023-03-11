package main
import ("fmt")

func main() {

	threshold := 7.0
	virustype := "sars2"
	sars2signature := "LSTV"

	switch {
		case virustype == "sars2":
		
			//sars2 := make(map[string]string)
			sars2SignCount := make(map[string]float64)
			// sars2SignCount["sars2" + " " + "LSTV"] = 7.0
			// sars2SignCount["sars2" + " " + "LSTV" + " " + "Count"] = 1
			// sars2SignCount["sars2" + " " + "LST"] = 7.4
			// sars2SignCount["sars2" + " " + "LST" + " " + "Count"] = 1
			
					 
			if _, ok := sars2SignCount[virustype + " " + sars2signature]; ok {
				if sars2SignCount[virustype + " " + sars2signature] != threshold {
					sars2SignCount[virustype + " " + sars2signature] = threshold
					sars2SignCount[virustype + " " + sars2signature + " " + "Count"] = 1
					
				} else {
					sars2SignCount[virustype + " " + sars2signature + " " + "Count"] = sars2SignCount[virustype + " " + sars2signature + " " + "Count"] + 1
				}
			} else {
				sars2SignCount[virustype + " " + sars2signature] = threshold
				sars2SignCount[virustype + " " + sars2signature + " " + "Count"] = 1
			}

			fmt.Println("SARS2", sars2SignCount)
			
		// case virustype == "mers":
		// 	merssignature := "LSTV"
		// 	//sars2 := make(map[string]string)
		// 	mersSignCount := make(map[string]int)
		// 	//mersSignCount["LSTV"] = 1
		// 	if _, ok := mersSignCount[merssignature]; ok {
		// 		mersSignCount[merssignature] = mersSignCount[merssignature] + 1
		// 	} else {
		// 		mersSignCount[merssignature] = 1
		// 	}

		// 	fmt.Println("MERS", mersSignCount)

	}






	//sars2["Type"] = 
	//sars2SignCount["LSTV"] = 1
	//sars2["Number"] = 
	//fmt.Println(sars2["Sign"])
	// if sars2signature != sars2SignCount[] {
	// 	//sars2["Sign"] = sars2["Sign"] + 1
	// 	sars2SignCount[sars2signature] = 1
	// } else {
	// 	sars2SignCount[sars2signature] =  sars2SignCount[sars2signature] + 1
	// }
}