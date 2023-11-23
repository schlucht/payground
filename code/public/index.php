<?php


require_once __DIR__ . '/inc/all.php';

$route = @(string) ($_GET['route'] ?? 'page');

if ($route === 'page') {
    $pagesController = new \App\Controller\PagesController(
        new \App\Pages\PagesRepository($pdo)
    );
    $pages = @(string) ($_GET['page'] ?? 'index');
    $pagesController->show($pages);    
} else {
    $notFoundController = new \App\Controller\NotFoundController(
        new \App\Pages\PagesRepository($pdo)
    );
    $notFoundController->error404();
}