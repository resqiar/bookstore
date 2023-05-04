let currentUser;

async function getCurrentUser() {
  const guestNavbar = document.getElementById("guest");
  const loggedNavbar = document.getElementById("logged");
  const loggedUsername = document.getElementById("logged-username");
  const loggedEmail = document.getElementById("logged-email");
  const loggedAdmin = document.getElementById("logged-admin");

  try {
    const req = await window.fetch("/api/user/current");
    const result = await req.json();

    if (result.status === 200) {
      currentUser = result.user;

      // show logged navbar
      loggedNavbar.className = "flex";
      // hide login register navbar
      guestNavbar.className = "hidden";
      // update logged in information
      loggedUsername.innerText = currentUser.username;
      loggedEmail.innerText = currentUser.email;
      loggedAdmin.className = "hidden";

      if (currentUser.isAdmin) {
        loggedAdmin.className = "flex";
      }

      return;
    }

    // show guest navbar
    loggedNavbar.className = "hidden";
    // hide logged navbar
    guestNavbar.className = "flex";
  } catch (error) {
    console.log(error);
  }
}

getCurrentUser();
