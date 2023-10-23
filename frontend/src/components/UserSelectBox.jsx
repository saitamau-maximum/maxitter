import { 
    FormControl,
    InputLabel,
    Select,
    MenuItem,
} from '@mui/material';


export  const UserSelectBox = ({users,selectedUser, handleChange}) =>{
  return(
  <>
  <FormControl fullWidth>
    <InputLabel id="user">User</InputLabel>
    <Select
      labelId="user"
      id="user"
      value={selectedUser}
      label="User"
      onChange={handleChange}
    >
    {users.map((user) => (
      <MenuItem value = {user.username}>{user.username}</MenuItem>
    ))}
    <MenuItem value={10}>Ten</MenuItem>
    <MenuItem value={20}>Twenty</MenuItem>
    <MenuItem value={30}>Thirty</MenuItem>
    </Select>
  </FormControl>
  </>
  )
}