function selectItem(element) {
  let items = document.querySelectorAll('.list-group-item');
  items.forEach(function (item) {
    item.classList.remove('selected');
  });
  element.classList.add('selected');
  let formElement = document.getElementById("selectedDocument");
  formElement.value = element.textContent;
}

function showThumbnail(input) {
  if (input.files && input.files[0]) {
    let reader = new FileReader();
    reader.onload = function (e) {
      let thumbnail = document.getElementById('thumbnail');
      thumbnail.src = e.target.result;
      thumbnail.style.display = 'block';
      thumbnail.style.padding = "10px 0";
    };
    reader.readAsDataURL(input.files[0]);
  }
}

