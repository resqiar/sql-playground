let currentPage = 1;

const loadMoreBtn = document.getElementById("load-more-btn");
const table = document.getElementById("item-list");

loadMoreBtn.addEventListener("click", loadMore);

async function loadMore() {
    try {
        const req = await fetch(`/api/load-more-city?page=${currentPage + 1}`);
        const data = await req.json();

        renderMore(data);
        currentPage++;
    } catch (e) {
        console.log(e);
    }
}

function renderMore(data) {
    const fragments = document.createDocumentFragment();

    for (let i = 0; i < data.length; i++) {
        const elem = data[i];

        const row = document.createElement("tr");

        row.innerHTML = `
            <tr>
                <td>${elem.ID}</td>
                <td>${elem.Name}</td>
                <td>${elem.CountryCode}</td>
                <td>${elem.District}</td>
                <td>${elem.Population}</td>
            </tr>
        `;

        fragments.append(row);
    }

    table.append(fragments);
}
