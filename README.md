# MovieHub API

## use the command:
```
go run .
// or 
go build .
```
To get started
visit the link here
[moviehub](https://movieapi-gcve.onrender.com)

# Docs

### Routes:
`/series/*`
Add the name of any series to the end of the route to get a JSON response
Example: 
```
{
  "Title": "Frieren: Beyond Journey's End",
  "Year": "2023–2024",
  "Rated": "TV-14",
  "Released": "29 Sep 2023",
  "Runtime": "N/A",
  "Genre": "Animation, Adventure, Drama",
  "Director": "N/A",
  "Writer": "Tsukasa Abe, Kanehito Yamada",
  "Actors": "Atsumi Tanezaki, Kana Ichinose, Mallorie Rodak",
  "Plot": "An elf and her friends defeat a demon king in a great war. But the war is over, and the elf must search for a new way of life.",
  "Language": "Japanese, Portuguese, Spanish, English",
  "Country": "Japan",
  "Awards": "7 wins & 2 nominations",
  "Poster": "https://m.media-amazon.com/images/M/MV5BMjVjZGU5ZTktYTZiNC00N2Q1LThiZjMtMDVmZDljN2I3ZWIwXkEyXkFqcGdeQXVyMTUzMTg2ODkz._V1_SX300.jpg",
  "Ratings": [
    {
      "Source": "Internet Movie Database",
      "Value": "8.9/10"
    }
  ],
  "Metascore": "N/A",
  "imdbRating": "8.9",
  "imdbVotes": "20,108",
  "imdbID": "tt22248376",
  "Type": "series",
  "totalSeasons": "1",
  "Response": "True"
}
```

If the series doesn't exist Response field will be False.

`/movies/*`
Add the name of any movie to the end of the route to get a JSON response
Example: 
```
{
  "Title": "The Godfather",
  "Year": "1972",
  "Rated": "R",
  "Released": "24 Mar 1972",
  "Runtime": "175 min",
  "Genre": "Crime, Drama",
  "Director": "Francis Ford Coppola",
  "Writer": "Mario Puzo, Francis Ford Coppola",
  "Actors": "Marlon Brando, Al Pacino, James Caan",
  "Plot": "Don Vito Corleone, head of a mafia family, decides to hand over his empire to his youngest son, Michael. However, his decision unintentionally puts the lives of his loved ones in grave danger.",
  "Language": "English, Italian, Latin",
  "Country": "United States",
  "Awards": "Won 3 Oscars. 31 wins & 31 nominations total",
  "Poster": "https://m.media-amazon.com/images/M/MV5BM2MyNjYxNmUtYTAwNi00MTYxLWJmNWYtYzZlODY3ZTk3OTFlXkEyXkFqcGdeQXVyNzkwMjQ5NzM@._V1_SX300.jpg",
  "Ratings": [
    {
      "Source": "Internet Movie Database",
      "Value": "9.2/10"
    },
    {
      "Source": "Rotten Tomatoes",
      "Value": "97%"
    },
    {
      "Source": "Metacritic",
      "Value": "100/100"
    }
  ],
  "Metascore": "100",
  "imdbRating": "9.2",
  "imdbVotes": "2,031,306",
  "imdbID": "tt0068646",
  "Type": "movie",
  "totalSeasons": "",
  "Response": "True"
}
```
If the movie doesn't exist Response field will be False.

`/search/*`
You can pass in incomplete values of series and movies to the end of the route and get a JSON response showing a mix of both series and movies. 
Example: (/search/the last):
```{
  "Search": [
    {
      "Title": "Indiana Jones and the Last Crusade",
      "Year": "1989",
      "imdbID": "tt0097576",
      "Type": "movie",
      "Poster": "https://m.media-amazon.com/images/M/MV5BY2Q0ODg4ZmItNDZiYi00ZWY5LTg2NzctNmYwZjA5OThmNzE1XkEyXkFqcGdeQXVyMjM4MzQ4OTQ@._V1_SX300.jpg"
    },
    {
      "Title": "Star Wars: Episode VIII - The Last Jedi",
      "Year": "2017",
      "imdbID": "tt2527336",
      "Type": "movie",
      "Poster": "https://m.media-amazon.com/images/M/MV5BMjQ1MzcxNjg4N15BMl5BanBnXkFtZTgwNzgwMjY4MzI@._V1_SX300.jpg"
    },
    {
      "Title": "The Last of Us",
      "Year": "2023–",
      "imdbID": "tt3581920",
      "Type": "series",
      "Poster": "https://m.media-amazon.com/images/M/MV5BZGUzYTI3M2EtZmM0Yy00NGUyLWI4ODEtN2Q3ZGJlYzhhZjU3XkEyXkFqcGdeQXVyNTM0OTY1OQ@@._V1_SX300.jpg"
    },
    {
      "Title": "X-Men: The Last Stand",
      "Year": "2006",
      "imdbID": "tt0376994",
      "Type": "movie",
      "Poster": "https://m.media-amazon.com/images/M/MV5BNDBhNDJiMWEtOTg4Yi00NTYzLWEzOGMtMjNmNjAxNTBlMzY3XkEyXkFqcGdeQXVyNTIzOTk5ODM@._V1_SX300.jpg"
    },
    {
      "Title": "The Last Samurai",
      "Year": "2003",
      "imdbID": "tt0325710",
      "Type": "movie",
      "Poster": "https://m.media-amazon.com/images/M/MV5BMzkyNzQ1Mzc0NV5BMl5BanBnXkFtZTcwODg3MzUzMw@@._V1_SX300.jpg"
    },
    {
      "Title": "Avatar: The Last Airbender",
      "Year": "2005–2008",
      "imdbID": "tt0417299",
      "Type": "series",
      "Poster": "https://m.media-amazon.com/images/M/MV5BODc5YTBhMTItMjhkNi00ZTIxLWI0YjAtNTZmOTY0YjRlZGQ0XkEyXkFqcGdeQXVyODUwNjEzMzg@._V1_SX300.jpg"
    },
    {
      "Title": "The Last King of Scotland",
      "Year": "2006",
      "imdbID": "tt0455590",
      "Type": "movie",
      "Poster": "https://m.media-amazon.com/images/M/MV5BYzZkNjNhN2QtMThmNC00YjE0LTg0ZmMtMmU5MWE4Y2NjN2RiXkEyXkFqcGdeQXVyMTUzMDUzNTI3._V1_SX300.jpg"
    },
    {
      "Title": "The Last of the Mohicans",
      "Year": "1992",
      "imdbID": "tt0104691",
      "Type": "movie",
      "Poster": "https://m.media-amazon.com/images/M/MV5BZDNiYmRkNDYtOWU1NC00NmMxLWFkNmUtMGI5NTJjOTJmYTM5XkEyXkFqcGdeQXVyNzQ1ODk3MTQ@._V1_SX300.jpg"
    },
    {
      "Title": "The Last Duel",
      "Year": "2021",
      "imdbID": "tt4244994",
      "Type": "movie",
      "Poster": "https://m.media-amazon.com/images/M/MV5BZGExZTUzYWQtYWJjZi00OTI4LTk4OGYtNTA2YzcwMmNiZTMxXkEyXkFqcGdeQXVyMTEyMjM2NDc2._V1_SX300.jpg"
    },
    {
      "Title": "Puss in Boots: The Last Wish",
      "Year": "2022",
      "imdbID": "tt3915174",
      "Type": "movie",
      "Poster": "https://m.media-amazon.com/images/M/MV5BNjMyMDBjMGUtNDUzZi00N2MwLTg1MjItZTk2MDE1OTZmNTYxXkEyXkFqcGdeQXVyMTQ5NjA0NDM0._V1_SX300.jpg"
    }
  ],
  "totalResults": "5554",
  "Response": "True"
}```

If you notice we have both movies and series mixed in, also an incomplete title passed `the last`

`/recommend`
This route takes in a query param called genre. and then returns a random response with that genre included.
For example with this `/recommend?genre=comedy` you get this JSON response 
```
{
  "Title": "Empire Records",
  "Year": "1995",
  "Rated": "PG-13",
  "Released": "20 Oct 1995",
  "Runtime": "90 min",
  "Genre": "Comedy, Drama, Music",
  "Director": "Allan Moyle",
  "Writer": "Carol Heikkinen",
  "Actors": "Anthony LaPaglia, Debi Mazar, Maxwell Caulfield",
  "Plot": "Twenty-four hours in the lives of the young employees at Empire Records when they all grow up and become young adults thanks to each other and the manager. They all face the store joining a chain store with strict rules.",
  "Language": "English",
  "Country": "United States",
  "Awards": "1 win",
  "Poster": "https://m.media-amazon.com/images/M/MV5BMjExMmQ1NmItOWEyOC00NzZjLWE0NWItNmE1NmNkY2NmODk5XkEyXkFqcGdeQXVyNDk3NzU2MTQ@._V1_SX300.jpg",
  "Ratings": [
    {
      "Source": "Internet Movie Database",
      "Value": "6.7/10"
    },
    {
      "Source": "Rotten Tomatoes",
      "Value": "31%"
    },
    {
      "Source": "Metacritic",
      "Value": "36/100"
    }
  ],
  "Metascore": "36",
  "imdbRating": "6.7",
  "imdbVotes": "62,315",
  "imdbID": "tt0112950",
  "Type": "movie",
  "totalSeasons": "",
  "Response": "True"
}
```
Note you can also pass in to 3 genres like this `/recommend?genre=action,animation`