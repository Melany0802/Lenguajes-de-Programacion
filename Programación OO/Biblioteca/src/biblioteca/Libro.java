
package biblioteca;

public class Libro {
    private int codigo;
    private String titulo;
    private String autor;
    private boolean disponibilidad;

    public Libro(int codigo, String titulo, String autor, boolean disponibilidad) {
        this.codigo = codigo;
        this.titulo = titulo;
        this.autor = autor;
        this.disponibilidad = disponibilidad;
    }

    public int getCodigo() {
        return codigo;
    }

    public String getTitulo() {
        return titulo;
    }

    public String getAutor() {
        return autor;
    }

    public boolean isDisponibilidad() {
        return disponibilidad;
    }

    public void setDisponibilidad(boolean disponibilidad) {
        this.disponibilidad = disponibilidad;
    }

    @Override
    public String toString() {
        final StringBuilder sb = new StringBuilder("Libro {");
        sb.append("Código del libro = ").append(codigo);
        sb.append(", Titulo = ").append(titulo).append('\'');
        sb.append(", Autor = ").append(autor).append('\'');
        sb.append(", Disponibilidad = ").append(disponibilidad);
        sb.append('}');
        return sb.toString();
    }
}