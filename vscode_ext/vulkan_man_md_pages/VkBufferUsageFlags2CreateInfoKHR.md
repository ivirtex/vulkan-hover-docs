# VkBufferUsageFlags2CreateInfoKHR(3) Manual Page

## Name

VkBufferUsageFlags2CreateInfoKHR - Extended buffer usage flags



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferUsageFlags2CreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance5
typedef struct VkBufferUsageFlags2CreateInfoKHR {
    VkStructureType           sType;
    const void*               pNext;
    VkBufferUsageFlags2KHR    usage;
} VkBufferUsageFlags2CreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `usage` is a bitmask of
  [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html) specifying
  allowed usages of the buffer.

## <a href="#_description" class="anchor"></a>Description

If this structure is included in the `pNext` chain of a buffer creation
structure, `usage` is used instead of the corresponding `usage` value
passed in that creation structure, allowing additional usage flags to be
specified. If this structure is included in the `pNext` chain of a
buffer query structure, the usage flags of the buffer are returned in
`usage` of this structure, and the usage flags representable in `usage`
of the buffer query structure are also returned in that field.

Valid Usage (Implicit)

- <a href="#VUID-VkBufferUsageFlags2CreateInfoKHR-sType-sType"
  id="VUID-VkBufferUsageFlags2CreateInfoKHR-sType-sType"></a>
  VUID-VkBufferUsageFlags2CreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BUFFER_USAGE_FLAGS_2_CREATE_INFO_KHR`

- <a href="#VUID-VkBufferUsageFlags2CreateInfoKHR-usage-parameter"
  id="VUID-VkBufferUsageFlags2CreateInfoKHR-usage-parameter"></a>
  VUID-VkBufferUsageFlags2CreateInfoKHR-usage-parameter  
  `usage` **must** be a valid combination of
  [VkBufferUsageFlagBits2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits2KHR.html) values

- <a href="#VUID-VkBufferUsageFlags2CreateInfoKHR-usage-requiredbitmask"
  id="VUID-VkBufferUsageFlags2CreateInfoKHR-usage-requiredbitmask"></a>
  VUID-VkBufferUsageFlags2CreateInfoKHR-usage-requiredbitmask  
  `usage` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkBufferUsageFlags2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags2KHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferUsageFlags2CreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
