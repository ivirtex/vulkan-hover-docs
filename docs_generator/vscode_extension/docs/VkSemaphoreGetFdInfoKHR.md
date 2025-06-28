# VkSemaphoreGetFdInfoKHR(3) Manual Page

## Name

VkSemaphoreGetFdInfoKHR - Structure describing a POSIX FD semaphore export operation



## [](#_c_specification)C Specification

The `VkSemaphoreGetFdInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_external_semaphore_fd
typedef struct VkSemaphoreGetFdInfoKHR {
    VkStructureType                          sType;
    const void*                              pNext;
    VkSemaphore                              semaphore;
    VkExternalSemaphoreHandleTypeFlagBits    handleType;
} VkSemaphoreGetFdInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `semaphore` is the semaphore from which state will be exported.
- `handleType` is a [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) value specifying the type of handle requested.

## [](#_description)Description

The properties of the file descriptor returned depend on the value of `handleType`. See [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) for a description of the properties of the defined external semaphore handle types.

Valid Usage

- [](#VUID-VkSemaphoreGetFdInfoKHR-handleType-01132)VUID-VkSemaphoreGetFdInfoKHR-handleType-01132  
  `handleType` **must** have been included in [VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportSemaphoreCreateInfo.html)::`handleTypes` when `semaphore`’s current payload was created
- [](#VUID-VkSemaphoreGetFdInfoKHR-semaphore-01133)VUID-VkSemaphoreGetFdInfoKHR-semaphore-01133  
  `semaphore` **must** not currently have its payload replaced by an imported payload as described below in [Importing Semaphore Payloads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-importing) unless that imported payload’s handle type was included in [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreProperties.html)::`exportFromImportedHandleTypes` for `handleType`
- [](#VUID-VkSemaphoreGetFdInfoKHR-handleType-01134)VUID-VkSemaphoreGetFdInfoKHR-handleType-01134  
  If `handleType` refers to a handle type with copy payload transference semantics, as defined below in [Importing Semaphore Payloads](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-importing), there **must** be no queue waiting on `semaphore`
- [](#VUID-VkSemaphoreGetFdInfoKHR-handleType-01135)VUID-VkSemaphoreGetFdInfoKHR-handleType-01135  
  If `handleType` refers to a handle type with copy payload transference semantics, `semaphore` **must** be signaled, or have an associated [semaphore signal operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-signaling) pending execution
- [](#VUID-VkSemaphoreGetFdInfoKHR-handleType-01136)VUID-VkSemaphoreGetFdInfoKHR-handleType-01136  
  `handleType` **must** be defined as a POSIX file descriptor handle
- [](#VUID-VkSemaphoreGetFdInfoKHR-handleType-03253)VUID-VkSemaphoreGetFdInfoKHR-handleType-03253  
  If `handleType` refers to a handle type with copy payload transference semantics, `semaphore` **must** have been created with a [VkSemaphoreType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreType.html) of `VK_SEMAPHORE_TYPE_BINARY`
- [](#VUID-VkSemaphoreGetFdInfoKHR-handleType-03254)VUID-VkSemaphoreGetFdInfoKHR-handleType-03254  
  If `handleType` refers to a handle type with copy payload transference semantics, `semaphore` **must** have an associated semaphore signal operation that has been submitted for execution and any semaphore signal operations on which it depends **must** have also been submitted for execution

Valid Usage (Implicit)

- [](#VUID-VkSemaphoreGetFdInfoKHR-sType-sType)VUID-VkSemaphoreGetFdInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_GET_FD_INFO_KHR`
- [](#VUID-VkSemaphoreGetFdInfoKHR-pNext-pNext)VUID-VkSemaphoreGetFdInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkSemaphoreGetFdInfoKHR-semaphore-parameter)VUID-VkSemaphoreGetFdInfoKHR-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handle
- [](#VUID-VkSemaphoreGetFdInfoKHR-handleType-parameter)VUID-VkSemaphoreGetFdInfoKHR-handleType-parameter  
  `handleType` **must** be a valid [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_KHR\_external\_semaphore\_fd](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore_fd.html), [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetSemaphoreFdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSemaphoreFdKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSemaphoreGetFdInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0