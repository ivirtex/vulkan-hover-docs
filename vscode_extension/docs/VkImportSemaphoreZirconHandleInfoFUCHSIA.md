# VkImportSemaphoreZirconHandleInfoFUCHSIA(3) Manual Page

## Name

VkImportSemaphoreZirconHandleInfoFUCHSIA - Structure specifying Zircon event handle to import to a semaphore



## [](#_c_specification)C Specification

The `VkImportSemaphoreZirconHandleInfoFUCHSIA` structure is defined as:

```c++
// Provided by VK_FUCHSIA_external_semaphore
typedef struct VkImportSemaphoreZirconHandleInfoFUCHSIA {
    VkStructureType                          sType;
    const void*                              pNext;
    VkSemaphore                              semaphore;
    VkSemaphoreImportFlags                   flags;
    VkExternalSemaphoreHandleTypeFlagBits    handleType;
    zx_handle_t                              zirconHandle;
} VkImportSemaphoreZirconHandleInfoFUCHSIA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `semaphore` is the semaphore into which the payload will be imported.
- `flags` is a bitmask of [VkSemaphoreImportFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreImportFlagBits.html) specifying additional parameters for the semaphore payload import operation.
- `handleType` is a [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) value specifying the type of `zirconHandle`.
- `zirconHandle` is the external handle to import.

## [](#_description)Description

The handle types supported by `handleType` are:

Table 1. Handle Types Supported by `VkImportSemaphoreZirconHandleInfoFUCHSIA`    Handle Type Transference Permanence Supported

`VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_ZIRCON_EVENT_BIT_FUCHSIA`

Reference

Temporary,Permanent

Valid Usage

- [](#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-handleType-04765)VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-handleType-04765  
  `handleType` **must** be a value included in the [Handle Types Supported by `VkImportSemaphoreZirconHandleInfoFUCHSIA`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphore-handletypes-fuchsia) table
- [](#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-zirconHandle-04766)VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-zirconHandle-04766  
  `zirconHandle` **must** obey any requirements listed for `handleType` in [external semaphore handle types compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#external-semaphore-handle-types-compatibility)
- [](#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-zirconHandle-04767)VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-zirconHandle-04767  
  `zirconHandle` **must** have `ZX_RIGHTS_BASIC` and `ZX_RIGHTS_SIGNAL` rights
- [](#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-semaphoreType-04768)VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-semaphoreType-04768  
  The [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreTypeCreateInfo.html)::`semaphoreType` field **must** not be `VK_SEMAPHORE_TYPE_TIMELINE`

Valid Usage (Implicit)

- [](#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-sType-sType)VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_ZIRCON_HANDLE_INFO_FUCHSIA`
- [](#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-pNext-pNext)VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-semaphore-parameter)VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handle
- [](#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-flags-parameter)VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-flags-parameter  
  `flags` **must** be a valid combination of [VkSemaphoreImportFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreImportFlagBits.html) values
- [](#VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-handleType-parameter)VUID-VkImportSemaphoreZirconHandleInfoFUCHSIA-handleType-parameter  
  `handleType` **must** be a valid [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) value

Host Synchronization

- Host access to `semaphore` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_FUCHSIA\_external\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_external_semaphore.html), [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkSemaphoreImportFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreImportFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkImportSemaphoreZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkImportSemaphoreZirconHandleFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportSemaphoreZirconHandleInfoFUCHSIA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0