<script>
  import { onMount } from 'svelte';
  import Navbar from './_components/navbar.svelte';
  import Footer from './_components/footer.svelte';
  import AddItem from './_components/add_item.svelte';
  import EditItem from './_components/edit_item.svelte';

  let isAddItemOpen = false;
  let isEditItemOpen = false;

  function openModalAddItem() { isAddItemOpen = true }
  function closeModalAddItem() { isAddItemOpen = false }
  function openModalEditItem( curTxt, tdoId ) {
    user.current_text = curTxt;
    user.todo_id = tdoId;
    isEditItemOpen = true;
  }
  function closeModalEditItem() { isEditItemOpen = false; }

  let Todos = [];
  let user = {
    user_id: 0,
    loggedIn: false,
    session_id: '',
    current_text: '',
    todo_id: '',
  };
  async function getTodos() {
    const resp = await fetch('/api/todo/list?page=1&limit=10', { method: 'GET' });
    Todos = await resp.json();
  }

  async function deleteTodos(tdoId) {
    user.todo_id = tdoId;
    const reqBody = JSON.stringify({
      id: user.todo_id,
    });
    try {
      const resp = await fetch('/api/todo/delete', {
        method: 'DELETE',
        headers: {
          'X-Session-ID': user.session_id,
          'Content-Type': 'application/json',
        },
        body: reqBody,
      });
      if (resp.ok) {
        const msg = await resp.json();
        console.log(msg.message);
        user.todo_id = '';
        location.reload();
      }
    } catch (error) {
      console.error(error);
      alert('Error delete item');
    }
  }

  onMount(() => {
    getTodos();
    const cookies = document.cookie.split(';');
    cookies.forEach(cookie => {
      const [name, value] = cookie.trim().split('=');
      if (name === 'session_id') {
        user.session_id = value;
        user.loggedIn = true;
      }
    });
    user.user_id = localStorage.getItem('user_id');
  });

  function formatTime(timestamp) {
    const optionsTime = { hour: 'numeric', minute: 'numeric', hour12: true };
    const date = new Date(timestamp);
    return date.toLocaleString('en-US', optionsTime);
  }

  function formatDate(timestamp) {
    const optionsDate = { day: 'numeric', month: 'long', year: 'numeric' };
    const date = new Date(timestamp);
    return date.toLocaleString('en-US', optionsDate);
  }
</script>

