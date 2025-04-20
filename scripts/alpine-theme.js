document.addEventListener("alpine:init", () => {
  Alpine.data("theme", () => ({
    isDark: false,
    toggleTheme() {
      this.isDark = !this.isDark;
      document.documentElement.classList.toggle("dark", this.isDark);
      localStorage.setItem("theme", this.isDark ? "dark" : "light");
    },
    init() {
      const savedTheme = localStorage.getItem("theme");
      if (savedTheme) {
        this.isDark = savedTheme === "dark";
        document.documentElement.classList.toggle("dark", this.isDark);
      }
    },
  }));
});
