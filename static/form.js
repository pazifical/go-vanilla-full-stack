document.addEventListener("submit", async (event) => {
  event.preventDefault();
  const form = event.target;

  const url = form.action;
  const data = parseFormData(form);

  if (form.method === "get") {
    await handleGet(url, data);
  } else if (form.method === "post") {
    await handleJsonPost(url, data);
  }
});

async function handleGet(url, data) {
  // TODO
}

async function handleJsonPost(url, data) {
  const response = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  console.log(response);
}

function parseFormData(form) {
  let data = {};

  for (const input of form.getElementsByTagName("input")) {
    switch (input.type) {
      case "text":
      case "password":
      case "tel":
      case "search":
      case "email":
        data[input.name] = String(input.value);
        input.value = null;
        break;
      case "number":
      case "range":
        data[input.name] = Number(input.value);
        input.value = null;
        break;
      case "checkbox":
        data[input.name] = input.checked;
        input.checked = false;
        break;
      case "radio":
        if (input.checked) {
          data[input.name] = String(input.value);
        }
        input.value = null;
        break;
      case "time":
        // TODO:
        break;
      case "date":
        // TODO:
        break;
      case "datetime-local":
        // TODO:
        break;
    }
  }

  for (const textarea of form.getElementsByTagName("textarea")) {
    data[textarea.name] = String(textarea.value);
    textarea.value = null;
  }

  for (const select of form.getElementsByTagName("select")) {
    data[select.name] = String(select.value);
    select.value = null;
  }

  return data;
}
