package com.github.jSierraB3991.script;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.Scanner;

import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
public class GameTest {

    @InjectMocks
    private Game game;

    @Mock
    private Scanner scanner;

    @Test
    public void when_writeQuit_then_exitGame() {
        when(scanner.nextLine())
                .thenReturn("Paper")
                .thenReturn("Rock")
                .thenReturn("Scissors")
                .thenReturn("Quit");
        game.play();
    }
}