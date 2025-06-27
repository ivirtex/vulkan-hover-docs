# VkImportMemoryWin32HandleInfoKHR(3) Manual Page

## Name

VkImportMemoryWin32HandleInfoKHR - Import Win32 memory created on the
same physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To import memory from a Windows handle, add a
[VkImportMemoryWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryWin32HandleInfoKHR.html)
structure to the `pNext` chain of the
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) structure.

The `VkImportMemoryWin32HandleInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_memory_win32
typedef struct VkImportMemoryWin32HandleInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalMemoryHandleTypeFlagBits    handleType;
    HANDLE                                handle;
    LPCWSTR                               name;
} VkImportMemoryWin32HandleInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleType` is a
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value specifying the type of `handle` or `name`.

- `handle` is `NULL` or the external handle to import.

- `name` is `NULL` or a null-terminated UTF-16 string naming the payload
  to import.

## <a href="#_description" class="anchor"></a>Description

Importing memory object payloads from Windows handles does not transfer
ownership of the handle to the Vulkan implementation. For handle types
defined as NT handles, the application **must** release handle ownership
using the `CloseHandle` system call when the handle is no longer needed.
For handle types defined as NT handles, the imported memory object holds
a reference to its payload.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Non-NT handle import operations do not add a reference to their
associated payload. If the original object owning the payload is
destroyed, all resources and handles sharing that payload will become
invalid.</p></td>
</tr>
</tbody>
</table>

Applications **can** import the same payload into multiple instances of
Vulkan, into the same instance from which it was exported, and multiple
times into a given Vulkan instance. In all cases, each import operation
**must** create a distinct `VkDeviceMemory` object.

Valid Usage

- <a href="#VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00658"
  id="VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00658"></a>
  VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00658  
  If `handleType` is not `0`, it **must** be supported for import, as
  reported by
  [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatProperties.html)
  or [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalBufferProperties.html)

- <a href="#VUID-VkImportMemoryWin32HandleInfoKHR-handle-00659"
  id="VUID-VkImportMemoryWin32HandleInfoKHR-handle-00659"></a>
  VUID-VkImportMemoryWin32HandleInfoKHR-handle-00659  
  The memory from which `handle` was exported, or the memory named by
  `name` **must** have been created on the same underlying physical
  device as `device`

- <a href="#VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00660"
  id="VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00660"></a>
  VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00660  
  If `handleType` is not `0`, it **must** be defined as an NT handle or
  a global share handle

- <a href="#VUID-VkImportMemoryWin32HandleInfoKHR-handleType-01439"
  id="VUID-VkImportMemoryWin32HandleInfoKHR-handleType-01439"></a>
  VUID-VkImportMemoryWin32HandleInfoKHR-handleType-01439  
  If `handleType` is not
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT`,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT`,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT`, or
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT`, `name` **must**
  be `NULL`

- <a href="#VUID-VkImportMemoryWin32HandleInfoKHR-handleType-01440"
  id="VUID-VkImportMemoryWin32HandleInfoKHR-handleType-01440"></a>
  VUID-VkImportMemoryWin32HandleInfoKHR-handleType-01440  
  If `handleType` is not `0` and `handle` is `NULL`, `name` **must**
  name a valid memory resource of the type specified by `handleType`

- <a href="#VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00661"
  id="VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00661"></a>
  VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00661  
  If `handleType` is not `0` and `name` is `NULL`, `handle` **must** be
  a valid handle of the type specified by `handleType`

- <a href="#VUID-VkImportMemoryWin32HandleInfoKHR-handle-01441"
  id="VUID-VkImportMemoryWin32HandleInfoKHR-handle-01441"></a>
  VUID-VkImportMemoryWin32HandleInfoKHR-handle-01441  
  If `handle` is not `NULL`, `name` **must** be `NULL`

- <a href="#VUID-VkImportMemoryWin32HandleInfoKHR-handle-01518"
  id="VUID-VkImportMemoryWin32HandleInfoKHR-handle-01518"></a>
  VUID-VkImportMemoryWin32HandleInfoKHR-handle-01518  
  If `handle` is not `NULL`, it **must** obey any requirements listed
  for `handleType` in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-memory-handle-types-compatibility"
  target="_blank" rel="noopener">external memory handle types
  compatibility</a>

- <a href="#VUID-VkImportMemoryWin32HandleInfoKHR-name-01519"
  id="VUID-VkImportMemoryWin32HandleInfoKHR-name-01519"></a>
  VUID-VkImportMemoryWin32HandleInfoKHR-name-01519  
  If `name` is not `NULL`, it **must** obey any requirements listed for
  `handleType` in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#external-memory-handle-types-compatibility"
  target="_blank" rel="noopener">external memory handle types
  compatibility</a>

Valid Usage (Implicit)

- <a href="#VUID-VkImportMemoryWin32HandleInfoKHR-sType-sType"
  id="VUID-VkImportMemoryWin32HandleInfoKHR-sType-sType"></a>
  VUID-VkImportMemoryWin32HandleInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR`

- <a href="#VUID-VkImportMemoryWin32HandleInfoKHR-handleType-parameter"
  id="VUID-VkImportMemoryWin32HandleInfoKHR-handleType-parameter"></a>
  VUID-VkImportMemoryWin32HandleInfoKHR-handleType-parameter  
  If `handleType` is not `0`, `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_memory_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_win32.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportMemoryWin32HandleInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
