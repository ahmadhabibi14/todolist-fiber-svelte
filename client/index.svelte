<script>
   import { onMount } from "svelte";
   import Navbar from "./_components/navbar.svelte";
   import Footer from "./_components/footer.svelte";

   let Todos = [];
   let user = {
      username: "",
      loggedIn: false
   }
   async function getTodos() {
      const resp = await fetch("/api/todo/list", { method: "GET" });
      Todos = await resp.json();
      console.log(Todos);
   }

   onMount(() => {
      getTodos();
      const cookies = document.cookie.split(";");
      cookies.forEach(cookie => {
         const [name, value] = cookie.trim().split("=");
         if (name === "session_id") {
            console.log("session_id:", value)
            user.loggedIn = true;
         }
      });
      user.username = localStorage.getItem("username");
   })

   function formatTime(timestamp) {
      const optionsTime = { hour: "numeric", minute: "numeric", hour12: true };
      const date = new Date(timestamp);
      return date.toLocaleString("en-US", optionsTime);
   }

   function formatDate(timestamp) {
      const optionsDate = { day: "numeric", month: "long", year: "numeric" };
      const date = new Date(timestamp);
      return date.toLocaleString("en-US", optionsDate);
   }
</script>

<Navbar></Navbar>
<main class="min-h-[80vh] mx-44 mt-20">

   {#if Todos.length > 0}
      <div class="flex flex-col w-9/12 space-y-3 justify-center mx-auto">
         {#if user.loggedIn}
            <button class="py-2 mx-auto px-auto rounded-xl text-lg text-zinc-50 font-semibold bg-gradient-to-br from-sky-500 hover:from-sky-600 to-fuchsia-500 hover:to-fuchsia-600 w-full">New Todo +</button>
         {/if}
         {#each Todos as todo}
            <div class="flex flex-row justify-between items-start space-y-2 rounded-xl shadow-md py-3 px-5 border border-zinc-300">
               <div class="flex flex-col space-y-1 grow">
                  <h3 class="text-2xl font-medium">{todo.text}</h3>
                  <p>{todo.username}</p>
                  <div class="text-xs flex flex-row space-x-1.5">
                     <svg viewBox="0 0 24 24" class="fill-current w-[14px] h-auto">
                        <path d="M12 22.75C6.07 22.75 1.25 17.93 1.25 12C1.25 6.07 6.07 1.25 12 1.25C17.93 1.25 22.75 6.07 22.75 12C22.75 17.93 17.93 22.75 12 22.75ZM12 2.75C6.9 2.75 2.75 6.9 2.75 12C2.75 17.1 6.9 21.25 12 21.25C17.1 21.25 21.25 17.1 21.25 12C21.25 6.9 17.1 2.75 12 2.75Z"/>
                        <path d="M15.7096 15.93C15.5796 15.93 15.4496 15.9 15.3296 15.82L12.2296 13.97C11.4596 13.51 10.8896 12.5 10.8896 11.61V7.51001C10.8896 7.10001 11.2296 6.76001 11.6396 6.76001C12.0496 6.76001 12.3896 7.10001 12.3896 7.51001V11.61C12.3896 11.97 12.6896 12.5 12.9996 12.68L16.0996 14.53C16.4596 14.74 16.5696 15.2 16.3596 15.56C16.2096 15.8 15.9596 15.93 15.7096 15.93Z"/>
                     </svg>
                     <p>{formatDate(todo.updated_at)} • {formatTime(todo.updated_at)}</p>
                  </div>
               </div>

               {#if todo.username === user.username}
                  <div class="flex flex-row space-x-2 w-fit text-zinc-50">
                     <button class="py-1 px-3 bg-gradient-to-br from-yellow-500 to-amber-600 rounded-lg flex flex-row space-x-2 items-center">
                        <svg viewBox="0 0 24 24" class="w-[18px] h-auto fill-current">
                           <path d="M5.53999 19.5201C4.92999 19.5201 4.35999 19.31 3.94999 18.92C3.42999 18.43 3.17999 17.69 3.26999 16.89L3.63999 13.65C3.70999 13.04 4.07999 12.23 4.50999 11.79L12.72 3.10005C14.77 0.930049 16.91 0.870049 19.08 2.92005C21.25 4.97005 21.31 7.11005 19.26 9.28005L11.05 17.97C10.63 18.42 9.84999 18.84 9.23999 18.9401L6.01999 19.49C5.84999 19.5 5.69999 19.5201 5.53999 19.5201ZM15.93 2.91005C15.16 2.91005 14.49 3.39005 13.81 4.11005L5.59999 12.8101C5.39999 13.0201 5.16999 13.5201 5.12999 13.8101L4.75999 17.05C4.71999 17.38 4.79999 17.65 4.97999 17.82C5.15999 17.99 5.42999 18.05 5.75999 18L8.97999 17.4501C9.26999 17.4001 9.74999 17.14 9.94999 16.93L18.16 8.24005C19.4 6.92005 19.85 5.70005 18.04 4.00005C17.24 3.23005 16.55 2.91005 15.93 2.91005Z"/>
                           <path d="M17.3404 10.9498C17.3204 10.9498 17.2904 10.9498 17.2704 10.9498C14.1504 10.6398 11.6404 8.26985 11.1604 5.16985C11.1004 4.75985 11.3804 4.37985 11.7904 4.30985C12.2004 4.24985 12.5804 4.52985 12.6504 4.93985C13.0304 7.35985 14.9904 9.21985 17.4304 9.45985C17.8404 9.49985 18.1404 9.86985 18.1004 10.2798C18.0504 10.6598 17.7204 10.9498 17.3404 10.9498Z"/>
                           <path d="M21 22.75H3C2.59 22.75 2.25 22.41 2.25 22C2.25 21.59 2.59 21.25 3 21.25H21C21.41 21.25 21.75 21.59 21.75 22C21.75 22.41 21.41 22.75 21 22.75Z"/>
                        </svg>                                                        
                        <span>Edit</span>
                     </button>
                     <button class="py-1 px-3 bg-gradient-to-br from-rose-400 to-red-500 rounded-lg flex flex-row space-x-2 items-center">
                        <svg viewBox="0 0 24 24" class="w-[18px] h-auto fill-current">
                           <path d="M20.9997 6.72998C20.9797 6.72998 20.9497 6.72998 20.9197 6.72998C15.6297 6.19998 10.3497 5.99998 5.11967 6.52998L3.07967 6.72998C2.65967 6.76998 2.28967 6.46998 2.24967 6.04998C2.20967 5.62998 2.50967 5.26998 2.91967 5.22998L4.95967 5.02998C10.2797 4.48998 15.6697 4.69998 21.0697 5.22998C21.4797 5.26998 21.7797 5.63998 21.7397 6.04998C21.7097 6.43998 21.3797 6.72998 20.9997 6.72998Z"/>
                           <path d="M8.50074 5.72C8.46074 5.72 8.42074 5.72 8.37074 5.71C7.97074 5.64 7.69074 5.25 7.76074 4.85L7.98074 3.54C8.14074 2.58 8.36074 1.25 10.6907 1.25H13.3107C15.6507 1.25 15.8707 2.63 16.0207 3.55L16.2407 4.85C16.3107 5.26 16.0307 5.65 15.6307 5.71C15.2207 5.78 14.8307 5.5 14.7707 5.1L14.5507 3.8C14.4107 2.93 14.3807 2.76 13.3207 2.76H10.7007C9.64074 2.76 9.62074 2.9 9.47074 3.79L9.24074 5.09C9.18074 5.46 8.86074 5.72 8.50074 5.72Z"/>
                           <path d="M15.2104 22.7501H8.79039C5.30039 22.7501 5.16039 20.8201 5.05039 19.2601L4.40039 9.19007C4.37039 8.78007 4.69039 8.42008 5.10039 8.39008C5.52039 8.37008 5.87039 8.68008 5.90039 9.09008L6.55039 19.1601C6.66039 20.6801 6.70039 21.2501 8.79039 21.2501H15.2104C17.3104 21.2501 17.3504 20.6801 17.4504 19.1601L18.1004 9.09008C18.1304 8.68008 18.4904 8.37008 18.9004 8.39008C19.3104 8.42008 19.6304 8.77007 19.6004 9.19007L18.9504 19.2601C18.8404 20.8201 18.7004 22.7501 15.2104 22.7501Z"/>
                           <path d="M13.6601 17.25H10.3301C9.92008 17.25 9.58008 16.91 9.58008 16.5C9.58008 16.09 9.92008 15.75 10.3301 15.75H13.6601C14.0701 15.75 14.4101 16.09 14.4101 16.5C14.4101 16.91 14.0701 17.25 13.6601 17.25Z"/>
                           <path d="M14.5 13.25H9.5C9.09 13.25 8.75 12.91 8.75 12.5C8.75 12.09 9.09 11.75 9.5 11.75H14.5C14.91 11.75 15.25 12.09 15.25 12.5C15.25 12.91 14.91 13.25 14.5 13.25Z"/>
                        </svg>                           
                        <span>Delete</span>
                     </button>
                  </div>
               {/if}
            </div>
         {/each}
      </div>
   {:else}
      <p>NO Content</p>
   {/if}

</main>
<Footer></Footer>