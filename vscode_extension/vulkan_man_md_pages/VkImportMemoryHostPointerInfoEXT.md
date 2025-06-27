# VkImportMemoryHostPointerInfoEXT(3) Manual Page

## Name

VkImportMemoryHostPointerInfoEXT - Import memory from a host pointer



## <a href="#_c_specification" class="anchor"></a>C Specification

To import memory from a host pointer, add a
[VkImportMemoryHostPointerInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryHostPointerInfoEXT.html)
structure to the `pNext` chain of the
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) structure. The
`VkImportMemoryHostPointerInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_external_memory_host
typedef struct VkImportMemoryHostPointerInfoEXT {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalMemoryHandleTypeFlagBits    handleType;
    void*                                 pHostPointer;
} VkImportMemoryHostPointerInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleType` is a
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value specifying the handle type.

- `pHostPointer` is the host pointer to import from.

## <a href="#_description" class="anchor"></a>Description

Importing memory from a host pointer shares ownership of the memory
between the host and the Vulkan implementation. The application **can**
continue to access the memory through the host pointer but it is the
application’s responsibility to synchronize device and non-device access
to the payload as defined in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-device-hostaccess"
target="_blank" rel="noopener">Host Access to Device Memory Objects</a>.

Applications **can** import the same payload into multiple instances of
Vulkan and multiple times into a given Vulkan instance. However,
implementations **may** fail to import the same payload multiple times
into a given physical device due to platform constraints.

Importing memory from a particular host pointer **may** not be possible
due to additional platform-specific restrictions beyond the scope of
this specification in which case the implementation **must** fail the
memory import operation with the error code
`VK_ERROR_INVALID_EXTERNAL_HANDLE_KHR`.

Whether device memory objects imported from a host pointer hold a
reference to their payload is undefined. As such, the application
**must** ensure that the imported memory range remains valid and
accessible for the lifetime of the imported memory object.

Implementations **may** support importing host pointers for memory types
which are not host-visible. In this case, after a successful call to
[vkAllocateMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateMemory.html), the memory range imported
from `pHostPointer` **must** not be accessed by the application until
the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) has been destroyed. Memory
contents for the host memory becomes undefined on import, and is left
undefined after the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) has been
destroyed. Applications **must** also not access host memory which is
mapped to the same physical memory as `pHostPointer`, but mapped to a
different host pointer while the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)
handle is valid. Implementations running on general-purpose operating
systems **should** not support importing host pointers for memory types
which are not host-visible.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Using host pointers to back non-host visible allocations is a
platform-specific use case, and applications should not attempt to do
this unless instructed by the platform.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkImportMemoryHostPointerInfoEXT-handleType-01747"
  id="VUID-VkImportMemoryHostPointerInfoEXT-handleType-01747"></a>
  VUID-VkImportMemoryHostPointerInfoEXT-handleType-01747  
  If `handleType` is not `0`, it **must** be supported for import, as
  reported in
  [VkExternalMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryProperties.html)

- <a href="#VUID-VkImportMemoryHostPointerInfoEXT-handleType-01748"
  id="VUID-VkImportMemoryHostPointerInfoEXT-handleType-01748"></a>
  VUID-VkImportMemoryHostPointerInfoEXT-handleType-01748  
  If `handleType` is not `0`, it **must** be
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT` or
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT`

- <a href="#VUID-VkImportMemoryHostPointerInfoEXT-pHostPointer-01749"
  id="VUID-VkImportMemoryHostPointerInfoEXT-pHostPointer-01749"></a>
  VUID-VkImportMemoryHostPointerInfoEXT-pHostPointer-01749  
  `pHostPointer` **must** be a pointer aligned to an integer multiple of
  `VkPhysicalDeviceExternalMemoryHostPropertiesEXT`::`minImportedHostPointerAlignment`

- <a href="#VUID-VkImportMemoryHostPointerInfoEXT-handleType-01750"
  id="VUID-VkImportMemoryHostPointerInfoEXT-handleType-01750"></a>
  VUID-VkImportMemoryHostPointerInfoEXT-handleType-01750  
  If `handleType` is
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT`,
  `pHostPointer` **must** be a pointer to `allocationSize` number of
  bytes of host memory, where `allocationSize` is the member of the
  `VkMemoryAllocateInfo` structure this structure is chained to

- <a href="#VUID-VkImportMemoryHostPointerInfoEXT-handleType-01751"
  id="VUID-VkImportMemoryHostPointerInfoEXT-handleType-01751"></a>
  VUID-VkImportMemoryHostPointerInfoEXT-handleType-01751  
  If `handleType` is
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT`,
  `pHostPointer` **must** be a pointer to `allocationSize` number of
  bytes of host mapped foreign memory, where `allocationSize` is the
  member of the `VkMemoryAllocateInfo` structure this structure is
  chained to

Valid Usage (Implicit)

- <a href="#VUID-VkImportMemoryHostPointerInfoEXT-sType-sType"
  id="VUID-VkImportMemoryHostPointerInfoEXT-sType-sType"></a>
  VUID-VkImportMemoryHostPointerInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMPORT_MEMORY_HOST_POINTER_INFO_EXT`

- <a href="#VUID-VkImportMemoryHostPointerInfoEXT-handleType-parameter"
  id="VUID-VkImportMemoryHostPointerInfoEXT-handleType-parameter"></a>
  VUID-VkImportMemoryHostPointerInfoEXT-handleType-parameter  
  `handleType` **must** be a valid
  [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html)
  value

- <a href="#VUID-VkImportMemoryHostPointerInfoEXT-pHostPointer-parameter"
  id="VUID-VkImportMemoryHostPointerInfoEXT-pHostPointer-parameter"></a>
  VUID-VkImportMemoryHostPointerInfoEXT-pHostPointer-parameter  
  `pHostPointer` **must** be a pointer value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_external_memory_host](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_external_memory_host.html),
[VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImportMemoryHostPointerInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
