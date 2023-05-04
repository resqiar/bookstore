let cartData = [];

async function initCart() {
  try {
    const req = await fetch("/api/cart/current");
    const result = await req.json();

    if (result.status === 200) {
      cartData = result.result;
      renderItems(result.result);
      initTotal();
    }
  } catch (error) {
    console.log(error);
  }
}

initCart();

const notFoundError = document.getElementById("404");

function renderItems(data) {
  if (!data || !data.length) {
    return notFoundError.innerHTML = `
      <h1 class="font-medium my-24">You don't have anything inside your cart</h1>
    `;
  }

  notFoundError.innerHTML = ``;

  // update the table body
  const body = document.getElementById("table-body");

  // reset the inner html
  body.innerHTML = "";

  for (let i = 0; i < data.length; i++) {
    const current = data[i];
    const newElem = document.createElement("tr");

    newElem.innerHTML = `
        <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
          <td class="flex justify-center p-4">
            <img src="${current.book.imageURL}" width="60" height="50" alt="Apple Watch">
          </td>
          <td class="px-6 py-4 font-semibold text-gray-900 dark:text-white">
            ${current.book.title}
          </td>
          <td class="px-6 py-4 font-semibold text-gray-900 dark:text-white">
            ${current.book.author}
          </td>
          <td class="px-6 py-4 font-semibold text-gray-900 dark:text-white">
            ${current.book.releaseDate}
          </td>
          <td class="px-6 py-4 font-semibold text-gray-900 dark:text-white">
            $${current.book.price}
          </td>
          <td class="px-6 py-4">
            <div>
                <input onfocusout="updateQuantity(${current.book.ID}, this.value)" type="text" class="px-4 bg-gray-50 w-14 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block px-2.5 py-1 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" value="${current.quantity}">
            </div>
          </td>
        </tr>
    `;

    body.append(newElem);
  }
}

function initTotal() {
  if (!cartData.length) return;

  const totalPriceElem = document.getElementById("total-price");

  let totalPrice = 0;

  for (let i = 0; i < cartData.length; i++) {
    const current = cartData[i];
    const currentTotal = current.quantity * current.book.price;
    totalPrice += currentTotal;
  }

  totalPriceElem.innerText = "$" + totalPrice;
}

async function updateQuantity(id, quantity) {
  if (!currentUser) return window.location.href = "/login";

  try {
    const req = await fetch("/api/cart/edit", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        userID: Number(currentUser.ID),
        bookID: Number(id),
        quantity: Number(quantity),
      }),
    });

    const result = await req.json();

    if (result.status === 200) {
      initCart();
    }
  } catch (error) {
    console.log(error);
  }
}
