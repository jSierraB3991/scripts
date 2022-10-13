package com.github.jSierraB3991.script;

public class ScoreDahBoard {
    private Integer ties;
    private Integer wins;
    private Integer lose;

    public ScoreDahBoard() {
        wins = 0;
        ties = 0;
        lose = 0;
    }

    public Integer getTies() {
        return ties;
    }

    public void incrementTies() {
        this.ties++;
    }

    public Integer getWins() {
        return wins;
    }

    public void incrementWins() {
        this.wins++;
    }

    public Integer getLose() {
        return lose;
    }

    public void incrementLose() {
        this.lose++;
    }
}