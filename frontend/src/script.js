async function fetchAsync(url) {
  let response = await fetch(url);
  let data = await response.json();
  return data;
}

const p = document.querySelector(".test");

console.log(p);
