const item = document.querySelectorAll('.step');
const lenitem = item.length;
const suivant = document.querySelector('.next');
const precedent = document.querySelector('.prev');
let count = 0



function slideNext(){
    console.log("slideNext")
    for (let i=0; i < lenitem; i++){
        if (item[i].classList.contains('active-step')){
            console.log(i,lenitem)
            item[i].classList.remove('active-step')
            if ((i+1) >= lenitem){
                item[0].classList.add('active-step')
                break
            }else{
                const index = i+1
                item[index].classList.add('active-step')
                break
            }
            
        }
    }
}

function slidePrev(){
    console.log("slidePrev")
    for (let i=0; i < lenitem; i++){
        if (item[i].classList.contains('active-step')){
            item[i].classList.remove('active-step')
            console.log(i,lenitem)
            if (i-1 < 0){
                console.log("lala")
                item[lenitem-1].classList.add('active-step')
                break
            }else{
                console.log("ici")
                item[i-1].classList.add('active-step')
                break
            }
        }
    }
}

suivant.addEventListener('click',slideNext)
precedent.addEventListener('click', slidePrev)