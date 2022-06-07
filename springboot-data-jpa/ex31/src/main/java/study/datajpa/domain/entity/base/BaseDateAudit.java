package study.datajpa.domain.entity.base;

import lombok.Getter;
import org.hibernate.envers.Audited;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import javax.persistence.EntityListeners;
import javax.persistence.MappedSuperclass;
import java.time.LocalDateTime;

@Audited
@EntityListeners(AuditingEntityListener.class)
@MappedSuperclass
@Getter
public abstract class BaseDateAudit extends CreatedDateAudit {
    @LastModifiedDate
    private LocalDateTime lastModifiedDate;
}