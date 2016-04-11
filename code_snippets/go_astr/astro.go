package main
/******************************************
*
* File Name : astro.go
* Purpose:
* Author: pggsx
* Creation Date: 28-03-2016
Last Modified: Mon 11 Apr 2016 01:04:26 PM EDT
* Choose-Your-Adventure Program for one my gen-ed classes
* Based on user input, story scrolls forwards and generates text-based adventure
* exploring various exo-planets 
* representing various scenes in the storyline.
*
********************************************/
//imports
import(
	"fmt"
	"log"
"os"
	)


// main function
/*
*
*
*
*/

func main(){
	var control int
	var start bool = false
	control = getMainMenu()
	switch control {
	case 1:
		start = true
		charString := getCharacterSetup()
		fmt.Println("You have selected: " + charString)
		fmt.Println("Generating Story.....")
		genStoryLine(start)
	case 2:
	case 3:
	fmt.Println("Welcome to ASTR121 Choose Your Own Adventure")
	fmt.Println("Please Choose one of the following Options:")
	fmt.Println("1 Start Adventure")
	fmt.Println("3 Shows this help Menu")
	fmt.Println("4 Exits Program")
	case 4:
	return
	default:
		fmt.Println("Error Unidentified value passed")
	}


}

/*
*
*
*/
func getMainMenu() int {
	var input int
	fmt.Println("Welcome to ASTR121 Choose Your Own Adventure")
	fmt.Println("Please Choose one of the following Options:")
	fmt.Println("1 Start Adventure")
	fmt.Println("3 Shows this help Menu")
	fmt.Println("4 Exits Program")
	if _,err := fmt.Scanf("%d",&input); err != nil{
		fmt.Println("Invalid Menu Selection Exiting now..")
		log.Print("Scanf for Main Menu failed due to",err)
	}
	return input
}



/*
*
*
*/
func getCharacterSetup() string{
	var input int
	var character string
	fmt.Println("Choose from one of the following premade characters or create your own character")
	fmt.Println(" 1 Ewoks from Planet Endor")
	fmt.Println(" 2 Borgs from Borg HiveMind looking to Assimilate")
	fmt.Println(" 3 Space Captain Zaphod BeebleBrox of the ship Heart of Gold Looking for the Ultimate Question in Life ")
	fmt.Println(" 4 Bobba Fett tracking down Han Solo and the Millenium Falcon")
	fmt.Println(" 5 Create your own Character")
	if _,err := fmt.Scanf("%d",&input); err != nil{
		fmt.Println("Invalid Character Setup Selection Exiting now..")
		log.Print("Scanf for Character Setup failed due to",err)
	}
	switch input{
		case 1: character = "Ewoks from Planet Endor"
		case 2: character = "Borgs from Borg HiveMInd Looking to Assmiilate"
		case 3: character = "Space Captain Zaphod BeebleBrox of the ship Heart of Gold Looking for the Ultimate Question in Life"
		case 4: character = "Bobba the Fett tracking down Han Solo and the Millenium Falcon"
	case 5:
		fmt.Println("Insert Character Name/Description")
		if _,err := fmt.Scanf("%s",character); err != nil{
			fmt.Println("Invalid Character Setup Selection Exiting now..")
			log.Print("Scanf for Character Setup failed due to",err)
		}
	}
	return character
}

