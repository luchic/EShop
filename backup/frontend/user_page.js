
const purchases = [
    { name: "Item 1", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 3", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 4", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 5", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 6", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 7", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 8", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 9", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 10", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 11", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    { name: "Item 2", price: "tbd", img: "blahaj.jpg" },
    // ... add as many as you want
];

let currentPage = 1;
const itemsPerPage = 5;

function renderTable() {
    const tbody = document.getElementById("purchase-body");
    tbody.innerHTML = "";

    const start = (currentPage - 1) * itemsPerPage;
    const end = start + itemsPerPage;

    const pageItems = purchases.slice(start, end);

    pageItems.forEach(item => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td><img src="${item.img}" alt=""></td>
            <td>${item.name}</td>
            <td>${item.price}</td>
        `;
        tbody.appendChild(row);
    });

    document.getElementById("page-info").textContent =
        `Page ${currentPage} of ${Math.ceil(purchases.length / itemsPerPage)}`;
}

document.getElementById("prev-btn").onclick = () => {
    if (currentPage > 1) {
        currentPage--;
        renderTable();
    }
};

document.getElementById("next-btn").onclick = () => {
    if (currentPage < Math.ceil(purchases.length / itemsPerPage)) {
        currentPage++;
        renderTable();
    }
	console.log(currentPage)
};

renderTable();