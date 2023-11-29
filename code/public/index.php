<?php


function show(string $stuff){
    echo "<pre>";
    print_r($stuff);
    echo "</pre>";
}

show("<h1>Hallo Lothar</h1>");