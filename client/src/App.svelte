<script lang="ts">
  

  let currentInput:string = ''
  let secondInput:string = ''
  let result:string = ''

  const setValue = (value:number) => {
    if(result !== ''){
      currentInput = ''
      result = ''
    }
    if((currentInput.length !== 2 || currentInput.includes('.')) && currentInput.length <= 6){
      currentInput += value
    }
  }

  const changeValue = () => {
    let numberInput:number = Number(currentInput)

    currentInput = (numberInput * -1).toString()
  }

  const handleModulus = () => {
    if (result !== '') {
      currentInput = ''
      result = ''
    }
    if (currentInput !== '' && secondInput === '') {
      secondInput = currentInput + ' %'
      currentInput = ''
    }
  }

  const handleDivision = () => {
    if (result !== '') {
      currentInput = ''
      result = ''
    }
    if (currentInput !== '' && secondInput === '') {
      secondInput = currentInput + ' /'
      currentInput = ''
    }
  }

  const handleMultiplication = () => {
    if (result !== '') {
      currentInput = ''
      result = ''
    }
    if (currentInput !== '' && secondInput === '') {
      secondInput = currentInput + ' *'
      currentInput = ''
    }
    }

  const handleSubtraction = () => {
    if (result !== '') {
      currentInput = ''
      result = ''
    }
    if (currentInput !== '' && secondInput === '') {
      secondInput = currentInput + '-'
      currentInput = ''
    }
  }

  const handleAddition = () => {
    if (result !== '') {
      currentInput = ''
      result = ''
    }
    if (currentInput !== '' && secondInput === '') {
      secondInput = currentInput + ' +'
      currentInput = ''
    }
  }

  const handleDecimal = () => {
    if (result !== '') {
      currentInput = ''
      result = ''
    }
    if (currentInput !== '' && currentInput.length < 3 && !currentInput.includes('.')) {
      currentInput += '.'
    }
  }

  const handleEqual = () => {
    if (currentInput !== '' && secondInput !== '') {
      let concatenation = `${secondInput}  ${currentInput}`
      result = eval(concatenation)

      const currentDate = new Date()
      const options = { timeZone: 'America/Argentina/Buenos_Aires' }
      const argentinaDate = currentDate.toLocaleString('es-AR', options)

      secondInput = ''
      currentInput = result
    }
  }

  const resetValueC = () => {
    currentInput = ''
    secondInput = ''
  }

  const resetValueCE = () => {
    currentInput = ''
  }
</script>

<div class="m-4 border-8 border-black w-[20rem] h-[28rem] flex flex-col items-center">

  <div class="bg-gray flex-grow mb-3 w-[304px] text-center">
    <h1 class="text-2xl font-bold font-serif">Calculadora Basica</h1>
  </div>

  <input 
  type="text" 
  name="result" 
  bind:value={secondInput}
  readonly 
  disabled
  class="bg-black text-white text-right w-[230px] text-sm h-4"
  />

  <input 
  type="text" 
  name="result" 
  bind:value={currentInput} 
  readonly 
  disabled
  class="bg-black text-white text-right w-[230px] text-lg font-bold"
  />
  
  <div class="buttons">
    <div>
      <button class="bg-orange" on:click={resetValueC}>C</button>
      <button class="bg-orange" on:click={resetValueCE}>CE</button>
      <button class="bg-bordo" on:click={changeValue}>+/-</button>
      <button class="bg-bordo" on:click={handleModulus}>%</button>
    </div>
    <div>
      <button class="bg-blue" on:click={()=> setValue(7)}>7</button>
      <button class="bg-blue" on:click={()=> setValue(8)}>8</button>
      <button class="bg-blue" on:click={()=> setValue(9)}>9</button>
      <button class="bg-bordo" on:click={handleDivision}>/</button>
    </div>
    <div>
      <button class="bg-blue" on:click={()=> setValue(4)}>4</button>
      <button class="bg-blue" on:click={()=> setValue(5)}>5</button>
      <button class="bg-blue" on:click={()=> setValue(6)}>6</button>
      <button class="bg-bordo" on:click={handleMultiplication}>X</button>
    </div>
    <div>
      <button class="bg-blue" on:click={()=> setValue(1)}>1</button>
      <button class="bg-blue" on:click={()=> setValue(2)}>2</button>
      <button class="bg-blue" on:click={()=> setValue(3)}>3</button>
      <button class="bg-bordo" on:click={handleSubtraction}>-</button>
    </div>
    <div>
      <button class="bg-blue" on:click={()=> setValue(0)}>0</button>
      <button class="bg-blue" on:click={handleDecimal}>.</button>
      <button class="bg-green" on:click={handleEqual}>=</button>
      <button class="bg-bordo" on:click={handleAddition}>+</button>
    </div>
  </div>
</div>

<style>
  .buttons button{
    @apply w-[52.5px] h-[52.5px] mr-[1px] ml-[1px] mt-[2px] mb-[2px] text-lg text-white font-bold;
  }

  .buttons{
    @apply w-[240px] flex flex-col items-center grow;
  }
</style>