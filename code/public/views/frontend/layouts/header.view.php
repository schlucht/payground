<header>
    <h1>Ein eigenes CMS</h1>
    <p>Hier finden Sie eine schöne CMS-App</p>
    <nav class="nav"> 
    <?php foreach ($navigation as $nav): ?>
        <a href="./?<?=http_build_query(['page' => $nav->pageKey])?>"><?=e($nav->pageTitle);?></a>
        <?php endforeach ?>
    </nav>
</header>