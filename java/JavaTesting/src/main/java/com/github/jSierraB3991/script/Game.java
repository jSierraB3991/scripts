package com.github.jSierraB3991.script;

import java.util.Random;
import java.util.Scanner;

public class Game {
    private Scanner input = new Scanner(System.in);
    private Random random = new Random();

    public void play() {
        ScoreDahBoard dahBoard = new ScoreDahBoard();
        GameOptions option = null;
        do {
            System.out.println("Let's play Rock, Paper, Scissors!");
            System.out.println("Say \"Rock\", \"Paper\", or \"Scissors\" to indicate your choice. Otherwise say \"Quit\" to quit.");
            option = findChoose(input.nextLine());
            if (option == null){
                System.out.println("Sorry, it looks like you didn't enter a correct input. Try again.");
            } else if (option != GameOptions.QUIT) {
                GameOptions optionPc = generateChoosePc();
                decideWins(dahBoard, option, optionPc);
                printResults(dahBoard);
            }
        }while (option != GameOptions.QUIT);
    }

    private static void decideWins(ScoreDahBoard dahBoard, GameOptions option, GameOptions optionPc) {
        if (option == optionPc) {
            playTie(dahBoard);
        } else if ((option == GameOptions.ROCK && optionPc == GameOptions.SCISSORS)
            || (option == GameOptions.SCISSORS && optionPc == GameOptions.PAPER)
            || (option == GameOptions.PAPER && optionPc == GameOptions.ROCK) ) {
            playerWins(dahBoard);
        } else {
            playerLose(dahBoard);
        }
    }

    private static void playerLose(ScoreDahBoard dahBoard) {
        System.out.println("you lose.");
        dahBoard.incrementLose();
    }

    private static void playerWins(ScoreDahBoard dahBoard) {
        System.out.println("you win!");
        dahBoard.incrementWins();
    }

    private static void playTie(ScoreDahBoard dahBoard) {
        System.out.println("It's a tie");
        dahBoard.incrementTies();
    }

    private GameOptions generateChoosePc() {
        int randomValue = random.nextInt(3);
        GameOptions option = GameOptions.values()[randomValue];
        System.out.println("Computer chose " + option.toString().toLowerCase());
        return option;
    }

    private GameOptions findChoose(String choose) {
        GameOptions options = null;
        try {
            options = GameOptions.valueOf(choose.toUpperCase());
        }catch (Exception ex){
            return null;
        }
        return options;
    }

    private static void printResults(ScoreDahBoard dahBoard) {
        System.out.println("wins:" + dahBoard.getWins() + "\nloses:" + dahBoard.getLose() + "\nties:" + dahBoard.getTies());
    }
}