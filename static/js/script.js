document.addEventListener("DOMContentLoaded", function() {
    const modal = document.getElementById("editModal");
    const span = document.getElementsByClassName("close")[0];
    const editForm = document.getElementById("editForm");

    // Обработка кнопок редактирования
    const editButtons = document.getElementsByClassName("edit-btn");
    Array.from(editButtons).forEach(button => {
        button.addEventListener("click", function() {
            const id = this.getAttribute("data-id");
            const row = this.parentElement.parentElement;
            document.getElementById("editId").value = id;
            document.getElementById("editBorrowerID").value = row.cells[1].innerText;
            document.getElementById("editFirstName").value = row.cells[2].innerText;
            document.getElementById("editLastName").value = row.cells[3].innerText;
            document.getElementById("editSurname").value = row.cells[4].innerText;
            document.getElementById("editPassport").value = row.cells[5].innerText;
            document.getElementById("editINN").value = row.cells[6].innerText;
            document.getElementById("editSNILS").value = row.cells[7].innerText;
            document.getElementById("editDriverLicense").value = row.cells[8].innerText;
            document.getElementById("editAdditionalDocs").value = row.cells[9].innerText;
            document.getElementById("editComment").value = row.cells[10].innerText;

            modal.style.display = "block";
        });
    });

    // Закрытие модального окна
    span.onclick = function() {
        modal.style.display = "none";
    }

    window.onclick = function(event) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
    }

    // Обработка кнопок удаления
    const deleteButtons = document.getElementsByClassName("delete-btn");
    Array.from(deleteButtons).forEach(button => {
        button.addEventListener("click", function() {
            const id = this.getAttribute("data-id");
            if (confirm("Вы уверены, что хотите удалить эту запись?")) {
                fetch(`/delete?id=${id}`, {
                    method: 'GET'
                }).then(response => {
                    if (response.ok) {
                        window.location.reload();
                    } else {
                        alert("Ошибка при удалении записи.");
                    }
                });
            }
        });
    });
});