/*
*
*
*/
func genStoryLine(start bool){
	var response int
	var FTLCounter int
	var continueStr string

	if start {
		fmt.Println("We begin our adventure on the green and blue Marble called Earth")
		fmt.Println("Our Adventurers begin preparation for their long journey ahead of them")
		fmt.Println("to explore far and distant planets in the far regions of space in hopes")
		fmt.Println("of finding a potential new homes for all of humankind.....")
		fmt.Println("Your spaceship is unlike most and relies on answering questions to power it")
		fmt.Println("to power itself throughout your journey. If you answer a question wrong then,")
		fmt.Println("your ship stops in its tracks. So make sure to answer questions gly")
		fmt.Println("God Luck and God Speed~! :)")
	} 

		fmt.Println("Your First Destination is  the Planet 51 Pegasi b")
		fmt.Println("Prior to FTL Jump, Please answer the following Question:")
		fmt.Println("--------------------------------------------------------")
			fmt.Println("What is an exoplanet?")
			fmt.Println("1 A planet beyond our solar system that most resemble our parent star")
			fmt.Println("2 A planet that eats other planets")
			fmt.Println("3 A totally normal planet with a weird name given to it")
			if _,err := fmt.Scanf("%d\n",&response); err != nil{
				fmt.Println("Invalid Question Response Selection Exiting now..")
				log.Print("Scanf for Question Response failed due to",err)
}
				if response == 1{
					fmt.Println("Correct~!")
					fmt.Print("Powering up FTL Drive")
					for FTLCounter <= 5{
						fmt.Print("*")
						FTLCounter++
					}
					fmt.Println("FTL Jump!")
				}else {
					fmt.Println("You have chosen an incorrect answer.Please try again..")
					fmt.Println("FTLDrive Powering Down...")

				}
	
		
		fmt.Println("You have now arrived at Planet 51 Pegasi b")
		fmt.Println("Current Living Conditions:")
		fmt.Println("Type : G Star")
		fmt.Println("Mass is Approximately half of Jupiter")
		fmt.Println("Tempearture is Approximately 1265 K Hot!")
		fmt.Println("Potentially Viable for Human Habitation given Teraforming Operations however Uncertain")
		fmt.Println("Would you like to stop your journey or continue to the next viable planet?")
		fmt.Println("Enter Y or N to continue or not.... ")
		if _,err := fmt.Scanf("%s",&continueStr); err != nil{
			fmt.Println("Invalid Character Setup Selection Exiting now..")
			log.Print("Scanf for Character Setup failed due to",err)
		}
		if continueStr == "Y"{

				fmt.Println("How does one classify an exoplanet?")
				fmt.Println("1 Planets that are just hanging around willy-nilly")
				fmt.Println("2 Planets that are surrounded by asteroids")
				fmt.Println("3 Planets must common around the sun-like stars")
				if _,err := fmt.Scanf("%d",&response); err != nil{
					fmt.Println("Invalid Question Response Selection Exiting now..")
					log.Print("Scanf for Question Response failed due to",err)
}
					if response == 3{
						fmt.Println("Correct~!")
						fmt.Print("Powering up FTL Drive")
						for FTLCounter <= 5{
							fmt.Print("*")
							FTLCounter++
						}
						fmt.Println("FTL Jump!")
					}else {
						fmt.Println("FTLDrive Powering Down...")
					}
		} else{
			fmt.Println("You now begin your settlement on Planet 51 Pegasi B")
			fmt.Println("While your initial strugles may seem to be daunting at first, you realize that your terraforming operations")
			fmt.Println("will help the future of mankind")
			os.Exit(1)
		}
	

	fmt.Println("You have now arrived at Planet HD 209458 b")
	fmt.Println("Current Living Conditions:")
	fmt.Println("Type : G0V Star")
	fmt.Println("Carbon-rich planet with lack of cloud cover")
	fmt.Println("Would you like to stop your journey or continue to the next viable planet?")
	fmt.Println("Enter Y or N to continue or not.... ")
	if _,err := fmt.Scanf("%s",&continueStr); err != nil{
		fmt.Println("Invalid Character Setup Selection Exiting now..")
		log.Print("Scanf for Character Setup failed due to",err)
	}
	if continueStr == "Y"{
			fmt.Println("What is the habitable zone and what purpose does it serve in context with space exploration")
			fmt.Println("1 Are you talking about the ChalkZone?")
			fmt.Println("2 A range of distances from a star where planet's temperature allows for liquid water oceans, critical for human life")
			fmt.Println("3 I have no idea what you are talking about")
			if _,err := fmt.Scanf("%d",&response); err != nil{
				fmt.Println("Invalid Question Response Selection Exiting now..")
				log.Print("Scanf for Question Response failed due to",err)
}
				if response == 2{
					fmt.Println("Correct~!")
					fmt.Print("Powering up FTL Drive")
					for FTLCounter <= 5{
						fmt.Print("*")
						FTLCounter++
					}
					fmt.Println("FTL Jump!")
				}else {
					fmt.Println("You have chosen an incorrect answer.Please try again..")
					fmt.Println("FTLDrive Powering Down...")
					fmt.Println("Exiting now....")
					os.Exit(-1)

				} 
}else{
		fmt.Println("You have chosen to colonize the HD 209548 b")
		fmt.Println("As you descend into the atmosphere you begin to realize the possiblities of living on this planet")
		fmt.Println("`This will be an interesting journey...` you think to yourself as you setup base-camp near your ship")
		fmt.Println("The Journey Continues......")
os.Exit(1)	
}

	fmt.Println("You have now 55 Cancri e")
	fmt.Println("Current Living Conditions:")
	fmt.Println("Type : G8V Star")
	fmt.Println("Year is considered to be onyl 17 Hours and 41 minutes long in Earth Time")
	fmt.Println("Would you like to stop your journey or continue to the next viable planet?")
	fmt.Println("Enter Y or N to continue or not.... ")
	if _,err := fmt.Scanf("%s",&continueStr); err != nil{
		fmt.Println("Invalid Character Setup Selection Exiting now..")
		log.Print("Scanf for Character Setup failed due to",err)
	}
	if continueStr == "Y"{

			fmt.Println("Why do astronomers rely on the sun as a criteria when looking for exo-planets?")
			fmt.Println("1 What a Sun?")
			fmt.Println("2 I have no idea what you are talking about")
			fmt.Println("3 Most of the time sun-like stars have been known to have habitable planets around sun-like stars")
			if _,err := fmt.Scanf("%d",&response); err != nil{
				fmt.Println("Invalid Question Response Selection Exiting now..")
				log.Print("Scanf for Question Response failed due to",err)
}
				if continueStr == "Y"{
				if response == 3{
					fmt.Println("Correct~!")
					fmt.Print("Powering up FTL Drive")
					for FTLCounter <= 5{
						fmt.Print("*")
						FTLCounter++
					}
					fmt.Println("------")
					fmt.Println("FTL Jump!")
				}
				}else {
					fmt.Println("You have chosen an incorrect answer.Please try again..")
					fmt.Println("FTLDrive Powering Down...")
					fmt.Println("Exiting now....")
					os.Exit(-1)
				}
			}else{
		fmt.Println("You have chosen to colonize the Planet 55 Cancri e")
		fmt.Println("As you descend into the atmosphere you begin to realize the possiblities of living on this planet")
		fmt.Println("`This will be an interesting journey...` you think to yourself as you setup base-camp near your ship")
		fmt.Println("The Journey Continues......")
		os.Exit(1)
	}
	fmt.Println("You have now arrived at Planet HD  80606 b")
	fmt.Println("Current Living Conditions:")
	fmt.Println("Type : G5V Star")
	fmt.Println("Considered to have a weird orbit (similar to Hailey's Comet around the Sun")
	fmt.Println("Given weird orbit, it has variable environment conditions")
	fmt.Println("Would you like to stop your journey or continue to the next viable planet?")
	fmt.Println("Enter Y or N to continue or not.... ")
	if _,err := fmt.Scanf("%s",&continueStr); err != nil{
		fmt.Println("Invalid Character Setup Selection Exiting now..")
		log.Print("Scanf for Character Setup failed due to",err)
	}
	if continueStr == "Y"{

			fmt.Println("How have exoplanets defined planetary research")	
			fmt.Println("1 Research into exoplanets have led to discoveries about sun-like stars including magnetic braking and other revelations")
			fmt.Println("2 Planetary observations into how stars are born and destroyed")
			fmt.Println("3 I have no idea what you are talking about")
			if _,err := fmt.Scanf("%d",&response); err != nil{
				fmt.Println("Invalid Question Response Selection Exiting now..")
				log.Print("Scanf for Question Response failed due to",err)
}
				if response == 1{
					fmt.Println("Correct~!")
					fmt.Print("Powering up FTL Drive")
					for FTLCounter <= 5{
						fmt.Print("*")
						FTLCounter++
					}
					fmt.Println("------")
					fmt.Println("FTL Jump!")
				}else {
					fmt.Println("You have chosen an incorrect answer.Please try again..")
					fmt.Println("FTLDrive Powering Down...")
					fmt.Println("Exiting now....")
					os.Exit(-1)
				}
		}else{
		fmt.Println("You have chosen to colonize the HD 80606 b")
		fmt.Println("As you descend into the atmosphere you begin to realize the possiblities of living on this planet")
		fmt.Println("`This will be an interesting journey...` you think to yourself as you setup base-camp near your ship")
		fmt.Println("The Journey Continues......")
		os.Exit(1)
	}

	fmt.Println("You have now arrived at Planet WASP-33b")
	fmt.Println("Current Living Conditions:")
	fmt.Println("Type : A5 Star")
	fmt.Println("Mass is Approximately half of Jupiter")
	fmt.Println("Tempearture is Approximately 1265 K Hot!")
	fmt.Println("Potentially Viable for Human Habitation given viable atmosphere")
	fmt.Println("Would you like to stop your journey or continue to the next viable planet?")
	fmt.Println("Enter Y or N to continue or not.... ")
	if _,err := fmt.Scanf("%s",&continueStr); err != nil{
		fmt.Println("Invalid Character Setup Selection Exiting now..")
		log.Print("Scanf for Character Setup failed due to",err)
	}
	if continueStr == "Y"{
			fmt.Println("How does the Kepler telescope search for exoplanets")
			fmt.Println("1 It looks for stars.....")
			fmt.Println("2 I have no idea what you are talking about")
			fmt.Println("3 The Kepler telescope uses the transit method to find stars by measuring how much a star's light dims when a planet passes in front of it")
			if _,err := fmt.Scanf("%d",&response); err != nil{
				fmt.Println("Invalid Question Response Selection Exiting now..")
				log.Print("Scanf for Question Response failed due to",err)
}
				if response == 3{
					fmt.Println("Correct~!")
					fmt.Print("Powering up FTL Drive")
					for FTLCounter <= 5{
						fmt.Print("*")
						FTLCounter++
					}
					fmt.Println("------")
					fmt.Println("FTL Jump!")
				}else {
					fmt.Println("You have chosen an incorrect answer Please try again..")
					fmt.Println("FTLDrive Powering Down...")
					fmt.Println("Exiting now....")
					os.Exit(-1)
				}
			}else{
		fmt.Println("You have chosen to colonize the Planet WASP-33b")
		fmt.Println("As you descend into the atmosphere you begin to realize the possiblities of living on this planet")
		fmt.Println("`This will be an interesting journey...` you think to yourself as you setup base-camp near your ship")
		fmt.Println("The Journey Continues......")
		os.Exit(1)
	}
}
func saveFile(){






}

func loadFile(){


}
