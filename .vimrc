set shell=/bin/bash
set tabstop=4
set belloff=all
set backspace=indent,eol,start

" Vundle менеджер пакетов 
source /Users/p141592/.vim/Vundle.sh

" Pathogen менеджер пакетов
execute pathogen#infect()
syntax on
filetype plugin indent on

" Плагины
Plugin 'vim-airline/vim-airline'
Plugin 'vim-airline/vim-airline-themes'

" Статус бар
" https://github.com/vim-airline/vim-airline
