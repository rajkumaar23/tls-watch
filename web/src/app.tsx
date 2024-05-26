import "./app.css";
import Login from "./pages/login";
import { ThemeProvider } from "./components/theme-provider";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { AuthProvider } from "./context/auth";
import { Domains } from "./pages/domains";
import { DashboardLayout } from "./components/layouts/dashboard";
import { Settings } from "./pages/settings";
import { RequireAuth } from "./components/require-auth";

function App() {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <BrowserRouter>
        <AuthProvider>
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route element={<RequireAuth />}>
              <Route element={<DashboardLayout />}>
                <Route path="/" element={<Domains />} />
                <Route path="/settings" element={<Settings />} />
              </Route>
            </Route>
          </Routes>
        </AuthProvider>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
