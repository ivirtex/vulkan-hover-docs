# VkMemoryBarrier(3) Manual Page

## Name

VkMemoryBarrier - Structure specifying a global memory barrier



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryBarrier` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkMemoryBarrier {
    VkStructureType    sType;
    const void*        pNext;
    VkAccessFlags      srcAccessMask;
    VkAccessFlags      dstAccessMask;
} VkMemoryBarrier;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcAccessMask` is a bitmask of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) specifying a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
  target="_blank" rel="noopener">source access mask</a>.

- `dstAccessMask` is a bitmask of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) specifying a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
  target="_blank" rel="noopener">destination access mask</a>.

## <a href="#_description" class="anchor"></a>Description

The first <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> is limited to access
types in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
target="_blank" rel="noopener">source access mask</a> specified by
`srcAccessMask`.

The second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-access-scopes"
target="_blank" rel="noopener">access scope</a> is limited to access
types in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-access-masks"
target="_blank" rel="noopener">destination access mask</a> specified by
`dstAccessMask`.

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryBarrier-sType-sType"
  id="VUID-VkMemoryBarrier-sType-sType"></a>
  VUID-VkMemoryBarrier-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_BARRIER`

- <a href="#VUID-VkMemoryBarrier-pNext-pNext"
  id="VUID-VkMemoryBarrier-pNext-pNext"></a>
  VUID-VkMemoryBarrier-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMemoryBarrier-srcAccessMask-parameter"
  id="VUID-VkMemoryBarrier-srcAccessMask-parameter"></a>
  VUID-VkMemoryBarrier-srcAccessMask-parameter  
  `srcAccessMask` **must** be a valid combination of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) values

- <a href="#VUID-VkMemoryBarrier-dstAccessMask-parameter"
  id="VUID-VkMemoryBarrier-dstAccessMask-parameter"></a>
  VUID-VkMemoryBarrier-dstAccessMask-parameter  
  `dstAccessMask` **must** be a valid combination of
  [VkAccessFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAccessFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdPipelineBarrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPipelineBarrier.html),
[vkCmdWaitEvents](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWaitEvents.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryBarrier"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
