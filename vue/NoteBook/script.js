new Vue({
    el: '#notebook',
    created() {
        this.content = localStorage.getItem('content') || 'you can write in **markdown**'
    },
    data() {
        return {
            content: 'This is a note',
            notes: []
        }
    },
    computed: {
        notePreview() {
            // markdown 转 HTML
            return marked(this.content)
        }
    },
    watch: {
        content: 'saveNote'
    },
    method: {
        saveNote(val) {
            localStorage.setItem('content', val)
        },
        addNote() {
            const time = Date.now();
            const note = {
                id: String(time),
                title: 'New Note' + (this.notes.length + 1),
                content: '',
                created: time,
                favorite: false
            };
            this.notes.push(note)
        }
    }
});
