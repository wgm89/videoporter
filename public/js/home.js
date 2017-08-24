var videoSearch = function(){}
$.extend(videoSearch.prototype, {
    videoUrl:"",
    title:"",
    m3: {},
    init: function() {
        this.setSearchBox();
        this.bindSearch();
        this.bindM3Download();
        this.toDownload();
        this.demo();
    },
    setSearchBox: function() {
        $(".search_box").focus(function(){
            $(".searchbox_tip").hide();
        }).blur(function(){
            if ($(".search_box").val()=="") {
                $(".searchbox_tip").show();
            }
        })
        $(".searchbox_tip").click(function(){
            $(".search_box").focus();
        })
        $(".search_box").focus();
    },
    bindSearch: function() {
        var self = this;
        $(".search_but").click(function(){
            $("#swfjs").siblings().remove();
            var videoUrl = $(".search_box").val();
            if ($.trim(videoUrl) == "" || !self.isUrl(videoUrl)) {
                BootstrapDialog.show({
                    title:'提示',
                    message: '请输入完整视频链接'
                });
                return;
            }
            $(".loading").show();
            this.videoUrl = videoUrl;
            $.ajax({
                url:'/video/get',
                data:{url:videoUrl},
                type:"POST",
                dataType:'json',
                success: function(data) {
                    self.title = data.title;
                    self.m3 = data.m3;
                    self.data = data;
                    $(".loading").hide();
                    if (data.domain_limit) {
                        self.downloadVideo(data);
                    } else {
                        self.playVideo(data);
                    }
                }
            })
        })
    },
    toDownload: function() {
        var self = this;
        $(document).on("click", ".toDownload", function(){
            self.downloadVideo(self.data);
            try {
                SewisePlayer.doStop();
            } catch(err) {

            }
        })
    },

    bindM3Download: function() {
        var self = this;
        $("body").on("click", ".download_m3", function() {
            var quality = $(this).attr("k");
            var m3_url = self.m3[quality];
            var content = "";
            if (m3_url) {
                if (self.data.m3_download) {
                    window.location.href=m3_url;
                } else {
                    content = self.getUrlContent(m3_url);
                    if (content) {
                        self.downloadData(self.title, content);
                    }
                }
            }
        })
    },

    playVideo: function(data) {
        $("#PlayVideoDialog").modal({
            backdrop: 'static',
            keyboard: false
        });
        $(".modal-title").html(data.title);
        m3List = Object.keys(data.m3);
        if (m3List.length > 0) {
            this._playVideo(data.m3[m3List[0]], this.title, "m3u8");
        } else if(data.flv.length > 0) {
            this._playVideo(data.flv[0], this.title, "mp4");
        }
    },

    downloadVideo: function(data) {
        var title = this.title;
        var message = "";

        if (data.flv.length == 1) {
            message = this.genDownloadMp4(data.flv);
        } else if (Object.keys(data.m3).length > 0) {
            message = this.genDownloadPanel(data.m3);
        } else if(data.flv.length > 0) {
            message = this.genDownloadMp4(data.flv);
        } else {
            message = "本视频暂不支持在线播放";
        }

        BootstrapDialog.show({
            title:title,
            message: message,
            closable: true,
            closeByBackdrop: false,
            closeByKeyboard: false,
            draggable: true
        });
    },

    _playVideo: function(url, title, vtype) {
         SewisePlayer.setup({
            server: "vod",
            type: vtype,
            videourl:url,
            skin: "vodWhite",
            title: title,
            lang: 'zh_CN',
            claritybutton:'disable',
            primary: 'flash'
        });
    },

    downloadData: function(filename, data) {
        var blob = new Blob([data], {type: 'application/vnd.apple.mpegurl'});
        if(window.navigator.msSaveOrOpenBlob) {
            window.navigator.msSaveBlob(blob, filename);
        }
        else{
            var elem = window.document.createElement('a');
            elem.href = window.URL.createObjectURL(blob);
            elem.download = filename;
            document.body.appendChild(elem);
            elem.click();
            document.body.removeChild(elem);
        }
    },
    getUrlContent: function(url) {
        var content;
        $.ajax({
            url:url,
            data:{},
            async: false,
            success: function(data) {
                content = data;
            }
        })
        return content;
    },
    genDownloadPanel: function(data) {
        var tpls = ["选择画质下载"];
        $.each(data, function(k, v){
            tpls.push("<a href='#' class='download_m3' k='"+k+"'>"+k+"</a>");
        })
        return tpls.join("<br/><br/>");
    },
    genDownloadMp4: function(data) {
        var tpls = ["全部视频地址"];
        $.each(data, function(k, v){
            tpls.push("<a href='"+v+"'>"+v+"</a>");
        })
        return tpls.join("<br/><br/>");
    },

    isUrl: function(str) {
        var regex = /(http|https):\/\/(\w+:{0,1}\w*)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%!\-\/]))?/;
        if(!regex .test(str)) {
            return false;
        } else {
            return true;
        }
    },

    demo: function() {
        $(".demo_but").click(function(){
            BootstrapDialog.show({
                size: BootstrapDialog.SIZE_NORMAL,
                title:"演示",
                message: '<img src="" style="width:600px;" class="demo_img" />',
                closable: true,
                closeByBackdrop: false,
                closeByKeyboard: false,
                draggable: true
            });
            setTimeout(function(){$(".demo_img").attr("src", "http://7mnokj.com1.z0.glb.clouddn.com/2017-07-24_1.gif")}, 500)
        })
    }

})
v = new videoSearch();
v.init();
