<script lang="ts">
    import Operations from './components/operations.svelte'
    import History from './components/history.svelte'
</script>

<div class="inline-flex w-full justify-center flex-wrap">
    <div class="m-4 border-8 border-black w-[20rem] h-[28rem] flex flex-col items-center">
    
      <div class="bg-gray flex-grow mb-3 w-[304px] text-center">
        <h1 class="text-2xl font-bold font-serif">Calculadora Basica</h1>
      </div>
      <Operations />
    </div>
    <div class="m-4 border-8 border-black w-[20rem] h-[28rem] flex flex-col items-center">
    
      <div class="bg-gray flex-grow mb-3 w-[304px] text-center">
        <h1 class="text-2xl font-bold font-serif">Calculadora Basica</h1>
      </div>
      <History />
    </div>
  </div>
