import "./App.css";
import Login from "./pages/login";
import { ThemeProvider } from "./components/theme-provider";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { AuthProvider, RequireAuth } from "./context/auth";
import { Dashboard } from "./pages/dashboard";
import { Footer } from "./components/footer";

function App() {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
        <AuthProvider>
          <BrowserRouter>
            <Routes>
              <Route path="/login" element={<Login />} />
              <Route element={<RequireAuth />}>
                <Route path="/" element={<Dashboard />} />
              </Route>
            </Routes>
          </BrowserRouter>
        </AuthProvider>
        <Footer className="border-t fixed bottom-0 left-0 z-20 w-full p-1 shadow md:flex md:items-center md:justify-between md:p-2" />
    </ThemeProvider>
  );
}

export default App;
