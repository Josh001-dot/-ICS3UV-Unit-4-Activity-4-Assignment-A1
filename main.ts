/**
 * @author Joshua Adeyemi
 * @version 1.0.0
 * @date 2025-12-26
 * @fileoverview Guessing game using predefined functions.
 */

console.log("Welcome to the Guessing Game!");
console.log("Guess a number between 1 and 10. Enter 0 to quit.");

const randomNumber: number = Math.floor(Math.random() * 10) + 1;
let guess: number = Number(prompt("Enter your guess:") || "0");

while (guess !== 0 && guess !== randomNumber) {
  console.log("Try again.");
  guess = Number(prompt("Enter your guess:") || "0");
}

if (guess === randomNumber) {
  console.log(`Congratulations! You guessed the correct number: ${randomNumber}`);
} else {
  console.log("You quit the game.");
}

console.log("\nDone.");
