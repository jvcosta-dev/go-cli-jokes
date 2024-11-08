# GO CLI Jokes

This Go application is a command-line tool that fetches jokes based on user-specified parameters. It utilizes flags to allow customization of the jokes fetched
The jokes are retrieved from an jokes-api and displayed in the console, showing both the setup and the punchline for each joke fetched.

- type (--type): Specifies the joke category.
- lang (--lang): Defines the language of the jokes.
- blacklist (--blacklist): Filters out jokes containing certain topics.
- range (--range): Sets the range of joke IDs to retrieve.
- amount (--amount): Determines the number of jokes to fetch.

## Table of Contents

- [Project Info](#project-info)
- [Running the Application](#running-the-application)
- [Examples](#examples)
  - [Random Joke](#random-joke)
  - [Multiple Jokes](#get-rate-from-currency)
  - [Joke With Different Arguments](#joke-with-different-arguments)

## Project Info

- **GO**: 1.22.2

## Running the Application

```bash
go build -o jokes cmd/main.go
jokes
```

## Examples

### Random Joke

```bash
./jokes
```

- **Output Example**

```bash
Joke #1
What do you call a group of 8 Hobbits?            
A Hobbyte.
```

### Multiple Jokes

```bash
./jokes -amount 3
```

- **Output Example**

```bash
Joke #1
What is the difference between the Constitutions of the USA and the USSR? Don't both of them guarantee freedom of speech?
Yes, but the Constitution of the USA also guarantees freedom after the speech.

Joke #2
Why did the functional programmer get thrown out of school?
Because he refused to take classes.               

Joke #3
What do you tell a woman with 2 black eyes?       
Nothing. You already told her twice.  
```

### Joke With Different Arguments

```bash
./jokes -type Misc -blacklist religious -lang fr  -amount 3
```

- **Output Example**

```bash
Joke #1
Qu'est-ce que deux fous avec des mitraillettes?   
Une fousillade.                                   

Joke #2
Qui court plus vite qu'Usain Bolt?                
Un Somalien avec un ticket restaurant             

Joke #3
Qu'est-ce qui se lève quand tu t'en sers, qui s'abaisse quand tu as fini de t'en servir et qui goutte après usage?
Un parapluie
```
