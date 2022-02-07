import {
  Flex,
  Box,
  FormControl,
  FormLabel,
  Input,
  InputGroup,
  InputRightElement,
  Stack,
  Button,
  Heading,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import { useEffect, useState } from 'react';
import axios from 'axios'
import { ViewIcon, ViewOffIcon } from '@chakra-ui/icons';

export default function Signup() {
  const [user, setUser] = useState({ email: "", password: "" })

  const onChange = (e) => {
    setUser({ ...user, [e.target.id]: e.target.value });
  };

  const onSubmit = async (e) => {
    e.preventDefault()
    console.log(user)

    axios.post(process.env.REACT_APP_USER_URL+"/signup", user, {
      headers: {
        "Content-Type": "application/json"
      }
    }).then(res => {
      localStorage.setItem('userid', res.data.id)
      console.log(res.data)
      window.location.reload();
    }).catch(err => {
      if (err.response.status == 400) {
        alert("Invalid email or password")
      } else if (err.response.status == 401) {
        alert("User is already registered but the password you are entering is wrong")
      } else {
        alert("Internal Server Error")
      }
    })

  };

  function validateEmail(str) {
    let re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(str)
  }
  function hasLowerCase(str) {
    return str.toUpperCase() !== str;
  }
  function hasUpperCase(str) {
    return str.toLowerCase() !== str;
  }
  function hasNumber(myString) {
    return /\d/.test(myString);
  }
  function hasSpecial(str) {
    var regex = /[ !@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/g;
    return regex.test(str);
  }

  const [showPassword, setShowPassword] = useState(false);
  const [disabled, setDisabled] = useState(true);

  useEffect(() => {
    if (hasLowerCase(user.password)
      && hasUpperCase(user.password)
      && hasNumber(user.password)
      && hasSpecial(user.password)
      && validateEmail(user.email)
      && user.password.length >= 8) {
      setDisabled(false)
    }
    else {
      setDisabled(true)
    }
  }, [user])

  return (
    <Flex
      minH={'90vh'}
      align={'center'}
      justify={'center'}
      bg={useColorModeValue('gray.50', 'gray.800')}>
      <Stack spacing={8} mx={'auto'} maxW={'lg'} py={12} px={6}>
        <Stack align={'center'}>
          <Heading fontSize={'4xl'} textAlign={'center'}>
            Sign up
          </Heading>
          <Text fontSize={'lg'} color={'gray.600'}>
            to enjoy all of our cool features ✌️
          </Text>
        </Stack>
        <Box
          rounded={'lg'}
          bg={useColorModeValue('white', 'gray.700')}
          boxShadow={'lg'}
          p={8}>
          <Stack spacing={4}>
            <FormControl
              id="email"
              isRequired
              type='email'
              value={user.email}
              onChange={onChange}>
              <FormLabel>Email address</FormLabel>
              <Input type="email" />
              <Text fontSize={'sm'} color={'red'}>
                {user.email.length !== 0 && !validateEmail(user.email) && <Text>Enter a valid email</Text>}
              </Text>
            </FormControl>
            <FormControl
              id="password"
              isRequired
              type='text'
              value={user.password}
              onChange={onChange}>
              <FormLabel>Password</FormLabel>
              <InputGroup>
                <Input type={showPassword ? 'text' : 'password'} />
                <InputRightElement h={'full'}>
                  <Button
                    variant={'ghost'}
                    onClick={() =>
                      setShowPassword((showPassword) => !showPassword)
                    }>
                    {showPassword ? <ViewIcon /> : <ViewOffIcon />}
                  </Button>
                </InputRightElement>
              </InputGroup>
              <Text fontSize={'sm'} color={'red'}>
                {user.password.length !== 0 && user.password.length < 8 && <Text>* Password must be at least 8 characters</Text>}
                {user.password.length !== 0 && !hasLowerCase(user.password) && <Text>* Password must have one lowercase letter</Text>}
                {user.password.length !== 0 && !hasUpperCase(user.password) && <Text>* Password must have one uppercase letter</Text>}
                {user.password.length !== 0 && !hasNumber(user.password) && <Text>* Password must have one number</Text>}
                {user.password.length !== 0 && !hasSpecial(user.password) && <Text>* Password must have one special character</Text>}
              </Text>
            </FormControl>
            <Stack spacing={10} pt={2}>
              <Button
                isDisabled={disabled}
                loadingText="Submitting"
                size="lg"
                bg={'blue.400'}
                color={'white'}
                onClick={onSubmit}
                _hover={{
                  bg: 'blue.500',
                }}>
                Sign up
              </Button>
            </Stack>
          </Stack>
        </Box>
      </Stack>
    </Flex>
  );
}