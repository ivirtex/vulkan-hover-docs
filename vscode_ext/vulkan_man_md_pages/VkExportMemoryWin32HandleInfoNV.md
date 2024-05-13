# VkExportMemoryWin32HandleInfoNV(3) Manual Page

## Name

VkExportMemoryWin32HandleInfoNV - Specify security attributes and access
rights for Win32 memory handles



## <a href="#_c_specification" class="anchor"></a>C Specification

When
[VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfoNV.html)::`handleTypes`
includes `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_NV`, add a
`VkExportMemoryWin32HandleInfoNV` structure to the `pNext` chain of the
[VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfoNV.html)
structure to specify security attributes and access rights for the
memory object’s external handle.

The `VkExportMemoryWin32HandleInfoNV` structure is defined as:

``` c
// Provided by VK_NV_external_memory_win32
typedef struct VkExportMemoryWin32HandleInfoNV {
    VkStructureType               sType;
    const void*                   pNext;
    const SECURITY_ATTRIBUTES*    pAttributes;
    DWORD                         dwAccess;
} VkExportMemoryWin32HandleInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pAttributes` is a pointer to a Windows `SECURITY_ATTRIBUTES`
  structure specifying security attributes of the handle.

- `dwAccess` is a `DWORD` specifying access rights of the handle.

## <a href="#_description" class="anchor"></a>Description

If this structure is not present, or if `pAttributes` is set to `NULL`,
default security descriptor values will be used, and child processes
created by the application will not inherit the handle, as described in
the MSDN documentation for “Synchronization Object Security and Access
Rights”<sup>1</sup>. Further, if the structure is not present, the
access rights will be

`DXGI_SHARED_RESOURCE_READ` \| `DXGI_SHARED_RESOURCE_WRITE`

1  
<a
href="https://docs.microsoft.com/en-us/windows/win32/sync/synchronization-object-security-and-access-rights"
class="bare">https://docs.microsoft.com/en-us/windows/win32/sync/synchronization-object-security-and-access-rights</a>

Valid Usage (Implicit)

- <a href="#VUID-VkExportMemoryWin32HandleInfoNV-sType-sType"
  id="VUID-VkExportMemoryWin32HandleInfoNV-sType-sType"></a>
  VUID-VkExportMemoryWin32HandleInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV`

- <a href="#VUID-VkExportMemoryWin32HandleInfoNV-pAttributes-parameter"
  id="VUID-VkExportMemoryWin32HandleInfoNV-pAttributes-parameter"></a>
  VUID-VkExportMemoryWin32HandleInfoNV-pAttributes-parameter  
  If `pAttributes` is not `NULL`, `pAttributes` **must** be a valid
  pointer to a valid `SECURITY_ATTRIBUTES` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_external_memory_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_external_memory_win32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMemoryWin32HandleInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
