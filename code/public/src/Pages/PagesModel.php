<?php

namespace App\Pages;

class PagesModel {
    public int $pageId;
    public string $pageTitle;
    public string $pageKey;
    public string $pageContent;

    public function __set($name, $value)
    {
        // empty
    }

}