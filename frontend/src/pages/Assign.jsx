import { Container, nativeSelectClasses } from "@mui/material";
import { Box, Button, FormLabel, TextField } from "@mui/material";
import { useState } from "react";
import { Navigate, useNavigate } from "react-router-dom";

export const Assign = () => {
    const [name, setName] = useState("");
    const [profileURL, setProfileURL] = useState("");
    const [bio, setBio] = useState("");
    const [isSending, setIsSending] = useState(false);//一応書いておくが、連続して登録はしない気がするので後で考える。

    const navigate = useNavigate();
    const sendPost = async (e) => {
        e.preventDefault();
        setIsSending(true);
        const res = await fetch(`/api/users/new`, {
          method: "POST",
          body: JSON.stringify({ name, profileURL, bio }),
          headers: {
            "Content-Type": "application/json",
          },
        });
        setIsSending(false);
        setName("");
        setProfileURL("");
        setBio(""); 
        
        if(res.ok){
          navigate("/");
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
              required
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
              required
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
                disabled={ !name && !bio }//ここの条件には後ですべてのボックスが空白でないことを追加する。
              >
                登録する
              </Button>
              
            </Box>
            
          </form>
          </Container>
        </>
      );
};