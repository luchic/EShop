function Placeholder() {
    alert("Placeholder")
}

async function AddItem() {
    let name = document.getElementById("Iname").value;
    if (!name.trim()) {
        alert('Please enter all fields!');
        return;
    }

    let logo_raw = document.getElementById("Ilogo").files;
    if (!logo_raw || logo_raw.length === 0) {
        alert('Please enter all fields!');
        return;
    }

    if (!logo_raw[0].type.startsWith("image")) {
        alert("Please attach an image file");
        return;
    }

    // Convert File to Base64 and wait for it
    const logo = await new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = () => resolve(reader.result.split(',')[1]);
        reader.onerror = () => reject(reader.error);
        reader.readAsDataURL(logo_raw[0]);
    });

    let desc = document.getElementById('Idesc').value;
    let price = document.getElementById('Iprice').value;

    if (!desc.trim() || !price.trim()) {
        alert('Please enter all fields!');
        return;
    }

    const item = { name, logo, desc, price };
    const jsonstr = JSON.stringify(item);
    const response = await fetch('http://localhost:8080/itmes/add', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify(item)
});
const data = await response.json();
    console.log(data)
}