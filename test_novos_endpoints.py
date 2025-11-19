#!/usr/bin/env python3
"""
Script de teste para os novos endpoints:
- Relatório Fotográfico
- Diário Semanal
"""

import requests
import json
from datetime import datetime, timedelta

# Configurações
BASE_URL = "http://localhost:9090"

def criar_usuario():
    """Cria um usuário de teste"""
    print("📝 Criando usuário de teste...")
    
    payload = {
        "nome": "Usuario Teste",
        "email": "teste@teste.com",
        "senha": "senha123",
        "perfil_acesso": "ADMIN"
    }
    
    response = requests.post(f"{BASE_URL}/usuarios", json=payload)
    
    if response.status_code == 201:
        print("✅ Usuário criado com sucesso!")
        return True
    elif response.status_code == 400:
        # Usuário já existe - isso é OK
        print("ℹ️  Usuário já existe (isso é normal)")
        return True
    else:
        print(f"❌ Erro ao criar usuário: {response.status_code}")
        print(response.text)
        return False

def fazer_login():
    """Faz login e retorna o token"""
    print("\n🔐 Fazendo login...")
    
    payload = {
        "email": "teste@teste.com",
        "senha": "senha123"
    }
    
    response = requests.post(f"{BASE_URL}/login", json=payload)
    
    if response.status_code == 200:
        data = response.json()
        token = data.get("access_token") or data.get("data", {}).get("token") or data.get("token")
        
        if token:
            print("✅ Login realizado com sucesso!")
            print(f"🎫 Token obtido: {token[:50]}...")
            return token
        else:
            print("❌ Token não encontrado na resposta")
            print(f"Resposta: {json.dumps(data, indent=2)}")
            return None
    else:
        print(f"❌ Erro ao fazer login: {response.status_code}")
        print(response.text)
        return None

def listar_obras(token):
    """Lista as obras disponíveis"""
    print("\n📋 Listando obras...")
    
    headers = {
        "Authorization": f"Bearer {token}",
        "Content-Type": "application/json"
    }
    
    response = requests.get(f"{BASE_URL}/obras", headers=headers)
    
    if response.status_code == 200:
        data = response.json()
        obras = data.get("data", [])
        
        if obras:
            print(f"✅ Encontradas {len(obras)} obra(s):")
            for obra in obras[:3]:  # Mostrar apenas as 3 primeiras
                # Tratar diferentes formatos de resposta
                obra_id = obra.get("id")
                if isinstance(obra_id, dict):
                    obra_id = obra_id.get('Int64', 'N/A')
                    
                nome = obra.get("nome")
                if isinstance(nome, dict):
                    nome = nome.get('String', 'N/A')
                
                print(f"   - ID: {obra_id} | Nome: {nome}")
            
            # Retornar o ID da primeira obra
            first_id = obras[0].get("id")
            if isinstance(first_id, dict) and first_id.get("Valid"):
                return first_id["Int64"]
            elif isinstance(first_id, int):
                return first_id
        else:
            print("⚠️  Nenhuma obra encontrada")
            return None
    else:
        print(f"❌ Erro ao listar obras: {response.status_code}")
        print(response.text)
        return None

def testar_relatorio_fotografico(token, obra_id):
    """Testa o endpoint de relatório fotográfico"""
    print(f"\n📸 Testando Relatório Fotográfico (Obra ID: {obra_id})...")
    
    headers = {
        "Authorization": f"Bearer {token}",
        "Content-Type": "application/json"
    }
    
    response = requests.get(
        f"{BASE_URL}/relatorios/fotografico/{obra_id}",
        headers=headers
    )
    
    print(f"Status: {response.status_code}")
    
    if response.status_code == 200:
        data = response.json()
        relatorio = data.get("data", {})
        
        print("✅ Relatório Fotográfico obtido com sucesso!")
        print(f"\n📊 Estrutura do Relatório:")
        print(f"   - Empresa: {relatorio.get('cabecalho_empresa', {}).get('nome_empresa')}")
        print(f"   - Obra: {relatorio.get('resumo_obra', {}).get('nome_obra')}")
        print(f"   - Localização: {relatorio.get('resumo_obra', {}).get('localizacao')}")
        print(f"   - Total de Fotos: {len(relatorio.get('fotos', []))}")
        
        # Mostrar detalhes de algumas fotos
        fotos = relatorio.get('fotos', [])
        if fotos:
            print(f"\n📷 Primeiras fotos:")
            for i, foto in enumerate(fotos[:3], 1):
                titulo = foto.get('titulo_legenda')
                if isinstance(titulo, dict):
                    titulo = titulo.get('String', 'Sem título')
                
                data = foto.get('data')
                if isinstance(data, dict):
                    data = data.get('String', 'N/A')
                
                print(f"   {i}. {titulo or 'Sem título'}")
                print(f"      Data: {data or 'N/A'}")
                print(f"      URL: {foto.get('url', '')[:60]}...")
        else:
            print("   ⚠️  Nenhuma foto encontrada para esta obra")
        
        return True
    elif response.status_code == 404:
        print("⚠️  Obra não encontrada")
        print(response.text)
        return False
    else:
        print(f"❌ Erro: {response.status_code}")
        print(response.text)
        return False

