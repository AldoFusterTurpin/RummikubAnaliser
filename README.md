### About
Simple program to analyse the score of [Rummikub](https://en.wikipedia.org/wiki/Rummikub) matches.

[Scoring rules](https://en.wikipedia.org/wiki/Rummikub#Scoring)

Score -432 is better than -508  and score -407 is better than -432.

From best to worst:
- -407
- -432
- -508

### Execute the program
```sh
go run *.go
```
or 
```sh
go build && ./RummikubAnaliser
```

### Notes
Date format: DD/MM/YYYY

The program assumes all the matches are played by the same number of players to calculate the total score of each player. Modify the program accordingly to your needs.

Cheers!