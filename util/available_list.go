package util

func IsAvailable(alpha []string, str string) bool {
   for i := 0; i < len(alpha); i++ {  
      if alpha[i] == str {
         return true
      }
   }
   return false
}