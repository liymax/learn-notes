// https://www.aavbook.club/
(()=>{
    let topic = document.querySelector('h1.chapter_title')
    let chapters = document.querySelectorAll('#chapter_content > p')
    let content  = topic.innerText
    for(let n of chapters) {
        content += '\r\n' + n.innerText
    }
    
    window.copy('\r\n' + content)
    console.log('copy finished ' + topic.innerText)

})()
