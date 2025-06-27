# VkMemoryHeapFlagBits(3) Manual Page

## Name

VkMemoryHeapFlagBits - Bitmask specifying attribute flags for a heap



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in [VkMemoryHeap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHeap.html)::`flags`,
indicating attribute flags for the heap, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkMemoryHeapFlagBits {
    VK_MEMORY_HEAP_DEVICE_LOCAL_BIT = 0x00000001,
  // Provided by VK_VERSION_1_1
    VK_MEMORY_HEAP_MULTI_INSTANCE_BIT = 0x00000002,
  // Provided by VK_KHR_device_group_creation
    VK_MEMORY_HEAP_MULTI_INSTANCE_BIT_KHR = VK_MEMORY_HEAP_MULTI_INSTANCE_BIT,
} VkMemoryHeapFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_MEMORY_HEAP_DEVICE_LOCAL_BIT` specifies that the heap corresponds
  to device-local memory. Device-local memory **may** have different
  performance characteristics than host-local memory, and **may**
  support different memory property flags.

- `VK_MEMORY_HEAP_MULTI_INSTANCE_BIT` specifies that in a logical device
  representing more than one physical device, there is a per-physical
  device instance of the heap memory. By default, an allocation from
  such a heap will be replicated to each physical device’s instance of
  the heap.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkMemoryHeapFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHeapFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryHeapFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
