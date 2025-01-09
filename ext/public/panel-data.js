chrome.devtools.network.onRequestFinished.addListener(request => { 
    request.getContent((body) => {
       if(request.request && request.request.url) {
        if (request.request.url.includes('.')) { 
        //   chrome.runtime.sendMessage({ 
        //       response: body 
        //   }); 
          console.log('body:',body); 
        } 
      } 
    }); 
  });