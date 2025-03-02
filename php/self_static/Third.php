<?php

class Third extends Second
{
    public static $name = 'third';

    public function parentName(): string
    {
        return parent::$name;
    }

    public static function createUsingSelf2(): self
    {
        return new self();
    }
}