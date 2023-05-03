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
          <td class="">
              <div class="flex items-center gap-2">
                  <button class="p-1 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-full focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700" type="button">
                      <svg class="w-4 h-4" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clip-rule="evenodd"></path></svg>
                  </button>
                  <div>
                      <input type="text" class="bg-gray-50 w-14 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block px-2.5 py-1 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" value="${current.quantity}" disabled>
                  </div>
                  <button class="p-1 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-full focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700" type="button">
                      <svg class="w-4 h-4" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" clip-rule="evenodd"></path></svg>
                  </button>
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

