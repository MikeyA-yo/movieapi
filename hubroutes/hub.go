package hubroutes

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

var Num = rand.Intn(100)

/*
 Implementation Details:
  takes to arguments:
  1.) the global omdb API Url with your api key
  2.) Name of search query
  ensures to return only type series (omdbapi.com)
  combines them into a url and fetch with http
  then return the response body
*/

// Function to identify a series based on it's name, returns an array of bytes, takes a url and name as parameters
func GetSeries(url, name string) []byte {
	var errR []byte
	combineUrl := fmt.Sprintf("%v&t=%v&type=series", url, name)
	res, err := http.Get(combineUrl)
	if err != nil {
		fmt.Print("Error\n")
		fmt.Println(err.Error())
		return errR
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

/*
 Implementation Details:
  takes to arguments:
  1.) the global omdb API Url with your api key
  2.) Name of search query
  ensures to return only type movie (omdbapi.com)
  combines them into a url and fetch with http
  then return the response body
*/

// Function to identify a movie based on it's name, returns an array of bytes, takes a url and name as parameters
func GetMovies(url, name string) []byte {
	combineUrl := fmt.Sprintf("%v&t=%v&type=movie", url, name)
	res, err := http.Get(combineUrl)
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

/*
 Implementation Details:
  takes to arguments:
  1.) the global omdb API Url with your api key
  2.) Name of search query
  combines them into a url and fetch with http
  then return the response body
*/

// function to Search for movies and series generally
func GetSearch(url, name string) []byte {
	combineUrl := fmt.Sprintf("%v&s=%v", url, name)
	res, err := http.Get(combineUrl)
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

/*
 Implementation Details:
 Create a struct type based of the json response for search
 check out omdbapi.com for search results
*/

type serials struct {
	Search []struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	} `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response     string `json:"Response"`
}

/*
 Implementation Details:
 use the GetSearch funtion to get the response body
 then store the count of possible page numbers from TotalResults property (read omdbapi.com docs)
 generate a random page number based  on the max possible page
 use that random number to fetch a random page and return the response body
*/

// Function to get a search page pseudo randomly
func GetSearchRand(url, name string) []byte {
	var dataJson serials
	randData := GetSearch(url, name)
	json.Unmarshal(randData, &dataJson)
	count, err := strconv.Atoi(dataJson.TotalResults)
	if err != nil {
		fmt.Println("Error")
		return randData
	}
	maxPageNumber := math.Round(float64(count / 10))
	num := rand.Intn(int(maxPageNumber)) + 1
	if num > 100 {
		num = 100
	}
	combineUrl := fmt.Sprintf("%v&s=%v&page=%v", url, name, num)
	res, err := http.Get(combineUrl)
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

/*
 Implementation Details:
 Get large list's of word's basically from an api and return it as a slice of strings
*/

// function to simplify process
func Words() []string {
	result := []string{}
	res, err := http.Get("https://random-word-api.herokuapp.com/all")
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &result)
	return result
}

// store all the words in this variable
var WordsArray = Words()

/*
 Implementation Details:
  use a random number to get an indiviual word from the slice
*/
//https://wordgenerator-api.herokuapp.com/api/v1/resources/words?lang=EN&amount=5
// function For a random word 178186
func GetWord() string {
	var WordNum = rand.Intn(len(WordsArray))
	result := WordsArray
	return result[WordNum]
}

/*
 Implementation Details:
   Below i'd be working on genre not much to be said on implementation details yet
*/

// Type for working with genre
type searchResult struct {
	Title string `json:"Title"`
	Type  string `json:"Type"`
}

// actually get a title, and type to search based on GetWord() function
func GetTitle() []searchResult {
	var title serials
	randomWord := GetWord()
	URL := fmt.Sprintf("https://movieapi-gcve.onrender.com/search/%v", randomWord)
	res, err := http.Get(URL)
	if err != nil {
		fmt.Print("Error\n")
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &title)
	if title.Response == "False" {
		return GetTitle()
	}
	arrLength := len(title.Search)
	titlesRes := []searchResult{}
	//titles := []string{}
	for i := 0; i < arrLength; i++ {
		results := searchResult{
			Title: title.Search[i].Title,
			Type:  title.Search[i].Type,
		}
		titlesRes = append(titlesRes, results)
		//titles = append(titles, title.Search[i].Title)
	}
	return titlesRes
}

type genreLike struct {
	Title string `json:"Title"`
	Genre string `json:"Genre"`
}

// huge slice of movies and genre for fallback cases (todo)
var breakingBad genreLike = genreLike{
	Title: "Breaking Bad",
	Genre: "Crime, Drama, Thriller",
}

var gameOfThrones genreLike = genreLike{
	Title: "Game of Thrones",
	Genre: "Action, Adventure, Drama",
}
var chernobyl genreLike = genreLike{
	Title: "Chernobyl",
	Genre: "Drama, History, Thriller",
}

var frieren genreLike = genreLike{
	Title: "Frieren: Beyond Journey's End",
	Genre: "Animation, Adventure, Drama",
}

var gintama genreLike = genreLike{
	Title: "Gintama",
	Genre: "Animation, Action, Comedy",
}

var attackOnTitan genreLike = genreLike{
	Title: "Attack on Titan",
	Genre: "Animation, Action, Adventure",
}

var hxh genreLike = genreLike{
	Title: "Hunter x Hunter",
	Genre: "Animation, Action, Adventure",
}

var demonSlayer genreLike = genreLike{
	Title: "Demon Slayer: Kimetsu no Yaiba",
	Genre: "Animation, Action, Adventure",
}

var rickNMorty genreLike = genreLike{
	Title: "Rick and Morty",
	Genre: "Animation, Adventure, Comedy",
}

var theOffice genreLike = genreLike{
	Title: "The Office",
	Genre: "Comedy",
}

var theWalkingDead genreLike = genreLike{
	Title: "The Walking Dead",
	Genre: "Drama, Horror, Thriller",
}

var theLastOfUs genreLike = genreLike{
	Title: "The Last of Us",
	Genre: "Action, Adventure, Drama",
}

var theUmbrellaAcademy genreLike = genreLike{
	Title: "The Umbrella Academy",
	Genre: "Action, Adventure, Comedy",
}

var supacell genreLike = genreLike{
	Title: "Supacell",
	Genre: "Action, Adventure, Drama",
}

var blackButler genreLike = genreLike{
	Title: "Black Butler",
	Genre: "Animation, Action, Comedy",
}

var blackish genreLike = genreLike{
	Title: "Black-ish",
	Genre: "Comedy",
}

var mixedish genreLike = genreLike{
	Title: "Mixed-ish",
	Genre: "Comedy, Family",
}

var grownish genreLike = genreLike{
	Title: "Grown-ish",
	Genre: "Comedy, Drama",
}

var theWitcher genreLike = genreLike{
	Title: "The Witcher",
	Genre: "Action, Adventure, Drama",
}

var dragonBallZ genreLike = genreLike{
	Title: "Dragon Ball Z",
	Genre: "Animation, Action, Adventure",
}

var dragonBallSuper genreLike = genreLike{
	Title: "Dragon Ball Super",
	Genre: "Animation, Action, Adventure",
}

var naruto genreLike = genreLike{
	Title: "Naruto",
	Genre: "Animation, Action, Adventure",
}

var bleach genreLike = genreLike{
	Title: "Bleach",
	Genre: "Animation, Action, Adventure",
}

var onePiece genreLike = genreLike{
	Title: "One Piece",
	Genre: "Animation, Action, Adventure",
}

var venom genreLike = genreLike{
	Title: "Venom",
	Genre: "Action, Adventure, Sci-Fi",
}

var letItShine genreLike = genreLike{
	Title: "Let It Shine",
	Genre: "Music, Reality-TV",
}

var suits genreLike = genreLike{
	Title: "Suits",
	Genre: "Comedy, Drama",
}

var loveNMon genreLike = genreLike{
	Title: "Love and Monsters",
	Genre: "Action, Adventure, Comedy",
}

var youngSheldon genreLike = genreLike{
	Title: "Young Sheldon",
	Genre: "Comedy",
}

var raisingDion genreLike = genreLike{
	Title: "Raising Dion",
	Genre: "Drama, Sci-Fi",
}

var community genreLike = genreLike{
	Title: "Community",
	Genre: "Comedy",
}

var hannibal genreLike = genreLike{
	Title: "Hannibal",
	Genre: "Crime, Drama, Horror",
}

var homeland = genreLike{
	Title: "Homeland",
	Genre: "Crime, Drama, Mystery",
}

var topChef = genreLike{
	Title: "Top Chef",
	Genre: "Game-Show, Reality-TV",
}

var theGoodFight = genreLike{
	Title: "The Good Fight",
	Genre: "Crime, Drama",
}

var blackMirror = genreLike{
	Title: "Black Mirror",
	Genre: "Crime, Drama, Mystery",
}

var iMayDestroyYou = genreLike{
	Title: "I May Destroy You",
	Genre: "Drama",
}

var stElsewhere = genreLike{
	Title: "St. Elsewhere",
	Genre: "Comedy, Drama",
}

var daria = genreLike{
	Title: "Daria",
	Genre: "Animation, Comedy, Drama",
}

var theCosbyShow = genreLike{
	Title: "The Cosby Show",
	Genre: "Comedy, Family, Romance",
}
var misterRoger = genreLike{
	Title: "Mister Rogers' Neighborhood",
	Genre: "Family, Fantasy, Music",
}

var genHospital = genreLike{
	Title: "General Hospital",
	Genre: "Crime, Drama, Romance",
}

var happyDays = genreLike{
	Title: "Happy Days",
	Genre: "Comedy, Family, Music",
}

var girls = genreLike{
	Title: "Girls",
	Genre: "Comedy, Drama",
}

var columbo = genreLike{
	Title: "Columbo",
	Genre: "Crime, Drama, Mystery",
}

var atlanta = genreLike{
	Title: "Atlanta",
	Genre: "Comedy, Drama, Music",
}

var strangerThings = genreLike{
	Title: "Stranger Things",
	Genre: "Drama, Fantasy, Horror",
}

var fleabag = genreLike{
	Title: "Fleabag",
	Genre: "Comedy, Drama",
}

var thirtySomething = genreLike{
	Title: "Thirtysomething",
	Genre: "Drama, Romance",
}

var scandal = genreLike{
	Title: "Scandal",
	Genre: "Drama, Thriller",
}

var theMuppetShow = genreLike{
	Title: "The Muppet Show",
	Genre: "Comedy, Family, Musical",
}

var dallas = genreLike{
	Title: "Dallas",
	Genre: "Drama, Romance",
}

var theFreshBel = genreLike{
	Title: "The Fresh Prince of Bel-Air",
	Genre: "Comedy",
}

var taxi = genreLike{
	Title: "Taxi",
	Genre: "Comedy",
}

var deadwood = genreLike{
	Title: "Deadwood",
	Genre: "Crime, Drama, History",
}

var nypdBlue = genreLike{
	Title: "NYPD Blue",
	Genre: "Crime, Drama, Mystery",
}

var theWonderYears = genreLike{
	Title: "The Wonder Years",
	Genre: "Comedy, Drama, Family",
}

var livingSinlge = genreLike{
	Title: "Living Single",
	Genre: "Comedy",
}

var ojSimpsonStory = genreLike{
	Title: "Inside Look: The People v. O.J. Simpson - American Crime Story",
	Genre: "Documentary, Short",
}

var roseanne = genreLike{
	Title: "Roseanne",
	Genre: "Comedy, Drama",
}

var greysAnatomy = genreLike{
	Title: "Grey's Anatomy",
	Genre: "Drama, Romance",
}

var draggggin = genreLike{
	Title: "RuPaul's Drag Race",
	Genre: "Game-Show, Reality-TV",
}

var bobNewhart = genreLike{
	Title: "The Bob Newhart Show",
	Genre: "Comedy",
}

var freaksNGeeks = genreLike{
	Title: "Freaks and Geeks",
	Genre: "Comedy, Drama",
}

var jeffersons = genreLike{
	Title: "The Jeffersons",
	Genre: "Comedy",
}

var angelsInAmerica = genreLike{
	Title: "Angels in America",
	Genre: "Drama, Fantasy, Romance",
}

var theComeback = genreLike{
	Title: "The Comeback",
	Genre: "Comedy",
}

var orangeBlack = genreLike{
	Title: "Orange Is the New Black",
	Genre: "Comedy, Crime, Drama",
}

var inLivingColor = genreLike{
	Title: "In Living Color",
	Genre: "Comedy, Music",
}

var southPark = genreLike{
	Title: "South Park",
	Genre: "Animation, Comedy",
}

var theGoodPlace = genreLike{
	Title: "The Good Place",
	Genre: "Comedy, Drama, Fantasy",
}

var chap = genreLike{
	Title: "Chappelle's Show",
	Genre: "Comedy, Music",
}

var lawnorderSpecial = genreLike{
	Title: "Law & Order: Special Victims Unit",
	Genre: "Crime, Drama, Mystery",
}

var bojack = genreLike{
	Title: "BoJack",
	Genre: "Animation, Comedy, Drama",
}

var gilmoreGirls = genreLike{
	Title: "Gilmore Girls",
	Genre: "Comedy, Drama",
}

var sixFeetUnder = genreLike{
	Title: "Six Feet Under",
	Genre: "Comedy, Drama",
}

var arrestedDev = genreLike{
	Title: "Arrested Development",
	Genre: "Comedy",
}

var soCalled = genreLike{
	Title: "My So-Called Life",
	Genre: "Comedy, Drama, Romance",
}

var watchmen = genreLike{
	Title: "Watchmen",
	Genre: "Crime, Drama, Mystery",
}

var theShield = genreLike{
	Title: "The Shield",
	Genre: "Crime, Drama, Thriller",
}

var fridayNightLights = genreLike{
	Title: "Friday Night Lights",
	Genre: "Drama, Sport",
}

var theLeftOvers = genreLike{
	Title: "The Leftovers",
	Genre: "Drama, Fantasy, Mystery",
}

var starTrek = genreLike{
	Title: "Star Trek: The Next Generation",
	Genre: "Action, Adventure, Drama",
}

var theAmericans = genreLike{
	Title: "The Americans",
	Genre: "Crime, Drama, Mystery",
}

var theRealWorld = genreLike{
	Title: "The Real World",
	Genre: "Drama, Reality-TV",
}

var buffy = genreLike{
	Title: "Buffy the Vampire Slayer",
	Genre: "Action, Adventure, Drama",
}

var theXFiles = genreLike{
	Title: "The X-Files",
	Genre: "Crime, Drama, Mystery",
}

var enlightened = genreLike{
	Title: "Enlightened",
	Genre: "Comedy, Drama",
}

var curbEnthusiasm = genreLike{
	Title: "Curb Your Enthusiasm",
	Genre: "Comedy",
}

var ER = genreLike{
	Title: "ER",
	Genre: "Drama, Romance",
}

var lost = genreLike{
	Title: "Lost",
	Genre: "Adventure, Drama, Fantasy",
}

var survivor = genreLike{
	Title: "Survivor",
	Genre: "Adventure, Game-Show, Reality-TV",
}

var hillStreetBlues = genreLike{
	Title: "Hill Street Blues",
	Genre: "Crime, Drama, Mystery",
}

var friends = genreLike{
	Title: "Friends",
	Genre: "Comedy, Romance",
}

var theCivilWar = genreLike{
	Title: "The Civil War",
	Genre: "Documentary, History, War",
}

var twinPeaks = genreLike{
	Title: "Twin Peaks",
	Genre: "Crime, Drama, Mystery",
}

var veep = genreLike{
	Title: "Veep",
	Genre: "Comedy",
}

var westWing = genreLike{
	Title: "The West Wing",
	Genre: "Drama",
}

var theCarolBunnShow = genreLike{
	Title: "The Carol Burnett Show",
	Genre: "Comedy, Family",
}

var thirtyRock = genreLike{
	Title: "30 Rock",
	Genre: "Comedy",
}

var sixtyMins = genreLike{
	Title: "60 Minutes",
	Genre: "News",
}

var theGoldenGirls = genreLike{
	Title: "The Golden Girls",
	Genre: "Comedy, Drama",
}

var theOphrahShow = genreLike{
	Title: "The Oprah Winfrey Show",
	Genre: "News, Talk-Show",
}

var allinfamily = genreLike{
	Title: "All in the Family",
	Genre: "Comedy, Drama",
}

var twZone = genreLike{
	Title: "The Twilight Zone",
	Genre: "Drama, Fantasy, Horror",
}

var succession = genreLike{
	Title: "Succession",
	Genre: "Comedy, Drama",
}

var sesameStreet = genreLike{
	Title: "Sesame Street",
	Genre: "Animation, Adventure, Family",
}

var cheers = genreLike{
	Title: "Cheers",
	Genre: "Comedy, Drama",
}

var roots = genreLike{
	Title: "Roots",
	Genre: "Biography, Drama, History",
}

var seinfield = genreLike{
	Title: "Seinfeld",
	Genre: "Comedy",
}

var theWire = genreLike{
	Title: "The Wire",
	Genre: "Crime, Drama, Thriller",
}

var sexNCity = genreLike{
	Title: "Sex and the City",
	Genre: "Comedy, Drama, Romance",
}

var theSimpsons = genreLike{
	Title: "The Simpsons",
	Genre: "Animation, Comedy",
}

var theSopranos = genreLike{
	Title: "The Sopranos",
	Genre: "Crime, Drama",
}

var madMen = genreLike{
	Title: "Mad Men",
	Genre: "Drama",
}

var iLoveLucy = genreLike{
	Title: "I Love Lucy",
	Genre: "Comedy",
}

var bluey = genreLike{
	Title: "Bluey",
	Genre: "Animation, Family",
}
var recommendations = []genreLike{
	bluey,
	iLoveLucy,
	madMen,
	theSopranos,
	theSimpsons,
	sexNCity,
	theWire,
	seinfield,
	roots,
	cheers,
	sesameStreet,
	succession,
	twZone,
	allinfamily,
	theOphrahShow,
	theGoldenGirls,
	sixtyMins,
	thirtyRock,
	theCarolBunnShow,
	westWing,
	veep,
	twinPeaks,
	theCivilWar,
	friends,
	hillStreetBlues,
	survivor,
	lost,
	ER,
	curbEnthusiasm,
	enlightened,
	theXFiles,
	buffy,
	theRealWorld,
	theAmericans,
	starTrek,
	theLeftOvers,
	fridayNightLights,
	theShield,
	watchmen,
	soCalled,
	arrestedDev,
	sixFeetUnder,
	gilmoreGirls,
	bojack,
	lawnorderSpecial,
	chap,
	theGoodPlace,
	southPark,
	inLivingColor,
	orangeBlack,
	theComeback,
	angelsInAmerica,
	jeffersons,
	freaksNGeeks,
	bobNewhart,
	draggggin,
	greysAnatomy,
	roseanne,
	ojSimpsonStory,
	livingSinlge,
	theWonderYears,
	nypdBlue,
	deadwood,
	taxi,
	theFreshBel,
	dallas,
	theMuppetShow,
	scandal,
	thirtySomething,
	fleabag,
	strangerThings,
	atlanta,
	columbo,
	girls,
	happyDays,
	genHospital,
	misterRoger,
	theCosbyShow,
	daria,
	stElsewhere,
	iMayDestroyYou,
	blackMirror,
	theGoodFight,
	topChef,
	homeland,
	hannibal,
	community,
	raisingDion,
	youngSheldon,
	loveNMon,
	suits,
	letItShine,
	venom,
	onePiece,
	bleach,
	naruto,
	dragonBallSuper,
	dragonBallZ,
	theWitcher,
	grownish,
	mixedish,
	breakingBad,
	gameOfThrones,
	gintama,
	frieren,
	chernobyl,
	attackOnTitan,
	hxh,
	theLastOfUs,
	theOffice,
	theWalkingDead,
	demonSlayer,
	rickNMorty,
	theUmbrellaAcademy,
	supacell,
	blackButler,
	blackish}

func fallBack(genre string) serialP {
	var searchResult serials
	var jsonData serialP
	n := rand.Intn(len(recommendations))
	var movieName string
	movieGenres := strings.Split(strings.ToLower(recommendations[n].Genre), ", ")
	if strings.Contains(genre, ",") {
		if ContainsSlice(movieGenres, strings.Split(genre, ",")) {
			movieName = recommendations[n].Title
		} else {
			return fallBack(genre)
		}
	} else {
		if slices.Contains(movieGenres, strings.ToLower(genre)) {
			movieName = recommendations[n].Title
		} else {
			return fallBack(genre)
		}

	}
	fetchUrl := fmt.Sprintf("https://movieapi-gcve.onrender.com/search/%v", movieName)
	res, err := http.Get(fetchUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &searchResult)
	if searchResult.Search[0].Type == "movie" {
		fetchUrl := fmt.Sprintf("https://movieapi-gcve.onrender.com/movies/%v", movieName)
		res, err := http.Get(fetchUrl)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &jsonData)
	} else if searchResult.Search[0].Type == "series" {
		fetchUrl := fmt.Sprintf("https://movieapi-gcve.onrender.com/series/%v", movieName)
		res, err := http.Get(fetchUrl)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &jsonData)
	}
	return jsonData
}

// func GetRandomRec() genreLike {
// 	n := rand.Intn(len(recommendations))
// 	return recommendations[n]
// }

// function to check if a slice is in a slice
func ContainsSlice(a, b []string) bool {
	for _, v := range b {
		if !slices.Contains(a, v) {
			return false
		}
	}
	return true
}

type serialP struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore    string `json:"Metascore"`
	ImdbRating   string `json:"imdbRating"`
	ImdbVotes    string `json:"imdbVotes"`
	ImdbID       string `json:"imdbID"`
	Type         string `json:"Type"`
	TotalSeasons string `json:"totalSeasons"`
	Response     string `json:"Response"`
}

// function to find and keep genres
func GetDetailedRecommendation(genre string) serialP {
	var jsonData serialP
	URL := "https://movieapi-gcve.onrender.com/"
	lists := GetTitle()
	length := len(lists)
	genreLower := strings.ToLower(genre)
	for i := 0; i < length; i++ {
		name := fmt.Sprintf("%v", lists[i].Title)
		fmt.Println(lists)
		var genreData genreLike
		if lists[i].Type == "movie" {
			urlContent := fmt.Sprintf("%vmovies/%v", URL, name)
			res, e := http.Get(urlContent)
			if e != nil {
				fmt.Println(e)
			}
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			json.Unmarshal(body, &genreData)

			genreSlice := strings.Split(strings.ToLower(genreData.Genre), ", ")
			if strings.Contains(genreLower, ",") {

				genLowSlice := strings.Split(genreLower, ",")
				if ContainsSlice(genreSlice, genLowSlice) {
					json.Unmarshal(body, &jsonData)
					break
				} else {
					jsonData = fallBack(genre)
					if jsonData.Response == "True" {
						break
					}
				}
			} else if slices.Contains(genreSlice, genreLower) {
				json.Unmarshal(body, &jsonData)
				break
			} else {
				jsonData = fallBack(genre)
				if jsonData.Response == "True" {
					break
				}
			}
		} else {
			urlContent := fmt.Sprintf("%vseries/%v", URL, name)
			res, e := http.Get(urlContent)
			if e != nil {
				fmt.Println(e)
			}
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			json.Unmarshal(body, &genreData)
			genreSlice := strings.Split(genreData.Genre, ", ")
			if strings.Contains(genreLower, ",") {
				genLowSlice := strings.Split(genreLower, ",")
				if ContainsSlice(genreSlice, genLowSlice) {
					json.Unmarshal(body, &jsonData)
					break
				} else {
					jsonData = fallBack(genre)
					if jsonData.Response == "True" {
						break
					}
					// res, err := http.Get("https://movieapi-gcve.onrender.com/series/frieren")
					// if err != nil {
					// 	fmt.Println(err)
					// }
					// defer res.Body.Close()
					// body, _ := io.ReadAll(res.Body)
					// json.Unmarshal(body, &jsonData)
				}
			} else if slices.Contains(genreSlice, genreLower) {
				json.Unmarshal(body, &jsonData)
				break
			} else {
				jsonData = fallBack(genre)
				if jsonData.Response == "True" {
					break
				}
				// res, err := http.Get("https://movieapi-gcve.onrender.com/series/frieren")
				// if err != nil {
				// 	fmt.Println(err)
				// }
				// defer res.Body.Close()
				// body, _ := io.ReadAll(res.Body)
				// json.Unmarshal(body, &jsonData)
			}
		}
	}
	return jsonData
}
