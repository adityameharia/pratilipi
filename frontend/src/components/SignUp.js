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
import { ViewIcon, ViewOffIcon } from '@chakra-ui/icons';

export default function Signup() {
  const [user, setUser] = useState({ email: '', password: '' })

  const onChange = (e) => {
    setUser({ ...user, [e.target.id]: e.target.value });
    console.log(user)
  };

  const onSubmit = async (e) => {
    e.preventDefault()
  };

  function validateEmail(str){
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
      onSubmit={onSubmit}
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