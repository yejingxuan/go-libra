//Create By JFUI Engine ver 3.0.1 2018-04-02 14:11:40
function BarrettMu(A) {
    this.modulus = biCopy(A);
    this.k = biHighIndex(this.modulus) + 1;
    var B = new BigInt();
    B.digits[2 * this.k] = 1;
    this.mu = biDivide(B, this.modulus);
    this.bkplus1 = new BigInt();
    this.bkplus1.digits[this.k + 1] = 1;
    this.modulo = BarrettMu_modulo;
    this.multiplyMod = BarrettMu_multiplyMod;
    this.powMod = BarrettMu_powMod
}

function BarrettMu_modulo(H) {
    var G = biDivideByRadixPower(H, this.k - 1);
    var E = biMultiply(G, this.mu);
    var D = biDivideByRadixPower(E, this.k + 1);
    var C = biModuloByRadixPower(H, this.k + 1);
    var I = biMultiply(D, this.modulus);
    var B = biModuloByRadixPower(I, this.k + 1);
    var A = biSubtract(C, B);
    if (A.isNeg) {
        A = biAdd(A, this.bkplus1)
    }
    var F = biCompare(A, this.modulus) >= 0;
    while (F) {
        A = biSubtract(A, this.modulus);
        F = biCompare(A, this.modulus) >= 0
    }
    return A
}

function BarrettMu_multiplyMod(A, C) {
    var B = biMultiply(A, C);
    return this.modulo(B)
}

function BarrettMu_powMod(B, E) {
    var A = new BigInt();
    A.digits[0] = 1;
    var C = B;
    var D = E;
    while (true) {
        if ((D.digits[0] & 1) != 0) {
            A = this.multiplyMod(A, C)
        }
        D = biShiftRight(D, 1);
        if (D.digits[0] == 0 && biHighIndex(D) == 0) {
            break
        }
        C = this.multiplyMod(C, C)
    }
    return A
}

var biRadixBase = 2;
var biRadixBits = 16;
var bitsPerDigit = biRadixBits;
var biRadix = 1 << 16;
var biHalfRadix = biRadix >>> 1;
var biRadixSquared = biRadix * biRadix;
var maxDigitVal = biRadix - 1;
var maxInteger = 9999999999999998;
var maxDigits;
var ZERO_ARRAY;
var bigZero, bigOne;

function setMaxDigits(B) {
    maxDigits = B;
    ZERO_ARRAY = new Array(maxDigits);
    for (var A = 0; A < ZERO_ARRAY.length; A++) {
        ZERO_ARRAY[A] = 0
    }
    bigZero = new BigInt();
    bigOne = new BigInt();
    bigOne.digits[0] = 1
}

setMaxDigits(20);
var dpl10 = 15;
var lr10 = biFromNumber(1000000000000000);

function BigInt(A) {
    if (typeof A == "boolean" && A == true) {
        this.digits = null
    } else {
        this.digits = ZERO_ARRAY.slice(0)
    }
    this.isNeg = false
}

function biFromDecimal(E) {
    var D = E.charAt(0) == "-";
    var C = D ? 1 : 0;
    var A;
    while (C < E.length && E.charAt(C) == "0") {
        ++C
    }
    if (C == E.length) {
        A = new BigInt()
    } else {
        var B = E.length - C;
        var F = B % dpl10;
        if (F == 0) {
            F = dpl10
        }
        A = biFromNumber(Number(E.substr(C, F)));
        C += F;
        while (C < E.length) {
            A = biAdd(biMultiply(A, lr10), biFromNumber(Number(E.substr(C, dpl10))));
            C += dpl10
        }
        A.isNeg = D
    }
    return A
}

function biCopy(B) {
    var A = new BigInt(true);
    A.digits = B.digits.slice(0);
    A.isNeg = B.isNeg;
    return A
}

function biFromNumber(C) {
    var A = new BigInt();
    A.isNeg = C < 0;
    C = Math.abs(C);
    var B = 0;
    while (C > 0) {
        A.digits[B++] = C & maxDigitVal;
        C = Math.floor(C / biRadix)
    }
    return A
}

function reverseStr(C) {
    var A = "";
    for (var B = C.length - 1; B > -1; --B) {
        A += C.charAt(B)
    }
    return A
}

var hexatrigesimalToChar = new Array("0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z");

function biToString(C, E) {
    var B = new BigInt();
    B.digits[0] = E;
    var D = biDivideModulo(C, B);
    var A = hexatrigesimalToChar[D[1].digits[0]];
    while (biCompare(D[0], bigZero) == 1) {
        D = biDivideModulo(D[0], B);
        digit = D[1].digits[0];
        A += hexatrigesimalToChar[D[1].digits[0]]
    }
    return (C.isNeg ? "-" : "") + reverseStr(A)
}

