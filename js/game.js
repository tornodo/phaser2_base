/*
 *  微信： tornodo
 *  博客：https://www.jianshu.com/u/1b05a5363c32
 *  微信公众号： 工匠前沿
 */

var game = new Phaser.Game(600, 600, Phaser.AUTO, 'game');

game.states = {};
// 引导
game.states.boot = function() {
    this.preload = function() {
        this.load.image('loading', 'assets/image/progress.png');
    },
    this.create = function() {
    }
}
// 用来显示资源加载进度
game.states.preloader = function() {
    this.preload = function() {
    },
    this.create = function() {
    }
}

// 游戏菜单
game.states.menu = function() {
    this.create = function() {
    },
    this.startGame = function() {
    }
}
//游戏界面
game.states.start = function() {
    this.preload = function() {
    },
    this.create = function() {
    }
    this.update = function() {
    }
}

game.state.add('boot', game.states.boot);
game.state.add('preloader', game.states.preloader);
game.state.add('menu', game.states.menu);
game.state.add('start', game.states.start);
game.state.start('boot');