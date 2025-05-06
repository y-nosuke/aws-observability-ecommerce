# アーキテクチャ図生成スクリプト 実行手順

この手順は、提供されたPythonスクリプト (`architecture.py`) を実行し、SVG形式のアーキテクチャ図を生成するためのものです。

**必要なソフトウェアとライブラリ:**

1. **Python:** バージョン 3.6 以上が必要です。
2. **pip:** Pythonパッケージインストーラー。通常、Python 3.4以降には同梱されています。
3. **Graphviz:** 図の描画エンジン。`diagrams`ライブラリが内部で利用します。
4. **diagrams (Pythonライブラリ):** 図を生成するためのPythonライブラリ。

**実行手順:**

1. **Pythonのインストール確認:**
    ターミナル（コマンドプロンプト、PowerShell等）を開き、以下を実行してPython 3.6以上がインストールされていることを確認します。

    ```bash
    python --version
    # または
    python3 --version
    # または (Windowsの場合)
    py --version
    ```

    インストールされていない場合は、[Python公式サイト](https://www.python.org/)からダウンロードしてインストールしてください。

2. **pipのインストール確認:**
    ターミナルで以下を実行してpipが利用可能か確認します。

    ```bash
    pip --version
    # または
    pip3 --version
    # または (Windowsの場合)
    py -m pip --version
    ```

    もしインストールされていない場合は、前の回答で説明した `ensurepip` または `get-pip.py` の方法でインストールしてください。

3. **Graphvizのインストール:**
    **重要:** `diagrams` ライブラリはこのソフトウェアに依存しています。
    * **macOS (Homebrew使用):**

        ```bash
        brew install graphviz
        ```

    * **Ubuntu/Debian:**

        ```bash
        sudo apt update && sudo apt install graphviz -y
        ```

    * **Windows:**
        * [Graphviz公式サイト](https://graphviz.org/download/) からインストーラーをダウンロードしてインストールします。
        * **インストール中に「Add Graphviz to the system PATH」系のオプションがあれば必ずチェックを入れてください。**
        * インストール後、コマンドプロンプト等で `dot -V` を実行し、バージョン情報が表示されることを確認します（PCの再起動が必要な場合もあります）。

4. **プロジェクトディレクトリの準備:**
    ターミナルで、提供されたPythonコード (`architecture.py`) を保存するためのディレクトリを作成し、そこに移動します。

    ```bash
    # 例
    mkdir my_diagram_project
    cd my_diagram_project
    ```

5. **Python仮想環境の作成とアクティベート:**
    プロジェクトごとに独立した環境を作ることを強く推奨します。

    ```bash
    # 仮想環境を作成 (.venvはディレクトリ名)
    python3 -m venv .venv
    # または python -m venv .venv

    # 仮想環境をアクティベート
    # macOS / Linux (bash, zsh):
    source .venv/bin/activate
    # Windows (コマンドプロンプト):
    # .\.venv\Scripts\activate.bat
    # Windows (PowerShell):
    # .\.venv\Scripts\Activate.ps1
    ```

    プロンプトの先頭に `(.venv)` が表示されればアクティベート成功です。

6. **`diagrams` ライブラリのインストール:**
    **仮想環境がアクティベートされた状態で**、以下を実行します。

    ```bash
    pip install diagrams
    # または pip3 install diagrams
    ```

7. **Pythonコードの配置:**
    提供されたPythonスクリプトの内容を `architecture.py` という名前で、現在のプロジェクトディレクトリ（`my_diagram_project`）内に保存します。

8. **スクリプトの実行:**
    **仮想環境がアクティベートされた状態で**、以下を実行します。

    ```bash
    python architecture.py
    # または python3 architecture.py
    ```

9. **結果の確認:**
    エラーが発生しなければ、同じディレクトリに `aws_mvp_architecture.svg` という名前のSVGファイルが生成されます。このファイルをWebブラウザなどで開いて図を確認してください。

10. **仮想環境の終了 (任意):**
    作業が終わったら、ターミナルで `deactivate` コマンドを実行して仮想環境を終了します。

## 参考

* [Diagrams](https://diagrams.mingrammer.com/)
