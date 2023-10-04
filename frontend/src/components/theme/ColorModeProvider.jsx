import { ThemeProvider, createTheme } from "@mui/material/styles";
import { useMemo, useState } from "react";
import { ColorModeContext } from "./ColorModeContext";

export const ColorModeProvider = ({ children }) => {
  const [mode, setMode] = useState("light");
  const colorMode = useMemo(
    () => ({
      toggleColorMode: () => {
        setMode((prevMode) => (prevMode === "light" ? "dark" : "light"));
      },
    }),
    []
  );

  const theme = useMemo(
    () =>
      createTheme({
        palette: {
          mode,
        },
      }),
    [mode]
  );

  return (
    <ThemeProvider theme={theme}>
      <ColorModeContext.Provider value={colorMode}>
        {children}
      </ColorModeContext.Provider>
    </ThemeProvider>
  );
}