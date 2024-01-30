const Controller = {
  getQuery: () => {
    const form = document.getElementById("form");
    const data = Object.fromEntries(new FormData(form));
    return data.query
  },
  loadMore: (ev) => {
    ev.preventDefault();
    const numberOfRows = document.querySelectorAll('#table-body tr').length
    fetch(`/search?q=${Controller.getQuery()}&limit=20&offset=${numberOfRows}`).then((response) => {
      response.json().then((results) => {
        Controller.updateTable(results);
      });
    });

  },
  search: (ev) => {
    ev.preventDefault();
    fetch(`/search?q=${Controller.getQuery()}`).then((response) => {
      response.json().then((results) => {
        Controller.updateTable(results);
      });
    });
  },

  updateTable: (results) => {
    const table = document.getElementById("table-body");
    const rows = [];
    for (let result of results) {
      rows.push(`<tr><td>${result}</td></tr>`);
    }
    table.innerHTML = rows;
  },
};

const form = document.getElementById("form");
form.addEventListener("submit", Controller.search);
document.getElementById("load-more").onclick = Controller.loadMore;
