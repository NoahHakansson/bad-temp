<script>
  import logo from "./assets/svelte.png";
  import TempList from "./lib/TempList.svelte";
  import { Image, Navbar } from "sveltestrap";
  import { Col, Container, Row } from "sveltestrap";

  const apiUrl = "http://localhost:5000/badtemp"; // local
  // const apiUrl = "http://localhost:5000/badtemp" // prod
</script>

<Container>
  <h1>Badtemperaturer Kalmar</h1>

  <TempList endpoint={apiUrl} />
</Container>

<style>
  h1 {
    text-align: center;
    font-size: 55px;
    padding: 45px;
  }
</style>
