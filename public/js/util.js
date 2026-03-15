// openLinksInNewTab は指定したコンテナ内のすべてのリンクを新しいタブで開くよう設定する。
function openLinksInNewTab(containerId) {
	var container = document.getElementById(containerId);
	if (!container) return;
	container.querySelectorAll('a').forEach(function(a) {
		a.setAttribute('target', '_blank');
		a.setAttribute('rel', 'noopener noreferrer');
	});
}
