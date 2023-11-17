let currentPage = 1;

const loadMoreBtn = document.getElementById("load-more-btn");
const table = document.getElementById("item-list");
const filterInputs = document.querySelectorAll("input[name^='filter_']");
const inputID = document.querySelector("input[name='filter_id']");
const inputName = document.querySelector("input[name='filter_name']");
const inputCountry = document.querySelector("input[name='filter_countrycode']");
const inputDistrict = document.querySelector("input[name='filter_district']");
const filterBtn = document.getElementById("filter-btn");
const clearFilterBtn = document.getElementById("clear-filter-btn");

loadMoreBtn.addEventListener("click", loadMore);
filterBtn.addEventListener("click", filterData);
clearFilterBtn.addEventListener("click", clearFilter);

for (let i = 0; i < filterInputs.length; i++) {
    const elem = filterInputs[i];
    elem.addEventListener("keyup", function(e) {
        if (e.key === "Enter") {
            filterData();
        }
    });
}

async function loadMore() {
    try {
        const req = await fetch(`/api/load-more-city?page=${currentPage + 1}`);
        const data = await req.json();

        render(data);
        currentPage++;
    } catch (e) {
        console.log(e);
    }
}

function render(data, clean) {
    if (clean) {
        table.innerHTML = `
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>CountryCode</th>
                <th>District</th>
                <th>Population</th>
            </tr>
        `;
    }

    if (!data) return;

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

function resetPage() {
    currentPage = 1;
}

async function filterData() {
    try {
        const req = await fetch(`
            /api/filter?id=${inputID.value}&name=${inputName.value}&country=${inputCountry.value}&district=${inputDistrict.value}
        `);
        const data = await req.json();

        render(data, true);
        resetPage();
    } catch (e) {
        console.log(e);
    }
}

function clearFilter() {
    inputID.value = "";
    inputName.value = "";
    inputCountry.value = "";
    inputDistrict.value = "";
    filterData();
}
