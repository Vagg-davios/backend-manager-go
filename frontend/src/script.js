var url = "http://127.0.0.1:8080/items";

async function fetchAsync() {
  let response = await fetch(url);
  let data = await response.json();
  return data;
}

var latestId;

const table = document.querySelector(".table");

function drawTable(obj) {
  var tableRow = document.createElement("tr");
  var itemId = document.createElement("td");
  var itemName = document.createElement("td");
  var itemPrice = document.createElement("td");
  var itemQuantity = document.createElement("td");

  latestId = obj.id;
  itemId.innerHTML = obj.id;
  itemName.innerHTML = obj.title;
  itemPrice.innerHTML = obj.price + "\u20AC";
  itemQuantity.innerHTML = obj.quantity;

  tableRow.appendChild(itemId);
  tableRow.appendChild(itemName);
  tableRow.appendChild(itemPrice);
  tableRow.appendChild(itemQuantity);

  table.appendChild(tableRow);
}

fetchAsync().then((items) => {
  for (let item of items) {
    drawTable(item);
  }
});

const allInputs = document.querySelectorAll("input");
const submitButton = document.querySelector(".submit-button");

submitButton.addEventListener("click", () => {
  if (allInputs[0].value && allInputs[1].value && allInputs[2].value != "") {
    latestId++;
    var xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(
      '{"id":"' +
        latestId +
        '","title":"' +
        allInputs[0].value +
        '", "price":' +
        allInputs[1].value +
        ', "quantity":' +
        allInputs[2].value +
        "}"
    );
    location.reload();
  }
});
