package main

// https://leetcode.cn/problems/implement-trie-prefix-tree/description/

// medium
// 思路很简单，每个节点就是26个字母的数组，26叉树，从根节点往下搜每个单词的每个字母即可
// 值得注意的是：apple和app这种情况，所以还需要一个isEnd来标记是否是单词的结尾

type Trie struct {
	Node  [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.Node[v-'a'] == nil {
			this.Node[v-'a'] = &Trie{}
		}
		this = this.Node[v-'a']
	}
	this.isEnd = true
}

func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.Node[v-'a'] == nil {
			return false
		}
		this = this.Node[v-'a']
	}
	return this.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.Node[v-'a'] == nil {
			return false
		}
		this = this.Node[v-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
