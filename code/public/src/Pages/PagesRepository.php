<?php

namespace App\Pages;

class PagesRepository {

    public function __construct(protected PDO $pdo) {
    }

    public function fetchPage(string $key) {
        $pageModel = new PagesModel();
        $pageModel->pageId = 1;
        $pageModel->pageKey = $key;
        $pageModel->pageTitle = 'Hallo Lothar';
        $pageModel->pageContent = '<h1 style="color: yellow;">Ich bin die Seite selber</h1>';
        return $pageModel;
    }
}