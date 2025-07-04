🎯 Number Guessing Game (Go)

This is a simple CLI game written in Go where the player has three chances to guess a randomly selected number between 1 and 10.

---

 🕹️ How It Works

- The program randomly picks a number between 1 and 10
- You get three tries to guess the correct number
- After each guess:
  - If your guess is too low, it will let you know
  - If your guess is too high, it will tell you
  - If you guess correctly, the game ends and congratulates you

---

 What I Learned

Through building this project, I learned how to:

- Take keyboard input in Go using `fmt.Scan`
- Work with Boolean logic and conditional statements
- Use Go packages like:
  - `math/rand` for generating random numbers
  - `time` to seed randomness properly

---

🚀 Why I Built This

This was one of my first Go projects and a great way to practice:
- Structuring a small CLI program
- Working with logic and control flow
- Building confidence with Go syntax

---

📂 Run It Yourself

To run the program:

1. Make sure you have Go installed
2. Clone this repo or download the `.go` file
3. Run it in your terminal:

```bash
go run NumbersGuessingGame.go
