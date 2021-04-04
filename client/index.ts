document.getElementById('lol').onclick = () => {
    const req = new Request('/api/note', {
        method: 'POST',
        body: JSON.stringify({
            title: 'HELLO',
            content: 'KSKSKKSKSKSKSK'
        })
    });
    fetch(req);
};

const conn = new WebSocket('ws://localhost:8080/ws');
conn.onopen = (evt) => {
  console.log('sddsddsddssd');
};
conn.onerror = (evt) => {
  console.log('sddsddsddssd');
};
conn.onclose = function (evt) {
  console.log('sddsddsddssd');

  var item = document.createElement('div');
  item.innerHTML = '<b>Connection closed.</b>';
  document.body.append(item);
};
conn.onmessage = function (evt) {
  console.log('sddsddsddssd');
  var messages = evt.data.split('\n');
  for (var i = 0; i < messages.length; i++) {
    var item = document.createElement('div');
    item.innerText = messages[i];
    document.body.append(item);
  }
};
