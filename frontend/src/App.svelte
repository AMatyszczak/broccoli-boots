<script>
  // import LayerCake from "./layercake/src/LayerCake.svelte";

  import { LayerCake, Svg, Canvas, Html } from "layercake";
  import Scatter from "./components/Scatter.svg.svelte";
  import Line from "./components/Line.svelte";
  import AxisX from "./components/AxisX.svelte";
  import AxisY from "./components/AxisY.svelte";

  let files;

  $: if (files) {
    // Note that `files` is of type `FileList`, not an Array:
    // https://developer.mozilla.org/en-US/docs/Web/API/FileList
    console.log(files);

    for (const file of files) {
      console.log(`${file.name}: ${file.size} bytes`);
    }
  }

  // Define some data
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
      <label for="many">{i}</label>
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
  <div class="input-box" id="input">
    <label for="many">Upload multiple files of any type:</label>
    <input bind:files id="many" multiple type="file" />
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
