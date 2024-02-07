<script>
    import { onMount } from 'svelte';
    import { router } from './router.js';

    import todoService from './apis/todoService'

    import Header from './Header.svelte';
    import Footer from './Footer.svelte';
    import Item from './Item.svelte';

    import "./app.css";
    import "todomvc-app-css/index.css";
    import "todomvc-common/base.css";

    let currentFilter = "all";
    let items = [];

    async function addItem(event) {
        const {response} = await todoService.newTodo({ content: event.detail.text })
        items.push(response.todo)
        items = items //forcing ui update
    }

    async function removeItem(index) {
        const item = items[index]
        items.splice(index, 1);
        items = items;
        await todoService.deleteTodo({ id: item.id })
    }

    async function toggleAllItems() {
        await todoService.completeAll() 
        items = items.map((item) => ({
            ...item,
            completed: true,
        }));
    }

    async function removeCompletedItems() {
        await todoService.clearCompleted()
        items = items.filter((item) => !item.completed);
    }
    
    onMount(async () => {
      router(route => currentFilter = route).init();

      const { response } = await todoService.getTodos({})
      items = response.todos
    });

    $: filtered = currentFilter === "all" ? items : currentFilter === "completed" ? items.filter((item) => item.completed) : items.filter((item) => !item.completed);
    $: numActive = items.filter((item) => !item.completed).length;
    $: numCompleted = items.filter((item) => item.completed).length;
</script>

<Header on:addItem={addItem} />

{#if items.length > 0}
    <main class="main">
        <div class="toggle-all-container">
            <input id="toggle-all" class="toggle-all" type="checkbox" on:change={toggleAllItems} checked={numCompleted === items.length} />
            <label for="toggle-all">Mark all as complete</label>
        </div>
        <ul class="todo-list">
            {#each filtered as item, index (item.id)}
                <Item bind:item on:removeItem={() => removeItem(index)} />
            {/each}
        </ul>

        <Footer {numActive} {currentFilter} {numCompleted} on:removeCompletedItems={removeCompletedItems} />
    </main>
{/if}