function biToDecimal(C) {
    var B = new BigInt();
    B.digits[0] = 10;
    var D = biDivideModulo(C, B);
    var A = String(D[1].digits[0]);
    while (biCompare(D[0], bigZero) == 1) {
        D = biDivideModulo(D[0], B);
        A += String(D[1].digits[0])
    }
    return (C.isNeg ? "-" : "") + reverseStr(A)
}

var hexToChar = new Array("0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f");

function digitToHex(C) {
    var B = 15;
    var A = "";
    for (i = 0; i < 4; ++i) {
        A += hexToChar[C & B];
        C >>>= 4
    }
    return reverseStr(A)
}

function biToHex(B) {
    var A = "";
    var D = biHighIndex(B);
    for (var C = biHighIndex(B); C > -1; --C) {
        A += digitToHex(B.digits[C])
    }
    return A
}

function charToHex(H) {
    var C = 48;
    var B = C + 9;
    var D = 97;
    var G = D + 25;
    var F = 65;
    var E = 65 + 25;
    var A;
    if (H >= C && H <= B) {
        A = H - C
    } else {
        if (H >= F && H <= E) {
            A = 10 + H - F
        } else {
            if (H >= D && H <= G) {
                A = 10 + H - D
            } else {
                A = 0
            }
        }
    }
    return A
}

function hexToDigit(D) {
    var B = 0;
    var A = Math.min(D.length, 4);
    for (var C = 0; C < A; ++C) {
        B <<= 4;
        B |= charToHex(D.charCodeAt(C))
    }
    return B
}

function biFromHex(E) {
    var B = new BigInt();
    var A = E.length;
    for (var D = A, C = 0; D > 0; D -= 4, ++C) {
        B.digits[C] = hexToDigit(E.substr(Math.max(D - 4, 0), Math.min(D, 4)))
    }
    return B
}

function biFromString(I, H) {
    var A = I.charAt(0) == "-";
    var D = A ? 1 : 0;
    var J = new BigInt();
    var B = new BigInt();
    B.digits[0] = 1;
    for (var C = I.length - 1; C >= D; C--) {
        var E = I.charCodeAt(C);
        var F = charToHex(E);
        var G = biMultiplyDigit(B, F);
        J = biAdd(J, G);
        B = biMultiplyDigit(B, H)
    }
    J.isNeg = A;
    return J
}

function biDump(A) {
    return (A.isNeg ? "-" : "") + A.digits.join(" ")
}

function biAdd(B, F) {
    var A;
    if (B.isNeg != F.isNeg) {
        F.isNeg = !F.isNeg;
        A = biSubtract(B, F);
        F.isNeg = !F.isNeg
    } else {
        A = new BigInt();
        var E = 0;
        var D;
        for (var C = 0; C < B.digits.length; ++C) {
            D = B.digits[C] + F.digits[C] + E;
            A.digits[C] = D % biRadix;
            E = Number(D >= biRadix)
        }
        A.isNeg = B.isNeg
    }
    return A
}

function biSubtract(B, F) {
    var A;
    if (B.isNeg != F.isNeg) {
        F.isNeg = !F.isNeg;
        A = biAdd(B, F);
        F.isNeg = !F.isNeg
    } else {
        A = new BigInt();
        var E, D;
        D = 0;
        for (var C = 0; C < B.digits.length; ++C) {
            E = B.digits[C] - F.digits[C] + D;
            A.digits[C] = E % biRadix;
            if (A.digits[C] < 0) {
                A.digits[C] += biRadix
            }
            D = 0 - Number(E < 0)
        }
        if (D == -1) {
            D = 0;
            for (var C = 0; C < B.digits.length; ++C) {
                E = 0 - A.digits[C] + D;
                A.digits[C] = E % biRadix;
                if (A.digits[C] < 0) {
                    A.digits[C] += biRadix
                }
                D = 0 - Number(E < 0)
            }
            A.isNeg = !B.isNeg
        } else {
            A.isNeg = B.isNeg
        }
    }
    return A
}

function biHighIndex(B) {
    var A = B.digits.length - 1;
    while (A > 0 && B.digits[A] == 0) {
        --A
    }
    return A
}

function biNumBits(C) {
    var E = biHighIndex(C);
    var D = C.digits[E];
    var B = (E + 1) * bitsPerDigit;
    var A;
    for (A = B; A > B - bitsPerDigit; --A) {
        if ((D & 32768) != 0) {
            break
        }
        D <<= 1
    }
    return A
}

function biMultiply(G, F) {
    var J = new BigInt();
    var E;
    var B = biHighIndex(G);
    var I = biHighIndex(F);
    var H, A, C;
    for (var D = 0; D <= I; ++D) {
        E = 0;
        C = D;
        for (j = 0; j <= B; ++j, ++C) {
            A = J.digits[C] + G.digits[j] * F.digits[D] + E;
            J.digits[C] = A & maxDigitVal;
            E = A >>> biRadixBits
        }
        J.digits[D + B + 1] = E
    }
    J.isNeg = G.isNeg != F.isNeg;
    return J
}

