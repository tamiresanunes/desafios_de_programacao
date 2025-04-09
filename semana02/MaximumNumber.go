class Solution {
	public:
			// Retorna a maior subsequência possível com tamanho k
			vector<int> getMaxSubsequence(vector<int>& nums, int k) {
					vector<int> stack;
					int drop = nums.size() - k;
					for (int num : nums) {
							while (!stack.empty() && drop > 0 && stack.back() < num) {
									stack.pop_back();
									drop--;
							}
							stack.push_back(num);
					}
					stack.resize(k);
					return stack;
			}
	
			// Compara se seq1 a partir de i é maior que seq2 a partir de j
			bool greater(vector<int>& seq1, int i, vector<int>& seq2, int j) {
					while (i < seq1.size() && j < seq2.size()) {
							if (seq1[i] > seq2[j]) return true;
							if (seq1[i] < seq2[j]) return false;
							i++; j++;
					}
					return (i != seq1.size());
			}
	
			// Intercala seq1 e seq2 para formar o maior número
			vector<int> merge(vector<int>& seq1, vector<int>& seq2) {
					vector<int> merged;
					int i = 0, j = 0;
					while (i < seq1.size() || j < seq2.size()) {
							if (greater(seq1, i, seq2, j)) {
									merged.push_back(seq1[i++]);
							} else {
									merged.push_back(seq2[j++]);
							}
					}
					return merged;
			}
	
			// Função principal
			vector<int> maxNumber(vector<int>& nums1, vector<int>& nums2, int k) {
					vector<int> best;
					int m = nums1.size(), n = nums2.size();
	
					for (int i = max(0, k - n); i <= min(k, m); ++i) {
							vector<int> seq1 = getMaxSubsequence(nums1, i);
							vector<int> seq2 = getMaxSubsequence(nums2, k - i);
							vector<int> candidate = merge(seq1, seq2);
							if (greater(candidate, 0, best, 0)) {
									best = candidate;
							}
					}
					return best;
			}
	};
	