import http from 'k6/http';
import { check, sleep, group } from 'k6';

// Carregando variáveis de ambiente
const IP_RATE_LIMIT = parseInt(__ENV.IP_RATE_LIMIT || 5); // Valor padrão de 5 se não definido
const BLOCK_TIME_SECONDS = parseInt(__ENV.BLOCK_TIME_SECONDS || 60); // Valor padrão de 60 segundos se não definido
let API_KEYS = __ENV.API_KEYS_RATE_LIMIT || [
  { key: 'key1', rate_limit: 1 },
  { key: 'key2', rate_limit: 2 },
];

// Parseia a string JSON das chaves de API se necessário
try {
  const parsedKeys = JSON.parse(__ENV.API_KEYS_RATE_LIMIT);
  API_KEYS = parsedKeys;
} catch (error) {
  console.error('Erro ao fazer parse das chaves de API:', error.message);
}

const WEB_SERVER_PORT = __ENV.WEB_SERVER_PORT || 8080;

// Configurações do teste
const BASE_URL = `http://app:${WEB_SERVER_PORT}`;

// Imprime um resumo dos parâmetros de teste
let summary = '\n=== Parâmetros do Teste ===\n';
summary += `URL Base: ${BASE_URL}\n`;
summary += `Limite de requisições por IP: ${IP_RATE_LIMIT} req/sec\n`;
summary += `Tempo de bloqueio: ${BLOCK_TIME_SECONDS} segundos\n`;
summary += '\nChaves de API configuradas:\n';

for (let i = 0; i < API_KEYS.length; i++) {
  summary += `- Chave: ${API_KEYS[i].key}, Limite: ${API_KEYS[i].rate_limit} req/sec\n`;
}
summary += '=======================\n\n';

export function setup() {
  console.log(summary);
  return;
}

// Opções do teste
export const options = {
  summaryTrendStats: ['count'],

  scenarios: {
    ip_test: {
      exec: 'testIpLimits',
      executor: 'shared-iterations',
      vus: 1,
      iterations: 1,
      startTime: '0s',
    },
    key_test: {
      exec: 'testApiKeyLimits',
      executor: 'shared-iterations',
      vus: 1,
      iterations: 1,
      startTime: `${BLOCK_TIME_SECONDS * 1}s`,
    },
    invalid_key_test: {
      exec: 'testInvalidApiKey',
      executor: 'shared-iterations',
      vus: 1,
      iterations: 1,
      startTime: `${BLOCK_TIME_SECONDS * 2}s`,
    },
    recovery_test: {
      exec: 'testRecovery',
      executor: 'shared-iterations',
      vus: 1,
      iterations: 1,
      startTime: `${BLOCK_TIME_SECONDS * 3}s`,
    },
  },
};

// Função auxiliar para fazer requisição e verificar o status
function makeRequest(label, headers = {}) {
  const response = http.get(BASE_URL, { headers });
  console.log(`${label} - Status: ${response.status}`);

  return response;
}

// Cenário 1: Testando limites por IP
export function testIpLimits() {
  group('Cenário 1: Testando limite por IP', () => {
    // Teste 1.1: Requisição inicial por IP
    let response = makeRequest('Teste 1.1: Requisição inicial por IP');
    check(response, {
      'Requisição inicial por IP bem-sucedida': (r) => r.status === 200,
    });

    sleep(1);

    // Teste 1.2: Bloqueando IP com múltiplas requisições
    const totalRequests = IP_RATE_LIMIT + 3;
    console.log(
      `Teste 1.2: Bloqueando IP com múltiplas requisições (${totalRequests})`
    );
    for (let i = 0; i < totalRequests; i++) {
      response = makeRequest(`Requisição ${i + 1} por IP`);
    }

    // Verifica se a última requisição foi bloqueada
    check(response, {
      'IP foi bloqueado após múltiplas requisições': (r) => r.status === 429,
    });

    sleep(1);

    // Teste 1.3: Verificando acesso com API_Key quando IP está bloqueado
    response = makeRequest(
      'Teste 1.3: Requisição com API_Key quando IP bloqueado',
      { API_Key: API_KEYS[0].key }
    );
    check(response, {
      'API_Key funciona quando IP está bloqueado': (r) => r.status === 200,
    });

    // Teste 1.4: Verificando acesso com API_Key alternativa
    response = makeRequest('Teste 1.4: Requisição com API_Key alternativa', {
      API_Key: API_KEYS[1].key,
    });
    check(response, {
      'API_Key alternativa também funciona': (r) => r.status === 200,
    });
  });
}

