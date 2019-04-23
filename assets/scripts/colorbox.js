let ppms = document.getElementsByClassName('colorable');

const great = 10;
const good = 20;
const poor = 34;
const danger = 35;
//great  #c3e6cb
//good   #d6d8db
//poor   #ffeeba
//danger #f5c6cb
let i;
for (i = 0; i < ppms.length; i++) {
	let partdata = parseInt(ppms[i].children[1].children[0].textContent.split(' ')[0]);
	console.log(partdata);
	if (partdata > danger) { //danger
		ppms[i].style.backgroundColor = '#f5c6cb';
	} else if (partdata < danger && partdata > poor) {
		ppms[i].style.backgroundColor = '#ffeeba';
	} else if (partdata < poor && partdata > good) {
		ppms[i].style.backgroundColor = '#d6d8db';
	} else if (partdata < good) {
		ppms[i].style.backgroundColor = '#c3e6cb';
	}
}






