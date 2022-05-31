const btnNav = document.querySelectorAll('.btn-nav')
const Onglet = document.querySelectorAll('.nav-body')


for (i = 0; i < btnNav.length; i++){ 
    var changeOnglet = function(){
        for (i = 0; i < btnNav.length; i++){
            if (btnNav[i] == this){
                this.classList.add('active-onglet')
                Onglet[i].classList.add('active-sub')
            } else {
                btnNav[i].classList.remove('active-onglet')
                Onglet[i].classList.remove('active-sub')
            }
        }
    }
    btnNav[i].addEventListener('click',changeOnglet)
}
