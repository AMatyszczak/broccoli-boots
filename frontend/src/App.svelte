<script>
  import * as rt from "../wailsjs/runtime/runtime.js"; // the runtime for Wails2

  import { LayerCake, Svg } from "layercake";
  import Line from "./components/Line.svelte";
  import AxisX from "./components/AxisX.svelte";
  import AxisY from "./components/AxisY.svelte";

  import { SelectFile, FetchOperationsForFilter } from "../wailsjs/go/main/App";

  function selectOperationsStatements(event) {
    SelectFile().catch((err) => {
      console.error("error:", err);
    });
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
  rt.EventsOn("operations-loaded", (message) => {
    FetchOperationsForFilter({ operationType: "type" }).then((result) => {
      console.log("FetchOperationsForFilter:", result);
    });
  });
</script>

<main>
  <label for="many" on:click={selectOperationsStatements}>Upload</label>

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
