# VkExportMemoryWin32HandleInfoKHR(3) Manual Page

## Name

VkExportMemoryWin32HandleInfoKHR - Structure specifying additional
attributes of Windows handles exported from a memory



## <a href="#_c_specification" class="anchor"></a>C Specification

To specify additional attributes of NT handles exported from a memory
object, add a
[VkExportMemoryWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryWin32HandleInfoKHR.html)
structure to the `pNext` chain of the
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) structure. The
`VkExportMemoryWin32HandleInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_memory_win32
typedef struct VkExportMemoryWin32HandleInfoKHR {
    VkStructureType               sType;
    const void*                   pNext;
    const SECURITY_ATTRIBUTES*    pAttributes;
    DWORD                         dwAccess;
    LPCWSTR                       name;
} VkExportMemoryWin32HandleInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pAttributes` is a pointer to a Windows `SECURITY_ATTRIBUTES`
  structure specifying security attributes of the handle.

- `dwAccess` is a `DWORD` specifying access rights of the handle.

- `name` is a null-terminated UTF-16 string to associate with the
  payload referenced by NT handles exported from the created memory.

## <a href="#_description" class="anchor"></a>Description

If [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html) is not
included in the same `pNext` chain, this structure is ignored.

If [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html) is
included in the `pNext` chain of
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) with a Windows
`handleType`, but either `VkExportMemoryWin32HandleInfoKHR` is not
included in the `pNext` chain, or it is included but `pAttributes` is
set to `NULL`, default security descriptor values will be used, and
child processes created by the application will not inherit the handle,
as described in the MSDN documentation for “Synchronization Object
Security and Access Rights”<sup>1</sup>. Further, if the structure is
not present, the access rights used depend on the handle type.

For handles of the following types:

- `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT`

The implementation **must** ensure the access rights allow read and
write access to the memory.

1  
<a
href="https://docs.microsoft.com/en-us/windows/win32/sync/synchronization-object-security-and-access-rights"
class="bare">https://docs.microsoft.com/en-us/windows/win32/sync/synchronization-object-security-and-access-rights</a>

Valid Usage

- <a href="#VUID-VkExportMemoryWin32HandleInfoKHR-handleTypes-00657"
  id="VUID-VkExportMemoryWin32HandleInfoKHR-handleTypes-00657"></a>
  VUID-VkExportMemoryWin32HandleInfoKHR-handleTypes-00657  
  If
  [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes`
  does not include `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT`, a
  `VkExportMemoryWin32HandleInfoKHR` structure **must** not be included
  in the `pNext` chain of
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html)

Valid Usage (Implicit)

- <a href="#VUID-VkExportMemoryWin32HandleInfoKHR-sType-sType"
  id="VUID-VkExportMemoryWin32HandleInfoKHR-sType-sType"></a>
  VUID-VkExportMemoryWin32HandleInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR`

- <a href="#VUID-VkExportMemoryWin32HandleInfoKHR-pAttributes-parameter"
  id="VUID-VkExportMemoryWin32HandleInfoKHR-pAttributes-parameter"></a>
  VUID-VkExportMemoryWin32HandleInfoKHR-pAttributes-parameter  
  If `pAttributes` is not `NULL`, `pAttributes` **must** be a valid
  pointer to a valid `SECURITY_ATTRIBUTES` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_memory_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_win32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMemoryWin32HandleInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
