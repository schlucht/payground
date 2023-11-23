<?php

namespace App\Controller;
use App\Pages\PagesRepository;

class PagesController extends AbstractController {   

  
    public function __construct(protected PagesRepository $pagesRepository) {
        parent::__construct($pagesRepository);
    }

    public function show(string $pageKey) {
        $page = $this->pagesRepository->fetchPage($pageKey);
        if (empty($page)) {
          $this->showError404();
        }
        $navigation = $this->pagesRepository->getNavigation();      
        $this->render('pages/showPage', [
            'page' => $page,
            'navigation' => $navigation,
        ]);
    }
}