document.addEventListener('DOMContentLoaded', async () => {
	await getLists();
	await getItems()
});

document.getElementById('item-form').addEventListener('submit', async () => {
	event.preventDefault();
	await addNewItem();
})

document.getElementById('new-list-form').addEventListener('submit', async () => {
	event.preventDefault();
	await addNewList();
})
document.getElementById('list-create-button').addEventListener('submit', async () => {
	event.preventDefault();
	await addNewList();
})

document.getElementById('list-name-select').addEventListener('change', async () => {
	event.preventDefault();
	await getItems();
})

document.getElementById('list-delete-button').addEventListener('click', async () => {
	event.preventDefault();
	await deleteList();
})

async function deleteList() {
	const listID = document.getElementById('list-name-select').value;
	try {
		const res = await fetch(`/api/groceries/lists/${listID}`, {
			method: 'DELETE'
		});
		if (!res.ok) {
			const data = await res.json()
			throw new Error(`Failed to delete list. Error: ${data.error}`);
		}
		await getLists()
		await getItems();
	} catch (err) {
		alert(`Error here: ${err.message}`);
	}
}

async function getLists() {
	try {
		const res = await fetch('/api/groceries/lists', {
			method: 'GET'
		});
		if (!res.ok) {
			const data = await res.json()
			throw new Error(`Failed to get lists. Error: ${data.error}`);
		}
		const lists = await res.json();
		const listsSelect = document.getElementById('list-name-select');
		listsSelect.innerHTML = null;
		for( const list of lists) {
			newListOpt = document.createElement("option");
			newListOpt.textContent = list.Name;
			newListOpt.value = list.ID;
			listsSelect.appendChild(newListOpt);
		}

	} catch(err) {
		alert(`Error: ${err.message}`);
	}
}

async function addNewList() {
	try {
		const res = await fetch('/api/groceries/lists', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				Name: document.getElementById('list-name-input').value
			}),
		});
		if(!res.ok) {
			const data = await res.json();
			throw new Error(`Failed to add list. Error: ${data.error}`);
		}
		await getLists();
	} catch(err) {
		alert("Error: " + err.message)
	}
	document.getElementById("list-name-input").value = "";
}

async function addNewItem() {
	try {
		const res = await fetch('/api/groceries/items', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				Name: document.getElementById('item-input-name').value,
				Amount: document.getElementById('item-input-amount').value,
				ListID:  document.getElementById('list-name-select').value
			}),
		});
		if (!res.ok) {
			const data = await res.json();
			throw new Error(`Failed to add item. Error: ${data.error}`);
		}

		await getItems();
	} catch(err) {
		alert(`Error: ${err.message}`);
	}
}

async function getItems() {
	try {
		const listID = document.getElementById('list-name-select').value;
		if(listID === "") {
			return;
		}
		const res = await fetch('/api/groceries/items/' + listID, {
			method: 'GET'
		});
		if (!res.ok) {
			const data = await res.json()
			throw new Error(`Failed to get items. Error: ${data.error}`);
		}
		const items = await res.json();
		const itemsList = document.getElementById('items-list');
		itemsList.innerHTML = '';
		for (const item of items) {
			const listItem = document.createElement("li");
			listItem.textContent = item.amount + "\t" + item.name;
			listItem.onclick = async () => await removeItem(item.id);
			listItem.id = "item-desc";
			itemsList.appendChild(listItem);

		}
	} catch (err) {
		alert(`Error: ${err.message}`);
	}
}

async function removeItem(id) {
	try {
		const res = await fetch(`/api/groceries/items/${id}`, {
			method: 'DELETE'
		});
		if (!res.ok) {
			const data = await res.json()
			throw new Error(`Failed to delete item. Error: ${data.error}`);
		}
		const listID = document.getElementById('list-name-select').value;
		await getItems();
	} catch(err) {
		alert(`Error: ${err.message}`);
	}
}