var url = "http://127.0.0.1:8080/items";

// Fetch the data from the given url
async function fetchAsync() {
  let response = await fetch(url);
  let data = await response.json();
  return data;
}

// Fetch data and execute the function above
fetchAsync().then((items) => {
  for (let item of items) {
    drawTable(item);
  }
});

// Dynamically create the table according to the data given
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

var latestId = 0;

const table = document.querySelector(".table");
const allInputs = document.querySelectorAll("input");
const submitButton = document.querySelector(".submit-button");
const delButton = document.querySelector(".del-button");
const editButton = document.querySelector(".edit-button");

submitButton.addEventListener("click", () => {
  if (allInputs[0].value && allInputs[1].value && allInputs[2].value != "") {
    // If all inputs are good
    latestId++; // Increase id by one
    var xhr = new XMLHttpRequest(); // Make an http request of POST and send the body
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
    location.reload(); // Reload for table to update. Could be done in react tbh but too much work for a small project. Ain't no frontend engineer
  }
});

delButton.addEventListener("click", () => {
  if (allInputs[4].value != "") {
    var xhr = new XMLHttpRequest(); // Same as above but now without a body, the id is sent on the url (url)/:id
    xhr.open("DELETE", url + "/" + allInputs[4].value, true); // PATCH because it works better than DELETE for some reason.
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send();
    console.log(xhr);

    location.reload();
  }
});

editButton.addEventListener("click", () => {
  if (allInputs[6].value && allInputs[7].value != "") {
    var xhr = new XMLHttpRequest();
    xhr.open(
      "PATCH",
      url + "/" + allInputs[6].value + "/" + allInputs[7].value,
      true
    );
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send();
    console.log(xhr);
    location.reload();
  }
});
