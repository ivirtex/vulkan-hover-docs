# VkPhysicalDeviceCopyMemoryIndirectPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceCopyMemoryIndirectPropertiesNV - Structure describing
supported queues for indirect copy



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceCopyMemoryIndirectPropertiesNV` structure is
defined as:

``` c
// Provided by VK_NV_copy_memory_indirect
typedef struct VkPhysicalDeviceCopyMemoryIndirectPropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    VkQueueFlags       supportedQueues;
} VkPhysicalDeviceCopyMemoryIndirectPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `supportedQueues` is a bitmask of
  [VkQueueFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFlagBits.html) indicating the queues on which
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#indirect-copies"
  target="_blank" rel="noopener">indirect copy commands</a> are
  supported.

## <a href="#_description" class="anchor"></a>Description

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-indirectCopy"
target="_blank" rel="noopener"><code>indirectCopy</code></a> feature is
supported, `supportedQueues` **must** return at least one supported
queue.

If the `VkPhysicalDeviceCopyMemoryIndirectPropertiesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceCopyMemoryIndirectPropertiesNV-sType-sType"
  id="VUID-VkPhysicalDeviceCopyMemoryIndirectPropertiesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceCopyMemoryIndirectPropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COPY_MEMORY_INDIRECT_PROPERTIES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_copy_memory_indirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_copy_memory_indirect.html),
[VkQueueFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceCopyMemoryIndirectPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
