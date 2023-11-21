<?php
try{
    $dns = 'mysql:host=' . DB_HOST . ';dbname=' . DB_NAME;
    
    $pdo = new PDO($dns, DB_USER, DB_PASS, [
        PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION
    ]);
}
catch(PDOException $e) {
    echo 'Probleme mit der Datenbanverbindung...' . $e;
    die();
}