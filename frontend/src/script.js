async function fetchAsync(url) {
  let response = await fetch(url);
  let data = await response.json();
  return data;
}

console.log(fetchAsync("http://127.0.0.1:8080/items/3"));

const p = document.querySelector(".test");

console.log(p);
