<?php

function sh($stuff, $die=false) {
    "<pre>{$stuff}</pre>";
    $die ?? die; 
}