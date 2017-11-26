package main

import (
	"bufio"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"
	"strconv"

)
type CEDICTDATA struct {
	Traditional string `bson:"Traditional" json:"Traditional"`
	Simplified string `bson:"Simplified" json:"Simplified"`
	PinyinNumbered string `bson:"PinyinNumbered" json:"PinyinNumbered"`
	Pinyin string `bson:"Pinyin" json:"Pinyin"`
	Definition string `bson:"Definition" json:"Definition"`
}

func replaceAtIndex(input string,  index int,replacement string) string {
	return input[:index] + string(replacement) + input[index+1:]
}

func detectTone(pinyinstring string) int {
	if(strings.Contains(pinyinstring,"1")){
		return 1;
	}else if(strings.Contains(pinyinstring,"2")){
		return 2;
	}else if(strings.Contains(pinyinstring,"3")){
		return 3;
	}else if(strings.Contains(pinyinstring,"4")){
		return 4;
	}else if(strings.Contains(pinyinstring,"5")){
		return 5;
	}
	return 0;

}
func createToneMarks(numberedPinyin string)string{

	var words = strings.Split(numberedPinyin," ");

	 atones := []string{"ā","á","ǎ","à","a"}
	 etones := []string{"ē","é","ě","è","e"}
	 itones := []string{"ī","í","ǐ","ì","i"}
	 otones := []string{"ō","ó","ǒ","ò","o"}
	 utones := []string{"ū","ú","ǔ","ù","u"}
	 udottones:= []string{"ǖ","ǘ","ǚ","ǜ","ü"}

	// Look for u values replace u: with ü
	for pos, word := range words{
		for _, char := range word {
			if string(char) == ":" {
				word = strings.Replace(word, "u:", "ü", -1)
				words[pos] = word
			}
		}
	}

	//replace a or e with its tone because thats one of the easiest
	//A and e trump all other vowels and always take the tone mark.
	// There are no Mandarin syllables in Hanyu Pinyin that contain both a and e.
	for pos, word := range words{

		if(strings.Contains(word,"a") || strings.Contains(word,"e")){
				strs := "a"
			if(strings.Contains(word,"e")){
				strs = "e"
			}
			var toneval = detectTone(word)
			if(toneval != 0){
				
			if(strs =="a"){
				word = strings.Replace(word,strs,atones[toneval-1],-1)
			}else{
				word = strings.Replace(word,strs,etones[toneval-1],-1)
			}
			
			}
				word = strings.Replace(word,strconv.Itoa(toneval),"",-1)
				words[pos] = word;
	    }
	}
	
	//In the combination ou, o takes the mark. So lets replace it
	for pos,word :=range words{

		if ( !strings.Contains(word,"a") && !strings.Contains(word,"e") && strings.Contains(word,"ou") ) {
			var toneval = detectTone(word)
			if(toneval != 0){
				word = strings.Replace(word,"o",otones[toneval-1],-1)
				

			}
			word = strings.Replace(word,strconv.Itoa(toneval),"",-1)
			//fmt.Println("Found ou lets print out the fixed word!")
			
			words[pos] = word;
			//fmt.Println(words);
		}


	}

	//Last case, we have to assign the tone mark to the last vowel of the word
	//We will just loop through to find the last vowel and replace it.
		for pos, word := range words{
			positionToBeReplaced := 0;
			vowelToReplace := ""
			voweExist := false;
			var toneval = detectTone(word)
			if ( !strings.Contains(word,"a") && !strings.Contains(word,"e") && !strings.Contains(word,"ou") ) {


			for wordpos, char := range word {
				switch string(char){
				case "i":
					positionToBeReplaced = wordpos;
					vowelToReplace = string(char)
					voweExist = true;
				case "o":
					positionToBeReplaced = wordpos;
					vowelToReplace = string(char)
					voweExist = true;
				case "u":
					positionToBeReplaced = wordpos;
					vowelToReplace = string(char)
					voweExist = true;
				case "ü":
					positionToBeReplaced = wordpos;
					vowelToReplace = string(char)
					voweExist = true;		
				case "I":
					positionToBeReplaced = wordpos;
					vowelToReplace ="i"
					voweExist = true;
				case "O":
					positionToBeReplaced = wordpos;
					vowelToReplace ="o"
					voweExist = true;
				case "U":
					positionToBeReplaced = wordpos;
					vowelToReplace ="u"
					voweExist = true;
				}


				
			}

			if voweExist && toneval != 0{
				switch vowelToReplace{
					case "i":
						words[pos] = replaceAtIndex(words[pos],positionToBeReplaced,itones[toneval-1])

					case "o":
						words[pos] = replaceAtIndex(words[pos],positionToBeReplaced,otones[toneval-1])
						
					case "u":
						words[pos] = replaceAtIndex(words[pos],positionToBeReplaced,utones[toneval-1])
						
					case "ü":
						words[pos] = replaceAtIndex(words[pos],positionToBeReplaced,udottones[toneval-1])

				}
			}
			if(toneval != 0){
				words[pos] = strings.Replace(words[pos],strconv.Itoa(toneval),"",-1)

			}
			
		}
		}

	return strings.Join(words," ");
}


func readLine(path string) {
	fmt.Println("Begin read");
	var newdb []CEDICTDATA;
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	  scanner.Split(bufio.ScanLines) 
	
	for scanner.Scan() {
	 
	

	    var TraditionalCharacters []rune
	    var SimplifiedCharacters []rune
	    var Pinyin []rune
	    var Definition []rune

	
	    runes := []rune(scanner.Text())
		var spaces = 0;
		var leftbracket = 0;
		var rightbracket = 0;
		var slashCount = 0;
		iscomment := false;

		for _,character := range runes {
			var currentChar = string(character)
			if(spaces <1 && currentChar == "#"){
				iscomment = true;
				break
			}

			if spaces < 1 && currentChar  != " " {
				TraditionalCharacters = append(TraditionalCharacters,character)
			}

			if spaces > 0 && spaces < 2 && currentChar  != " " {
				SimplifiedCharacters = append(SimplifiedCharacters,character)
			}

			if(currentChar  == " ") {
				spaces++;
			}

			if( spaces > 1 && leftbracket > 0 && rightbracket < 1 && currentChar  != "[" && currentChar  != "]") {
				Pinyin = append(Pinyin,character)	
			}

			if(currentChar == "[") {
				leftbracket++;
			}
			if(currentChar == "]") {
				rightbracket++;
			}

			if(slashCount >0 && spaces > 0 && leftbracket > 0) {
				if(currentChar == "/") {
				    trune := rune(';')
				    Definition = append(Definition,trune)
				} else{
					Definition = append(Definition,character)
				}

			}

			if(currentChar == "/") {
				slashCount++;
			}


		}

		var hskstruct = CEDICTDATA{
			string(TraditionalCharacters),
			string(SimplifiedCharacters),
			string(Pinyin),
			createToneMarks(string(Pinyin)),
			string(Definition)}

			if(!iscomment){
				newdb = append(newdb, hskstruct)
			}


	}
	fmt.Println("Getting ready to write file")
	pagesJson, err := json.Marshal(newdb)
	if err != nil {
		fmt.Println("Error")
		 return
			 }
			
	ioutil.WriteFile("./cedict.json",pagesJson,0644)
  }


func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide file name.")
		fmt.Println("example ./cedicttojson cedict_ts.u8")
        return
    }
    filename := os.Args[1]
    fmt.Println("File content is:",filename)
	readLine(filename);

}

