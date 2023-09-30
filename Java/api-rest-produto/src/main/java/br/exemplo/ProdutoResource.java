package br.exemplo;

import jakarta.ws.rs.*;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

@Path("/produto")
public class ProdutoResource {

    List<Produto> itens = new ArrayList<>();

    @GET
    public List<Produto> getTodosProdutos(){
        return itens;
    }

    @POST
    public String postNovoProduto(Produto produto){
        itens.add(produto);
        return "Produto cadastrado com sucesso!";
    }

    @PUT
    @Path("/{id}")
    public String putProduto(long id, Produto produtoAtualizado){
        itens.stream()
                .filter(produto -> produto.getId() == id)
                .findFirst()
                .ifPresent(produto -> {
                    produto.setNome(produtoAtualizado.getNome());
                    produto.setPreco(produtoAtualizado.getPreco());
                });
        return "Atualização feita com sucesso!";
    }

    @DELETE
    @Path("/{id}")
    public String deleteProduto(long id){
        itens = itens.stream()
                .filter(produto -> produto.getId() != id)
                .collect(Collectors.toList());
        return "Exclusão feita com sucesso!";
    }
}
