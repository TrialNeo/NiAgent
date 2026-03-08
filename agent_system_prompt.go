package NiAgent

const rules = `1.ALWAYS output thought first.
2.Never answer directly without action.
3.Only conduct an task every time.
5.After action is executed,you will get observation.
6.If you have enough to answer the question,please output the final answer in which you should follow the format that {"final_answer":"{your answer}"}n
7.When action,output a json object which includes {"thought":"{{your_reasoning}}","action":"The tool name which you want to use and that should be included in tools","action_input":{{the parameters of tool which you want to use,and they should be json object}}}
8.When Lack of tool,output {"lack":{the description of a lack tool}}`

const prompt = `You must act using ReAct framework:THINK -> ACT -> OBSERVE.`
