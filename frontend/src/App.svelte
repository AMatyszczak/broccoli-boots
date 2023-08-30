<script>
  import * as rt from "../wailsjs/runtime/runtime.js"; // the runtime for Wails2

  import { LayerCake, Svg } from "layercake";
  import Line from "./components/Line.svelte";
  import AxisX from "./components/AxisX.svelte";
  import AxisY from "./components/AxisY.svelte";

  import { SelectFile, FetchOperationsForFilter } from "../wailsjs/go/main/App";
  import { scaleTime } from "d3-scale";

  function selectOperationsStatements(event) {
    SelectFile().catch((err) => {
      console.error("error:", err);
    });
  }

  let allPoints = [];
  let tableOperations = [];

  FetchOperationsForFilter({ operationType: "type" }).then((operations) => {
    const points = [];
    tableOperations = operations;
    console.log("tableOperations:", tableOperations);
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

  rt.EventsOn("operations-loaded", (message) => {
    FetchOperationsForFilter({ operationType: "type" }).then((operations) => {
      const points = [];
      tableOperations = operations;
      console.log("tableOperations:", tableOperations);
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
  <meta charset="UTF-8" />
  <label for="many" on:click={selectOperationsStatements}>Upload</label>
  <div class="charts-container">
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
  <div class="table-container">
    <!-- Native Table Element -->
    <table class="table table-hover">
      <thead>
        <tr>
          <th>Position</th>
          <th>Name</th>
          <th>Symbol</th>
        </tr>
      </thead>
      <tbody>
        {#each tableOperations as operation, i}
          <tr>
            <td>{i}</td>
            <td>{operation["Data operacji"]}</td>
            <td>{operation["Data waluty"]}</td>
            <td>{operation["Kwota"]}</td>
            <td>{operation["Opis transakcji"]}</td>
            <td>{operation["Saldo po transakcji"]}</td>
            <td>{operation["Typ transakcji"]}</td>
          </tr>
        {/each}
      </tbody>
      <tfoot>
        <tr>
          <th colspan="3">Calculated Total Weight</th>
          <td>10</td>
        </tr>
      </tfoot>
    </table>
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
