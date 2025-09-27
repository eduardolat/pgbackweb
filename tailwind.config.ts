import type { Config } from "tailwindcss";
import daisyui from "daisyui";
import * as daisyuiThemes from "daisyui/src/theming/themes";

export default {
  content: ["./internal/view/web/**/*.go"],
  plugins: [daisyui as any],
  daisyui: {
    logs: false,
    themes: [
      {
        light: {
          ...daisyuiThemes.light,
          primary: "#2be7c8",
          "success-content": "#ffffff",
          "error-content": "#ffffff",
        },
        dark: {
          ...daisyuiThemes.dracula,
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
} satisfies Config;
