<?php

class First
{
    public static $name = 'first';

    public static function createUsingSelf(): self
    {
        return new self();
    }

    public static function createUsingStatic(): static
    {
        return new static();
    }
}