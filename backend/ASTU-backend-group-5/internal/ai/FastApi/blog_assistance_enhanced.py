from langchain_community.chat_models import ChatOpenAI
from langchain_core.tools import tool
from langchain.agents import initialize_agent,AgentExecutor, create_openai_tools_agent
from langchain.agents import AgentType
from langchain import hub

from langchain.schema import SystemMessage

from langchain_google_genai import ChatGoogleGenerativeAI

from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder
from langchain_core.messages import AIMessage, HumanMessage



import os
from dotenv import load_dotenv
load_dotenv()

os.environ["OPENAI_API_KEY"] = os.getenv("OPENAI_API_KEY")



     

llm = ChatOpenAI(model="gpt-3.5-turbo", temperature=0)

# os.environ["GOOGLE_API_KEY"] = os.getenv("GEMINI_API_KEY")
# llm = ChatGoogleGenerativeAI(
#     model="gemini-pro",
#     convert_system_message_to_human=True,
#     handle_parsing_errors=True,
#     # temperature=0,
#     max_tokens= 2000,
# )




class BlogAssistant:
    def __init__(self, tools):
        self.tools = [tools]
        
        self.prompt = ChatPromptTemplate.from_messages(
    [
        ("system", "You are hepful Blog Writter for any given topic You crafts perfect blog. you use various tools to make the perfect blog. you have acces to blog object which has title and content. title is a string adn content is an array. the content is an array because you can add differecnt contents iterativelt, like images, codes and descriptive texts. you have acces tools to add content to the blog, to sent tiltle, to generate image, and others. you never respond to prompts other that chats. to create a blog first you generate an outline of hat should be covered and follow that also you consider what the size of the blog should be Never generate images more than 4 times. you blog should not be short at leat it has to contain 4000 characters"),
        MessagesPlaceholder("chat_history", optional=True),
        ("human", "{input}"),
        MessagesPlaceholder("agent_scratchpad"),
    ]
)
        self.agent = create_openai_tools_agent(
            tools = tools,
            llm = llm,
            prompt = self.prompt,

        )
        self.agent_executor = AgentExecutor(agent=self.agent, tools=tools, verbose=True, max_iterations=30, return_intermediate_steps=True)

    def run(self, query):
        self.agent_executor.invoke({"input": query.query })


