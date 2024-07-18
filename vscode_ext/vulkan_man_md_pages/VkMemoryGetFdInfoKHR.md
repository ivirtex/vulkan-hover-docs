# VkMemoryGetFdInfoKHR(3) Manual Page

## Name

VkMemoryGetFdInfoKHR - Structure describing a POSIX FD memory export
operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryGetFdInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_external_memory_fd
typedef struct VkMemoryGetFdInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkDeviceMemory                        memory;
    VkExternalMemoryHandleTypeFlagBits    handleType;
} VkMemoryGetFdInfoKHR;
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

The properties of the file descriptor exported depend on the value of
`handleType`. See
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
for a description of the properties of the defined external memory
handle types.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The size of the exported file <strong>may</strong> be larger than the
size requested by <a
href="VkMemoryAllocateInfo.html">VkMemoryAllocateInfo</a>::<code>allocationSize</code>.
If <code>handleType</code> is
<code>VK_EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT</code>, then the
application <strong>can</strong> query the file’s actual size with <a
href="https://man7.org/linux/man-pages/man2/lseek.2.html"><code>lseek</code></a>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkMemoryGetFdInfoKHR-handleType-00671"
  id="VUID-VkMemoryGetFdInfoKHR-handleType-00671"></a>
  VUID-VkMemoryGetFdInfoKHR-handleType-00671  
  `handleType` **must** have been included in
  [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes`
  when `memory` was created

- <a href="#VUID-VkMemoryGetFdInfoKHR-handleType-00672"
  id="VUID-VkMemoryGetFdInfoKHR-handleType-00672"></a>
  VUID-VkMemoryGetFdInfoKHR-handleType-00672  
  `handleType` **must** be
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT` or
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_DMA_BUF_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryGetFdInfoKHR-sType-sType"
  id="VUID-VkMemoryGetFdInfoKHR-sType-sType"></a>
  VUID-VkMemoryGetFdInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_GET_FD_INFO_KHR`

- <a href="#VUID-VkMemoryGetFdInfoKHR-pNext-pNext"
  id="VUID-VkMemoryGetFdInfoKHR-pNext-pNext"></a>
  VUID-VkMemoryGetFdInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMemoryGetFdInfoKHR-memory-parameter"
  id="VUID-VkMemoryGetFdInfoKHR-memory-parameter"></a>
  VUID-VkMemoryGetFdInfoKHR-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
  handle

- <a href="#VUID-VkMemoryGetFdInfoKHR-handleType-parameter"
  id="VUID-VkMemoryGetFdInfoKHR-handleType-parameter"></a>
  VUID-VkMemoryGetFdInfoKHR-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_memory_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_fd.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetMemoryFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryFdKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryGetFdInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
