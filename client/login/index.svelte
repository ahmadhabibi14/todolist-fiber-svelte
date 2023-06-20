<script>
   let data = {
      username: "",
      password: ""
   }
   const handleSubmit = e => {
      // Getting the action URL
      const actionURL = e.target.action;

      console.log(data);
      console.log(actionURL);

      data.username = ""
      data.password = ""
   }
</script>

<div class="w-full h-fit flex flex-col items-center justify-center pt-16">
   <form
      action="localhost:3000/api/login"
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