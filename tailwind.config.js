/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/view/web/**/*.go"],
  plugins: [require("daisyui")],
  daisyui: {
    logs: false,
    themes: [
      {
        light: {
          ...require("daisyui/src/theming/themes").light,
          primary: "#2be7c8",
          "success-content": "#ffffff",
          "error-content": "#ffffff",
        },
        dark: {
          ...require("daisyui/src/theming/themes").dracula,
          primary: "#2be7c8",
        },
      },
    ],
    darkTheme: "dark",
  },
  theme: {
    screens: {
      desk: "768px", // only one breakpoint to keep it simple
    },
    extend: {},
  },
};
