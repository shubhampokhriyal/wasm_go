(async function () {
  // Step 0: Registering web assembly
  const go = new Go();
  const result = await WebAssembly.instantiateStreaming(fetch("lib.wasm"), go.importObject)
  go.run(result.instance);

  document.getElementById("loader").classList.add('hide');
  document.getElementById("main").classList.remove('hide');

  // Step 1: Default hashed value
  let hashedValue = "";

  // Step 2: Setting up event listeners
  document.getElementById('check-button').addEventListener('click', () => {
    const str = document.getElementById('check').value;
    hashedValue = hash(str);
    document.getElementById('hashed-value').innerText = hashedValue;
  });
  document.getElementById('verify-button').addEventListener('click', () => {
    const str = document.getElementById('verify').value;
    const matched = verify(hashedValue, str)
    document.getElementById('verified-value').innerText = matched;
  });
})()