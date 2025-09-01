# VkSemaphoreGetZirconHandleInfoFUCHSIA(3) Manual Page

## Name

VkSemaphoreGetZirconHandleInfoFUCHSIA - Structure describing a Zircon event handle semaphore export operation



## [](#_c_specification)C Specification

The `VkSemaphoreGetZirconHandleInfoFUCHSIA` structure is defined as:

```c++
// Provided by VK_FUCHSIA_external_semaphore
typedef struct VkSemaphoreGetZirconHandleInfoFUCHSIA {
    VkStructureType                          sType;
    const void*                              pNext;
    VkSemaphore                              semaphore;
    VkExternalSemaphoreHandleTypeFlagBits    handleType;
} VkSemaphoreGetZirconHandleInfoFUCHSIA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `semaphore` is the semaphore from which state will be exported.
- `handleType` is a [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) value specifying the type of handle requested.

## [](#_description)Description

The properties of the Zircon event handle returned depend on the value of `handleType`. See [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) for a description of the properties of the defined external semaphore handle types.

Valid Usage

- [](#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04758)VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04758  
  `handleType` **must** have been included in [VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportSemaphoreCreateInfo.html)::`handleTypes` when `semaphore`’s current payload was created
- [](#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-04759)VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-04759  
  `semaphore` **must** not currently have its payload replaced by an imported payload as described below in [Importing Semaphore Payloads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-importing) unless that imported payload’s handle type was included in [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreProperties.html)::`exportFromImportedHandleTypes` for `handleType`
- [](#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04760)VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04760  
  If `handleType` refers to a handle type with copy payload transference semantics, as defined below in [Importing Semaphore Payloads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-importing), there **must** be no queue waiting on `semaphore`
- [](#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04761)VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04761  
  If `handleType` refers to a handle type with copy payload transference semantics, `semaphore` **must** be signaled, or have an associated [semaphore signal operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-signaling) pending execution
- [](#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04762)VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-04762  
  `handleType` **must** be defined as a Zircon event handle
- [](#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-04763)VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-04763  
  `semaphore` **must** have been created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY`

Valid Usage (Implicit)

- [](#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-sType-sType)VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_GET_ZIRCON_HANDLE_INFO_FUCHSIA`
- [](#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-pNext-pNext)VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-parameter)VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handle
- [](#VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-parameter)VUID-VkSemaphoreGetZirconHandleInfoFUCHSIA-handleType-parameter  
  `handleType` **must** be a valid [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_FUCHSIA\_external\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_external_semaphore.html), [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetSemaphoreZirconHandleFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSemaphoreZirconHandleFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSemaphoreGetZirconHandleInfoFUCHSIA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0