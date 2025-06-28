# VkImportSemaphoreWin32HandleInfoKHR(3) Manual Page

## Name

VkImportSemaphoreWin32HandleInfoKHR - Structure specifying Windows handle to import to a semaphore



## [](#_c_specification)C Specification

The `VkImportSemaphoreWin32HandleInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_external_semaphore_win32
typedef struct VkImportSemaphoreWin32HandleInfoKHR {
    VkStructureType                          sType;
    const void*                              pNext;
    VkSemaphore                              semaphore;
    VkSemaphoreImportFlags                   flags;
    VkExternalSemaphoreHandleTypeFlagBits    handleType;
    HANDLE                                   handle;
    LPCWSTR                                  name;
} VkImportSemaphoreWin32HandleInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `semaphore` is the semaphore into which the payload will be imported.
- `flags` is a bitmask of [VkSemaphoreImportFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreImportFlagBits.html) specifying additional parameters for the semaphore payload import operation.
- `handleType` is a [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) value specifying the type of `handle`.
- `handle` is `NULL` or the external handle to import.
- `name` is `NULL` or a null-terminated UTF-16 string naming the underlying synchronization primitive to import.

## [](#_description)Description

The handle types supported by `handleType` are:

Table 1. Handle Types Supported by `VkImportSemaphoreWin32HandleInfoKHR`    Handle Type Transference Permanence Supported

`VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT`

Reference

Temporary,Permanent

`VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT`

Reference

Temporary,Permanent

`VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT`

Reference

Temporary,Permanent

Valid Usage

- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01140)VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01140  
  `handleType` **must** be a value included in the [Handle Types Supported by `VkImportSemaphoreWin32HandleInfoKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphore-handletypes-win32) table
- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01466)VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01466  
  If `handleType` is not `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT` or `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT`, `name` **must** be `NULL`
- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01467)VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01467  
  If `handle` is `NULL`, `name` **must** name a valid synchronization primitive of the type specified by `handleType`
- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01468)VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-01468  
  If `name` is `NULL`, `handle` **must** be a valid handle of the type specified by `handleType`
- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-handle-01469)VUID-VkImportSemaphoreWin32HandleInfoKHR-handle-01469  
  If `handle` is not `NULL`, `name` **must** be `NULL`
- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-handle-01542)VUID-VkImportSemaphoreWin32HandleInfoKHR-handle-01542  
  If `handle` is not `NULL`, it **must** obey any requirements listed for `handleType` in [external semaphore handle types compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#external-semaphore-handle-types-compatibility)
- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-name-01543)VUID-VkImportSemaphoreWin32HandleInfoKHR-name-01543  
  If `name` is not `NULL`, it **must** obey any requirements listed for `handleType` in [external semaphore handle types compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#external-semaphore-handle-types-compatibility)
- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-03261)VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-03261  
  If `handleType` is `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT` or `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT`, the [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html)::`flags` field **must** match that of the semaphore from which `handle` or `name` was exported
- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-03262)VUID-VkImportSemaphoreWin32HandleInfoKHR-handleType-03262  
  If `handleType` is `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT` or `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT`, the [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreTypeCreateInfo.html)::`semaphoreType` field **must** match that of the semaphore from which `handle` or `name` was exported
- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-flags-03322)VUID-VkImportSemaphoreWin32HandleInfoKHR-flags-03322  
  If `flags` contains `VK_SEMAPHORE_IMPORT_TEMPORARY_BIT`, the [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreTypeCreateInfo.html)::`semaphoreType` field of the semaphore from which `handle` or `name` was exported **must** not be `VK_SEMAPHORE_TYPE_TIMELINE`

Valid Usage (Implicit)

- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-sType-sType)VUID-VkImportSemaphoreWin32HandleInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR`
- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-pNext-pNext)VUID-VkImportSemaphoreWin32HandleInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-semaphore-parameter)VUID-VkImportSemaphoreWin32HandleInfoKHR-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handle
- [](#VUID-VkImportSemaphoreWin32HandleInfoKHR-flags-parameter)VUID-VkImportSemaphoreWin32HandleInfoKHR-flags-parameter  
  `flags` **must** be a valid combination of [VkSemaphoreImportFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreImportFlagBits.html) values

Host Synchronization

- Host access to `semaphore` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_external\_semaphore\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore_win32.html), [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkSemaphoreImportFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreImportFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkImportSemaphoreWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkImportSemaphoreWin32HandleKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportSemaphoreWin32HandleInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0