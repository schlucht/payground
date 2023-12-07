<?php
use Ots\Api\DB;

require_once "../Api/autoload.php";
session_start();



?>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Home Page</title>
</head>
<body>
    <a href="./signup.php">Signup</a>
    <?php $db = DB::connectDB();?>
</body>
</html>

