import { Food } from "./types";
function addElementsFood() {
  const listLocations = document.getElementById("location-list");
  const food = document.getElementById("comida");
  fetch("/api/")
    .then((r) => r.json())
    .then((d: Record<string, Food[]>) => {
      Object.keys(d).forEach((i) => {
        listLocations!.innerHTML += `
               <a href="#${i.replace(" ", "-")}">${i}</a>

                `;
        try {
          food!.innerHTML += `<div id="${i.replace(" ", "-")}">

        <h1>${i}</h1>
        ${d[i]
          .map(
            (i) => `<a href="${document.URL+"&buy="+i["food-id"]}" class="post">
        <p>
            ${i["where-it-is"]}
        </p>
        <div class="image-superposition">
            <div class="image-food" style="background-image: url(${i["image-url"]})">
                <p>
                ${i.name}
                </p>
            </div>
        </div>
        <p>
            ${i.cost}$
        </p>

    </a>`
          )
          .join("")}</div>`;
        } catch (e) {
          console.log(e);
        }
      });
    });
}
addElementsFood();
