import { useContext } from "react";
import { ColorModeContext } from "./ColorModeContext";
import { useTheme } from "@emotion/react";
import { Box } from "@mui/system";
import { IconButton } from "@mui/material";
import { Brightness4, Brightness7 } from "@mui/icons-material";


export const ToggleTheme = () => {
  const theme = useTheme();
  const colorMode = useContext(ColorModeContext);
  const icon =
    theme.palette.mode === "dark" ? <Brightness7 /> : <Brightness4 />;
  return (
    <Box sx={{ ml: 1 }}>
      <IconButton onClick={colorMode.toggleColorMode} color="inherit">
        {icon}
      </IconButton>
    </Box>
  );
};
