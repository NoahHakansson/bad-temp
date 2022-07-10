<script>
  import logo from "./assets/svelte.png";
  import TempList from "./lib/TempList.svelte";
  import { Image, Navbar } from "sveltestrap";
  import { Col, Container, Row } from "sveltestrap";

  const apiUrl = "http://localhost:5000/badtemp" // local
  // const apiUrl = "http://yourdomain.example:5000/badtemp" // prod
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
