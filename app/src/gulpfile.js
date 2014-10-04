var gulp    = require('gulp')
, shell     = require('gulp-shell')
, slim      = require("gulp-slim");

gulp.task('slim', function(){
    gulp.src("views/*.slim")
        .pipe(slim({
            pretty: true
    }))
    .pipe(gulp.dest("./dist/html"));
});

gulp.task('duo', shell.task([
    'duo index.js dist',
    // 'duo assets/coffee/*.{coffee} dist/js',
    'duo assets/scss/*.{scss} dist/css'
]))

gulp.task('build', [
    'slim',
    'duo'
]);
