#include<bits/stdc++.h>
#include<string.h>

using namespace std;

struct jsonFormat {
    string key, value;
};

string formateSentence( string id, string header , string value) {
    string val = value;
    if(value[0] != '_') val = "\"" + val + "\"";
    return id + " <" + header + "> " + val + " .\n";
}

vector<string> separate(string a, string separator) {
    vector<string>ans;
    int n = a.size();
    for(int i = 0;i <n;i+=3){
        string ac;
        while(a[i] != separator[0]){
            string c;
            c.push_back(a[i]);
            c.push_back(a[i+1]);
            if(c == separator) break;
            ac.push_back(a[i]);
            i++;
        }
        if(ac.size()){
            ans.push_back(ac);
        }
    }
    return ans;
}

string currentDate() {
    time_t now = time(0);
    tm *ltm = localtime(&now);
    char ans[100];
    sprintf(ans,"%04d-%02d-%02d",1900+ltm->tm_year,1+ ltm->tm_mon, ltm->tm_mday);
    return ans;
}

vector<string> TransfromCsv(vector<string>lines, char oldSeparator, char newSeparator) {
    for(string& i:lines) {
        bool inString = false;
        for(char& actual: i) {
            if(actual == '"') inString = !inString;
            if(!inString && oldSeparator == actual) {
                actual = newSeparator;
            }
        }
    }
    return lines;
}


vector<jsonFormat> separateJson(string json) {
    vector<jsonFormat> ans;
    int i = 0;
    int n = json.size();
    set<char>skipable = {'[', '"', ']', ',','{'};
    set<char>terminator = {'}',',',':'};
    while(i < n) {
        if(skipable.count(json[i])){
            i++;
            continue;
        }
        jsonFormat actual;
        for(int j = 0;j < 2;j++){
            string temp;
            while(i < n && !terminator.count(json[i])) {
                if(json[i] == '"'){
                    i++;
                    continue;
                }
                temp.push_back(json[i]);
                i++;
            }
            i++;
            if( j == 0) {
                actual.key = temp;
            } else {
                actual.value = temp;
            }
        }
        ans.push_back(actual);
    }
    return ans;
}


vector<string> getProducts(string a) {
    vector<string>ans;
    for(int i = 1;i < (int)a.size()-1;i++) {
        string ac;
        while(i < (int)a.size() -1 && a[i] != ','){
            ac.push_back(a[i]);
            i++;
        }
        ans.push_back(ac);
    }
    return ans;
}

void buyer(ofstream &out, string date) {
    ifstream file;
    file.open("buyer.json");
    string data;
    file >> data;
    auto jsonParser = separateJson(data);
    int n = jsonParser.size();
    for(int i = 0;i < n/3;i++) {
        string id = "_:" + jsonParser[i*3].value;
        for(int j = 1;j < 3;j++) {
            string val = jsonParser[i*3+j].value;
            out << "    " << formateSentence(id,jsonParser[i*3+j].key,jsonParser[i*3+j].value); 
        }

        out << "    " << formateSentence(id,"DType", "Buyer");
        out << "    " << formateSentence(id, "Date", date);
    }
    file.close();
}

vector<string> split(string s,char delimiter) {
    vector<string>ans;
    for(int i = 0;i < (int)s.size();i++) {
        string temp;
        while(i < (int)s.size() && s[i] != delimiter){
            if(s[i] == '"'){
                i++;
                continue;
            }
            temp.push_back(s[i]);
            i++;
        }
        ans.push_back(temp);
    }
    return ans;
}

vector<vector<string>> getProducts() {
    ifstream file;
    file.open("Products.csv");
    vector<string>raw_data;
    string line;
    while(getline(file,line)) {
        raw_data.push_back(line);
    }

    auto data = TransfromCsv(raw_data, '\'', ',');

    vector<vector<string>> ans;
    for(string& i: data){
        ans.push_back(split(i,','));
    }
    file.close();
    return ans;
}

void products(ofstream &out, string date) {
    vector<vector<string>> data = getProducts();
    vector<string> headers = { "id", "name", "price" };

    for(vector<string>&i: data){
        string id = "_:" + i[0];
        for(int j = 1; j < 3;j++) {
            out << "    " + formateSentence(id,headers[j],i[j]);
        }
        out << "    " + formateSentence(id,"DType","Product");
        out << "    " + formateSentence(id,"Date",date);
    }
}

void transaction(ofstream &out, string date) {
    ifstream file;
    file.open("transaction.txt");
    string actual;
    int count = 0;
    vector<string> names = {"id","buyer_id", "ip", "device", "products"};
    while(file >> actual) {
        string separator = "�";
        vector<string> ans = separate(actual,separator);
        cout << ans.size() << endl;
        int n = ans.size();
        for(int i = 0; i < n/5;i++){
            string id = "_:"+ ans[i*5].substr(1,ans[i*5].size()-1);
            for(int j = 1;j < 4;j++) {
                if(j == 1) {
                    out <<"             "<< formateSentence(id, names[j], "_:" + ans[i*5+j]);
                } else {
                    out <<"             "<< formateSentence(id, names[j], ans[i*5+j]);
                }
            }
            out << "             " << formateSentence(id, "DType", "Transaction");
            out << "             " << formateSentence(id, "Date", date);
            auto products = getProducts(ans[i*5+4]);
            for(string& i: products) {
                i = "_:" + i;
                out <<"             "<< formateSentence(id, "Products", i);
            }
        }
        count++;
    }
    file.close();
}


int main() {
    ofstream query;
    string date = currentDate();
    query.open("query.json");
    query << "{ set \n     {" << endl;
    buyer(query, date);
    products(query, date);
    transaction(query, date);
    query << "  }" << endl << "}";
    query.close();
}
