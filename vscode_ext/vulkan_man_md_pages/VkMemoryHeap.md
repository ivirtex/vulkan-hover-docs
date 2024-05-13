# VkMemoryHeap(3) Manual Page

## Name

VkMemoryHeap - Structure specifying a memory heap



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryHeap` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkMemoryHeap {
    VkDeviceSize         size;
    VkMemoryHeapFlags    flags;
} VkMemoryHeap;
```

## <a href="#_members" class="anchor"></a>Members

- `size` is the total memory size in bytes in the heap.

- `flags` is a bitmask of
  [VkMemoryHeapFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHeapFlagBits.html) specifying attribute
  flags for the heap.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkMemoryHeapFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHeapFlags.html),
[VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryHeap"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
