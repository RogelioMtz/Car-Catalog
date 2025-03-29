const apiUrl = 'http://localhost:8080/api/cars';

document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('data-form');
    const resultTable = document.getElementById('result-table');

    // Fetch and display the catalog on page load
    fetchCatalog();

    form.addEventListener('submit', async (event) => {
        event.preventDefault();

        const formData = new FormData(form);
        const data = Object.fromEntries(formData.entries());
        data.año = parseInt(data.año, 10);
        data.precio = parseFloat(data.precio);
        data.dateAdded = new Date().toLocaleString(); // Add the current date and time

        try {
            const response = await fetch(apiUrl, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });

            if (!response.ok) {
                throw new Error('Failed to add car');
            }

            // Clear the input fields
            form.reset();

            // Refresh the catalog after adding a car
            fetchCatalog();
        } catch (error) {
            alert('Error: ' + error.message);
        }
    });

    async function fetchCatalog() {
        try {
            const response = await fetch(apiUrl);
            if (!response.ok) {
                throw new Error('Failed to fetch catalog');
            }

            const catalog = await response.json();

            // Clear the table before populating
            resultTable.innerHTML = `
                <tr>
                    <th>Marca</th>
                    <th>Modelo</th>
                    <th>Color</th>
                    <th>Año</th>
                    <th>Precio</th>
                    <th>Fecha y Hora de Registro</th>
                </tr>
            `;

            // Populate the table with catalog data
            catalog.forEach((car) => {
                const row = document.createElement('tr');
                row.innerHTML = `
                    <td>${car.marca}</td>
                    <td>${car.modelo}</td>
                    <td>${car.color}</td>
                    <td>${car.año}</td>
                    <td>${car.precio}</td>
                    <td>${car.dateAdded || 'N/A'}</td>
                `;
                resultTable.appendChild(row);
            });
        } catch (error) {
            alert('Error: ' + error.message);
        }
    }
});