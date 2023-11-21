import { Container, Snackbar } from "@mui/material";
import { Box, Button, FormLabel, TextField } from "@mui/material";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

export const Registration = () => {
    const [name, setName] = useState("");
    const [profileURL, setProfileURL] = useState("");
    const [bio, setBio] = useState("");
    const [snackbarOpen, setSnackbarOpen] = useState(false);

    const navigate = useNavigate();
    const handleCloseSnackbar = () =>{
      setSnackbarOpen(false);
    };
    const sendPost = async (e) => {
        e.preventDefault();
        const res = await fetch(`/api/users/new`, {
          method: "POST",
          body: JSON.stringify({ name, profileURL, bio }),
          headers: {
            "Content-Type": "application/json",
          },
        });
        setName("");
        setProfileURL("");
        setBio(""); 
        
        if(res.ok){
          setSnackbarOpen(true);

          setTimeout(() => {
            setSnackbarOpen(false);
            navigate("/");
          },2000);
          
        }
      };

      return (
        <>
          <Container
            maxWidth="md"
            sx={{
              py: 3,
            }}
          >
          <form onSubmit={sendPost}>
            <FormLabel htmlFor="body">ユーザー登録</FormLabel>
            
            <TextField 
              variant="outlined"
              margin="normal"
              required
              fullWidth
              multiline
              rows={1}
              id="name"
              label="ユーザー名"
              name="name"
              autoFocus
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
            <TextField 
              variant="outlined"
              margin="normal"
              fullWidth
              multiline
              rows={1}
              id="profileURL"
              label="プロフィール画像"
              name="profileURL"
              autoFocus
              value={profileURL}
              onChange={(e) => setProfileURL(e.target.value)}
            />
             <TextField 
              variant="outlined"
              margin="normal"
              fullWidth
              multiline
              rows={4}
              id="bio"
              label="自己紹介"
              name="bio"
              autoFocus
              value={bio}
              onChange={(e) => setBio(e.target.value)}
            />
            
            <Box textAlign="center" m={3}>
            
              <Button
                type="submit"
                variant="contained"
                color="primary"
                disabled={ !name  }
              >
                登録する
              </Button>
              
            </Box>
            
          </form>

          <Snackbar
            anchorOrigin={{
              vertical: 'bottom',
              horizontal: 'center'
            }}
            open = {snackbarOpen}
            autoHideDuration={2000}
            onClose = {handleCloseSnackbar}
            message="登録完了しました"
          />
          </Container>
        </>
      );
};