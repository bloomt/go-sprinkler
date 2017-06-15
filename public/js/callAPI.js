  function createNode(element) {
      return document.createElement(element);
  }

  function append(parent, el) {
    return parent.appendChild(el);
  }

  function s1(checkboxElem) {
  if (checkboxElem.checked) {
    fetch(open1)
  .catch(function(error) {
    console.log(JSON.stringify(error));
  }); 
    document.getElementById("valve2").classList.add('disabled');
    document.getElementById("c2").disabled= true;
    document.getElementById("valve3").classList.add('disabled');
    document.getElementById("c3").disabled= true;
    document.getElementById("valve4").classList.add('disabled');
    document.getElementById("c4").disabled= true;
  } else {
    fetch(close1)
  .catch(function(error) {
    console.log(JSON.stringify(error));
  }); 
    document.getElementById("valve2").classList.remove('disabled');
    document.getElementById("c2").disabled= false;
    document.getElementById("valve3").classList.remove('disabled');
    document.getElementById("c3").disabled= false;
    document.getElementById("valve4").classList.remove('disabled');
    document.getElementById("c4").disabled= false;
  }
}

  function s2(checkboxElem) {
  if (checkboxElem.checked) {
    fetch(open2)
  .catch(function(error) {
    console.log(JSON.stringify(error));
  }); 
    document.getElementById("valve1").classList.add('disabled');
    document.getElementById("c1").disabled= true;
    document.getElementById("valve3").classList.add('disabled');
    document.getElementById("c3").disabled= true;
    document.getElementById("valve4").classList.add('disabled');
    document.getElementById("c4").disabled= true;
  } else {
    fetch(close2)
  .catch(function(error) {
    console.log(JSON.stringify(error));
  });
    document.getElementById("valve1").classList.remove('disabled');
    document.getElementById("c1").disabled= false;
    document.getElementById("valve3").classList.remove('disabled');
    document.getElementById("c3").disabled= false;
    document.getElementById("valve4").classList.remove('disabled');
    document.getElementById("c4").disabled= false; 
  }
}

  function s3(checkboxElem) {
  if (checkboxElem.checked) {
    fetch(open3)
  .catch(function(error) {
    console.log(JSON.stringify(error));
  }); 
    document.getElementById("valve1").classList.add('disabled');
    document.getElementById("c1").disabled= true;
    document.getElementById("valve2").classList.add('disabled');
    document.getElementById("c2").disabled= true;
    document.getElementById("valve4").classList.add('disabled');
    document.getElementById("c4").disabled= true;
  } else {
    fetch(close3)
  .catch(function(error) {
    console.log(JSON.stringify(error));
  }); 
    document.getElementById("valve1").classList.remove('disabled');
    document.getElementById("c1").disabled= false;
    document.getElementById("valve2").classList.remove('disabled');
    document.getElementById("c2").disabled= false;
    document.getElementById("valve4").classList.remove('disabled');
    document.getElementById("c4").disabled= false; 
  }
}

  function s4(checkboxElem) {
  if (checkboxElem.checked) {
    fetch(open4)
  .catch(function(error) {
    console.log(JSON.stringify(error));
  }); 
    document.getElementById("valve1").classList.add('disabled');
    document.getElementById("c1").disabled= true;
    document.getElementById("valve2").classList.add('disabled');
    document.getElementById("c2").disabled= true;
    document.getElementById("valve3").classList.add('disabled');
    document.getElementById("c3").disabled= true;
  } else {
    fetch(close4)
  .catch(function(error) {
    console.log(JSON.stringify(error));
  }); 
    document.getElementById("valve1").classList.remove('disabled');
    document.getElementById("c1").disabled= false;
    document.getElementById("valve2").classList.remove('disabled');
    document.getElementById("c2").disabled= false;
    document.getElementById("valve3").classList.remove('disabled');
    document.getElementById("c3").disabled= false; 
  }
}
  
  const open1 = 'http://192.168.88.72:8080/api/valve/1/open';
  const close1 = 'http://192.168.88.72:8080/api/valve/1/close';
  
  const open2 = 'http://192.168.88.72:8080/api/valve/2/open';
  const close2 = 'http://192.168.88.72:8080/api/valve/2/close';
  
  const open3 = 'http://192.168.88.72:8080/api/valve/3/open';
  const close3 = 'http://192.168.88.72:8080/api/valve/3/close';

  const open4 = 'http://192.168.88.72:8080/api/valve/4/open';
  const close4 = 'http://192.168.88.72:8080/api/valve/4/close';
