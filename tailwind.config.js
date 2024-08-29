/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./routes/**/*.templ", "./components/**/*.templ"],
  theme: {
    extend: {
      lineHeight: {
        'extra-loose': '2.5',
        '12': '3rem',
      }
    },
  },
  plugins: [],
};
