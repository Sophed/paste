function post() {
  const content = document.getElementById("input").value;
  const type = document.getElementById("type").value;
  fetch("/post", {
    method: "POST",
    body: content,
    headers: {
      "Post-Type": type,
    },
  }).then((response) => {
    if (response.status == 200) {
      response.text().then((id) => {
        window.location.href = "/p/" + id;
      });
    } else {
      alert("Error: " + response.statusText);
    }
  });
}
