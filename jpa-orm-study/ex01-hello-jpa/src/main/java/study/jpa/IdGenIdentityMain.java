package study.jpa;

import study.jpa.domain.MemberIdentity;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.EntityTransaction;
import javax.persistence.Persistence;

public class IdGenIdentityMain {
    public static void main(String[] args) {
        EntityManagerFactory emf = Persistence.createEntityManagerFactory("hello");
        EntityManager em = emf.createEntityManager();
        EntityTransaction tx = em.getTransaction();
        tx.begin();

        try {
            // 비영속 (new/transient)
            MemberIdentity member = new MemberIdentity();
            member.setUsername("HelloJPA");

            // 영속 (managed)
            System.out.println("persist ==================================================");
            em.persist(member);

            // 동일성
            System.out.println("find 1 ==================================================");
            MemberIdentity findMember1 = em.find(MemberIdentity.class, 1L);

            System.out.println("find 2 ==================================================");
            MemberIdentity findMember2 = em.find(MemberIdentity.class, 1L);

            System.out.println(findMember1 == findMember2);

            // 준영속, 회원 엔티티를 영속성 컨텍스트에서 분리
            System.out.println("detach ==================================================");
            em.detach(member);

            // 삭제
            System.out.println("remove ==================================================");
            em.remove(member);

            System.out.println("commit ==================================================");
            tx.commit();
        } catch (Exception e) {
            tx.rollback();
            e.printStackTrace();
        } finally {
            em.close();
            emf.close();
        }
    }
}
