<script>
   import { onMount } from "svelte";
   import Navbar from "../_components/navbar.svelte";
   import Footer from "../_components/footer.svelte";

   let user = { loggedIn: false }

   onMount(() => {
      const cookies = document.cookie.split(";");
      cookies.forEach(cookie => {
         const [name, value] = cookie.trim().split("=");
         if (name === "session_id") {
            console.log("session_id:", value)
            user.loggedIn = true;
         }
      })
   });
</script>

<Navbar></Navbar>
<main class="min-h-[78vh] mx-44 mt-20">
   <div class="w-full h-fit flex flex-col items-center justify-center pt-16">
      {#if user.loggedIn}
      <form
         action="/api/change-password"
         class="flex flex-col space-y-4 justify-start w-5/12 p-5 border border-zinc-200 rounded-2xl shadow-lg"
      >
         <h2 class="text-3xl font-bold text-sky-600 text-center">Change Password</h2>
         <input
            class="border-zinc-400 border rounded-xl bg-transparent focus:border-blue-500 caret-blue-500 py-2 px-4 outline-0"
            type="password"
            id="password"
            label="password"
            placeholder="Password"
            required
         />
         <button class="shadow-md hover:shadow-xl rounded-xl py-2 px-6 text-zinc-50 font-medium bg-gradient-to-br from-sky-500 hover:from-sky-600 to-fuchsia-500 hover:to-fuchsia-600" type="submit">
            Submit
         </button>
      </form>
      {:else}
         <div>
            <h2 class="text-3xl font-bold text-sky-600 text-center">You have to <a href="/login" class="text-sky-600 hover:underline">login</a> first</h2>
         </div>
      {/if}
   </div>
</main>
<Footer></Footer>