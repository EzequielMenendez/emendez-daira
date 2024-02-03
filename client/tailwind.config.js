/** @type {import('tailwindcss').Config} */
export default {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      backgroundColor: {
        'blue': '#0000ff',
        'orange': '#ff4500',
        'green': '#008000',
        'bordo': '#800000',
        'gray': '#f1f1f1'
      },
    },
  },
  plugins: [],
}

