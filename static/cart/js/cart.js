async function initCart() {
  try {
    const req = await fetch("/api/cart/current");
    const result = await req.json();

    if (result.status === 200) {
      renderItems(result.result);
    }
  } catch (error) {
    console.log(error);
  }
}

initCart();

const notFoundError = document.getElementById("404");
const tableBody = document.getElementById("table-body");

function renderItems(data) {
  if (!data || !data.length) {
    return notFoundError.innerHTML = `
      <h1 class="font-medium my-24">You don't have anything inside your cart</h1>
    `;
  }

  notFoundError.innerHTML = ``;

  // update the table body
}
