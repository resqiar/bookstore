async function listBook() {
  try {
    const req = await window.fetch("/api/adm/book/list");

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
          <td class="px-6 py-4 font-semibold text-gray-900 dark:text-white">
            <a href="/admin/book/edit/${current.ID}" class="px-3 py-2 text-xs font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Edit</a>
            <button type="button" class="px-3 py-2 text-xs text-white bg-red-700 hover:bg-red-800 focus:outline-none focus:ring-4 focus:ring-red-300 rounded-lg font-medium text-center mr-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900">Delete</button>
          </td>
        </tr>
    `;

    body.append(newElem);
  }
}
