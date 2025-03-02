<?php

require_once "vendor/autoload.php";

echo "Part 1. Create new object of \"Third\" using self " . PHP_EOL;
$third1 = Third::createUsingSelf();
var_dump(
    $third1
);
echo 'get_class: ' . get_class($third1) . PHP_EOL;

echo "Part 2. Create new object of \"Third\" using static" . PHP_EOL;

$third2 = Third::createUsingStatic();
var_dump(
    $third2
);
echo 'get_class: ' . get_class($third2) . PHP_EOL;
echo 'parentName: ' . $third2->parentName() . PHP_EOL;

echo "Part 3. Create new object of \"Third\" using self but method placed in Third class" . PHP_EOL;

$third3 = Third::createUsingSelf2();
var_dump(
    $third3
);
echo 'get_class: ' . get_class($third3) . PHP_EOL;
echo 'parentName: ' . $third3->parentName() . PHP_EOL;

