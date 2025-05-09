class Solution {
public:
    bool canJump(vector<int>& nums) {
        int maxReach = 0; // Armazena o índice máximo que podemos alcançar
        int n = nums.size();
        
        for (int i = 0; i < n; i++) {
            if (i > maxReach) return false; // Se chegarmos a um ponto inalcançável, retornamos falso
            maxReach = max(maxReach, i + nums[i]); // Atualiza o alcance máximo
            if (maxReach >= n - 1) return true; // Se conseguirmos alcançar ou ultrapassar o último índice, retornamos verdadeiro
        }
        
        return false;
    }
};