// Cenário 2: Testando limites por API_Key
export function testApiKeyLimits() {
  group('Cenário 2: Testando limite por API_Key', () => {
    // Teste 2.1: Requisição inicial com API_Key
    let response = makeRequest('Teste 2.1: Requisição inicial com API_Key', {
      API_Key: API_KEYS[0].key,
    });
    check(response, {
      'Requisição inicial com API_Key bem-sucedida': (r) => r.status === 200,
    });

    sleep(1);

    // Teste 2.2: Bloqueando API_Key com múltiplas requisições
    const totalRequests = API_KEYS[0].rate_limit + 3;
    console.log(
      `Teste 2.2: Bloqueando API_Key com múltiplas requisições (${totalRequests})`
    );
    for (let i = 0; i < totalRequests; i++) {
      response = makeRequest(`Requisição ${i + 1} com API_Key`, {
        API_Key: API_KEYS[0].key,
      });
    }

    // Verifica se a última requisição foi bloqueada
    check(response, {
      'API_Key foi bloqueada após múltiplas requisições': (r) =>
        r.status === 429,
    });

    sleep(1);

    // Teste 2.3: Verificando acesso com API_Key alternativa
    response = makeRequest('Teste 2.3: Requisição com API_Key alternativa', {
      API_Key: API_KEYS[1].key,
    });
    check(response, {
      'API_Key alternativa funciona quando a primeira está bloqueada': (r) =>
        r.status === 200,
    });

    // Teste 2.4: Verificando acesso apenas por IP
    response = makeRequest('Teste 2.4: Requisição apenas por IP');
    check(response, {
      'Acesso por IP funciona quando API_Key está bloqueada': (r) =>
        r.status === 200,
    });
  });
}

// Cenário 3: Testando API_Key inválida
export function testInvalidApiKey() {
  group('Cenário 3: Testando API_Key inválida', () => {
    // Teste 3.1: Testando com API_Key inválida
    let response = makeRequest('Teste 3.1: Requisição com API_Key inválida', {
      API_Key: 'invalid',
    });
    check(response, {
      'Requisição com API_Key inválida usa limite de IP': (r) =>
        r.status === 200,
    });

    sleep(1);

    // Teste 3.2: Esgotando limite de IP usando API_Key inválida
    const totalRequests = IP_RATE_LIMIT + 3;
    console.log(
      `Teste 3.2: Esgotando limite de IP com API_Key inválida (${totalRequests})`
    );
    for (let i = 0; i < totalRequests; i++) {
      response = makeRequest(`Requisição ${i + 1} com API_Key inválida`, {
        API_Key: 'invalid',
      });
    }

    // Verifica se a última requisição foi bloqueada
    check(response, {
      'IP foi bloqueado usando API_Key inválida': (r) => r.status === 429,
    });
    // Verifica se ip foi bloqueado
    response = makeRequest('Teste 3.3: Requisição com IP bloqueado');
    check(response, {
      'IP foi bloqueado': (r) => r.status === 429,
    });

    sleep(1);

    // Teste 3.3: Verificando que API_Key válida funciona quando IP está bloqueado
    response = makeRequest(
      'Teste 3.3: Requisição com API_Key válida após bloqueio de IP',
      { API_Key: API_KEYS[0].key }
    );
    check(response, {
      'API_Key válida funciona quando IP está bloqueado por requisições com chave inválida':
        (r) => r.status === 200,
    });
  });
}

// Cenário 4: Verificando recuperação após tempo de espera
export function testRecovery() {
  group('Cenário 4: Verificando desbloqueio', () => {
    // Primeiro bloqueia IP com múltiplas requisições
    const totalRequests = IP_RATE_LIMIT + 3;
    console.log(
      `Teste 4.1: Bloqueando IP com múltiplas requisições (${totalRequests})`
    );
    for (let i = 0; i < totalRequests; i++) {
      makeRequest(`Setup: Requisição ${i + 1}`);
    }

    // Espera um tempo para desbloqueio
    console.log('Aguardando desbloqueio do IP...');
    sleep(BLOCK_TIME_SECONDS);

    // Teste 4.1: Verificando acesso por IP após desbloqueio
    let response = makeRequest('Teste 4.1: Requisição por IP após desbloqueio');
    check(response, {
      'Acesso por IP liberado': (r) => r.status === 200,
    });

    // Bloqueia API_Key com múltiplas requisições
    const totalRequestsAPiKey = API_KEYS[0].rate_limit + 3;
    console.log(
      `Teste 4.2: Bloqueando API_Key com múltiplas requisições (${totalRequestsAPiKey})`
    );
    for (let i = 0; i < totalRequestsAPiKey; i++) {
      makeRequest(`Setup: Requisição ${i + 1} com API_Key`, {
        API_Key: API_KEYS[0].key,
      });
    }

    // Espera um tempo para desbloqueio
    console.log('Aguardando desbloqueio da API_Key...');
    sleep(BLOCK_TIME_SECONDS);

    // Teste 4.2: Verificando acesso por API_Key após desbloqueio
    response = makeRequest(
      'Teste 4.2: Requisição com API_Key após desbloqueio',
      { API_Key: API_KEYS[0].key }
    );
    check(response, {
      'Acesso por API_Key liberado': (r) => r.status === 200,
    });
  });
}
