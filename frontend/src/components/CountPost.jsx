import { Box } from "@mui/material";


export const CountPost = ({amount}) => {
    // console.log(amount);
    // if(amount !=0)
    // {
    //     amount =  1
    // }
  
  return (
    <>
      <Box textAlign="left" m={1}>
        投稿件数　{amount}
      </Box>
    </>
  );
};
