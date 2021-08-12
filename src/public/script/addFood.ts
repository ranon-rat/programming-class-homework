
function addFood(){
  const name = document.getElementById("name-of-the-food") as HTMLInputElement;
  const whereIsIt = document.getElementById("where-is-it") as HTMLInputElement;
  const imageURL = document.getElementById("image-url") as HTMLInputElement;
  const cost = document.getElementById("cost") as HTMLInputElement;
  
  fetch("/api/", {
    method: "POST",
    body: JSON.stringify({
        name: name.value,
        "where-it-is": whereIsIt.value,
        "image-url": imageURL.value,
        cost: parseFloat(cost.value),
      }),
  });
}
