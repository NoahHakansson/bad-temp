<script>
  import TempList from "./lib/TempList.svelte";
  import { Container } from "sveltestrap";

  const apiUrl = "http://localhost:5000/badtemp"; // local
  // const apiUrl = "https://api.yourdomain.com/badtemp" // prod
</script>

<Container>
  <h1>Badtemperaturer Kalmar</h1>

  <TempList endpoint={apiUrl} />
</Container>

<style>
  h1 {
    text-align: center;
    font-size: 8vw;
    margin: 45px;
  }
  @media screen and (min-width: 800px) {
    h1 {
      font-size: 65px;
    }
  }
</style>
