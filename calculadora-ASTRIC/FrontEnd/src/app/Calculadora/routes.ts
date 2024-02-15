
// @ts-nocheck
import { wrap } from 'svelte-spa-router/wrap'

const routes = { '/Calculadora': wrap({ asyncComponent: () => import('./Calculadora/Modul.svelte') }),};

export default routes;
