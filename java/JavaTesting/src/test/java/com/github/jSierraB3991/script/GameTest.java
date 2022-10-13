package com.github.jSierraB3991.script;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.io.ByteArrayOutputStream;
import java.io.PrintStream;
import java.util.Random;
import java.util.Scanner;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.ArgumentMatchers.anyInt;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
public class GameTest {

    private static final String USER_QUIT = "Quit";
    private static final Integer OPTION_ROCK = 0;
    private static final String USER_ROCK = "Rock";
    private static final Integer OPTION_PAPER = 1;
    private static final String USER_PAPER = "Paper";
    private static final Integer OPTION_SCISSORS = 2;
    private static final String USER_SCISSORS = "Scissors";

    @InjectMocks
    private Game game;

    @Mock
    private Scanner scanner;
    @Mock
    private Random random;

    private ByteArrayOutputStream outputStream;

    @BeforeEach
    public void setup(){
        outputStream = new ByteArrayOutputStream();
        System.setOut(new PrintStream(outputStream));
    }

    @Test
    public void when_writeAnyOptions_then_notErrorsForExit() {
        when(scanner.nextLine())
                .thenReturn(USER_PAPER)
                .thenReturn(USER_ROCK)
                .thenReturn(USER_SCISSORS)
                .thenReturn(USER_QUIT);
        game.play();
    }

    @Test
    public void when_writeQuite_then_exitOfGame() {
        when(scanner.nextLine())
                .thenReturn(USER_QUIT);

        game.play();

        String expect = "Let's play Rock, Paper, Scissors!\nSay \"Rock\", \"Paper\", or \"Scissors\" to indicate your choice. Otherwise say \"Quit\" to quit.\n";
        assertEquals(expect, outputStream.toString());
    }

    @Test
    public void when_chooseRockAndComputerChooseScissors_then_youWinToPlay() {
        when(scanner.nextLine())
                .thenReturn(USER_ROCK)
                .thenReturn(USER_QUIT);
        when(random.nextInt(anyInt()))
                .thenReturn(OPTION_SCISSORS);

        game.play();

        assertTrue(outputStream.toString().contains("Computer chose scissors"));
        assertTrue(outputStream.toString().contains("wins:1"));
        assertTrue(outputStream.toString().contains("ties:0"));
        assertTrue(outputStream.toString().contains("loses:0"));
        assertTrue(outputStream.toString().contains("you win!"));
    }

    @Test
    public void when_chooseScissorsAndComputerChoosePaper_then_youWinToPlay() {
        when(scanner.nextLine())
                .thenReturn(USER_SCISSORS)
                .thenReturn(USER_QUIT);
        when(random.nextInt(anyInt()))
                .thenReturn(OPTION_PAPER);

        game.play();

        assertTrue(outputStream.toString().contains("Computer chose paper"));
        assertTrue(outputStream.toString().contains("wins:1"));
        assertTrue(outputStream.toString().contains("ties:0"));
        assertTrue(outputStream.toString().contains("loses:0"));
        assertTrue(outputStream.toString().contains("you win!"));
    }

    @Test
    public void when_choosePaperAndComputerChooseRock_then_youWinToPlay() {
        when(scanner.nextLine())
                .thenReturn(USER_PAPER)
                .thenReturn(USER_QUIT);
        when(random.nextInt(anyInt()))
                .thenReturn(OPTION_ROCK);

        game.play();

        assertTrue(outputStream.toString().contains("Computer chose rock"));
        assertTrue(outputStream.toString().contains("wins:1"));
        assertTrue(outputStream.toString().contains("ties:0"));
        assertTrue(outputStream.toString().contains("loses:0"));
        assertTrue(outputStream.toString().contains("you win!"));
    }

    @Test
    public void when_AnyChooseRock_then_tie() {
        when(scanner.nextLine())
                .thenReturn(USER_ROCK)
                .thenReturn(USER_QUIT);
        when(random.nextInt(anyInt()))
                .thenReturn(OPTION_ROCK);

        game.play();

        assertTrue(outputStream.toString().contains("It's a tie"));
        assertTrue(outputStream.toString().contains("wins:0"));
        assertTrue(outputStream.toString().contains("ties:1"));
        assertTrue(outputStream.toString().contains("loses:0"));
    }

    @Test
    public void when_AnyChoosePaper_then_tie() {
        when(scanner.nextLine())
                .thenReturn(USER_PAPER)
                .thenReturn(USER_QUIT);
        when(random.nextInt(anyInt()))
                .thenReturn(OPTION_PAPER);

        game.play();

        assertTrue(outputStream.toString().contains("It's a tie"));
        assertTrue(outputStream.toString().contains("wins:0"));
        assertTrue(outputStream.toString().contains("ties:1"));
        assertTrue(outputStream.toString().contains("loses:0"));
    }

    @Test
    public void when_AnyChooseScissor_then_tie() {
        when(scanner.nextLine())
                .thenReturn(USER_SCISSORS)
                .thenReturn(USER_QUIT);
        when(random.nextInt(anyInt()))
                .thenReturn(OPTION_SCISSORS);

        game.play();

        assertTrue(outputStream.toString().contains("It's a tie"));
        assertTrue(outputStream.toString().contains("wins:0"));
        assertTrue(outputStream.toString().contains("ties:1"));
        assertTrue(outputStream.toString().contains("loses:0"));
    }


    @Test
    public void when_chooseRockAndComputerChoosePaper_then_youLoseToPlay() {
        when(scanner.nextLine())
                .thenReturn(USER_ROCK)
                .thenReturn(USER_QUIT);
        when(random.nextInt(anyInt()))
                .thenReturn(OPTION_PAPER);

        game.play();

        assertTrue(outputStream.toString().contains("Computer chose paper"));
        assertTrue(outputStream.toString().contains("wins:0"));
        assertTrue(outputStream.toString().contains("ties:0"));
        assertTrue(outputStream.toString().contains("loses:1"));
        assertTrue(outputStream.toString().contains("you lose."));
    }

    @Test
    public void when_chooseScissorsAndComputerChooseRock_then_youLoseToPlay() {
        when(scanner.nextLine())
                .thenReturn(USER_SCISSORS)
                .thenReturn(USER_QUIT);
        when(random.nextInt(anyInt()))
                .thenReturn(OPTION_ROCK);

        game.play();

        assertTrue(outputStream.toString().contains("Computer chose rock"));
        assertTrue(outputStream.toString().contains("wins:0"));
        assertTrue(outputStream.toString().contains("ties:0"));
        assertTrue(outputStream.toString().contains("loses:1"));
        assertTrue(outputStream.toString().contains("you lose."));
    }

    @Test
    public void when_choosePaperAndComputerChooseScissors_then_youLoseToPlay() {
        when(scanner.nextLine())
                .thenReturn(USER_PAPER)
                .thenReturn(USER_QUIT);
        when(random.nextInt(anyInt()))
                .thenReturn(OPTION_SCISSORS);

        game.play();

        assertTrue(outputStream.toString().contains("Computer chose scissor"));
        assertTrue(outputStream.toString().contains("wins:0"));
        assertTrue(outputStream.toString().contains("ties:0"));
        assertTrue(outputStream.toString().contains("loses:1"));
        assertTrue(outputStream.toString().contains("you lose."));
    }

    @Test
    public void when_playerNoChooseValid_then_anotherChance() {
        when(scanner.nextLine())
                .thenReturn("Bad")
                .thenReturn(USER_PAPER)
                .thenReturn(USER_QUIT);

        game.play();

        assertTrue(outputStream.toString().contains("Sorry, it looks like you didn't enter a correct input. Try again."));
    }
}