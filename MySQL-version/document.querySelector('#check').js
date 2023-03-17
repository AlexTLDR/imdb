document.querySelector('#check').addEventListener('click', function(){
                   let input = document.querySelector('input');
                   if(input.value == 'omega'){
                       input.style.backgroundColor = 'green';
                       document.querySelector('#feedback2').innerHTML = 'Correct!';
                   } else {
                       input.style.backgroundColor = 'red';
                       document.querySelector('#feedback2').innerHTML = 'Incorrect';
                   }
                });

                document.querySelector('#check').addEventListener('click', function(){
                    let input = document.querySelector('input');
                    //let lowerCase = input.toLowerCase();
                    if(input.value === 'omega'){
                        input.style.backgroundColor = 'green';
                        document.querySelector('#feedback2').innerHTML = 'Correct!';
                    } else {
                        input.style.backgroundColor = 'red';
                        document.querySelector('#feedback2').innerHTML = 'Incorrect!';
                    }
                });