function biMultiplyDigit(A, F) {
    var E, D, C;
    result = new BigInt();
    E = biHighIndex(A);
    D = 0;
    for (var B = 0; B <= E; ++B) {
        C = result.digits[B] + A.digits[B] * F + D;
        result.digits[B] = C & maxDigitVal;
        D = C >>> biRadixBits
    }
    result.digits[1 + E] = D;
    return result
}

function arrayCopy(E, H, C, G, F) {
    var A = Math.min(H + F, E.length);
    for (var D = H, B = G; D < A; ++D, ++B) {
        C[B] = E[D]
    }
}

var highBitMasks = new Array(0, 32768, 49152, 57344, 61440, 63488, 64512, 65024, 65280, 65408, 65472, 65504, 65520, 65528, 65532, 65534, 65535);

function biShiftLeft(B, H) {
    var D = Math.floor(H / bitsPerDigit);
    var A = new BigInt();
    arrayCopy(B.digits, 0, A.digits, D, A.digits.length - D);
    var G = H % bitsPerDigit;
    var C = bitsPerDigit - G;
    for (var E = A.digits.length - 1, F = E - 1; E > 0; --E, --F) {
        A.digits[E] = ((A.digits[E] << G) & maxDigitVal) | ((A.digits[F] & highBitMasks[G]) >>> (C))
    }
    A.digits[0] = ((A.digits[E] << G) & maxDigitVal);
    A.isNeg = B.isNeg;
    return A
}

var lowBitMasks = new Array(0, 1, 3, 7, 15, 31, 63, 127, 255, 511, 1023, 2047, 4095, 8191, 16383, 32767, 65535);

function biShiftRight(B, H) {
    var C = Math.floor(H / bitsPerDigit);
    var A = new BigInt();
    arrayCopy(B.digits, C, A.digits, 0, B.digits.length - C);
    var F = H % bitsPerDigit;
    var G = bitsPerDigit - F;
    for (var D = 0, E = D + 1; D < A.digits.length - 1; ++D, ++E) {
        A.digits[D] = (A.digits[D] >>> F) | ((A.digits[E] & lowBitMasks[F]) << G)
    }
    A.digits[A.digits.length - 1] >>>= F;
    A.isNeg = B.isNeg;
    return A
}

function biMultiplyByRadixPower(B, C) {
    var A = new BigInt();
    arrayCopy(B.digits, 0, A.digits, C, A.digits.length - C);
    return A
}

function biDivideByRadixPower(B, C) {
    var A = new BigInt();
    arrayCopy(B.digits, C, A.digits, 0, A.digits.length - C);
    return A
}

function biModuloByRadixPower(B, C) {
    var A = new BigInt();
    arrayCopy(B.digits, 0, A.digits, 0, C);
    return A
}

function biCompare(A, C) {
    if (A.isNeg != C.isNeg) {
        return 1 - 2 * Number(A.isNeg)
    }
    for (var B = A.digits.length - 1; B >= 0; --B) {
        if (A.digits[B] != C.digits[B]) {
            if (A.isNeg) {
                return 1 - 2 * Number(A.digits[B] > C.digits[B])
            } else {
                return 1 - 2 * Number(A.digits[B] < C.digits[B])
            }
        }
    }
    return 0
}

