/** @type {import('tailwindcss').Config} */

const colors = require("tailwindcss/colors");

module.exports = {
  // mode: "jit",

  content: ["./src/**/*.{html, js, ts, vue, tsx, jsx}", "./pkg/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [],
  // presets: [require("@acmecorp/tailwind-base")],
};
