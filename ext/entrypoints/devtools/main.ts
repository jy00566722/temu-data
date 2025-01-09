browser.devtools.panels.create(
    'TEMU',
    'icon/128.png',
    'devtools-panel.html',
  );
  
  browser.devtools.panels.elements
    .createSidebarPane('Mihu')
    .then((pane) => {
      pane.setPage('devtools-pane.html');
    });