function biDivideModulo(F, E) {
    var A = biNumBits(F);
    var D = biNumBits(E);
    var C = E.isNeg;
    var K, J;
    if (A < D) {
        if (F.isNeg) {
            K = biCopy(bigOne);
            K.isNeg = !E.isNeg;
            F.isNeg = false;
            E.isNeg = false;
            J = biSubtract(E, F);
            F.isNeg = true;
            E.isNeg = C
        } else {
            K = new BigInt();
            J = biCopy(F)
        }
        return new Array(K, J)
    }
    K = new BigInt();
    J = F;
    var H = Math.ceil(D / bitsPerDigit) - 1;
    var G = 0;
    while (E.digits[H] < biHalfRadix) {
        E = biShiftLeft(E, 1);
        ++G;
        ++D;
        H = Math.ceil(D / bitsPerDigit) - 1
    }
    J = biShiftLeft(J, G);
    A += G;
    var N = Math.ceil(A / bitsPerDigit) - 1;
    var S = biMultiplyByRadixPower(E, N - H);
    while (biCompare(J, S) != -1) {
        ++K.digits[N - H];
        J = biSubtract(J, S)
    }
    for (var Q = N; Q > H; --Q) {
        var I = (Q >= J.digits.length) ? 0 : J.digits[Q];
        var R = (Q - 1 >= J.digits.length) ? 0 : J.digits[Q - 1];
        var P = (Q - 2 >= J.digits.length) ? 0 : J.digits[Q - 2];
        var O = (H >= E.digits.length) ? 0 : E.digits[H];
        var B = (H - 1 >= E.digits.length) ? 0 : E.digits[H - 1];
        if (I == O) {
            K.digits[Q - H - 1] = maxDigitVal
        } else {
            K.digits[Q - H - 1] = Math.floor((I * biRadix + R) / O)
        }
        var M = K.digits[Q - H - 1] * ((O * biRadix) + B);
        var L = (I * biRadixSquared) + ((R * biRadix) + P);
        while (M > L) {
            --K.digits[Q - H - 1];
            M = K.digits[Q - H - 1] * ((O * biRadix) | B);
            L = (I * biRadix * biRadix) + ((R * biRadix) + P)
        }
        S = biMultiplyByRadixPower(E, Q - H - 1);
        J = biSubtract(J, biMultiplyDigit(S, K.digits[Q - H - 1]));
        if (J.isNeg) {
            J = biAdd(J, S);
            --K.digits[Q - H - 1]
        }
    }
    J = biShiftRight(J, G);
    K.isNeg = F.isNeg != C;
    if (F.isNeg) {
        if (C) {
            K = biAdd(K, bigOne)
        } else {
            K = biSubtract(K, bigOne)
        }
        E = biShiftRight(E, G);
        J = biSubtract(E, J)
    }
    if (J.digits[0] == 0 && biHighIndex(J) == 0) {
        J.isNeg = false
    }
    return new Array(K, J)
}

function biDivide(A, B) {
    return biDivideModulo(A, B)[0]
}

function biModulo(A, B) {
    return biDivideModulo(A, B)[1]
}

function biMultiplyMod(B, C, A) {
    return biModulo(biMultiply(B, C), A)
}

function biPow(B, D) {
    var A = bigOne;
    var C = B;
    while (true) {
        if ((D & 1) != 0) {
            A = biMultiply(A, C)
        }
        D >>= 1;
        if (D == 0) {
            break
        }
        C = biMultiply(C, C)
    }
    return A
}

function biPowMod(C, F, B) {
    var A = bigOne;
    var D = C;
    var E = F;
    while (true) {
        if ((E.digits[0] & 1) != 0) {
            A = biMultiplyMod(A, D, B)
        }
        E = biShiftRight(E, 1);
        if (E.digits[0] == 0 && biHighIndex(E) == 0) {
            break
        }
        D = biMultiplyMod(D, D, B)
    }
    return A
}

function RSAKeyPair(B, C, A) {
    this.e = biFromHex(B);
    this.d = biFromHex(C);
    this.m = biFromHex(A);
    this.chunkSize = 2 * biHighIndex(this.m);
    this.radix = 16;
    this.barrett = new BarrettMu(this.m)
}

function twoDigit(A) {
    return (A < 10 ? "0" : "") + String(A)
}

function encryptPass( K) {
    var key1 = new RSAKeyPair("10001", "", "b582bfab21f625687e985e19d")
    return encryptedString(key1, K)
}

function encryptedString(H, K) {
    var G = new Array();
    var A = K.length;
    var E = 0;
    while (E < A) {
        G[E] = K.charCodeAt(E);
        E++
    }
    while (G.length % H.chunkSize != 0) {
        G[E++] = 0
    }
    var F = G.length;
    var L = "";
    var D, C, B;
    for (E = 0; E < F; E += H.chunkSize) {
        B = new BigInt();
        D = 0;
        for (C = E; C < E + H.chunkSize; ++D) {
            B.digits[D] = G[C++];
            B.digits[D] += G[C++] << 8
        }
        var J = H.barrett.powMod(B, H.e);
        var I = H.radix == 16 ? biToHex(J) : biToString(J, H.radix);
        L += I + " "
    }
    return L.substring(0, L.length - 1)
}

function decryptedString(E, F) {
    var H = F.split(" ");
    var A = "";
    var D, C, G;
    for (D = 0; D < H.length; ++D) {
        var B;
        if (E.radix == 16) {
            B = biFromHex(H[D])
        } else {
            B = biFromString(H[D], E.radix)
        }
        G = E.barrett.powMod(B, E.d);
        for (C = 0; C <= biHighIndex(G); ++C) {
            A += String.fromCharCode(G.digits[C] & 255, G.digits[C] >> 8)
        }
    }
    if (A.charCodeAt(A.length - 1) == 0) {
        A = A.substring(0, A.length - 1)
    }
    return A
}

function doLogin() {
    setMaxDigits(131);
    document.getElementById("password").value = encryptedString(new RSAKeyPair("10001", "", "b582bfab21f625687e985e19d"), document.getElementById("password_show").value);
    document.logForm.submit()
}
