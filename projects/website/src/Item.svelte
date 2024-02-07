<script>
    import { createEventDispatcher, tick } from 'svelte';
    import todoService from './apis/todoService'

    export let item;

    const dispatch = createEventDispatcher();
    
    let editing = false;

    function removeItem() {
        dispatch('removeItem');
    }

    function startEdit() {
       editing = true;
    }

    function handleEdit(event) {
        if (event.key === "Enter")
            event.target.blur();
        else if (event.key === "Escape") 
            editing = false;
    }

    async function updateItem(event) {
        if (!editing) return;
        const { value } = event.target;
        if (value.length) {
            item.content = value;
            await todoService.editTodo(item)
        } else {
          removeItem();
        }
        editing = false;
    }

    async function focusInput(element) {
        await tick();
        element.focus();
    }

    async function onCheck(event) {
        item.completed = event.target.checked
        await todoService.editTodo(item)
    }
</script>

<li class:completed={item.completed} class:editing>
    <div class="view">
        <input class="toggle" type="checkbox" on:change={onCheck} checked={item.completed} />
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label on:dblclick={startEdit}>{item.content}</label>
        <button on:click={removeItem} class="destroy" />
    </div>

    {#if editing}
        <div class="input-container">
            <input value={item.content} id="edit-todo-input" class="edit" on:keydown={handleEdit} on:blur={updateItem} use:focusInput />
            <label class="visually-hidden" for="edit-todo-input">Edit Todo Input</label>
        </div>
    {/if}
</li>
