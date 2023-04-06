const logoutBtn = document.getElementById("logout-btn");
logoutBtn.addEventListener('click', logout);

async function logout() {
  try {
    const req = await window.fetch("/api/auth/logout");
    const result = await req.json();

    if (result.status === 200) {
      window.location.href = "/";
    }
  } catch (error) {
    console.log(error);
  }
}
