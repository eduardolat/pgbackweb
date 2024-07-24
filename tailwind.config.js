/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./internal/view/web/**/*.go'],
  plugins: [require('daisyui')],
  daisyui: {
    logs: false,
    themes: ['light', 'dim'],
    darkTheme: 'dim'
  },
  theme: {
    screens: {
      desk: '768px' // only one breakpoint to keep it simple
    },
    extend: {}
  }
}
