#include <vector>
#include <queue>
using namespace std;

/**
 * Definition for singly-linked list.
 */
struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
    // Comparator para a min-heap
    struct compare {
        bool operator()(ListNode* a, ListNode* b) {
            return a->val > b->val;
        }
    };
    
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        priority_queue<ListNode*, vector<ListNode*>, compare> minHeap;

        // Adiciona os nós iniciais de cada lista
        for (auto node : lists) {
            if (node) minHeap.push(node);
        }

        ListNode* dummy = new ListNode(0);
        ListNode* current = dummy;

        while (!minHeap.empty()) {
            ListNode* smallest = minHeap.top();
            minHeap.pop();

            current->next = smallest;
            current = current->next;

            if (smallest->next) {
                minHeap.push(smallest->next);
            }
        }

        return dummy->next;
    }
};
