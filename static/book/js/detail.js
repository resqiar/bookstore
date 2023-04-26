async function getDetail() {
  const URL = window.location.href;
  const splitURL = URL.split("/");
  const id = splitURL.pop();

  try {
    console.log(id);
    const req = await window.fetch(`/api/book/${id}`);
    const result = await req.json();

    renderData(result.result);
  } catch (error) {
    console.log(error);
  }
}

getDetail();

function renderData(data) {
  const title = document.getElementById("book-title");
  const desc = document.getElementById("book-desc");
  const price = document.getElementById("book-price");
  const img = document.getElementById("book-img");
  const author = document.getElementById("book-author");
  const release = document.getElementById("book-date");

  title.innerText = data.title;
  desc.innerText = data.description;
  price.innerText = "$" + data.price;
  img.setAttribute("src", data.imageURL);
  img.setAttribute("alt", data.title);
  author.innerText = data.author;
  release.innerText = data.releaseDate;
}
