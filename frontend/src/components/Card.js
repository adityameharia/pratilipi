import {
    Box,
    Heading,
    Text,
    HStack,
    Icon,
    Spacer,
    Button,
    Center,
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalCloseButton,
    ModalBody,
    ModalFooter,
    useDisclosure
} from '@chakra-ui/react';
import { AiOutlineHeart, AiFillHeart } from "react-icons/ai";
import { useRef } from 'react';
import axios from 'axios';
export default function Card({ callbackLike, id, title, likes, date, story, liked }) {

    const { isOpen, onOpen, onClose } = useDisclosure()
    const btnRef = useRef()
    const onLike = async () => {
        try {
            if (liked) {
                await axios.post("http://localhost:8001/like/remove/" + localStorage.getItem("userid") + "/" + id)
            } else {
                await axios.post("http://localhost:8001/like/add/" + localStorage.getItem("userid") + "/" + id)
            }
            callbackLike(id, liked)
        } catch (err) {
            alert("Unable to update likes")
        }
    }
    return (
        <Center>
            <Box
                w="300px"
                padding="2vw"
                boxShadow={'lg'}
                rounded={'lg'} >

                <Heading
                    onClick={onOpen}
                    _hover={{
                        cursor: 'pointer'
                    }}
                    fontSize="xl"
                    marginTop="2">
                    {title}
                </Heading>
                <br />
                <HStack>
                    <Text>{date}</Text>
                    <Spacer />
                    <Button bg="white" onClick={() => {
                        onLike()
                    }}>
                        {liked ? <Icon as={AiFillHeart} color="red.500" boxSize={6} /> : <Icon as={AiOutlineHeart} boxSize={6} />}
                    </Button>
                    <Text>{likes}</Text>
                </HStack>
                <Modal
                    onClose={onClose}
                    finalFocusRef={btnRef}
                    isOpen={isOpen}
                    size={'xl'}
                    scrollBehavior={'inside'}
                    isCentered
                >
                    <ModalOverlay />
                    <ModalContent>
                        <ModalHeader>{title}</ModalHeader>
                        <ModalCloseButton />
                        <ModalBody>
                            {story}
                        </ModalBody>
                        <ModalFooter>

                            <Text textAlign={'left'}>{date}</Text>
                            <Spacer />
                            <Button bg="white" onClick={() => {
                                onLike()
                                callbackLike(id, liked)
                            }}>
                                {liked ? <Icon as={AiFillHeart} color="red.500" boxSize={6} /> : <Icon as={AiOutlineHeart} boxSize={6} />}
                            </Button>
                            <Text marginLeft={'2'}>{likes}</Text>

                        </ModalFooter>
                    </ModalContent>
                </Modal>
            </Box>
        </Center>
    )
}