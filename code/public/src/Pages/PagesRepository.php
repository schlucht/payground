<?php

namespace App\Pages;

use PDO;
use App\Pages\PagesModel;


class PagesRepository {

    public function __construct(protected PDO $pdo) {
    }

    public function fetchPage(string $key): ?PagesModel {
        $stmt = $this->pdo->prepare("SELECT * FROM pages WHERE pageKey = :pkey");
        $stmt->execute(['pkey' => $key]);
        $stmt->setFetchMode(PDO::FETCH_CLASS, PagesModel::class);
        $pageModel = $stmt->fetch();        

        if (empty($pageModel)) {
            return null;
        } else {
            return $pageModel;
        }
    }

    public function getNavigation(): array {
        $stmt = $this->pdo->query('SELECT * FROM pages');
        $names = $stmt->fetchAll(PDO::FETCH_CLASS, PagesModel::class);
        if (empty($names)) {
            return null;
        }else {
            return $names;
        }
        
    }
}