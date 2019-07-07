const el = document.querySelector('table');
el.addEventListener('click', function(e) {
    let targetEl = e.target;
    if ( targetEl.classList.contains("edit")) {
        document.getElementById("edit").style.display = "flex";   
        console.log(targetEl);   
        console.log(targetEl.parentElement.parentElement); 
        row = targetEl.parentElement.parentElement;
        let id = parseInt(row.children[0].innerHTML);
        let code = (row.children[2].innerHTML);
        let name = (row.children[1].innerHTML);
        document.getElementById('editId').defaultValue = id;
        document.getElementById('editName').defaultValue = name;
        document.getElementById('editCode').defaultValue = code;
      
    };

});