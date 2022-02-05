import {
    Box,
    Heading,
    Link,
    Text,
    HStack,
    Icon,
    Spacer,
    Center
} from '@chakra-ui/react';
import { AiOutlineHeart,AiFillHeart } from "react-icons/ai";
export default function Card({title,likes,date}) {
    return (
        <Center>
            <Box w="300px"
                padding="2vw"
                // bg="gray.100"
                boxShadow={'lg'}
                rounded={'lg'} >
                <Heading fontSize="xl" marginTop="2">
                    <Link textDecoration="none" _hover={{ textDecoration: 'none' }}>
                        {title}
                    </Link>
                </Heading>
                <br />
                <HStack>
                <Text>{date}</Text>
                <Spacer />
                <Icon as={AiOutlineHeart} boxSize={6}/>
                <Text>{likes}</Text>
                </HStack>
            </Box>
        </Center>
    )
}