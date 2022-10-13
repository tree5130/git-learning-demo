#include <iostream>
#include <cstring>
#include <vector>
using namespace std;

struct Trie
{
    /* data */
    bool is_finished;
    Trie* child[26];
};


class MagicDictionary{
public:
    MagicDictionary(){
        root = new Trie();
    }
    void buildDict(vector<string> dictionary){
        for (auto word : dictionary){
            Trie* cur = root;
            for (auto ch : word){
                int idx = ch - 'a';
                if (!cur->child[idx]){
                    cur->child[idx] = new Trie();
                }
                cur = cur->child[idx];
            }
            cur->is_finished = true;
        }
    }
    bool search(string searchWord){
        function<bool(Trie*, int, bool)> dfs = [&](Trie* node, int pos, bool modified){
            if (pos == searchWord.size()){
                return modified && node->is_finished;
            }
            int idx = searchWord[pos] - 'a';
            if (node->child[idx]){
                if (dfs(node->child[idx], pos + 1, modified)){
                    return true;
                }
            }
            if (!modified){
                for (int i = 0; i < 26; i ++ ){
                    if (i != idx && node->child[i]){
                        if (dfs(node->child[i], pos + 1, !modified))
                            return true;
                    }
                }
            }
            return false;
        };
        return dfs(root, 0, false);
    }
private:
    Trie* root;
};