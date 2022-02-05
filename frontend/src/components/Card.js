import {
    Box,
    Heading,
    Link,
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
import { ChangeEvent, useCallback, useEffect, useState, useRef } from 'react';
export default function Card({ title, likes, date, story, liked }) {

    const { isOpen, onOpen, onClose } = useDisclosure()
    const btnRef = useRef()
    const onClick = () => {
        console.log("test")
    }
    const onLike = () => {
        console.log("dont know")
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
                    <Button bg="white" onClick={onLike}>
                        <Icon as={AiOutlineHeart} boxSize={6} />
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
                            <Button bg="white" onClick={onLike}>
                                <Icon as={AiOutlineHeart} boxSize={6} />
                            </Button>
                            <Text>{likes}</Text>

                        </ModalFooter>
                    </ModalContent>
                </Modal>
            </Box>
        </Center>
    )
}