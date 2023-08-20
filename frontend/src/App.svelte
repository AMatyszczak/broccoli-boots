<script>
  import { LayerCake, Svg } from "layercake";
  import Line from "./components/Line.svelte";
  import AxisX from "./components/AxisX.svelte";
  import AxisY from "./components/AxisY.svelte";

  import { SelectFile } from "../wailsjs/go/main/App";

  function doIt(event) {
    try {
      SelectFile()
        .then((result) => {
          console.log("result:", result);
        })
        .catch((err) => {
          console.error("error:", err);
        });
    } catch (err) {
      console.error(err);
    }
  }

  const allPoints = [
    {
      points: [
        { x: 0, y: 0 },
        { x: 5, y: 10 },
        { x: 10, y: 20 },
        { x: 15, y: 30 },
        { x: 20, y: 3 },
      ],
    },
    {
      points: [
        { x: 0, y: 0 },
        { x: 5, y: 10 },
        { x: 10, y: 20 },
        { x: 15, y: 20 },
        { x: 20, y: 10 },
      ],
    },
  ];
</script>

<main>
  <div class="charts-container">
    {#each allPoints as { points }, i}
      <label>{i}</label>
      <div class="chart-container">
        <LayerCake x="x" y="y" data={points}>
          <Svg>
            <AxisX />
            <AxisY />
            <Line stroke="#abaaaa" />
          </Svg>
        </LayerCake>
      </div>
    {/each}
  </div>
  <div id="lul" />
  <div class="input-box" id="input">
    <label for="many" on:click={doIt}>Upload</label>
  </div>
</main>

<style>
  .charts-container {
    padding: 2em;
  }
  .chart-container {
    height: 500px;
    background-color: rgba(10, 38, 54, 1);
    margin: 1em;
  }
</style>
