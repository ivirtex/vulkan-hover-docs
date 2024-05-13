# VkMemoryGetWin32HandleInfoKHR(3) Manual Page

## Name

VkMemoryGetWin32HandleInfoKHR - Structure describing a Win32 handle
memory export operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryGetWin32HandleInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_memory_win32
typedef struct VkMemoryGetWin32HandleInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkDeviceMemory                        memory;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkMemoryGetWin32HandleInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memory` is the memory object from which the handle will be exported.

- `handleType` is a
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value specifying the type of handle requested.

## <a href="#_description" class="anchor"></a>Description

The properties of the handle returned depend on the value of
`handleType`. See
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
for a description of the properties of the defined external memory
handle types.

Valid Usage

- <a href="#VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00662"
  id="VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00662"></a>
  VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00662  
  `handleType` **must** have been included in
  [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes`
  when `memory` was created

- <a href="#VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00663"
  id="VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00663"></a>
  VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00663  
  If `handleType` is defined as an NT handle,
  [vkGetMemoryWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryWin32HandleKHR.html) **must**
  be called no more than once for each valid unique combination of
  `memory` and `handleType`

- <a href="#VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00664"
  id="VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00664"></a>
  VUID-VkMemoryGetWin32HandleInfoKHR-handleType-00664  
  `handleType` **must** be defined as an NT handle or a global share
  handle

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryGetWin32HandleInfoKHR-sType-sType"
  id="VUID-VkMemoryGetWin32HandleInfoKHR-sType-sType"></a>
  VUID-VkMemoryGetWin32HandleInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR`

- <a href="#VUID-VkMemoryGetWin32HandleInfoKHR-pNext-pNext"
  id="VUID-VkMemoryGetWin32HandleInfoKHR-pNext-pNext"></a>
  VUID-VkMemoryGetWin32HandleInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMemoryGetWin32HandleInfoKHR-memory-parameter"
  id="VUID-VkMemoryGetWin32HandleInfoKHR-memory-parameter"></a>
  VUID-VkMemoryGetWin32HandleInfoKHR-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a href="#VUID-VkMemoryGetWin32HandleInfoKHR-handleType-parameter"
  id="VUID-VkMemoryGetWin32HandleInfoKHR-handleType-parameter"></a>
  VUID-VkMemoryGetWin32HandleInfoKHR-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_memory_win32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_win32.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetMemoryWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryWin32HandleKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryGetWin32HandleInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
