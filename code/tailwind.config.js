/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{html,tmpl}"],
  theme: {
    extend: {
      width: {
        128: "32rem",
      },
      spacing: {
        128: "32rem",
        144: "36rem",
      },
      colors: {
        "desert-sand": {
          50: "#fcf5f0",
          100: "#f7e9dd",
          200: "#ecc8af",
          300: "#e4ae8d",
          400: "#d8875f",
          500: "#cf6940",
          600: "#c15335",
          700: "#a0412e",
          800: "#81362b",
          900: "#682f26",
          950: "#381512",
        },
        "river-bed": {
          50: "#f4f6f7",
          100: "#e2e7eb",
          200: "#c8d2d9",
          300: "#a2b1be",
          400: "#758a9b",
          500: "#5a6f80",
          600: "#495867",
          700: "#434f5b",
          800: "#3c444e",
          900: "#353b44",
          950: "#20242c",
        },
        contessa: {
          50: "#fcf5f4",
          100: "#f8eae8",
          200: "#f4d9d4",
          300: "#eabeb7",
          400: "#dd978c",
          500: "#ce796b",
          600: "#b75a4b",
          700: "#99493c",
          800: "#803f34",
          900: "#6b3931",
          950: "#391b16",
        },
        dingley: {
          50: "#f6f8f5",
          100: "#eaf0e8",
          200: "#d5e1d1",
          300: "#b3c9ac",
          400: "#88a97f",
          500: "#678a5d",
          600: "#587a4f",
          700: "#415a3b",
          800: "#364932",
          900: "#2e3c2b",
          950: "#151f14",
        },
        "el-salva": {
          50: "#faf5f2",
          100: "#f4e8e0",
          200: "#e9cfbf",
          300: "#daaf97",
          400: "#cb896c",
          500: "#c06c4f",
          600: "#b25944",
          700: "#99493c",
          800: "#783b34",
          900: "#61322d",
          950: "#341816",
        },
        "golden-grass": {
          50: "#fdfbe9",
          100: "#fbf7c6",
          200: "#f9ec8f",
          300: "#f5da4f",
          400: "#f0c51f",
          500: "#e0ad12",
          600: "#c1860d",
          700: "#9a600e",
          800: "#804d13",
          900: "#6d3f16",
          950: "#3f2009",
        },
      },
    },
  },
  plugins: [],
};
