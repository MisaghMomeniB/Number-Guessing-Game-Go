# 🎯 **Guess the Number Game** 🎲  
Welcome to the **Guess the Number Game**! This is a fun and interactive console-based game where you can test your luck and guessing skills. Choose a difficulty level, guess the number, and beat the challenge! 🚀

---

## 🧠 **How It Works**  
1. **Choose Your Difficulty**  
   - Easy: Guess a number between `1-50`.  
   - Medium: Guess a number between `1-100`.  
   - Hard: Guess a number between `1-200`.  

2. **Make Your Guess**  
   - Enter your number to see if it's too high, too low, or spot-on!  
   - You get up to **10 attempts** to guess the correct number.  

3. **Time & History**  
   - Track how long it took you to guess the number.  
   - See all your guesses at the end of the game.  

4. **Exit Anytime**  
   - Type `exit` to gracefully leave the game.  

---

## 🚀 **Features**  
- 🧩 **Three Difficulty Levels**: Choose Easy, Medium, or Hard for different ranges of numbers.  
- ⏱ **Timer**: Measures the time taken to guess correctly.  
- 📜 **Guess Tracking**: Keeps a record of all your guesses for post-game analysis.  
- 🔄 **Error Handling**: Robust handling of invalid inputs (e.g., letters, out-of-range numbers).  
- 🎮 **Interactive Gameplay**: Smooth, user-friendly experience.  

---

## 🛠 **Setup & Installation**  

### Prerequisites  
- Install [Go](https://go.dev/) (version 1.16 or later).  

### Clone the Repository  
```bash
git clone https://github.com/yourusername/guess-the-number-game.git
cd guess-the-number-game
```

### Run the Game  
```bash
go run main.go
```

---

## 📝 **Code Overview**  

### 1️⃣ **Main Components**  
- **Difficulty Selection**  
  The game prompts the user to choose one of three difficulty levels.  

- **Random Number Generation**  
  Generates a random number within the range based on the chosen difficulty.  

- **Guess Validation**  
  Ensures all guesses are valid integers and within acceptable limits.  

- **Game Loop**  
  Iterates through guesses until the user either wins or runs out of attempts.  

### 2️⃣ **Functions**  
- `selectDifficulty()`  
  Prompts the user for a difficulty level and validates their input.  

- `generateRandomNumber(difficulty int)`  
  Generates a random number within a range based on the chosen difficulty.  

- `getUserGuess()`  
  Prompts the user for a guess, validates the input, and handles exit commands.  

- `clearBuffer()`  
  Clears any residual input from the buffer to ensure smooth gameplay.  

---

## 📷 **Preview**  

```plaintext
Select Difficulty Level: 1 (Easy), 2 (Medium), 3 (Hard)
Enter your choice: 2
Enter your guess (or type 'exit' to quit): 50
Your guess is too low. Try again!
Enter your guess (or type 'exit' to quit): 75
Your guess is too high. Try again!
Congratulations! You guessed the number 67 correctly in 3 attempts.
It took you 15.23s to guess the correct number.
Your guesses: [50, 75, 67]
```

---

## 🏆 **Game Highlights**  
- Play solo or challenge your friends to see who can guess the fastest.  
- Learn to improve your estimation and deduction skills.  
- Works seamlessly in any terminal or command-line interface.  

---

## 🖋 **Contributing**  

We welcome contributions! 🎉  
1. Fork the repository.  
2. Create your feature branch: `git checkout -b feature/YourFeature`.  
3. Commit your changes: `git commit -m "Add some feature"`.  
4. Push to the branch: `git push origin feature/YourFeature`.  
5. Open a pull request.  

---

## 📜 **License**  

This project is licensed under the MIT License.  

---

### 🌟 **Enjoy the Game?**  
Drop a star ⭐ on the repo and share it with your friends! 😊
