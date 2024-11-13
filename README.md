# Go Quiz Game

This is a command-line quiz game written in Go. The game reads questions from a CSV file, asks each question to the player, and tracks their score based on correct answers. The game includes a timer that limits the duration of the quiz, and an option to shuffle the questions for a fresh experience each time.

## How to Play
- Clone or download this repository.
- Build the Go application:

```bash
go build -o quiz
```
- Run the quiz game from the command line:

```bash
./quiz
```

By default, the quiz will read questions from problems.csv, and you'll have 30 seconds to answer as many questions as possible.

## Features

1. CSV-based Questions: The game reads questions and answers from a CSV file. Each row in the CSV file should contain a question in the first column and the answer in the second column, e.g.:

```csv
What is the capital of France?,Paris
5+7,12
```

2. **Timed Quiz**: The quiz is time-limited. You can set a custom time limit using the -timer flag.
3. **Shuffle Option**: Randomly shuffle the order of questions with the -shuffle flag.

## Flags
The game has a few command-line flags that allow you to customize the experience:

**-filename**: Specify a custom CSV file with questions. If not provided, it defaults to problems.csv.

```bash
./quiz -filename=science.csv
```

**-timer**: Set a custom time limit (in seconds) for the quiz. The default is 30 seconds.

```bash
./quiz -timer=60
```

**-shuffle**: Shuffle the order of questions. This flag does not require any additional input; it just needs to be added to the command.

```bash
./quiz -shuffle
```

## Example Usage
Here are a few ways to run the quiz with different options:

### Basic Quiz with Default Settings:

```bash
./quiz
```
### Using a Custom CSV File:

```bash
./quiz -filename=science.csv
```
### Setting a Custom Time Limit:

```bash
./quiz -timer=45
```
### Shuffling the Questions:

```bash
./quiz -shuffle
```


### Combining Flags:

```bash
./quiz -filename=history.csv -timer=60 -shuffle=true
```
## Scoring
At the end of the quiz, the game displays your score by showing the number of correct answers out of the total questions attempted.

Example Output:
```plaintext
Press Enter to start the quiz...

Question 1: What is the capital of France? = Paris
Correct!

Question 2: 5 + 7 = 13
Incorrect. The correct answer is 12.

...

Time's up!
You scored 5 out of 10.
```

Enjoy testing your knowledge with this Go quiz game!
