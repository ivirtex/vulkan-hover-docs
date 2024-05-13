# VkMemoryHostPointerPropertiesEXT(3) Manual Page

## Name

VkMemoryHostPointerPropertiesEXT - Properties of external memory host
pointer



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryHostPointerPropertiesEXT` structure is defined as:

``` c
// Provided by VK_EXT_external_memory_host
typedef struct VkMemoryHostPointerPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           memoryTypeBits;
} VkMemoryHostPointerPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `memoryTypeBits` is a bitmask containing one bit set for every memory
  type which the specified host pointer **can** be imported as.

## <a href="#_description" class="anchor"></a>Description

The value returned by `memoryTypeBits` **should** only include bits that
identify memory types which are host visible. Implementations **may**
include bits that identify memory types which are not host visible.
Behavior for imported pointers of such types is defined by <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#host-memory-import-non-visible-type"
target="_blank" rel="noopener"></a>[VkImportMemoryHostPointerInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryHostPointerInfoEXT.html).

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryHostPointerPropertiesEXT-sType-sType"
  id="VUID-VkMemoryHostPointerPropertiesEXT-sType-sType"></a>
  VUID-VkMemoryHostPointerPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_MEMORY_HOST_POINTER_PROPERTIES_EXT`

- <a href="#VUID-VkMemoryHostPointerPropertiesEXT-pNext-pNext"
  id="VUID-VkMemoryHostPointerPropertiesEXT-pNext-pNext"></a>
  VUID-VkMemoryHostPointerPropertiesEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_external_memory_host](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_external_memory_host.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetMemoryHostPointerPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryHostPointerPropertiesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryHostPointerPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