def testar_diario_semanal(token, obra_id):
    """Testa o endpoint de diário semanal"""
    print(f"\n📅 Testando Diário Semanal (Obra ID: {obra_id})...")
    
    # Calcular datas (últimos 30 dias)
    data_fim = datetime.now()
    data_inicio = data_fim - timedelta(days=30)
    
    payload = {
        "obra_id": obra_id,
        "data_inicio": data_inicio.strftime("%Y-%m-%d"),
        "data_fim": data_fim.strftime("%Y-%m-%d")
    }
    
    headers = {
        "Authorization": f"Bearer {token}",
        "Content-Type": "application/json"
    }
    
    print(f"   Período: {payload['data_inicio']} até {payload['data_fim']}")
    
    response = requests.post(
        f"{BASE_URL}/diarios/semanal",
        headers=headers,
        json=payload
    )
    
    print(f"Status: {response.status_code}")
    
    if response.status_code == 200:
        data = response.json()
        diario = data.get("data", {})
        
        print("✅ Diário Semanal obtido com sucesso!")
        print(f"\n📊 Estrutura do Diário:")
        
        dados_obra = diario.get('dados_obra', {})
        print(f"   - Obra: {dados_obra.get('nome_obra')}")
        print(f"   - Localização: {dados_obra.get('localizacao')}")
        
        contratante = dados_obra.get('contratante')
        if isinstance(contratante, dict):
            contratante = contratante.get('String', 'N/A')
        print(f"   - Contratante: {contratante or 'N/A'}")
        
        contratada = dados_obra.get('contratada')
        if isinstance(contratada, dict):
            contratada = contratada.get('String', 'N/A')
        print(f"   - Contratada: {contratada or 'N/A'}")
        
        semanas = diario.get('semanas', [])
        print(f"\n📆 Total de Semanas: {len(semanas)}")
        
        if semanas:
            print(f"\n📋 Detalhes das Semanas:")
            for semana in semanas[:5]:  # Mostrar apenas as 5 primeiras
                print(f"\n   Semana {semana.get('numero')}:")
                print(f"   - Período: {semana.get('data_inicio')} a {semana.get('data_fim')}")
                print(f"   - Dias de trabalho: {len(semana.get('dias_trabalho', []))}")
                
                descricao = semana.get('descricao', {}).get('String', '')
                if descricao:
                    # Mostrar apenas as primeiras 100 caracteres
                    desc_preview = descricao[:100] + "..." if len(descricao) > 100 else descricao
                    print(f"   - Descrição: {desc_preview}")
                else:
                    print(f"   - Descrição: (vazia)")
        else:
            print("   ⚠️  Nenhuma semana com atividades registradas no período")
        
        return True
    elif response.status_code == 404:
        print("⚠️  Obra não encontrada")
        print(response.text)
        return False
    else:
        print(f"❌ Erro: {response.status_code}")
        print(response.text)
        return False

def main():
    """Função principal"""
    print("=" * 70)
    print("🧪 TESTE DOS NOVOS ENDPOINTS")
    print("   - Relatório Fotográfico")
    print("   - Diário Semanal")
    print("=" * 70)
    
    # 1. Criar usuário
    if not criar_usuario():
        print("\n❌ Falha ao criar usuário. Abortando testes.")
        return
    
    # 2. Fazer login
    token = fazer_login()
    if not token:
        print("\n❌ Falha ao fazer login. Abortando testes.")
        return
    
    # 3. Listar obras e pegar o ID da primeira
    obra_id = listar_obras(token)
    if not obra_id:
        print("\n⚠️  Nenhuma obra disponível para teste.")
        print("💡 Crie uma obra primeiro para testar os relatórios.")
        return
    
    # 4. Testar Relatório Fotográfico
    sucesso_foto = testar_relatorio_fotografico(token, obra_id)
    
    # 5. Testar Diário Semanal
    sucesso_diario = testar_diario_semanal(token, obra_id)
    
    # Resumo final
    print("\n" + "=" * 70)
    print("📊 RESUMO DOS TESTES")
    print("=" * 70)
    print(f"✅ Relatório Fotográfico: {'PASSOU' if sucesso_foto else 'FALHOU'}")
    print(f"✅ Diário Semanal: {'PASSOU' if sucesso_diario else 'FALHOU'}")
    print("=" * 70)
    
    if sucesso_foto and sucesso_diario:
        print("\n🎉 Todos os testes passaram com sucesso!")
    else:
        print("\n⚠️  Alguns testes falharam. Verifique os detalhes acima.")

if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\n\n⚠️  Teste interrompido pelo usuário")
    except Exception as e:
        print(f"\n❌ Erro inesperado: {e}")
        import traceback
        traceback.print_exc()
