<script>
   import { writable } from "svelte/store";
   // Create writable store for the loading state
   let isLoading = writable(false);
   let isError = writable(false);
   let errorMessage = writable("");

   let data = {
      username: "",
      password: ""
   }

   const handleSubmit = async (e) => {
      isLoading.set(true)
      const reqBody = JSON.stringify({
         username: data.username,
         password: data.password
      })
      const actionURL = e.target.action;
      try {
         const resp = await fetch(actionURL, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: reqBody
         });
         
         if (resp.ok) {
            const cookieHeader = resp.headers.get("Set-Cookie");
            document.cookie = cookieHeader;
         } else {
            const errorData = await resp.json();
            errorMessage.set(errorData.message || "Login Failed");
         }
      } catch (error) {
         // Handle any network or server errors
         console.error(error);
         isError.set(true)
         errorMessage.set("Unexpected error occurred during login")
      }
      data.username = ""
      data.password = ""
      isLoading.set(false)
   }
</script>

<div class="w-full h-fit flex flex-col items-center justify-center pt-16">
   <div class="{isLoading === true
      ?  'block absoulte inset-0 h-full w-full bg-zinc-800 bg-opacity-75'
      :  'hidden absoulte inset-0 h-full w-full bg-zinc-800 bg-opacity-75 blur-md'
   }">
      <div class="py-2 px-6 bg-gradient-to-br from-sky-500 to-fuchsia-500 rounded-xl">Loading...</div>
   </div>
   <form
      action="/api/login"
      on:submit|preventDefault={handleSubmit}
      class="flex flex-col space-y-4 justify-start w-5/12 p-5 border border-zinc-200 rounded-2xl shadow-lg"
      >
      <h2 class="text-3xl font-bold text-sky-600 text-center">Login</h2>
      <input
         class="border-zinc-400 border rounded-xl bg-transparent focus:border-blue-500 caret-blue-500 py-2 px-4 outline-0"
         type="text"
         id="username"
         label="username"
         placeholder="Username"
         bind:value={data.username}
         required
      />
      <input
         class="border-zinc-400 border rounded-xl bg-transparent focus:border-blue-500 caret-blue-500 py-2 px-4 outline-0"
         type="password"
         id="password"
         label="password"
         placeholder="Password"
         bind:value={data.password}
         required
      />
      <button class="shadow-md hover:shadow-xl rounded-xl py-2 px-6 text-zinc-50 font-medium bg-gradient-to-br from-sky-500 hover:from-sky-600 to-fuchsia-500 hover:to-fuchsia-600" type="submit">
         Login
      </button>
   </form>
</div>