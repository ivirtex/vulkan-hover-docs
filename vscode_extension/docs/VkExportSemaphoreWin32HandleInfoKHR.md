# VkExportSemaphoreWin32HandleInfoKHR(3) Manual Page

## Name

VkExportSemaphoreWin32HandleInfoKHR - Structure specifying additional attributes of Windows handles exported from a semaphore



## [](#_c_specification)C Specification

To specify additional attributes of NT handles exported from a semaphore, add a `VkExportSemaphoreWin32HandleInfoKHR` structure to the `pNext` chain of the [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html) structure. The `VkExportSemaphoreWin32HandleInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_external_semaphore_win32
typedef struct VkExportSemaphoreWin32HandleInfoKHR {
    VkStructureType               sType;
    const void*                   pNext;
    const SECURITY_ATTRIBUTES*    pAttributes;
    DWORD                         dwAccess;
    LPCWSTR                       name;
} VkExportSemaphoreWin32HandleInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pAttributes` is a pointer to a Windows `SECURITY_ATTRIBUTES` structure specifying security attributes of the handle.
- `dwAccess` is a `DWORD` specifying access rights of the handle.
- `name` is a null-terminated UTF-16 string to associate with the underlying synchronization primitive referenced by NT handles exported from the created semaphore.

## [](#_description)Description

If [VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportSemaphoreCreateInfo.html) is not included in the same `pNext` chain, this structure is ignored.

If [VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportSemaphoreCreateInfo.html) is included in the `pNext` chain of [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html) with a Windows `handleType`, but either `VkExportSemaphoreWin32HandleInfoKHR` is not included in the `pNext` chain, or it is included but `pAttributes` is `NULL`, default security descriptor values will be used, and child processes created by the application will not inherit the handle, as described in the MSDN documentation for “Synchronization Object Security and Access Rights”1. Further, if the structure is not present, the access rights used depend on the handle type.

For handles of the following types:

`VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT`

The implementation **must** ensure the access rights allow both signal and wait operations on the semaphore.

For handles of the following types:

`VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT`

The access rights **must** be:

`GENERIC_ALL`

1

[https://docs.microsoft.com/en-us/windows/win32/sync/synchronization-object-security-and-access-rights](https://docs.microsoft.com/en-us/windows/win32/sync/synchronization-object-security-and-access-rights)

Valid Usage

- [](#VUID-VkExportSemaphoreWin32HandleInfoKHR-handleTypes-01125)VUID-VkExportSemaphoreWin32HandleInfoKHR-handleTypes-01125  
  If [VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportSemaphoreCreateInfo.html)::`handleTypes` does not include `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_OPAQUE_WIN32_BIT` or `VK_EXTERNAL_SEMAPHORE_HANDLE_TYPE_D3D12_FENCE_BIT`, `VkExportSemaphoreWin32HandleInfoKHR` **must** not be included in the `pNext` chain of [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html)

Valid Usage (Implicit)

- [](#VUID-VkExportSemaphoreWin32HandleInfoKHR-sType-sType)VUID-VkExportSemaphoreWin32HandleInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR`
- [](#VUID-VkExportSemaphoreWin32HandleInfoKHR-pAttributes-parameter)VUID-VkExportSemaphoreWin32HandleInfoKHR-pAttributes-parameter  
  If `pAttributes` is not `NULL`, `pAttributes` **must** be a valid pointer to a valid `SECURITY_ATTRIBUTES` value

## [](#_see_also)See Also

[VK\_KHR\_external\_semaphore\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore_win32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportSemaphoreWin32HandleInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0