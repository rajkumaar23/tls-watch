import "./app.css";
import Login from "./pages/login";
import { ThemeProvider } from "./components/theme-provider";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { AuthProvider, RequireAuth } from "./context/auth";
import { Domains } from "./pages/domains";
import { DashboardLayout } from "./components/layouts/dashboard";

function App() {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <AuthProvider>
        <BrowserRouter>
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route element={<RequireAuth />}>
              <Route element={<DashboardLayout />}>
                <Route path="/" element={<Domains />} />
              </Route>
            </Route>
          </Routes>
        </BrowserRouter>
      </AuthProvider>
    </ThemeProvider>
  );
}

export default App;