<Navbar />
<main class="min-h-[80vh] mx-44 mt-24">
  <!-- Tambah Item Todo -->
  <AddItem isOpen={isAddItemOpen} session_id={user.session_id} on:close={closeModalAddItem} />
  <!-- Edit Item Todo -->
  <EditItem
    isOpen={isEditItemOpen}
    session_id={user.session_id}
    currentText={user.current_text}
    todoId={user.todo_id}
    on:close={closeModalEditItem}
  />

  <div class="flex flex-col w-9/12 space-y-3 justify-center mx-auto">
    {#if user.loggedIn}
      <button
        on:click={openModalAddItem}
        class="py-2 mx-auto px-auto rounded-xl text-lg text-zinc-50 font-semibold bg-gradient-to-br from-sky-500 hover:from-sky-600 to-fuchsia-500 hover:to-fuchsia-600 w-full"
        >New Todo +</button
      >
    {/if}
    {#if Todos.length > 0}
      {#each Todos as todo}
        <div
          class="flex flex-row justify-between items-start space-x-3 rounded-xl shadow-md py-3 px-5 border border-zinc-300"
        >
          <div class="flex flex-col space-y-1.5 grow">
            <h3 class="text-2xl font-medium">{todo.text}</h3>
            <!-- User and Updated time -->
            <div class="flex flex-row space-x-2.5 text-xs">
              <!-- User -->
              <div class="flex flex-row space-x-1.5">
                <svg viewBox="0 0 24 24" class="fill-current w-[14px] h-auto">
                  <path
                    d="M12 12.75C8.83 12.75 6.25 10.17 6.25 7C6.25 3.83 8.83 1.25 12 1.25C15.17 1.25 17.75 3.83 17.75 7C17.75 10.17 15.17 12.75 12 12.75ZM12 2.75C9.66 2.75 7.75 4.66 7.75 7C7.75 9.34 9.66 11.25 12 11.25C14.34 11.25 16.25 9.34 16.25 7C16.25 4.66 14.34 2.75 12 2.75Z"
                  />
                  <path
                    d="M20.5901 22.75C20.1801 22.75 19.8401 22.41 19.8401 22C19.8401 18.55 16.3202 15.75 12.0002 15.75C7.68015 15.75 4.16016 18.55 4.16016 22C4.16016 22.41 3.82016 22.75 3.41016 22.75C3.00016 22.75 2.66016 22.41 2.66016 22C2.66016 17.73 6.85015 14.25 12.0002 14.25C17.1502 14.25 21.3401 17.73 21.3401 22C21.3401 22.41 21.0001 22.75 20.5901 22.75Z"
                  />
                </svg>
                <span>{todo.username}</span>
              </div>
              <!-- Date -->
              <div class="flex flex-row space-x-1.5">
                <svg viewBox="0 0 24 24" class="fill-current w-[14px] h-auto">
                  <path
                    d="M8 5.75C7.59 5.75 7.25 5.41 7.25 5V2C7.25 1.59 7.59 1.25 8 1.25C8.41 1.25 8.75 1.59 8.75 2V5C8.75 5.41 8.41 5.75 8 5.75Z"
                  />
                  <path
                    d="M16 5.75C15.59 5.75 15.25 5.41 15.25 5V2C15.25 1.59 15.59 1.25 16 1.25C16.41 1.25 16.75 1.59 16.75 2V5C16.75 5.41 16.41 5.75 16 5.75Z"
                  />
                  <path
                    d="M20.5 9.83984H3.5C3.09 9.83984 2.75 9.49984 2.75 9.08984C2.75 8.67984 3.09 8.33984 3.5 8.33984H20.5C20.91 8.33984 21.25 8.67984 21.25 9.08984C21.25 9.49984 20.91 9.83984 20.5 9.83984Z"
                  />
                  <path
                    d="M16 22.75H8C4.35 22.75 2.25 20.65 2.25 17V8.5C2.25 4.85 4.35 2.75 8 2.75H16C19.65 2.75 21.75 4.85 21.75 8.5V17C21.75 20.65 19.65 22.75 16 22.75ZM8 4.25C5.14 4.25 3.75 5.64 3.75 8.5V17C3.75 19.86 5.14 21.25 8 21.25H16C18.86 21.25 20.25 19.86 20.25 17V8.5C20.25 5.64 18.86 4.25 16 4.25H8Z"
                  />
                  <path
                    d="M8.5 14.5C8.37 14.5 8.24 14.47 8.12 14.42C8 14.37 7.89001 14.3 7.79001 14.21C7.70001 14.11 7.62999 14 7.57999 13.88C7.52999 13.76 7.5 13.63 7.5 13.5C7.5 13.24 7.61001 12.98 7.79001 12.79C7.89001 12.7 8 12.63 8.12 12.58C8.3 12.5 8.50001 12.48 8.70001 12.52C8.76001 12.53 8.82 12.55 8.88 12.58C8.94 12.6 9 12.63 9.06 12.67C9.11 12.71 9.15999 12.75 9.20999 12.79C9.24999 12.84 9.29999 12.89 9.32999 12.94C9.36999 13 9.40001 13.06 9.42001 13.12C9.45001 13.18 9.47001 13.24 9.48001 13.3C9.49001 13.37 9.5 13.43 9.5 13.5C9.5 13.76 9.38999 14.02 9.20999 14.21C9.01999 14.39 8.76 14.5 8.5 14.5Z"
                  />
                  <path
                    d="M12 14.4999C11.74 14.4999 11.48 14.3899 11.29 14.2099C11.25 14.1599 11.21 14.1099 11.17 14.0599C11.13 13.9999 11.1 13.9399 11.08 13.8799C11.05 13.8199 11.03 13.7599 11.02 13.6999C11.01 13.6299 11 13.5699 11 13.4999C11 13.3699 11.03 13.2399 11.08 13.1199C11.13 12.9999 11.2 12.8899 11.29 12.7899C11.57 12.5099 12.02 12.4199 12.38 12.5799C12.51 12.6299 12.61 12.6999 12.71 12.7899C12.89 12.9799 13 13.2399 13 13.4999C13 13.5699 12.99 13.6299 12.98 13.6999C12.97 13.7599 12.95 13.8199 12.92 13.8799C12.9 13.9399 12.87 13.9999 12.83 14.0599C12.79 14.1099 12.75 14.1599 12.71 14.2099C12.61 14.2999 12.51 14.3699 12.38 14.4199C12.26 14.4699 12.13 14.4999 12 14.4999Z"
                  />
                  <path
                    d="M8.5 17.9999C8.37 17.9999 8.24 17.9699 8.12 17.9199C8 17.8699 7.89001 17.7999 7.79001 17.7099C7.70001 17.6099 7.62999 17.5099 7.57999 17.3799C7.52999 17.2599 7.5 17.1299 7.5 16.9999C7.5 16.7399 7.61001 16.4799 7.79001 16.2899C7.89001 16.1999 8 16.1299 8.12 16.0799C8.49 15.9199 8.92999 16.0099 9.20999 16.2899C9.24999 16.3399 9.29999 16.3899 9.32999 16.4399C9.36999 16.4999 9.40001 16.5599 9.42001 16.6199C9.45001 16.6799 9.47001 16.7399 9.48001 16.8099C9.49001 16.8699 9.5 16.9399 9.5 16.9999C9.5 17.2599 9.38999 17.5199 9.20999 17.7099C9.01999 17.8899 8.76 17.9999 8.5 17.9999Z"
                  />
                </svg>
                <p>{formatDate(todo.updated_at)}</p>
              </div>

              <!-- Time -->
              <div class="flex flex-row space-x-1.5">
                <svg viewBox="0 0 24 24" class="fill-current w-[14px] h-auto">
                  <path
                    d="M12 22.75C6.07 22.75 1.25 17.93 1.25 12C1.25 6.07 6.07 1.25 12 1.25C17.93 1.25 22.75 6.07 22.75 12C22.75 17.93 17.93 22.75 12 22.75ZM12 2.75C6.9 2.75 2.75 6.9 2.75 12C2.75 17.1 6.9 21.25 12 21.25C17.1 21.25 21.25 17.1 21.25 12C21.25 6.9 17.1 2.75 12 2.75Z"
                  />
                  <path
                    d="M15.7096 15.93C15.5796 15.93 15.4496 15.9 15.3296 15.82L12.2296 13.97C11.4596 13.51 10.8896 12.5 10.8896 11.61V7.51001C10.8896 7.10001 11.2296 6.76001 11.6396 6.76001C12.0496 6.76001 12.3896 7.10001 12.3896 7.51001V11.61C12.3896 11.97 12.6896 12.5 12.9996 12.68L16.0996 14.53C16.4596 14.74 16.5696 15.2 16.3596 15.56C16.2096 15.8 15.9596 15.93 15.7096 15.93Z"
                  />
                </svg>
                <p>{formatTime(todo.updated_at)}</p>
              </div>
            </div>
          </div>

          {#if todo.user_id == user.user_id}
            <div class="flex flex-row space-x-2 w-fit text-zinc-50 text-sm">
              <button
                on:click={openModalEditItem(todo.text, todo.id)}
                class="py-1 px-3 bg-gradient-to-br from-yellow-400 to-amber-600 hover:from-yellow-500 hover:to-amber-600 rounded-lg flex flex-row space-x-2 items-center"
              >
                <svg viewBox="0 0 24 24" class="w-[17px] h-auto fill-current">
                  <path
                    d="M5.53999 19.5201C4.92999 19.5201 4.35999 19.31 3.94999 18.92C3.42999 18.43 3.17999 17.69 3.26999 16.89L3.63999 13.65C3.70999 13.04 4.07999 12.23 4.50999 11.79L12.72 3.10005C14.77 0.930049 16.91 0.870049 19.08 2.92005C21.25 4.97005 21.31 7.11005 19.26 9.28005L11.05 17.97C10.63 18.42 9.84999 18.84 9.23999 18.9401L6.01999 19.49C5.84999 19.5 5.69999 19.5201 5.53999 19.5201ZM15.93 2.91005C15.16 2.91005 14.49 3.39005 13.81 4.11005L5.59999 12.8101C5.39999 13.0201 5.16999 13.5201 5.12999 13.8101L4.75999 17.05C4.71999 17.38 4.79999 17.65 4.97999 17.82C5.15999 17.99 5.42999 18.05 5.75999 18L8.97999 17.4501C9.26999 17.4001 9.74999 17.14 9.94999 16.93L18.16 8.24005C19.4 6.92005 19.85 5.70005 18.04 4.00005C17.24 3.23005 16.55 2.91005 15.93 2.91005Z"
                  />
                  <path
                    d="M17.3404 10.9498C17.3204 10.9498 17.2904 10.9498 17.2704 10.9498C14.1504 10.6398 11.6404 8.26985 11.1604 5.16985C11.1004 4.75985 11.3804 4.37985 11.7904 4.30985C12.2004 4.24985 12.5804 4.52985 12.6504 4.93985C13.0304 7.35985 14.9904 9.21985 17.4304 9.45985C17.8404 9.49985 18.1404 9.86985 18.1004 10.2798C18.0504 10.6598 17.7204 10.9498 17.3404 10.9498Z"
                  />
                  <path
                    d="M21 22.75H3C2.59 22.75 2.25 22.41 2.25 22C2.25 21.59 2.59 21.25 3 21.25H21C21.41 21.25 21.75 21.59 21.75 22C21.75 22.41 21.41 22.75 21 22.75Z"
                  />
                </svg>
                <span>Edit</span>
              </button>
              <button
                on:click={deleteTodos(todo.id)}
                class="py-1 px-3 bg-gradient-to-br from-rose-400 to-red-500 hover:from-rose-500 hover:to-red-600 rounded-lg flex flex-row space-x-2 items-center"
              >
                <svg viewBox="0 0 24 24" class="w-[17px] h-auto fill-current">
                  <path
                    d="M20.9997 6.72998C20.9797 6.72998 20.9497 6.72998 20.9197 6.72998C15.6297 6.19998 10.3497 5.99998 5.11967 6.52998L3.07967 6.72998C2.65967 6.76998 2.28967 6.46998 2.24967 6.04998C2.20967 5.62998 2.50967 5.26998 2.91967 5.22998L4.95967 5.02998C10.2797 4.48998 15.6697 4.69998 21.0697 5.22998C21.4797 5.26998 21.7797 5.63998 21.7397 6.04998C21.7097 6.43998 21.3797 6.72998 20.9997 6.72998Z"
                  />
                  <path
                    d="M8.50074 5.72C8.46074 5.72 8.42074 5.72 8.37074 5.71C7.97074 5.64 7.69074 5.25 7.76074 4.85L7.98074 3.54C8.14074 2.58 8.36074 1.25 10.6907 1.25H13.3107C15.6507 1.25 15.8707 2.63 16.0207 3.55L16.2407 4.85C16.3107 5.26 16.0307 5.65 15.6307 5.71C15.2207 5.78 14.8307 5.5 14.7707 5.1L14.5507 3.8C14.4107 2.93 14.3807 2.76 13.3207 2.76H10.7007C9.64074 2.76 9.62074 2.9 9.47074 3.79L9.24074 5.09C9.18074 5.46 8.86074 5.72 8.50074 5.72Z"
                  />
                  <path
                    d="M15.2104 22.7501H8.79039C5.30039 22.7501 5.16039 20.8201 5.05039 19.2601L4.40039 9.19007C4.37039 8.78007 4.69039 8.42008 5.10039 8.39008C5.52039 8.37008 5.87039 8.68008 5.90039 9.09008L6.55039 19.1601C6.66039 20.6801 6.70039 21.2501 8.79039 21.2501H15.2104C17.3104 21.2501 17.3504 20.6801 17.4504 19.1601L18.1004 9.09008C18.1304 8.68008 18.4904 8.37008 18.9004 8.39008C19.3104 8.42008 19.6304 8.77007 19.6004 9.19007L18.9504 19.2601C18.8404 20.8201 18.7004 22.7501 15.2104 22.7501Z"
                  />
                  <path
                    d="M13.6601 17.25H10.3301C9.92008 17.25 9.58008 16.91 9.58008 16.5C9.58008 16.09 9.92008 15.75 10.3301 15.75H13.6601C14.0701 15.75 14.4101 16.09 14.4101 16.5C14.4101 16.91 14.0701 17.25 13.6601 17.25Z"
                  />
                  <path
                    d="M14.5 13.25H9.5C9.09 13.25 8.75 12.91 8.75 12.5C8.75 12.09 9.09 11.75 9.5 11.75H14.5C14.91 11.75 15.25 12.09 15.25 12.5C15.25 12.91 14.91 13.25 14.5 13.25Z"
                  />
                </svg>
                <span>Delete</span>
              </button>
            </div>
          {/if}
        </div>
      {/each}
    {:else}
      <div
        class="w-full h-fit flex justify-center py-20 px-auto border-[2px] border-dashed border-zinc-400 bg-zinc-200/90 rounded-xl shadow-lg"
      >
        <p class="text-lg font-bold italic text-zinc-600">No Content</p>
      </div>
    {/if}
  </div>
</main>
<Footer />
