module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}", 
  ],
  theme: {
    extend: {
      screens: { // Ensure it's 'screens' not 'screen'
        sm: '360px',    // iPhone 15 Pro, 15, and older iPhones (smallest screens)
        md: '768px',    // Medium devices like tablets
        lg: '1440px',   // Large devices like smaller laptops
        xl: '1920px',   // Larger laptops and desktops
      }
    },
  },
  plugins: [],
};
