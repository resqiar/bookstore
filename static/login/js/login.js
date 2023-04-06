const usernameInput = document.getElementById("username-input");
const passwordInput = document.getElementById("password-input");
const loginButton = document.getElementById("login-button");
const errorElem = document.getElementById("error");

loginButton.addEventListener("click", login);

async function login() {
  // reset error
  renderError("");

  const username = usernameInput.value;
  const password = passwordInput.value;

  if (!username || !password) return renderError("username / password is required");

  try {
    const req = await window.fetch("/api/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        username: username,
        password: password,
      }),
    });

    const result = await req.json();

    if (result.status === 400) {
      return renderError(result.message);
    }

    if (result.status === 200) window.location = "/";
  } catch (error) {
    renderError(error.message);
  }
}

function renderError(message) {
  if (message === "") return errorElem.innerHTML = "";

  errorElem.innerHTML = `
    <div class="flex p-4 mb-4 text-sm text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400" role="alert">
      <svg aria-hidden="true" class="flex-shrink-0 inline w-5 h-5 mr-3" fill="currentColor" viewBox="0 0 20 20"
        xmlns="http://www.w3.org/2000/svg">
        <path fill-rule="evenodd"
          d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
          clip-rule="evenodd"></path>
      </svg>
      <span class="sr-only">Info</span>
      <div>
        ${message}
      </div>
    </div>
  `;
}
