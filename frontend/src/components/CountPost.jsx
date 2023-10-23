import { Box } from "@mui/material";


export const CountPost = ({amount}) => {
    
  return (
    <>
      <Box textAlign="left" m={1}>
        投稿件数　{amount}
      </Box>
    </>
  );
};
