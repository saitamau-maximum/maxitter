import { 
    FormControl,
    InputLabel,
    Select,
    MenuItem,
} from '@mui/material';


export  const UserSelectBox = () =>{
  return(
  <FormControl fullWidth>
    <InputLabel id="usertable">User</InputLabel>
    <Select
      labelId="usertable"
      id="usertable"
      value={user}
      label="User"
      onChange={handleChange}
    >
      {users.map((user) => (
          <MenuItem value={10}>{user.username}</MenuItem>
      ))}
    </Select>
  </FormControl>
  )
}