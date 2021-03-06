import {
    Box,
    Heading,
    SimpleGrid,
    Center,
    Alert,
    AlertIcon,
} from '@chakra-ui/react';
import Card from './Card'
import { useCallback, useEffect, useState } from 'react';
import Navbar from './Navbar';
import Spin from './Spinner';
import {
    Pagination,
    usePagination,
    PaginationNext,
    PaginationPage,
    PaginationPrevious,
    PaginationContainer,
    PaginationPageGroup,
} from "@ajna/pagination";
import axios from 'axios'


export default function Books() {
    const [topcon, setTopcon] = useState(false)
    const [books, setBooks] = useState()
    const [totalPage, setTotalPage] = useState(0)
    const [loading, setLoading] = useState(true)
    const {
        currentPage,
        setCurrentPage,
        pagesCount,
        pages
    } = usePagination({
        total: totalPage,
        initialState: { currentPage: 1, pageSize: 9 },
        limits: {
            outer: 2,
            inner: 2,
        },
    });

    useEffect(() => {
        async function fetchData() {
            try {
                setLoading(true)
                if (topcon === false) {
                    let resp = await axios.get(process.env.REACT_APP_CONTENT_URL + "/books/" + localStorage.getItem('userid') + "/" + currentPage)
                    convertDate(resp.data.books.data)
                    setTotalPage(resp.data.books.count)
                    console.log(resp.data.books.count)
                } else {
                    let resp = await axios.get(process.env.REACT_APP_CONTENT_URL + "/getmostliked/" + localStorage.getItem('userid'))
                    setBooks(resp.data.mostLiked)
                    convertDate(resp.data.mostLiked)
                }

                setLoading(false)
            } catch (err) {
                console.log(err)
                setLoading(false)
                if (!err.response) {
                    alert("Network Error")
                } else {
                    if (err.response.status === 400) {
                        alert("Invalid userid")
                        localStorage.removeItem("userid");
                        window.location.reload();
                    } else {
                        alert("Internal Server Error")
                    }
                }
            }

        }
        fetchData()
    }, [currentPage, topcon, totalPage])

    const convertDate = (bookArray) => {
        bookArray?.forEach((ele) => {
            let date = new Date(ele.date)
            ele.date = (date.getMonth() + 1) + '/' + date.getDate() + '/' + date.getFullYear()
        })
        setBooks(bookArray)
    }

    const callbackTopCon = useCallback((topcon) => {
        setTopcon(topcon)
    }, [])

    const callbackLike = useCallback((id, liked) => {
        let index = books.findIndex((ele) => ele.id === id)
        let newarr = [...books]
        if (liked) {
            newarr[index].likes--
            newarr[index].liked = false
            setBooks(newarr)
        } else {
            newarr[index].likes++
            newarr[index].liked = true
            setBooks(newarr)
        }
    }, [books])

    return (
        <>
            <Navbar callbackTopCon={callbackTopCon} />
            {loading ? <Spin /> :
                <Box marginTop="5vh"
                    marginLeft="15vw"
                    marginRight="15vw">
                    {totalPage == 0 ?
                        <Alert status='info'>
                            <AlertIcon />
                            Currently no Books to display
                        </Alert> :
                        <Box>
                            <Box>
                                {!topcon ?
                                    <Center>
                                        <Heading as="h2">
                                            Some Books By Us
                                        </Heading>
                                    </Center> :
                                    <Center>
                                        <Heading as="h2">
                                            Top Books By Us
                                        </Heading>
                                    </Center>}
                                <br />
                                <SimpleGrid
                                    minChildWidth='300px'
                                    spacingX='50px'
                                    spacingY='5vh'
                                    justifyContent="center">
                                    {books?.map((b) => (
                                        <Card key={b.id}
                                            callbackLike={callbackLike}
                                            id={b.id}
                                            title={b.title}
                                            likes={b.likes}
                                            date={b.date}
                                            story={b.story}
                                            liked={b.liked}
                                        />
                                    ))}
                                </SimpleGrid>
                            </Box>
                            <br />
                            <br />
                            {!topcon && <Center>
                                <Pagination
                                    pagesCount={pagesCount}
                                    currentPage={currentPage}
                                    onPageChange={setCurrentPage}
                                >
                                    <PaginationContainer>
                                        <PaginationPrevious p={4} mx={2}>Previous</PaginationPrevious>
                                        <PaginationPageGroup>
                                            {pages.map((page) => (
                                                <PaginationPage
                                                    p={4}
                                                    mx={1}
                                                    _hover={{
                                                        bg: "blue.500"
                                                    }}
                                                    _current={{
                                                        bg: "blue.100"
                                                    }}
                                                    key={`pagination_page_${page}`}
                                                    page={page}
                                                />
                                            ))}
                                        </PaginationPageGroup>
                                        <PaginationNext p={4} mx={2}>Next</PaginationNext>
                                    </PaginationContainer>
                                </Pagination>
                            </Center>
                            }
                        </Box>
                    }
                </Box>
            }
        </>
    )

}