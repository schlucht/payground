<?php 
    
    if($_SERVER['REQUEST_METHOD'] == "POST") {
        print_r($_POST['email']);
        $email = $_POST['email'];
        if (preg_match('/^[\w\-]+@[\w\-]+.[\w\-]+$/', $email)) {
            echo "<strong>{$email}</strong>";
        } else {
            echo "<strong>Nix gefunden</strong>";
        }
    }
?>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Signup</title>
</head>
<body>
    <style>
        form {
            border: solid thin #ccc;
            margin: auto;
            padding: .6em;
            margin-top: 2rem;
            width: 40%;
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            gap: .5rem;
        }
    </style>
    <form method="POST">
        <h2>Signup</h2>
        <input type="email" name="email" required>
        <input type="password" name="password" required><br>
        <input type="submit" value="Submit">
    </form>
</body>
</html>