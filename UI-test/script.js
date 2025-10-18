document.addEventListener('DOMContentLoaded', () => {
  const tableBody = document.getElementById('menu-body');

  fetch('http://localhost:9090/products')
    .then(response => {
      if (!response.ok) throw new Error('Failed to fetch data');
      return response.json();
    })
    .then(data => {
      if (data.length === 0) {
        tableBody.innerHTML = '<tr><td colspan="3">No products available</td></tr>';
        return;
      }

      data.forEach(item => {
        const row = document.createElement('tr');
        row.innerHTML = `
          <td>${item.name}</td>
          <td>${item.price}</td>
          <td>${item.sku}</td>
        `;
        tableBody.appendChild(row);
      });
    })
    .catch(err => {
      console.error(err);
      tableBody.innerHTML = '<tr><td colspan="3">Error loading data</td></tr>';
    });
});
