<?php

namespace Ots\Api;

use PDO;

class DB {

    static function connectDB() {
        try {
            $pdo = new PDO('mysql:host=database;dbname=mydb', 'root', 'schlucht');
            return $pdo;
        }
        catch (Exception $e) {
            die($e);
        }
    }
}