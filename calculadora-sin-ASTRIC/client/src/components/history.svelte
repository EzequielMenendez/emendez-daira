<script lang="ts">
  import type { HistoryData } from "../types";

    export let info: HistoryData[]
</script>

<h2 class="text-4xl font-semibold">Historial</h2>
<div class="h-72 w-[265px] bg-black mb-8 overflow-hidden">
    <div class="flex place-content-between">
        <div class=" w-20">
            <p class="text-white text-sm ml-1">Fecha</p>
        </div>
        <div class=" w-20">
            <p class="text-white text-sm">Operacion</p>
        </div>
        <div class=" w-20">
            <p class="text-white text-sm">Resultado</p>
        </div>
    </div>
    {#each info as i}
        <div class="flex place-content-between mt-1">
            <div class=" w-20">
                <p class="text-white text-xs ml-1">{i.Date}</p>
            </div>
            <div class=" w-20">
                <p class="text-white text-xs">{i.Operation}</p>
            </div>
            <div class=" w-20 overflow-hidden">
                <p class="text-white text-xs">{i.Result}</p>
            </div>
        </div>
    {/each}
</div>