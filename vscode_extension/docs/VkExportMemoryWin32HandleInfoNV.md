# VkExportMemoryWin32HandleInfoNV(3) Manual Page

## Name

VkExportMemoryWin32HandleInfoNV - Specify security attributes and access rights for Win32 memory handles



## [](#_c_specification)C Specification

When [VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfoNV.html)::`handleTypes` includes `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_NV`, add a `VkExportMemoryWin32HandleInfoNV` structure to the `pNext` chain of the [VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfoNV.html) structure to specify security attributes and access rights for the memory object’s external handle.

The `VkExportMemoryWin32HandleInfoNV` structure is defined as:

```c++
// Provided by VK_NV_external_memory_win32
typedef struct VkExportMemoryWin32HandleInfoNV {
    VkStructureType               sType;
    const void*                   pNext;
    const SECURITY_ATTRIBUTES*    pAttributes;
    DWORD                         dwAccess;
} VkExportMemoryWin32HandleInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pAttributes` is a pointer to a Windows `SECURITY_ATTRIBUTES` structure specifying security attributes of the handle.
- `dwAccess` is a `DWORD` specifying access rights of the handle.

## [](#_description)Description

If this structure is not present, or if `pAttributes` is `NULL`, default security descriptor values will be used, and child processes created by the application will not inherit the handle, as described in the MSDN documentation for “Synchronization Object Security and Access Rights”1. Further, if the structure is not present, the access rights will be

`DXGI_SHARED_RESOURCE_READ` | `DXGI_SHARED_RESOURCE_WRITE`

1

[https://docs.microsoft.com/en-us/windows/win32/sync/synchronization-object-security-and-access-rights](https://docs.microsoft.com/en-us/windows/win32/sync/synchronization-object-security-and-access-rights)

Valid Usage (Implicit)

- [](#VUID-VkExportMemoryWin32HandleInfoNV-sType-sType)VUID-VkExportMemoryWin32HandleInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV`
- [](#VUID-VkExportMemoryWin32HandleInfoNV-pAttributes-parameter)VUID-VkExportMemoryWin32HandleInfoNV-pAttributes-parameter  
  If `pAttributes` is not `NULL`, `pAttributes` **must** be a valid pointer to a valid `SECURITY_ATTRIBUTES` value

## [](#_see_also)See Also

[VK\_NV\_external\_memory\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_memory_win32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportMemoryWin32HandleInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0