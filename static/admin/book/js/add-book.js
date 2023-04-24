const btnSubmit = document.getElementById("btn_submit");

btnSubmit.addEventListener("click", addBook);

function addBook() {
  const bookURL = document.getElementById("book_img").value;
  const bookTitle = document.getElementById("book_title").value;
  const bookDesc = document.getElementById("book_desc").value;
  const bookDate = document.getElementById("book_date").value;
  const bookAuthor = document.getElementById("book_author").value;
  const bookPrice = document.getElementById("book_price").value;

  console.log({
    bookURL, bookTitle, bookDesc, bookDate, bookAuthor, bookPrice
  })
}
