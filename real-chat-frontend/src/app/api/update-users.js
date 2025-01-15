

export default async function activeChatHandler(req, res) {
    if (req.method === 'POST') {
        const { userIndex } = req.body
        console.log(userIndex)
    }
}