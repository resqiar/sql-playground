let currentPage = 1;

const loadMoreBtn = document.getElementById("load-more-btn");
const table = document.getElementById("item-list");
const filterInputs = document.querySelectorAll("input[name^='filter_']");
const inputCode = document.querySelector("input[name='filter_code']");
const inputName = document.querySelector("input[name='filter_name']");
const inputCapital = document.querySelector("input[name='filter_capital']");
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
        const req = await fetch(`/api/country-capital?page=${currentPage + 1}`);
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
                <th>Code</th>
                <th>Name</th>
                <th>Capital</th>
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
                <td>${elem.Code}</td>
                <td>${elem.Name}</td>
                <td>${elem.Capital}</td>
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
            /api/filter-country?code=${inputCode.value}&name=${inputName.value}&capital=${inputCapital.value}
        `);
        const data = await req.json();

        render(data, true);
        resetPage();
    } catch (e) {
        console.log(e);
    }
}

function clearFilter() {
    inputCode.value = "";
    inputName.value = "";
    inputCapital.value = "";
    filterData();
}
