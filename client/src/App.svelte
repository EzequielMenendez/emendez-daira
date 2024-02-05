<script lang="ts">
  import Calculator from "./components/calculator.svelte";
  import History from "./components/history.svelte";
  import axios from "axios"
  import type {HistoryData} from './types'

  let info:HistoryData[] = []

  const postCalculator = async(date:string, operation:string, result:number) => {
    try {
      let {data} = await axios.post('http://localhost:3000/history', {
      date,
      operation,
      result
    })
    
    info = data

    info.reverse()

    } catch (error) {
      console.error(error)
    }
  }

  const getHistory = async() => {
    try {
    let {data} = await axios.get('http://localhost:3000/history')
    
    info = data

    info.reverse()

    } catch (error) {
      console.error(error)
    }
  }

  getHistory()
  
</script>

<div class="inline-flex w-full justify-center flex-wrap">
  <div class="m-4 border-8 border-black w-[20rem] h-[28rem] flex flex-col items-center">
  
    <div class="bg-gray flex-grow mb-3 w-[304px] text-center">
      <h1 class="text-2xl font-bold font-serif">Calculadora Basica</h1>
    </div>
    <Calculator postCalculator={postCalculator}/>
  </div>
  <div class="m-4 border-8 border-black w-[20rem] h-[28rem] flex flex-col items-center">
  
    <div class="bg-gray flex-grow mb-3 w-[304px] text-center">
      <h1 class="text-2xl font-bold font-serif">Calculadora Basica</h1>
    </div>
    <History info={info}/>
  </div>
</div>