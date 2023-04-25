async function listBook() {
  try {
    const req = await window.fetch("/api/book/list");

    const result = await req.json();

    if (result.status === 200) {
      renderData(result.result);
    };
  } catch (error) {
    console.log(error);
  }
}

listBook();

function renderData(data) {
  const body = document.getElementById("table-body");

  // reset the inner html
  body.innerHTML = "";

  for (let i = 0; i < data.length; i++) {
    const current = data[i];
    const newElem = document.createElement("tr");

    newElem.innerHTML = `
        <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
          <td class="flex justify-center p-4">
            <img src="${current.imageURL}" width="60" height="50" alt="Apple Watch">
          </td>
          <td class="px-6 py-4 font-semibold text-gray-900 dark:text-white">
            ${current.title}
          </td>
          <td class="px-6 py-4 font-semibold text-gray-900 dark:text-white">
            ${current.author}
          </td>
          <td class="px-6 py-4 font-semibold text-gray-900 dark:text-white">
            ${current.releaseDate}
          </td>
          <td class="px-6 py-4 font-semibold text-gray-900 dark:text-white">
            $${current.price}
          </td>
        </tr>
    `;

    body.append(newElem);
  }
}
