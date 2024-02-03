<script lang="ts">
  let currentInput:string = '';
  let result:string = ''

  const setValue = (value:number) => {
    if(currentInput.length !== 2){
      currentInput += value
    }
  }

  const resetValueC = () => {
    currentInput = ''
    result = ''
  }

  const resetValueCE = () => {
    currentInput = ''
  }
</script>

<div class="m-4 border-8 border-black w-[20rem] h-[28rem] flex flex-col items-center">
  <div class="bg-gray flex-grow mb-3 w-[304px] text-center">
    <h1 class="text-2xl font-bold font-serif">Calculadora Basica</h1>
  </div>
  <input type="text" name="result" bind:value={result} readonly disabled
  class="bg-black text-white text-right w-[230px] text-sm h-4"/>
  <input type="text" name="result" bind:value={currentInput} readonly disabled
  class="bg-black text-white text-right w-[230px] text-lg font-bold"/>
  <div class="buttons">
    <div>
      <button class="bg-orange text-white font-bold" on:click={resetValueC}>C</button>
      <button class="bg-orange text-white font-bold" on:click={resetValueCE}>CE</button>
      <button class="bg-bordo text-white font-bold">+/-</button>
      <button class="bg-bordo text-white font-bold">%</button>
    </div>
    <div>
      <button class="bg-blue text-white font-bold" on:click={()=> setValue(7)}>7</button>
      <button class="bg-blue text-white font-bold" on:click={()=> setValue(8)}>8</button>
      <button class="bg-blue text-white font-bold" on:click={()=> setValue(9)}>9</button>
      <button class="bg-bordo text-white font-bold">/</button>
    </div>
    <div>
      <button class="bg-blue text-white font-bold" on:click={()=> setValue(4)}>4</button>
      <button class="bg-blue text-white font-bold" on:click={()=> setValue(5)}>5</button>
      <button class="bg-blue text-white font-bold" on:click={()=> setValue(6)}>6</button>
      <button class="bg-bordo text-white font-bold">X</button>
    </div>
    <div>
      <button class="bg-blue text-white font-bold" on:click={()=> setValue(1)}>1</button>
      <button class="bg-blue text-white font-bold" on:click={()=> setValue(2)}>2</button>
      <button class="bg-blue text-white font-bold" on:click={()=> setValue(3)}>3</button>
      <button class="bg-bordo text-white font-bold">-</button>
    </div>
    <div>
      <button class="bg-blue text-white font-bold" on:click={()=> setValue(0)}>0</button>
      <button class="bg-blue text-white font-bold">.</button>
      <button class="bg-green text-white font-bold">=</button>
      <button class="bg-bordo text-white font-bold">+</button>
    </div>
  </div>
</div>

<style>
  .buttons button{
    @apply w-[52.5px] h-[52.5px] mr-[1px] ml-[1px] mt-[2px] mb-[2px] text-lg;
  }

  .buttons{
    @apply w-[240px] flex flex-col items-center grow;
  }
</style>