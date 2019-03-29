# INFDAT02-1

Voor het vak data science moesten wij de verschillend algorithmes in de praktijk toepassen.

## Folder structure

In folders kan je helpers.go file tegen komen deze ondersteund functies die hergebruikt worden of die er zijn ter ondersteuning.

In de map algortihms vind je de algorithmes: Cosine, Eaclidean & Pearson.

In de map assets zit de structuur om de resultaten netjes naar de commandline te printen.


### Installation

1. Ga naar [golang](https://golang.org/doc/install) - en download go door op de Download Go button te klikken.
2. Kies vervolgens je gewenste besturing systeem en volg de installatie
3. Eenmaal geinstalleerd kan je met de gewenste IDEA de uitgepakte zip folder openen.
4. Als je de folder geopend hebt moet je jouw geinstalleerde GO SDK koppelen aan het project.   
5. Open dan een shell in de folder direcotory en run:

```
go get ./...
```

### Run exercises
Run het onderstaande commands om de Exercise uit te voeren
* Exercise 1: User-Item
 ```
 go build userItem.go
 ```
* Exercise 2: Item-Item
```
go build itemItem.go
```
* Exercise 3: Apriori
 ```
 go build apriori.go
 ```

## Running the tests

Explain how to run the automated tests for this system

### Awnser questions

Exercise 1:
1. hoe zorg je er voor dat het bepalen van uiteindelijke advies (niet het inlezen) snel gaat
	Door het gebruik van:
	- Hasmap met string values voor directe benadering
	- _Pointers_ (reference naar geheugen)
	- Hergebruik van _data_ in een loop zie bijvoorbeeld het splitten van unieke & equal items in _assets.Data_ in de _findUsersWithMoreUniqueRatings_ method
	- Inlezen groupLens via Bufferd 1/0 voor het snelle lezen van data in bytes. 

2. Laden van de movieset data file u.data bevat de volgende gegevens:
	 * User id 
	 * Item id 
	 * Rating 
	 * Timestamp is verwijderd.

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Built with

* [GoLang](https://golang.org/doc/install) - The programming language

## Author

* **Raymon** - *Initial work* - [Raymonr](https://github.com/Raymonr)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Raymon
* Hogeschool Rotterdam for providing the exercise.
