<script lang="ts">
  import { onMount } from "svelte";
  import { Card, Col, Badge, Row } from "sveltestrap";
  import axios from "axios";
  export let endpoint = "";

  let temps = [];

  const options = {
    method: "GET",
    url: endpoint,
  };

  onMount(async () => {
    axios
      .request(options)
      .then(function (response) {
        temps = response.data;
        // assign a name and time
        temps.forEach((temp) => {
          let name = temp.id.split(":")[2];
          temp["name"] = name;
          let time = temp.dateObserved.value.split("T");
          let date = time[0];
          time = time[1];
          time = time.split(":", 2);
          time = `${time[0]}:${time[1]}`;
          temp["time"] = time;
          temp["date"] = date;
        });
      })
      .catch(function (error) {
        console.error(error);
      });
  });
</script>

{#if temps.length > 0}
  <Row cols={2}>
    {#each temps as temp (temp.id)}
      <Col>
        <Card body class="mb-3">
          <h3>
            {temp.name}
          </h3>
          <p>
            Temperatur: {temp.temperature.value}°C
            <br />Tid: {temp.time}, {temp.date}
          </p>
        </Card>
      </Col>
    {/each}
  </Row>
{:else}
  <p style="text-align: center;">Data kunde inte hämtas...</p>
{/if}

<style>
  h3 {
    font-size: 5vw;
  }
  @media screen and (min-width: 800px) {
    h3 {
      font-size: 30px;
    }
  }
  @media screen and (max-width: 400px) {
    h3 {
      font-size: 20px;
    }
  }
</style>
