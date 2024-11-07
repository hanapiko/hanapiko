// Smooth Scrolling
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function(e) {
        e.preventDefault();
        document.querySelector(this.getAttribute('href')).scrollIntoView({
            behavior: 'smooth'
        });
    });
});

// Form Validation
document.querySelector('form').addEventListener('submit', function(e) {
    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;

    if (name === '' || email === '') {
        e.preventDefault();
        alert('Please fill in all fields.');
    } else {
        alert('Message sent successfully!');
    }
});

// Project Descriptions (Simple Alert Example)
const projectDescriptions = {
    'Project Title 1': "Detailed description for Project 1.",
    'Project Title 2': "Detailed description for Project 2.",
    'Project Title 3': "Detailed description for Project 3."
};

document.querySelectorAll('.project h3').forEach(project => {
    project.addEventListener('click', function() {
        alert(projectDescriptions[this.innerText]);
    });
});