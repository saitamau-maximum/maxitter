import { CssBaseline } from "@mui/material";
import { GlobalStyles } from "@mui/material";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { Header } from "./components/Header";
import { ColorModeProvider } from "./components/theme/ColorModeProvider.jsx";
import { ToggleTheme } from "./components/theme/ToggleTheme.jsx";
import { Home } from "./pages/Home.jsx";

function App() {
  return (
    <>
      <ColorModeProvider>
        <CssBaseline />
        <Header />
        <ToggleTheme />
        <GlobalStyles
          styles={{
            body: {
              margin: 0,
            },
          }}
        />
        <Router>
          <Routes>
            <Route path="/" element={<Home />}></Route>
          </Routes>
        </Router>
      </ColorModeProvider>
    </>
  );
}

export default App;
