<script>
  import * as rt from "../wailsjs/runtime/runtime.js"; // the runtime for Wails2

  import { LayerCake, Svg } from "layercake";
  import Line from "./components/Line.svelte";
  import AxisX from "./components/AxisX.svelte";
  import AxisY from "./components/AxisY.svelte";

  import { SelectFile, FetchOperationsForFilter } from "../wailsjs/go/main/App";
  import { scaleBand, scaleTime } from "d3-scale";

  function selectOperationsStatements(event) {
    SelectFile().catch((err) => {
      console.error("error:", err);
    });
  }

  FetchOperationsForFilter({ operationType: "type" }).then((operations) => {
    const points = [];
    operations.forEach((operation) => {
      points.push({
        x: Date.parse(operation["Data operacji"]),
        y: parseFloat(operation["Kwota"]),
      });
    });
    allPoints.push({ points: points });
    allPoints = allPoints;
    console.log("allPoints:", allPoints);
  });

  let allPoints = [];
  rt.EventsOn("operations-loaded", (message) => {
    FetchOperationsForFilter({ operationType: "type" }).then((operations) => {
      const points = [];
      operations.forEach((operation) => {
        points.push({
          x: Date.parse(operation["Data operacji"]),
          y: parseFloat(operation["Kwota"]),
        });
      });
      allPoints.push({ points: points });
      allPoints = allPoints;
      console.log("allPoints:", allPoints);
    });
  });
</script>

<main>
  <label for="many" on:click={selectOperationsStatements}>Upload</label>
  <div class="charts-container">
    <div class="text-sky-500 dark:text-sky-400">Sarah Dayan</div>

    <button class="button border-orange-900 border">Greetings</button>

    {#each allPoints as { points }, i}
      <label>{i}</label>
      <div class="chart-container">
        <LayerCake x="x" y="y" data={points} xScale={scaleTime()}>
          <Svg>
            <AxisX formatTick={(d) => new Date(d).toLocaleDateString()} />
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
