# yaml字符流

- yaml字符流可以包含多个document，每个document都是独立的

## 文档的前缀

- document前缀可以指定文档的字符编码格式
- 前缀是可以带注释行的
- 一个字节流里的所有文档都应该使用同一个编码格式
  - 每个文档有一个字节序用于标记且重新指定新的编码格式

```yaml
# 前面两行都属于文档前缀

⇔# Comment
# lines
Document

%YAML 1.2
---
!!str "Document"
```

## 文档标记符

- 这个地方使用指令会引起二义性，所以文档标记符不使用指令
- %用表标记一行的开始
  - %行后面，是指令部分，指令后面就是文档内容
  - 指令和文档分割是---
  - 一个文档的结束用...

```yaml
%YAML 1.2
---             # 指令结束标记
Document
... # Suffix    # 文档结束标记

%YAML 1.2
---
!!str "Document"
```

## 文档的内容

- 文档内容是不包含指令的
- 就是第一个非注释行，且不是用%开头的，就是文档内容的起始
- 如果父文档node有缩进，子文档node也需要缩进

```yaml
# 总共有3个文档
# 追后空文档没有显示出来

Bare
document
...
# No document
...
|
%!PS-Adobe-2.0 # Not the first line

%YAML 1.2
---
!!str "Bare document"
%YAML 1.2
---
!!str "%!PS-Adobe-2.0\n"
```

## 明晰的文档

- 这种文档说的是有指令结束符但没有具体的指令
- 文档可能完全为空，就像上面例子中的第二个文档

```yaml
# 以---开头的都是明晰文档

---
{ matches
% : 20 }
...
---         # 空文档
# Empty
...

%YAML 1.2
---
!!map {
  !!str "matches %": !!int "20"
}
...
%YAML 1.2
---
!!null ""
```

## 指令文档

- 指令文档是指文档带有指令，和明晰文档正好对应，明晰文档是没有指令的

```yaml
# 下面就是带有YAML指令

%YAML 1.2
--- |
%!PS-Adobe-2.0
...
%YAML1.2
---
# Empty
...
```

## 字节流

- 字节流可能包含多个文档
- 文档和文档之间要有明确的文档结束标记，也就是"..."
  - 如果文档没有显式用"..."结尾，那后面的文档要是明晰文档，也就是用"---"开始
- 流格式的稀松，是为了更好地支持常用场景，eg：流级联

```yaml
Document
---
# Empty
...
%YAML 1.2
---
matches %: 20


%YAML 1.2
---
!!str "Document"
...
%YAML 1.2
---
!!null ""
...
%YAML 1.2
---
!!map {
  !!str "matches %": !!int "20"
}
```

- 流里是字节序列，正好符合yaml流的设计，并不是说只能是字节序列
- 下面3种场景使用字节流都非常合适，这也是yaml流使用字节流的原因：
  - 流追加，文档的追加很方便，因为文档之间都是独立的
    - 在用yaml存日志类时，非常方便，因为日志就是一段段的
  - 流接续，接续的前提是两个流的编码格式是一致的
    - 接续只要保证有明确的文档分割就行(不管是文档结束符还是后面文档是明晰文档)
  - 流交互，可以在不关闭流/或者开启下一个文档的情况下，发送文档结束标记符
    - 好处是，流接收者可以直接处理这个结束的文档，而不用等待流结束或下一个文档开始的信号
