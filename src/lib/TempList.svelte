<script>
  import { onMount } from "svelte";
  import axios from "axios";
  export let endpoint = "";

  let tempData = {};

  const options = {
    method: "GET",
    url: endpoint,
  };

  onMount(async () => {
    axios
      .request(options)
      .then(function (response) {
        console.log(response.data);
      })
      .catch(function (error) {
        console.error(error);
      });
  });

  console.log(tempData);
</script>

<div class="main">test</div>

<style>
  .main {
    font-family: inherit;
    font-size: inherit;
  }
</style>
