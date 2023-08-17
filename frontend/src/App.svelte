<script>
  import logo from "./assets/images/logo-universal.png";

  let resultText = "Please upload your expenses 👇";
  let files;

  $: if (files) {
    // Note that `files` is of type `FileList`, not an Array:
    // https://developer.mozilla.org/en-US/docs/Web/API/FileList
    console.log(files);

    for (const file of files) {
      console.log(`${file.name}: ${file.size} bytes`);
    }
  }
</script>

<main>
  <img alt="Wails logo" id="logo" src={logo} />
  <div class="result" id="result">{resultText}</div>
  <div class="input-box" id="input">
    <label for="many">Upload multiple files of any type:</label>
    <input bind:files id="many" multiple type="file" />
  </div>
</main>

<style>
  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }
</style>
